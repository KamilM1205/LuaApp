package button

import (
	"../../utils"
	"github.com/gotk3/gotk3/gtk"
	lua "github.com/yuin/gopher-lua"
)

var (
	f *lua.LFunction
)

//Button is struct of LuaApp button
type Button struct {
	Button    *gtk.Button
}

//RegisterButtonType is func for register Button
func RegisterButtonType(L *lua.LState) int {
	mt := L.NewTypeMetatable("Button")
	L.SetGlobal("Button", mt)

	L.SetField(mt, "new", L.NewFunction(newButton))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), buttonMethods))
	return 1
}

var buttonMethods = map[string]lua.LGFunction{
	"setText":   luaSetText,
	"setClick":  luaSetClick,
	"setHAlign": luaSetHAlign,
	"setVAlign": luaSetVAlign,
	"getParent": luaGetParent,
	"setEnable": luaSetEnable,
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

func newButton(L *lua.LState) int {
	gtkButton, err := gtk.ButtonNew()
	if err != nil {
		utils.FatalMessage(err.Error())
	}
	gtkButton.SetLabel(L.ToString(1))

	button := &Button{gtkButton}
	ud := L.NewUserData()
	ud.Value = button
	L.SetMetatable(ud, L.GetTypeMetatable("Button"))
	L.Push(ud)
	return 1
}

func checkButton(L *lua.LState) *Button {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Button); ok {
		return v
	}
	L.ArgError(1, "Button expected")
	return nil
}

func luaSetText(L *lua.LState) int {
	b := checkButton(L)
	b.Button.SetLabel(L.ToString(1))
	return 1
}

func luaSetClick(L *lua.LState) int {
	b := checkButton(L)
	f = L.ToFunction(2)
	b.Button.Connect("clicked", func() {
		if err := L.CallByParam(lua.P{
			Fn:      f,
			NRet:    0,
			Protect: true,
		}, lua.LString("click")); err != nil {
			utils.FatalMessage(err.Error())
		}
	})
	return 1
}

func luaSetEnable(L *lua.LState) int {
	b := checkButton(L)
	b.Button.SetSensitive(L.ToBool(2))
	return 1
}

func luaSetHAlign(L *lua.LState) int {
	b := checkButton(L)
	ud := L.ToUserData(2)
	b.Button.SetHAlign(ud.Value.(gtk.Align))
	b.Button.SetHExpand(true)
	return 1
}

func luaSetVAlign(L *lua.LState) int {
	b := checkButton(L)
	ud := L.ToUserData(2)
	b.Button.SetVAlign(ud.Value.(gtk.Align))
	b.Button.SetVExpand(true)
	return 1
}

func luaGetParent(L *lua.LState) int {
	b := checkButton(L)
	ud := L.NewUserData()
	ud.Value = b.Button
	L.Push(ud)
	return 1
}
