package build

import (
	"os"
	"testing"
)

func Test_test(t *testing.T) {
	tests := []struct {
		name    string
		kspsrc  string
		mission *Mission
	}{
		{"No Args", "g:\\kerboscripting", &Mission{"LKO_rescue.ks", nil}},
		{"Args", "g:\\kerboscripting", &Mission{"LKO_rescue.ks", []string{"80000", "true"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test(os.Stdout, tt.kspsrc, tt.mission)
		})
	}
}
