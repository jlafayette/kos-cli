package build

import (
	"io"
	"io/ioutil"
	"path/filepath"
	"text/template"

	"github.com/jlafayette/kos-cli/plan"
)

// Boot writes a customized .ks file.
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

// Mission writes a customized .ks file.
func Mission(w io.Writer, name, templateDir, planFile string) error {
	tmpl, err := template.ParseGlob(filepath.Join(templateDir, "*.ks"))
	if err != nil {
		return err
	}
	planData, err := ioutil.ReadFile(planFile)
	if err != nil {
		return err
	}
	var p plan.Plan
	err = plan.Unmarshal(planData, &p)
	if err != nil {
		return err
	}
	for _, p := range p.Parts {
		err = tmpl.ExecuteTemplate(w, p.Name, p.Data)
		if err != nil {
			return err
		}
	}
	return nil
}
