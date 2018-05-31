package build

import (
	"fmt"
	"html/template"
	"io"
	"path/filepath"
)

// Mission contains information need to customize boot template .ks files
type Mission struct {
	Filename string
	Args     []string
}

func test(w io.Writer, kspsrc string, mission *Mission) {
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
