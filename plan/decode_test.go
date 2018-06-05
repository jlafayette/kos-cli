package plan

import (
	"reflect"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    Plan
		wantErr bool
	}{
		{
			"1 part",
			[]byte("part.ks\na=1\nb=2"),
			Plan{[]Part{
				Part{"part.ks", map[string]string{"a": "1", "b": "2"}},
			}},
			false,
		}, {
			"2 parts",
			[]byte("part.ks\na = 1\nb = 2\npart2.ks\na =    12"),
			Plan{[]Part{
				Part{"part.ks", map[string]string{"a": "1", "b": "2"}},
				Part{"part2.ks", map[string]string{"a": "12"}},
			}},
			false,
		}, {
			"crlf",
			[]byte("part.ks\r\na=1\r\nb=2\r\n"),
			Plan{[]Part{
				Part{"part.ks", map[string]string{"a": "1", "b": "2"}},
			}},
			false,
		}, {
			"no part",
			[]byte("\na=1\nb=2\n"),
			Plan{[]Part{
				Part{"", map[string]string{"a": "1", "b": "2"}},
			}},
			true,
		}, {
			"multiple assigns on one line",
			[]byte("\na=1b=3\nb=2\n"),
			Plan{},
			true,
		}, {
			"blank lines",
			[]byte("part.ks\n\n\n\n\n\n\r\n\r\na = 1\n\nb = 2\npart2.ks\na =    12\n\n\r\n"),
			Plan{[]Part{
				Part{"part.ks", map[string]string{"a": "1", "b": "2"}},
				Part{"part2.ks", map[string]string{"a": "12"}},
			}},
			false,
		}, {
			"repeat",
			[]byte("part.ks \n a = 1st \n a = 2nd"),
			Plan{[]Part{
				Part{"part.ks", map[string]string{"a": "2nd"}},
			}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p Plan
			err := Unmarshal(tt.data, &p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("Unmarshal() error = %v", err)
			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("Unmarshal() = %v, want %v", p, tt.want)
			} else {
				t.Logf("Unmarshal() = %v", p)
			}
		})
	}
}
