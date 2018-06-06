package plan

import (
	"fmt"
	"sort"
)

// Marshal returns the PLAN encoding of p.
func Marshal(p Plan) ([]byte, error) {
	b := make([]byte, 0)
	for _, part := range p.Parts {
		b = append(b, []byte(part.Name+"\n")...)

		// Write the data in sorted order
		var keys []string
		for k := range part.Data {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			b = append(b, []byte(fmt.Sprintf("\t%s = %s\n", k, part.Data[k]))...)
		}
	}
	return b, nil
}
