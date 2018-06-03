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
func MakeMission(w io.Writer, templateDir string) {
	type LaunchWait struct {
		TgtVessel string
		TgtAngle  int
	}
	type Launch struct {
		TgtDir int
		TgtAlt int
	}
	type InitSpace struct {
		Antenna string
	}
	type Circularize struct {
		TgtDir int
	}
	type Intercept struct {
		ApproachRange int
	}
	type Approach struct {
		TgtVessel    string
		TgtRange     int
		OffsetVector string
		Agility      float64
	}
	type WaitForCrew struct {
		TgtCrew int
	}
	type Deorbit struct {
		LKO          bool
		TgtPeriapsis int
	}
	type ReEntry struct {
		DragchuteAlt int
		ParachuteAlt int
	}

	stages := [9]stage{
		stage{
			"launchWait.ks",
			LaunchWait{"A Ship", 35},
		},
		stage{
			"launch.ks",
			Launch{90, 80000},
		},
		stage{
			"initSpace.ks",
			InitSpace{"Communotron 16"},
		},
		stage{
			"circularize.ks",
			Circularize{90},
		},
		stage{
			"intercept.ks",
			Intercept{8000},
		},
		stage{
			"approach.ks",
			Approach{"A Ship", 100, "V(0,0,0)", 1.2},
		},
		stage{
			"waitForCrew.ks",
			WaitForCrew{1},
		},
		stage{
			"deorbit.ks",
			Deorbit{true, 30000},
		},
		stage{
			"reentry.ks",
			ReEntry{2500, 1200},
		},
	}
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
