package build

import (
	"fmt"
	"html/template"
	"io"
	"path/filepath"
)

// Boot will write a customized boot.ks file.
func Boot(w io.Writer, boot string, name string) error {
	tmpl, err := template.ParseFiles(boot)
	if err != nil {
		return err
	}
	type data struct {
		Filename string
	}
	b := data{name + ".ks"}
	err = tmpl.Execute(w, b)
	if err != nil {
		return err
	}
	return nil
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
