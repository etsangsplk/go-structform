// This file has been generated from 'unfold_refl.yml', do not edit
package gotype

import "reflect"

func (u *unfolderReflSlice) OnNil(ctx *unfoldCtx) error {
	u.prepare(ctx)
	return nil
}

func (u *unfolderReflSlice) OnByte(ctx *unfoldCtx, v byte) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	return ctx.unfolder.current.OnByte(ctx, v)
}

func (u *unfolderReflSlice) OnBool(ctx *unfoldCtx, v bool) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	return ctx.unfolder.current.OnBool(ctx, v)
}

func (u *unfolderReflSlice) OnString(ctx *unfoldCtx, v string) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	return ctx.unfolder.current.OnString(ctx, v)
}

func (u *unfolderReflSlice) OnUint(ctx *unfoldCtx, v uint) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	return ctx.unfolder.current.OnUint(ctx, v)
}

func (u *unfolderReflSlice) OnUint8(ctx *unfoldCtx, v uint8) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	return ctx.unfolder.current.OnUint8(ctx, v)
}

func (u *unfolderReflSlice) OnUint16(ctx *unfoldCtx, v uint16) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	return ctx.unfolder.current.OnUint16(ctx, v)
}

func (u *unfolderReflSlice) OnUint32(ctx *unfoldCtx, v uint32) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	return ctx.unfolder.current.OnUint32(ctx, v)
}

func (u *unfolderReflSlice) OnUint64(ctx *unfoldCtx, v uint64) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	return ctx.unfolder.current.OnUint64(ctx, v)
}

func (u *unfolderReflSlice) OnInt(ctx *unfoldCtx, v int) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	return ctx.unfolder.current.OnInt(ctx, v)
}

func (u *unfolderReflSlice) OnInt8(ctx *unfoldCtx, v int8) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	return ctx.unfolder.current.OnInt8(ctx, v)
}

func (u *unfolderReflSlice) OnInt16(ctx *unfoldCtx, v int16) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	return ctx.unfolder.current.OnInt16(ctx, v)
}

func (u *unfolderReflSlice) OnInt32(ctx *unfoldCtx, v int32) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	return ctx.unfolder.current.OnInt32(ctx, v)
}

func (u *unfolderReflSlice) OnInt64(ctx *unfoldCtx, v int64) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	return ctx.unfolder.current.OnInt64(ctx, v)
}

func (u *unfolderReflSlice) OnFloat32(ctx *unfoldCtx, v float32) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	return ctx.unfolder.current.OnFloat32(ctx, v)
}

func (u *unfolderReflSlice) OnFloat64(ctx *unfoldCtx, v float64) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	return ctx.unfolder.current.OnFloat64(ctx, v)
}

func (u *unfolderReflMapOnElem) OnNil(ctx *unfoldCtx) error {
	ptr := ctx.value.current
	m := ptr.Elem()
	v := reflect.Zero(m.Type().Elem())
	m.SetMapIndex(reflect.ValueOf(ctx.key.pop()), v)

	ctx.unfolder.current = u.shared.waitKey
	return nil
}

func (u *unfolderReflMapOnElem) OnByte(ctx *unfoldCtx, v byte) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	err := ctx.unfolder.current.OnByte(ctx, v)
	if err == nil {
		u.process(ctx)
	}
	return err
}

func (u *unfolderReflMapOnElem) OnBool(ctx *unfoldCtx, v bool) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	err := ctx.unfolder.current.OnBool(ctx, v)
	if err == nil {
		u.process(ctx)
	}
	return err
}

func (u *unfolderReflMapOnElem) OnString(ctx *unfoldCtx, v string) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	err := ctx.unfolder.current.OnString(ctx, v)
	if err == nil {
		u.process(ctx)
	}
	return err
}

func (u *unfolderReflMapOnElem) OnUint(ctx *unfoldCtx, v uint) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	err := ctx.unfolder.current.OnUint(ctx, v)
	if err == nil {
		u.process(ctx)
	}
	return err
}

func (u *unfolderReflMapOnElem) OnUint8(ctx *unfoldCtx, v uint8) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	err := ctx.unfolder.current.OnUint8(ctx, v)
	if err == nil {
		u.process(ctx)
	}
	return err
}

func (u *unfolderReflMapOnElem) OnUint16(ctx *unfoldCtx, v uint16) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	err := ctx.unfolder.current.OnUint16(ctx, v)
	if err == nil {
		u.process(ctx)
	}
	return err
}

func (u *unfolderReflMapOnElem) OnUint32(ctx *unfoldCtx, v uint32) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	err := ctx.unfolder.current.OnUint32(ctx, v)
	if err == nil {
		u.process(ctx)
	}
	return err
}

func (u *unfolderReflMapOnElem) OnUint64(ctx *unfoldCtx, v uint64) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	err := ctx.unfolder.current.OnUint64(ctx, v)
	if err == nil {
		u.process(ctx)
	}
	return err
}

func (u *unfolderReflMapOnElem) OnInt(ctx *unfoldCtx, v int) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	err := ctx.unfolder.current.OnInt(ctx, v)
	if err == nil {
		u.process(ctx)
	}
	return err
}

func (u *unfolderReflMapOnElem) OnInt8(ctx *unfoldCtx, v int8) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	err := ctx.unfolder.current.OnInt8(ctx, v)
	if err == nil {
		u.process(ctx)
	}
	return err
}

func (u *unfolderReflMapOnElem) OnInt16(ctx *unfoldCtx, v int16) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	err := ctx.unfolder.current.OnInt16(ctx, v)
	if err == nil {
		u.process(ctx)
	}
	return err
}

func (u *unfolderReflMapOnElem) OnInt32(ctx *unfoldCtx, v int32) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	err := ctx.unfolder.current.OnInt32(ctx, v)
	if err == nil {
		u.process(ctx)
	}
	return err
}

func (u *unfolderReflMapOnElem) OnInt64(ctx *unfoldCtx, v int64) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	err := ctx.unfolder.current.OnInt64(ctx, v)
	if err == nil {
		u.process(ctx)
	}
	return err
}

func (u *unfolderReflMapOnElem) OnFloat32(ctx *unfoldCtx, v float32) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	err := ctx.unfolder.current.OnFloat32(ctx, v)
	if err == nil {
		u.process(ctx)
	}
	return err
}

func (u *unfolderReflMapOnElem) OnFloat64(ctx *unfoldCtx, v float64) error {
	elem := u.prepare(ctx)
	u.elem.initState(ctx, elem)
	err := ctx.unfolder.current.OnFloat64(ctx, v)
	if err == nil {
		u.process(ctx)
	}
	return err
}
