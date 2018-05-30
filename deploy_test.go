package main

import "testing"

func Test_deploy(t *testing.T) {
	type args struct {
		kspsrc    string
		kspscript string
		verbose   bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test", args{"G:\\kerboscripting", "G:\\kerboscripting\\deploytest", true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := deploy(tt.args.kspsrc, tt.args.kspscript, tt.args.verbose); (err != nil) != tt.wantErr {
				t.Errorf("deploy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
