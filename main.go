package main

import (
	"./pc"
	"./utils"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	utils.Init()
	utils.InfoMessage("Initializing...")
	L := lua.NewState()
	L.PreloadModule("pc", pc.Loader)
	utils.InfoMessage("Initialized!")
	if err := L.DoFile("Projects/Test/PC/main.lua"); err != nil {
		utils.FatalMessage(err.Error())
	}
	defer L.Close()
}
