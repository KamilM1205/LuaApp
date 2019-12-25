package projectsystem

import (
	"io/ioutil"
	"os"
	"../utils"
)

type ProjectSystem struct {
	projectList []string
}

type FileWorker interface {
	GetProjectList() []string
}

func (p ProjectSystem)GetProjectList() []string{
	return p.projectList
}

func searchProjects(dir string) []string {
	
}