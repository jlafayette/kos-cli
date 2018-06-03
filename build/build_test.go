package build

import (
	"os"
	"path/filepath"
	"testing"
)

func Test_test(t *testing.T) {
	tests := []struct {
		name    string
		kspsrc  string
		mission *Boot
	}{
		{"No Args", "g:\\kerboscripting", &Boot{"LKO_rescue.ks", nil}},
		{"Args", "g:\\kerboscripting", &Boot{"LKO_rescue.ks", []string{"80000", "true"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Test(os.Stdout, tt.kspsrc, tt.mission)
		})
	}
}

func TestMakeMission(t *testing.T) {
	i, _ := Structify("launchWait.ks", []string{"Jesley's Capsule", "35"})
	lw := stage{"launchWait.ks", i}
	i, _ = Structify("launch.ks", []string{"90", "80000"})
	l := stage{"launch.ks", i}
	i, _ = Structify("initSpace.ks", []string{"Communotron 16"})
	is := stage{"initSpace.ks", i}
	i, _ = Structify("circularize.ks", []string{"90"})
	c := stage{"circularize.ks", i}
	i, _ = Structify("intercept.ks", []string{"8000"})
	in := stage{"intercept.ks", i}
	i, _ = Structify("approach.ks", []string{"Jesley's Capsule", "100", "V(0,0,0)", "1.2"})
	a := stage{"approach.ks", i}
	i, _ = Structify("waitForCrew.ks", []string{"1"})
	w := stage{"waitForCrew.ks", i}
	i, _ = Structify("deorbit.ks", []string{"true", "30000"})
	d := stage{"deorbit.ks", i}
	i, _ = Structify("reentry.ks", []string{"2500", "1200"})
	r := stage{"reentry.ks", i}

	tests := []struct {
		name   string
		kspsrc string
		stages []stage
	}{
		{
			"LKO_rescue",
			filepath.Join("..", "templates"),
			[]stage{lw, l, is, c, in, a, w, d, r},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MakeMission(os.Stdout, tt.kspsrc, tt.stages)
		})
	}
}
