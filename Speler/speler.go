package Speler

import "fmt"

type Speler struct {
	positie int
	naam    string
}

func New(n string) Speler {
	s := Speler{positie: 0, naam: n}
	return s
}

func (s *Speler) Naam() string {
	return s.naam
}

func (s *Speler) Positie() int {
	return s.positie
}

func (s *Speler) SetPositie(p int) {
	newPos := s.positie + p
	if newPos > 63 {
		fmt.Println("Te ver! Zet de stappen terug")
		s.positie = 63 - (newPos - 63)
	} else {
		s.positie = newPos
	}
}

func (s *Speler) ToStart() {
	s.positie = 0
}
