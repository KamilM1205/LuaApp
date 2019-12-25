package screen

import (
	"../../utils"
	"github.com/gotk3/gotk3/gtk"
	lua "github.com/yuin/gopher-lua"
)

type Screen struct {
	Grid *gtk.Grid
}

func RegisterScreenType(L *lua.LState) int {
	mt := L.NewTypeMetatable("Screen")
	L.SetGlobal("Screen", mt)

	L.SetField(mt, "new", L.NewFunction(newScreen))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), screenMethods))
	return 1
}

var screenMethods = map[string]lua.LGFunction{
	"addWidget": luaAddWidget,
}

func newScreen(L *lua.LState) int {
	grid, err := gtk.GridNew()
	if err != nil {
		utils.FatalMessage(err.Error())
	}
	grid.SetOrientation(gtk.ORIENTATION_VERTICAL)

	screen := &Screen{grid}
	ud := L.NewUserData()
	ud.Value = screen
	L.SetMetatable(ud, L.GetTypeMetatable("Screen"))
	L.Push(ud)
	return 1
}

func checkScreen(L *lua.LState) *Screen {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Screen); ok {
		return v
	}
	L.ArgError(1, "Screen expected")
	return nil
}

func luaAddWidget(L *lua.LState) int {
	s := checkScreen(L)
	wt := L.CheckUserData(2).Value.(gtk.IWidget)
	s.Grid.Add(wt)
	s.Grid.Show()
	return 1
}
