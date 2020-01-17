package menu

import (
	"../../utils"
	"github.com/gotk3/gotk3/gtk"
	"github.com/yuin/gopher-lua"
)

type Menu struct {
	menuBar *gtk.MenuBar
}

func RegisterMenuType(L *lua.LState) int {
	mt := L.NewTypeMetatable("Menu")
	L.SetGlobal("Menu", mt)
	
	L.SetField(mt, "new", L.NewFunction(newMenu))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), menuMethods))
	return 1
}

var menuMethods = map[string]lua.LGFunction{
	"addMenuItem": luaAddMenuItem,
	"getParent": luaGetParent,
}

func newMenu(L *lua.LState) int {
	gtkMenuBar, err := gtk.MenuBarNew()
	if err != nil {
		utils.FatalMessage(err.Error())
	}

	menuBar := &Menu{gtkMenuBar}
	ud := L.NewUserData()
	ud.Value = menuBar
	L.SetMetatable(ud, L.GetTypeMetatable("Menu"))
	L.Push(ud)
	return 1
}

func checkMenu(L *lua.LState) *Menu{
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Menu); ok {
		return v
	}
	L.ArgError(1, "Menu expected")
	return nil
}

func luaAddMenuItem(L *lua.LState) int {
	m := checkMenu(L)
	menu, err := gtk.MenuItemNewWithLabel(L.ToString(2))
	if err != nil {
		utils.FatalMessage(err.Error())
	}
	
	m.menuBar.Append(menu)
	return 1
}

func luaGetParent(L *lua.LState) int {
	m := checkMenu(L)
	ud := L.NewUserData()
	ud.Value =  m.menuBar
	L.Push(ud)
	return 1
}