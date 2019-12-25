package gui

import (
	"../utils"
	"./screen"
	"github.com/gotk3/gotk3/gtk"
	lua "github.com/yuin/gopher-lua"
)

var (
	window *gtk.Window
)

//Loader стандартный загрузчик функций lua в модуль
func Loader(L *lua.LState) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), exports)
	// register other stuff
	L.SetField(mod, "name", lua.LString("value"))
	L.Push(mod)
	return 1
}

//Экспортируемые функции
var exports = map[string]lua.LGFunction{
	"Init":      luaInit,
	"popup":     luaPopup,
	"setScreen": luaSetScreen,
}

//Init получение экземпляра окна
func Init(win *gtk.Window) {
	window = win
}

//Функция вызова lua функции
func call(L *lua.LState, funcName string) {
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal(funcName),
		NRet:    0,
		Protect: true,
	}); err != nil {
		utils.FatalMessage(err.Error())
	}
}

//LuaApp функция всплывающего сообщения
func luaPopup(L *lua.LState) int {
	return 1
}

//LuaApp функция инициализации
func luaInit(L *lua.LState) int {
	screen.RegisterScreenType(L)
	call(L, "Init")
	return 1
}

//LuaApp функция установки экрана
func luaSetScreen(L *lua.LState) int {
	scrn := L.ToUserData(1).Value.(*screen.Screen)
	window.Add(scrn.Grid)
	window.ShowAll()
	return 1
}
