package build

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Read from a mission plan and return Mission struct
func Read(r io.Reader) (Mission, error) {
	m := Mission{make([]part, 0)}

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	p := part{"", make(map[string]string)}
	for scanner.Scan() {
		line := scanner.Text()
		ps := strings.Split(strings.TrimSpace(line), "=")
		fmt.Printf("line: %q, ps: %q\n", line, ps)

		// len(ps) will never be 0 because the separator is not an empty string
		switch len(ps) {
		case 1:
			if ps[0] == "" {
				continue
			}
			// This means it's a new part, add the current part and initialize a new one
			if p.name != "" {
				m.parts = append(m.parts, p)
			}
			p.name = strings.TrimSpace(ps[0])
			p.data = make(map[string]string)
		case 2:
			p.data[strings.TrimSpace(ps[0])] = strings.TrimSpace(ps[1])
		default:
			err := fmt.Errorf("invalid part-line: '%s' too many '=' separators", line)
			return m, err
		}
	}
	m.parts = append(m.parts, p)
	fmt.Printf("%q\n", m)
	if p.name == "" {
		err := fmt.Errorf("no part name found")
		return m, err
	}
	return m, nil
}
