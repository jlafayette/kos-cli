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
	}{
		{"Test1", filepath.Join("g:\\kerboscripting", "missions", "templates")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MakeMission(os.Stdout, tt.kspsrc)
		})
	}
}
