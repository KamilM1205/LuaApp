package main

import (
	"os"
	"./pc"
	"./utils"
	"./projectsystem"
	lua "github.com/yuin/gopher-lua"
)

var (
	isOpen bool
)

func newProject(projName string) {
	for true {

	}
}

func runProject(projName string) {
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

func main() {
	if (os.Args[1] == "list") {
		list, err := projectsystem.SearchProjects()
		if err != nil {
			utils.FatalMessage(err.Error())
		}
		for _, v := range list {
			println(v)
		}
	} else if (os.Args[1] == "new") {
		if (os.Args[2] != "") {
			newProject(os.Args[2])
		}
	}
}
