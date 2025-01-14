package models

type Colony struct {
	NumberofAnts int
	Rooms        struct{}
	Link         struct{}
	Start        string
	End          string
}

type Link struct {
	From string
	To   string
}

type Path struct {
	Rooms []string
}

type Room struct {
	Name    string
	X       int
	Y       int
	IsStart bool
	IsEnd   bool
}
