package pc

import (
	"../core"
	"../utils"
	"github.com/gotk3/gotk3/gtk"
	lua "github.com/yuin/gopher-lua"
)

var (
	window *gtk.Window
)

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)
	L.SetField(mod, "name", lua.LString("value"))
	L.Push(mod)
	return 1
}

var exports = map[string]lua.LGFunction{
	//"drawTriangle":  drawTriangle,
	"Init":     luaInit,
	"SetSize":  luaSetSize,
	"SetTitle": luaSetTitle,
	"Run":      luaRun,
}

func call(L *lua.LState, funcName string) {
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal(funcName),
		NRet:    0,
		Protect: true,
	}); err != nil {
		utils.FatalMessage(err.Error())
	}
}

func luaInit(L *lua.LState) int {
	gtk.Init(nil)
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		utils.FatalMessage(err.Error())
	}
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.Connect("destroy", func() {
		core.Stop()
		gtk.MainQuit()
	})
	window = win
	call(L, "Init")
	return 0
}

func luaSetSize(L *lua.LState) int {
	window.SetDefaultSize(L.ToInt(1), L.ToInt(2))
	return 0
}

func luaSetTitle(L *lua.LState) int {
	window.SetTitle(L.ToString(1))
	return 0
}

func luaRun(L *lua.LState) int {
	window.ShowAll()
	go core.Run(window)
	gtk.Main()
	return 0
}
