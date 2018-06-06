package plan

// Plan for a mission, consisting of a series of parts.
type Plan struct {
	Parts []Part
}

// Part of a mission, like launch or circularize.
type Part struct {
	name string
	data map[string]string
}

// Boot contains information need to customize boot template .ks files
type Boot struct {
	Filename string
}
