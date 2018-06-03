package build

import (
	"fmt"
	"html/template"
	"io"
	"path/filepath"
)

// Boot contains information need to customize boot template .ks files
type Boot struct {
	Filename string
	Args     []string
}

// Test boot script template
func Test(w io.Writer, kspsrc string, mission *Boot) {
	boot := filepath.Join(kspsrc, "boot", "templates", "simple.ks")
	tmpl, err := template.ParseFiles(boot)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tmpl.Execute(w, mission)
	if err != nil {
		fmt.Println(err)
		return
	}
}

type stage struct {
	name string
	data interface{}
}

// MakeMission genrates a .ks file for a mission.
func MakeMission(w io.Writer, templateDir string, stages []stage) {
	tmpl, err := template.ParseGlob(filepath.Join(templateDir, "*.ks"))
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, s := range stages {
		err = tmpl.ExecuteTemplate(w, s.name, s.data)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
