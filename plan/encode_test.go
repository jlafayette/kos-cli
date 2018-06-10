package plan

import (
	"reflect"
	"testing"
)

func TestMarshal(t *testing.T) {
	tests := []struct {
		name    string
		plan    Plan
		want    []byte
		wantErr bool
	}{
		{
			"1 part",
			Plan{[]Part{
				Part{"part.ks", map[string]string{"a": "1", "b": "2"}},
			}},
			[]byte("part.ks\n\ta = 1\n\tb = 2\n"),
			false,
		}, {
			"2 parts",
			Plan{[]Part{
				Part{"part.ks", map[string]string{"a": "1", "b": "2"}},
				Part{"part2.ks", map[string]string{"a": "12"}},
			}},
			[]byte("part.ks\n\ta = 1\n\tb = 2\npart2.ks\n\ta = 12\n"),
			false,
		}, {
			"no part",
			Plan{[]Part{
				Part{"", map[string]string{"a": "1", "b": "2"}},
			}},
			[]byte("\n\ta = 1\n\tb = 2\n"),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.plan)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("Marshal() error = %v", err)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\n got: %v\nwant: %v", got, tt.want)
			} else {
				t.Logf("got:\n%v", got)
			}
		})
	}
}
