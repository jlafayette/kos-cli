package build

import (
	"fmt"
	"html/template"
	"io"
	"path/filepath"
)

// // Part interface
// type part interface {
// 	write(w io.Writer)
// }

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
	path string
	data interface{}
}

// MakeMission genrates a .ks file for a mission.
func MakeMission(w io.Writer, templateDir string) {
	type LaunchWait struct {
		TgtVessel string
		TgtAngle  int
	}
	type Launch struct {
		TgtDir int
		TgtAlt int
	}

	var stages [2]stage
	stages[0] = stage{
		"launchWait.ks",
		filepath.Join(templateDir, "launchWait.ks"),
		LaunchWait{"A Ship", 35},
	}
	stages[1] = stage{
		"launch.ks",
		filepath.Join(templateDir, "launch.ks"),
		Launch{90, 80000},
	}
	tmpl, err := template.ParseFiles(stages[0].path, stages[1].path)
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

	/*
		initSpace
		circularize
		intercept
		approach
		waitForCrew
		reentry

		stage initSpace
			antenna string

		stage circularize
			tgtDir int or string

		stage intercept (for matching orbit)
			approachRange int or string

		stage approach
			tgtVessel    string
			tgtDistance  int or string
			offsetVector vector or string
			agility      float or string

		stage waitForCrew
			tgtCrew  int or string

		stage deorbit
			startLoc  (LKO or HKO)  string
			tgtPeriapsis   int or string

		stage reentry
			dragAlt
			parachuteAlt

	*/
}
