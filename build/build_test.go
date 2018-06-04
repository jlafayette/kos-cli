package build

import (
	"os"
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
