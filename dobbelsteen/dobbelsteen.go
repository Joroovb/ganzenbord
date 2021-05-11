package dobbelsteen

import (
	"fmt"
	"math/rand"
	"time"
)

type dobbelsteen struct {
	laatsteWorp int
}

func New() dobbelsteen {
	db := dobbelsteen{laatsteWorp: 0}
	return db
}

func (db *dobbelsteen) LaatsteWorp() int {
	return db.laatsteWorp
}

func (db *dobbelsteen) Rollen() int {
	rand.Seed(time.Now().UnixNano())
	db.laatsteWorp = rand.Intn(6) + 1
	fmt.Printf("Je hebt %v gegooid\n", db.laatsteWorp)
	return db.laatsteWorp
}
