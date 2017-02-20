package gotype

import (
	"reflect"
	"unsafe"

	structform "github.com/urso/go-structform"
)

type liftedReflUnfolder struct{ unfolder ptrUnfolder }

type unfolderReflSlice struct {
	elem reflUnfolder
}

type unfolderReflSliceStart struct {
	unfolderErrArrayStart
}

type unfolderReflMap struct {
	shared unfolderReflMapShared
}

type unfolderReflMapShared struct {
	waitKey  *unfolderReflMapOnKey
	waitElem *unfolderReflMapOnElem
}

type unfolderReflMapStart struct {
	unfolderErrObjectStart
}

type unfolderReflMapOnKey struct {
	unfolderErrExpectKey
	shared *unfolderReflMapShared
}

type unfolderReflMapOnElem struct {
	shared *unfolderReflMapShared
	elem   reflUnfolder
}

var (
	_singletonUnfolderReflSliceStart = &unfolderReflSliceStart{}
	_singletonUnfolderReflMapStart   = &unfolderReflMapStart{}
)

func liftGoUnfolder(u ptrUnfolder) *liftedReflUnfolder { return &liftedReflUnfolder{u} }

func (u *liftedReflUnfolder) initState(ctx *unfoldCtx, v reflect.Value) {
	ptr := unsafe.Pointer(v.Pointer())
	u.unfolder.initState(ctx, ptr)
}

func newUnfolderReflSlice(elem reflUnfolder) *unfolderReflSlice {
	return &unfolderReflSlice{elem}
}

func (u *unfolderReflSlice) initState(ctx *unfoldCtx, v reflect.Value) {
	ctx.value.push(v)
	ctx.unfolder.push(u)
	ctx.idx.push(0)
	ctx.unfolder.push(_singletonUnfolderReflSliceStart)
}

func (u *unfolderReflSlice) cleanup(ctx *unfoldCtx) {
	ctx.idx.pop()
	ctx.value.pop()
	ctx.unfolder.pop()
}

func (u *unfolderReflSliceStart) cleanup(ctx *unfoldCtx) {
	ctx.unfolder.pop()
}

func (u *unfolderReflSliceStart) OnArrayStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	ptr := ctx.value.current
	v := ptr.Elem()

	if l < 0 {
		l = 0
	}

	if v.IsNil() && l > 0 {
		v.Set(reflect.MakeSlice(v.Type(), l, l))
	} else if !v.IsNil() && l < v.Len() {
		v.SetLen(l)
	}

	u.cleanup(ctx)
	return nil
}

func (u *unfolderReflSlice) OnArrayFinished(ctx *unfoldCtx) error {
	u.cleanup(ctx)
	return nil
}

func (u *unfolderReflSlice) prepare(ctx *unfoldCtx) reflect.Value {
	// make space for some more element
	ptr := ctx.value.current
	idx := &ctx.idx
	v := ptr.Elem()

	switch {
	case v.Len() > idx.current:

	case v.Cap() > idx.current:
		v.SetLen(idx.current + 1)

	default:
		v.Set(reflect.Append(v, reflect.Zero(v.Type().Elem())))
	}

	elem := v.Index(idx.current).Addr()
	idx.current++

	return elem
}

func (u *unfolderReflSlice) OnArrayStart(ctx *unfoldCtx, l int, bt structform.BaseType) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	return ctx.unfolder.current.OnArrayStart(ctx, l, bt)
}

func (u *unfolderReflSlice) OnChildArrayDone(ctx *unfoldCtx) error {
	return nil
}

func (u *unfolderReflSlice) OnObjectStart(ctx *unfoldCtx, l int, bt structform.BaseType) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	return ctx.unfolder.current.OnObjectStart(ctx, l, bt)
}

func (u *unfolderReflSlice) OnChildObjectDone(_ *unfoldCtx) error { return nil }

func (u *unfolderReflSlice) OnObjectFinished(_ *unfoldCtx) error {
	return errUnsupported
}

func (u *unfolderReflSlice) OnKey(_ *unfoldCtx, _ string) error {
	return errUnsupported
}

func newUnfolderReflMap(elem reflUnfolder) *unfolderReflMap {
	u := &unfolderReflMap{}
	u.shared.waitKey = &unfolderReflMapOnKey{shared: &u.shared}
	u.shared.waitElem = &unfolderReflMapOnElem{shared: &u.shared, elem: elem}
	return u
}

func (u *unfolderReflMap) initState(ctx *unfoldCtx, v reflect.Value) {
	ctx.value.push(v)
	ctx.unfolder.push(u.shared.waitKey)
	ctx.unfolder.push(_singletonUnfolderReflMapStart)
}

func (u *unfolderReflMapStart) OnObjectStart(ctx *unfoldCtx, l int, bt structform.BaseType) error {
	ctx.unfolder.pop()
	return nil
}

func (u *unfolderReflMapOnKey) OnKey(ctx *unfoldCtx, key string) error {
	ctx.key.push(key)
	ctx.unfolder.current = u.shared.waitElem
	return nil
}

func (u *unfolderReflMapOnKey) OnObjectFinished(ctx *unfoldCtx) error {
	ctx.unfolder.pop()
	ctx.value.pop()
	return nil
}

func (u *unfolderReflMapOnElem) prepare(ctx *unfoldCtx) reflect.Value {
	ptr := ctx.value.current
	v := ptr.Elem()

	target := reflect.New(v.Type().Elem())
	ctx.value.push(target)
	return target
}

func (u *unfolderReflMapOnElem) process(ctx *unfoldCtx) {
	ptr := ctx.value.pop()
	v := ptr.Elem()

	ptr = ctx.value.current
	m := ptr.Elem()
	m.SetMapIndex(reflect.ValueOf(ctx.key.pop()), v)

	ctx.unfolder.current = u.shared.waitKey
}

func (u *unfolderReflMapOnElem) OnArrayStart(ctx *unfoldCtx, l int, bt structform.BaseType) error {

	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	return ctx.unfolder.current.OnArrayStart(ctx, l, bt)
}

func (u *unfolderReflMapOnElem) OnChildArrayDone(ctx *unfoldCtx) error {
	u.process(ctx)
	return nil
}

func (u *unfolderReflMapOnElem) OnArrayFinished(_ *unfoldCtx) error {
	return errUnsupported
}

func (u *unfolderReflMapOnElem) OnObjectStart(ctx *unfoldCtx, l int, bt structform.BaseType) error {

	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	return ctx.unfolder.current.OnObjectStart(ctx, l, bt)
}

func (u *unfolderReflMapOnElem) OnChildObjectDone(ctx *unfoldCtx) error {
	u.process(ctx)
	return nil
}

func (u *unfolderReflMapOnElem) OnKey(_ *unfoldCtx, _ string) error {
	return errExpectedObjectValue
}

func (u *unfolderReflMapOnElem) OnObjectFinished(_ *unfoldCtx) error {
	return errExpectedObjectValue
}
