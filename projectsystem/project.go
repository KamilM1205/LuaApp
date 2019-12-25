package projectsystem

import (
	"encoding/xml"
	"os"
	"io/ioutil"
)

type Use struct {
	Name string `xml:"name,attr"`
	Value bool
}

type Project struct {
	XMLName xml.Name `xml:"project"`
	Name string `xml:"name,attr"`
	ProjVersion string `xml:"version,attr"`
	App LuaApp
}

type LuaApp struct {
	luaAppVersion string `xml:"version,attr"`
	Settings []Use  `xml:"use"`
}

//NewProject функция для создания нового LuaApp проекта
func NewProject(filename string) (error) {
	
	return nil
}

//OpenProject функция для открытия LuaApp проекта
func OpenProject(filename string) (*Project, error) {
	file, err := os.Open(filename)
	return nil, nil
}

func IsProject(projectName string) (bool, error){
	f, err := ioutil.
	if f.IsDir() {
		file, err := os.Open("Projects/" + f.Name() + "/project.xml")
		file.is
		if err != nil {
			return false, err
		}
	}
}

func projectUnmarshal(code string) (*Project, error) {
	v := Project{Name: "none"}
	err := xml.Unmarshal([]byte(code), &v)
	return &v, err
}