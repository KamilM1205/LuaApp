package projectsystem

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"

	"../utils"
)

//Use xml структура элемента настройки проекта движка
type Use struct {
	Name  string `xml:"name,attr"`
	Value bool
}

//Project xml структура LuaApp проекта
type Project struct {
	XMLName     xml.Name `xml:"project"`
	Name        string   `xml:"name,attr"`
	ProjVersion string   `xml:"version,attr"`
	LuaApp      App
}

//Uses массив структур Use
type Uses []Use

//App структура xml которая хранит настройки движка
type App struct {
	LuaAppVersion string `xml:"version,attr"`
	Settings      *Uses  `xml:"use"`
}

const (
	androidMain string = ``
	pcMain      string = `pc = require("pc")

	function Init()
		pc.SetSize(800, 600)
		pc.SetTitle("Test")
		pc.Run()
	end
	
	pc.Init()`
	coreMain string = `gui = require("GUI")
	widgets = require("Widgets")
	
	function btn_click(event)
		print(event)
	end
	
	function Init() 
		screen = Screen.new()
		v = Vertical.new()
		v:setHAlign(Align.Fill)
		v:setVAlign(Align.Fill)
		label = Label.new("Hello, world")
		label:setHAlign(Align.Center)
		label:setVAlign(Align.Center)
		v:addWidget(label:getParent())
		screen:addWidget(v:getParent())
		gui.setScreen(screen)
	end
	gui.Init()`
)

//NewProject функция для создания нового LuaApp проекта
func NewProject(projectName string, projectVersion string, android bool, pc bool, assets bool) error {
	err := os.Mkdir("Projects/"+projectName, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.Mkdir("Projects/"+projectName+"/Core", os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.Create("Projects/" + projectName + "/Core/" + "Core.lua")
	if err != nil {
		return err
	}
	file.WriteString(coreMain)
	file.Close()

	if android {
		err := os.Mkdir("Projects/"+projectName+"/Android", os.ModePerm)
		if err != nil {
			return err
		}
		file, err := os.Create("Projects/" + projectName + "/Android/" + "main.lua")
		if err != nil {
			return err
		}
		file.WriteString(androidMain)
		file.Close()
	}

	if pc {
		err := os.Mkdir("Projects/"+projectName+"/PC", os.ModePerm)
		if err != nil {
			return err
		}
		file, err := os.Create("Projects/" + projectName + "/PC/" + "main.lua")
		if err != nil {
			return err
		}
		file.WriteString(pcMain)
		file.Close()
	}

	if assets {
		err := os.Mkdir("Projects/"+projectName+"/Assets", os.ModePerm)
		if err != nil {
			return err
		}
	}

	file, err = os.Create("Projects/" + projectName + "/project.xml")
	if err != nil {
		return err
	}
	defer file.Close()

	marshal, err := projectMarshal(projectName, projectVersion, android, pc, assets)
	if err != nil {
		return err
	}
	_, err = file.Write([]byte(marshal))
	if err != nil {
		return err
	}
	return nil
}

//OpenProject функция для открытия LuaApp проекта
func OpenProject(filename string) (*Project, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return nil, err
	}

	project, err := projectUnmarshal(string(bs))
	return project, err
}

//IsProject функция проверяющая является ли структура проектом
func IsProject(projectPath string) (bool, error) {
	f, err := ioutil.ReadDir(projectPath)
	if err != nil {
		return false, err
	}
	for _, file := range f {
		if !file.IsDir() && file.Name() == "project.xml" {
			return true, nil
		}
	}
	return false, errors.New("The " + projectPath + " directory does not have a project.xml file.")
}

//SearchProjects функция, которая возвращает список проектов
func SearchProjects(dir string) (*[]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var projectList []string

	for _, v := range files {
		if v.IsDir() {
			isProject, err := IsProject(dir + v.Name())
			if err != nil {
				return nil, err
			}
			if isProject {
				projectList = append(projectList, v.Name())
			}
		}
	}
	return &projectList, nil
}

func projectUnmarshal(code string) (*Project, error) {
	v := &Project{Name: "none"}
	err := xml.Unmarshal([]byte(code), v)
	return v, err
}

func projectMarshal(projectName string, projectVersion string, android bool, pc bool, assets bool) (string, error) {
	useAndroid := Use{Name: utils.Android, Value: android}
	usePC := Use{Name: utils.PC, Value: pc}
	useAssets := Use{Name: utils.Assets, Value: assets}
	var uses *Uses = new(Uses)
	*uses = append(*uses, useAndroid)
	*uses = append(*uses, usePC)
	*uses = append(*uses, useAssets)
	luaApp := App{LuaAppVersion: utils.GetEngineVersion(), Settings: uses}
	v := &Project{Name: projectName, ProjVersion: projectVersion, LuaApp: luaApp}
	out, err := xml.MarshalIndent(v, " ", "    ")
	return string(out), err
}
