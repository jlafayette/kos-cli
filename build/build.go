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
func Test(w io.Writer, kspsrc string, b *Boot) {
	boot := filepath.Join(kspsrc, "boot", "templates", "simple.ks")
	tmpl, err := template.ParseFiles(boot)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tmpl.Execute(w, b)
	if err != nil {
		fmt.Println(err)
		return
	}
}

type Mission struct {
	parts []part
}

type part struct {
	name string
	data map[string]string
}

// MakeMission genrates a .ks file for a mission.
func MakeMission(w io.Writer, templateDir string, m Mission) {
	tmpl, err := template.ParseGlob(filepath.Join(templateDir, "*.ks"))
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, p := range m.parts {
		err = tmpl.ExecuteTemplate(w, p.name, p.data)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
