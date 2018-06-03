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
	tests := []struct {
		name   string
		kspsrc string
		stages []stage
	}{
		{
			"LKO_rescue",
			filepath.Join("..", "templates"),
			[]stage{
				stage{"launchWait.ks", LaunchWait{"Jesley's Capsule", "35"}},
				stage{"launch.ks", Launch{"90", "80000"}},
				stage{"initSpace.ks", InitSpace{"Communotron 16"}},
				stage{"circularize.ks", Circularize{"90"}},
				stage{"intercept.ks", Intercept{"8000"}},
				stage{"approach.ks", Approach{"Jesley's Capsule", "100", "V(0,0,0)", "1.2"}},
				stage{"waitForCrew.ks", WaitForCrew{"1"}},
				stage{"deorbit.ks", Deorbit{"true", "30000"}},
				stage{"reentry.ks", ReEntry{"2500", "1200"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MakeMission(os.Stdout, tt.kspsrc, tt.stages)
		})
	}
}
