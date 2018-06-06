package build

import (
	"io"
	"io/ioutil"
	"path/filepath"
	"text/template"

	"github.com/jlafayette/kos-cli/plan"
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

// Mission writes a custom .ks file for a mission.
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

// // MakeMission genrates a .ks file for a mission.
// func MakeMission(w io.Writer, templateDir string, m Mission) {
// 	tmpl, err := template.ParseGlob(filepath.Join(templateDir, "*.ks"))
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	for _, p := range m.parts {
// 		err = tmpl.ExecuteTemplate(w, p.Name, p.Data)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 	}
// }
