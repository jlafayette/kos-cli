package build

import (
	"fmt"
)

// Structify converts a template name and args into the corresponding struct
func Structify(templateName string, args []string) (stage interface{}, err error) {
	switch templateName {
	case "launchWait.ks":
		stage = LaunchWait{args[0], args[1]}
	case "launch.ks":
		stage = Launch{args[0], args[1]}
	case "initSpace.ks":
		stage = InitSpace{args[0]}
	case "circularize.ks":
		stage = Circularize{args[0]}
	case "intercept.ks":
		stage = Intercept{args[0]}
	case "approach.ks":
		stage = Approach{args[0], args[1], args[2], args[3]}
	case "waitForCrew.ks":
		stage = WaitForCrew{args[0]}
	case "deorbit.ks":
		stage = Deorbit{args[0], args[1]}
	case "reentry.ks":
		stage = ReEntry{args[0], args[1]}
	default:
		err = fmt.Errorf("No template found with name: %s", templateName)
	}
	return stage, err
}

type LaunchWait struct {
	TgtVessel string // string
	TgtAngle  string // int
}
type Launch struct {
	TgtDir string // int 0 - 360
	TgtAlt string // int
}
type InitSpace struct {
	Antenna string // string
}
type Circularize struct {
	TgtDir string // int 0 - 360
}
type Intercept struct {
	ApproachRange string // int
}
type Approach struct {
	TgtVessel    string // string
	TgtRange     string // int
	OffsetVector string // Vector string V(0,0,0)
	Agility      string //float64
}
type WaitForCrew struct {
	TgtCrew string // int
}
type Deorbit struct {
	LKO          string // bool
	TgtPeriapsis string // int
}
type ReEntry struct {
	DragchuteAlt string // int
	ParachuteAlt string // int
}
