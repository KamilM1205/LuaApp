package layout

import (
	"../../utils"
	"github.com/gotk3/gotk3/gtk"
	lua "github.com/yuin/gopher-lua"
)

//Vertical is struct of LuaApp vertical layout
type Vertical struct {
	lay   *gtk.Grid
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
	gtkVertical, err := gtk.GridNew()
	if err != nil {
		utils.FatalMessage(err.Error())
	}
	gtkVertical.SetOrientation(gtk.ORIENTATION_VERTICAL)

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

func luaVGetParent(L *lua.LState) int {
	v := checkVertical(L)
	ud := L.NewUserData()
	ud.Value = v.lay
	L.Push(ud)
	return 1
}