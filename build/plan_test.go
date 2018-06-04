package build

import (
	"reflect"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	tests := []struct {
		name    string
		plan    string
		want    Mission
		wantErr bool
	}{
		{
			"1 part",
			"part.ks\na=1\nb=2",
			Mission{[]part{part{"part.ks", map[string]string{"a": "1", "b": "2"}}}},
			false,
		}, {
			"2 parts",
			"part.ks\na = 1\nb = 2\npart2.ks\na =    12",
			Mission{[]part{
				part{"part.ks", map[string]string{"a": "1", "b": "2"}},
				part{"part2.ks", map[string]string{"a": "12"}},
			}},
			false,
		}, {
			"crlf",
			"part.ks\r\na=1\r\nb=2\r\n",
			Mission{[]part{
				part{"part.ks", map[string]string{"a": "1", "b": "2"}},
			}},
			false,
		}, {
			"no part",
			"\na=1\nb=2\n",
			Mission{[]part{
				part{"", map[string]string{"a": "1", "b": "2"}},
			}},
			true,
		}, {
			"multiple assigns on one line",
			"\na=1b=3\nb=2\n",
			Mission{[]part{}},
			true,
		}, {
			"blank lines",
			"part.ks\n\n\n\n\n\n\r\n\r\na = 1\n\nb = 2\npart2.ks\na =    12\n\n\r\n",
			Mission{[]part{
				part{"part.ks", map[string]string{"a": "1", "b": "2"}},
				part{"part2.ks", map[string]string{"a": "12"}},
			}},
			false,
		}, {
			"repeat",
			"part.ks \n a = 1st \n a = 2nd",
			Mission{[]part{
				part{"part.ks", map[string]string{"a": "2nd"}},
			}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.plan)
			got, err := Read(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("Read() error = %v", err)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() = %v, want %v", got, tt.want)
			} else {
				t.Logf("Read() = %v", got)
			}
		})
	}
}
