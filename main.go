package main

import (
	"fmt"
	"os"

	"./pc"
	"./projectsystem"
	"./utils"
	lua "github.com/yuin/gopher-lua"
)

var (
	openProjectName string = ""
)

func newProject(projName string) {
	var (
		version string
		v       string
		android bool
		pc      bool
		assets  bool
	)
	fmt.Print("Project version: ")
	fmt.Scanln(&version)
	fmt.Print("Use android(y/n): ")
	fmt.Scanln(&v)
	if v == "y" {
		android = true
	} else if v == "n" {
		android = false
	} else {
		fmt.Println("Error: invalid answer")
		os.Exit(1)
	}
	fmt.Print("Use PC(y/n): ")
	v = ""
	fmt.Scanln(&v)
	if v == "y" {
		pc = true
	} else if v == "n" {
		pc = false
	} else {
		fmt.Println("Error: invalid answer")
		os.Exit(1)
	}
	fmt.Print("Use assets(y/n): ")
	v = ""
	fmt.Scanln(&v)
	if v == "y" {
		assets = true
	} else if v == "n" {
		assets = false
	} else {
		fmt.Println("Error: invalid answer")
		os.Exit(1)
	}
	err := projectsystem.NewProject(projName, version, android, pc, assets)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func runProject(projName string) {
	utils.InfoMessage("Initializing...")
	L := lua.NewState()
	L.PreloadModule("pc", pc.Loader)
	utils.InfoMessage("Initialized!")
	if err := L.DoFile("Projects/" + projName + "/PC/main.lua"); err != nil {
		utils.FatalMessage(err.Error())
	}
	defer L.Close()
}

func openProject() {
	var cmd string
	for true {
		fmt.Print(">>")
		fmt.Scanln(&cmd)
		if cmd == "close" {
			break
		} else if cmd == "run" {
			runProject(openProjectName)
		}
	}
}

func main() {
	utils.Init()
	if (len(os.Args) > 1) {
		if os.Args[1] == "list" {
			list, err := projectsystem.SearchProjects()
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			for _, v := range *list {
				fmt.Println(v)
			}
		} else if os.Args[1] == "new" {
			if len(os.Args) > 2 && os.Args[2] != "" {
				newProject(os.Args[2])
			} else {
				fmt.Println("Error: LuaApp new project_name")
			}
		} else if os.Args[1] == "open" {
			if len(os.Args) > 2 && os.Args[2] != "" {
				isProj, err := projectsystem.IsProject(os.Args[2])
				if err != nil {
					fmt.Println(err.Error())
					os.Exit(1)
				}
				if isProj {
					openProjectName = os.Args[2]
					openProject()
				}
			} else {
				fmt.Println("Error: LuaApp open project_name")
			}
		}
	}
}
