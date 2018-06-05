package plan

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

// Unmarshal parses the PLAN-encoded data and stores the result in the Plan pointed to by p.
func Unmarshal(data []byte, p *Plan) error {
	r := bytes.NewReader(data)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	part := Part{"", make(map[string]string)}
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
			if part.name != "" {
				p.Parts = append(p.Parts, part)
			}
			part.name = strings.TrimSpace(ps[0])
			part.data = make(map[string]string)
		case 2:
			part.data[strings.TrimSpace(ps[0])] = strings.TrimSpace(ps[1])
		default:
			err := fmt.Errorf("invalid part-line: '%s' too many '=' separators", line)
			return err
		}
	}
	p.Parts = append(p.Parts, part)
	fmt.Printf("%v\n", p)
	if part.name == "" {
		err := fmt.Errorf("no part name found")
		return err
	}
	return nil
}
