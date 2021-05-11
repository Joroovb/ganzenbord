package main

import (
	"bufio"
	"fmt"
	"ganzenbord/dobbelsteen"
	Speler "ganzenbord/speler"
	"os"
)

func main() {

	db := dobbelsteen.New()
	sp := Speler.New("Jaap")
	so := Speler.New("Jan")
	spelers := []*Speler.Speler{&sp, &so}
	klaar := false
	for !klaar {
	game:
		for i := range spelers {
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("%v is aan de beurt. q om te stoppen, d om te dobbelen\n", spelers[i].Naam())
			char, _, err := reader.ReadRune()

			if err != nil {
				fmt.Println(err)
			}

			if char == 'q' {
				os.Exit(0)
			}
			spelers[i].SetPositie(db.Rollen())
			fmt.Printf("%v staat op %v\n\n", spelers[i].Naam(), spelers[i].Positie())
			switch spelers[i].Positie() {
			case 63:
				fmt.Printf("%v heeft gewonnen!\n", spelers[i].Naam())
				klaar = true
				break game
			case 10, 20, 30, 40, 50, 60:
				fmt.Println("Dubbele stappen!")
				spelers[i].SetPositie(db.LaatsteWorp())
				fmt.Printf("Je staat nu op %v\n", spelers[i].Positie())
			case 23:
				fmt.Println("Je gaat naar de gevangenis!\nGame over!")
				klaar = true
				break game
			case 25, 45:
				fmt.Println("Terug naar start!")
				sp.ToStart()
			}

		}
	}
}
