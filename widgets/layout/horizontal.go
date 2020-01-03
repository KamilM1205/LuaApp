package layout

import (
	"../../utils"
	"github.com/gotk3/gotk3/gtk"
	lua "github.com/yuin/gopher-lua"
)

//Horizontal is struct of LuaApp vertical layout
type Horizontal struct {
	lay *gtk.Grid
}

//RegisterHorizontalType is func for register vertical layout
func RegisterHorizontalType(L *lua.LState) int {
	mt := L.NewTypeMetatable("Horizontal")
	L.SetGlobal("Horizontal", mt)

	L.SetField(mt, "new", L.NewFunction(newHorizontal))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), horizontalMethods))
	return 1
}

var horizontalMethods = map[string]lua.LGFunction{
	"getParent": luaHGetParent,
	"setHAlign": luaHSetHAlign,
	"setVAlign": luaHSetVAlign,
	"addWidget": luaHAddWidget,
}

func newHorizontal(L *lua.LState) int {
	gtkHorizontal, err := gtk.GridNew()
	if err != nil {
		utils.FatalMessage(err.Error())
	}
	gtkHorizontal.SetOrientation(gtk.ORIENTATION_HORIZONTAL)

	horizontal := &Horizontal{gtkHorizontal}
	ud := L.NewUserData()
	ud.Value = horizontal
	L.SetMetatable(ud, L.GetTypeMetatable("Horizontal"))
	L.Push(ud)
	return 1
}

func checkHorizontal(L *lua.LState) *Horizontal {
	ud := L.CheckUserData(1)
	if h, ok := ud.Value.(*Horizontal); ok {
		return h
	}
	L.ArgError(1, "Horizontal expected")
	return nil
}

func luaHAddWidget(L *lua.LState) int {
	h := checkHorizontal(L)
	wt := L.CheckUserData(2).Value.(gtk.IWidget)
	h.lay.Add(wt)
	h.lay.Show()
	return 1
}

func luaHSetHAlign(L *lua.LState) int {
	h := checkHorizontal(L)
	ud := L.ToUserData(2)
	h.lay.SetHAlign(ud.Value.(gtk.Align))
	return 1
}

func luaHSetVAlign(L *lua.LState) int {
	h := checkHorizontal(L)
	ud := L.ToUserData(2)
	h.lay.SetVAlign(ud.Value.(gtk.Align))
	return 1
}

func luaHGetParent(L *lua.LState) int {
	h := checkHorizontal(L)
	ud := L.NewUserData()
	ud.Value = h.lay
	L.Push(ud)
	return 1
}
