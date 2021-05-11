package bord

import "fmt"

type vakje struct {
}

type dubbelvakje struct {
	vak vakje
}

type bord struct {
	Vakjes [63]*vakje
}

func New() *bord {
	b := new(bord)
	for i, _ := range b.Vakjes {
		switch i {
		case 10, 20, 30, 40, 50, 60:
			fmt.Println(i)
		default:
			b.Vakjes[i] = new(vakje)

		}
	}
	return b
}
