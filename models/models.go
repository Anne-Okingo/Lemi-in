package models

type Graph struct {
	AntCount  int
	AntNames  []string
	Rooms     map[string]*ARoom
	AllPaths  [][]string
	StartRoom string
	EndRoom   string
}

type ARoom struct {
	Name        string
	XCoordinate int
	YCoordinate int
	Links       []string
	Visited     bool
}
