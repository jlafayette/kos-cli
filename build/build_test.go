package build

import (
	"bytes"
	"testing"
)

func TestBoot(t *testing.T) {
	type args struct {
		boot string
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		{
			"Test",
			args{"testboot.ks", "a"},
			"Filename = a.ks",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := Boot(w, tt.args.boot, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Boot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotW := w.String()
			if gotW != tt.wantW {
				t.Errorf("Boot() = %v, want %v", gotW, tt.wantW)
			}
			t.Logf("Boot(): %v", gotW)
		})
	}
}
