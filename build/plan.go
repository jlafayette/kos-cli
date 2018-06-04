package build

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Read from a mission plan and return Mission struct
func Read(r io.Reader) (Mission, error) {
	m := Mission{make([]part, 0)}

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	p := part{"", make(map[string]string)}
	for scanner.Scan() {
		line := scanner.Text()
		ps := strings.Split(strings.TrimSpace(line), "=")
		fmt.Printf("line: %q, ps: %q\n", line, ps)

		// len(ps) will never be 0 because the separator is not an empty string
		switch len(ps) {
		case 1:
			if ps[0] == "" {
				continue
			}
			// This means it's a new part, add the current part and initialize a new one
			if p.name != "" {
				m.parts = append(m.parts, p)
			}
			p.name = strings.TrimSpace(ps[0])
			p.data = make(map[string]string)
		case 2:
			p.data[strings.TrimSpace(ps[0])] = strings.TrimSpace(ps[1])
		default:
			err := fmt.Errorf("invalid part-line: '%s' too many '=' separators", line)
			return m, err
		}
	}
	m.parts = append(m.parts, p)
	fmt.Printf("%q\n", m)
	if p.name == "" {
		err := fmt.Errorf("no part name found")
		return m, err
	}
	return m, nil
}

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
