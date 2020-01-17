package layout

import (
	"../../utils"
	"github.com/gotk3/gotk3/gtk"
	lua "github.com/yuin/gopher-lua"
)

//Vertical is struct of LuaApp vertical layout
type Vertical struct {
	lay   *gtk.Box
}

//RegisterVerticalType is func for register vertical layout
func RegisterVerticalType(L *lua.LState) int {
	mt := L.NewTypeMetatable("Vertical")
	L.SetGlobal("Vertical", mt)

	L.SetField(mt, "new", L.NewFunction(newVertical))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), verticalMethods))
	return 1
}

var verticalMethods = map[string]lua.LGFunction{
	"getParent": luaVGetParent,
	"setHAlign": luaVSetHAlign,
	"setVAlign": luaVSetVAlign,
	"setMarginTop": luaVSetMarginTop,
	"setMarginBottom": luaVSetMarginBottom,
	"setMarginStart": luaVSetMarginStart,
	"setMarginEnd": luaVSetMarginEnd,
	"setWidth": luaVSetWidth,
	"setHeight": luaVSetHeight,
	"addWidget": luaVAddWidget,
}

func call(L *lua.LState, funcName string, arg string) {
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal(funcName),
		NRet:    0,
		Protect: true,
	}, lua.LString(arg)); err != nil {
		utils.FatalMessage(err.Error())
	}
}

func newVertical(L *lua.LState) int {
	gtkVertical, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 2)
	if err != nil {
		utils.FatalMessage(err.Error())
	}

	vertical := &Vertical{gtkVertical}
	ud := L.NewUserData()
	ud.Value = vertical
	L.SetMetatable(ud, L.GetTypeMetatable("Vertical"))
	L.Push(ud)
	return 1
}

func checkVertical(L *lua.LState) *Vertical {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Vertical); ok {
		return v
	}
	L.ArgError(1, "Vertical expected")
	return nil
}

func luaVAddWidget(L *lua.LState) int {
	v := checkVertical(L)
	wt := L.CheckUserData(2).Value.(gtk.IWidget)
	v.lay.Add(wt)
	v.lay.Show()
	return 1
}

func luaVSetHAlign(L *lua.LState) int {
	v := checkVertical(L)
	ud := L.ToUserData(2)
	v.lay.SetHAlign(ud.Value.(gtk.Align))
	v.lay.SetHExpand(true)
	return 1
}

func luaVSetVAlign(L *lua.LState) int {
	v := checkVertical(L)
	ud := L.ToUserData(2)
	v.lay.SetVAlign(ud.Value.(gtk.Align))	
	v.lay.SetVExpand(true)
	return 1
}

func luaVSetMarginTop(L *lua.LState) int {
	v := checkVertical(L)
	v.lay.SetMarginTop(L.ToInt(2))
	return 1
}

func luaVSetMarginBottom(L *lua.LState) int {
	v := checkVertical(L)
	v.lay.SetMarginBottom(L.ToInt(2))
	return 1
}

func luaVSetMarginStart(L *lua.LState) int {
	v := checkVertical(L)
	v.lay.SetMarginStart(L.ToInt(2))
	return 1
}

func luaVSetMarginEnd(L *lua.LState) int {
	v := checkVertical(L)
	v.lay.SetMarginEnd(L.ToInt(2))
	return 1
}

func luaVSetWidth(L *lua.LState) int {
	v := checkVertical(L)
	v.lay.GetAllocation().SetWidth(L.ToInt(2))
	return 1
}

func luaVSetHeight(L *lua.LState) int {
	v := checkVertical(L)
	v.lay.GetAllocation().SetHeight(L.ToInt(2))
	return 1
}

func luaVGetParent(L *lua.LState) int {
	v := checkVertical(L)
	ud := L.NewUserData()
	ud.Value = v.lay
	L.Push(ud)
	return 1
}