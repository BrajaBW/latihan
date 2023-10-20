package main

import "fmt"

type Player struct {
	Name  string
	Score int
}

func (p *Player) Addscore(Score int) {
	p.Score += Score

}

func (p Player) DisplayInfo() {

	fmt.Printf("nama pemain :%s Skor:%d", p.Name, p.Score)
}

func main() {
	player := Player{
		Name:  "John",
		Score: 0,
	}
	player.Addscore(5)
	player.Addscore(10)
	player.DisplayInfo()
}
