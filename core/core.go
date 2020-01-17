package core

import (
	"time"

	"../gui"
	"../utils"
	"../widgets"
	"github.com/gotk3/gotk3/gtk"
	lua "github.com/yuin/gopher-lua"
)

var (
	window *gtk.Window
	isRun  bool
)

func Stop() {
	isRun = false
}

func Run(win *gtk.Window, path string) {
	window = win
	isRun = true

	utils.InfoMessage("Initializing Core...")
	core := lua.NewState()
	core.PreloadModule("Widgets", widgets.Loader)
	gui.Init(win)
	core.PreloadModule("GUI", gui.Loader)
	utils.InfoMessage("Initialized!")
	if err := core.DoFile(path + "/Core/Core.lua"); err != nil {
		utils.FatalMessage(err.Error())
	}
	for isRun {
		time.Sleep(8 * time.Millisecond)
	}
	defer core.Close()
}
