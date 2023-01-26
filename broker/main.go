package main

import (
	"fmt"
	"strings"
)

type thingsWithNames interface {
	is() string
	name() string
}

type stuffInMyHouse struct {
	kind   string
	called string
}

func (s *stuffInMyHouse) name() string {
	return s.called
}

func (s *stuffInMyHouse) is() string {
	return s.kind
}

func yellName(thing thingsWithNames) (string, error) {
	if len(thing.name()) > 8 {
		return "", fmt.Errorf("That name's too long for me to yell: %s", thing.name())
	}
	up := strings.ToUpper(thing.name())
	fmt.Println("thing's name:", up)
	return up, nil
}

func main() {
	things := []stuffInMyHouse{{called: "Osório"}, {kind: "plant", called: "Henrietta"}, {kind: "miniature", called: "beholder"}}
	for _, thing := range things {
		if _, err := yellName(&thing); err != nil {
			fmt.Println("Got an error during execution but I'm still going, errors are just values anyways")
		}
	}
}
