package build

import (
	"fmt"
	"html/template"
	"io"
	"path/filepath"
)

type stage interface {
	src() string
}

type mission struct {
	stages []stage
}

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

// MakeMission genrates a .ks file for a mission.
func MakeMission(w io.Writer, kspsrc string) {
	// wait
	/* Things needed:
	src string
	data struct (different for each)
	*/
	waitSrc := filepath.Join(kspsrc, "missions", "templates", "launchWait.ks")
	type LaunchWait struct {
		TgtVessel string
		TgtAngle  int
	}
	lw := LaunchWait{"A Ship", 35}
	tmpl, err := template.ParseFiles(waitSrc)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tmpl.Execute(w, lw)
	if err != nil {
		fmt.Println(err)
	}

	/* stage launch
	tgtDir  int or string
	tgtAlt  int or string
	*/
	launchSrc := filepath.Join(kspsrc, "missions", "templates", "launch.ks")
	type Launch struct {
		TgtDir int
		TgtAlt int
	}
	l := Launch{90, 80000}
	tmpl, err = template.ParseFiles(launchSrc)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tmpl.Execute(w, l)
	if err != nil {
		fmt.Println(err)
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
