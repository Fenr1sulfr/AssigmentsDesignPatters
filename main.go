package gotour

import (
	"fmt"
)

type Game struct {
	name    string
	players []Player
	isStart bool
}
type Gamer interface {
	accept(*Game)
}

type Player struct {
	name   string
	id     int
	choice bool
}
type LeadPlayer struct {
	name             string
	id               int
	thisChoiceMatter bool
}

func (p *Player) accept(g *Game) {
	if p.choice {
		g.players = append(g.players, *p)
	}
}
func (l *LeadPlayer) accept(g *Game) {
	if l.thisChoiceMatter {
		for i := 0; i < len(g.players); i++ {
			g.players[i].choice = true
		}
	}
}

func Start(g Gamer, gm *Game) {
	switch g.(type) {
	case *LeadPlayer:
		fmt.Println("Your party leader start the game")

	case *Player:
		for i := 0; i < len(gm.players); i++ {
			if !gm.players[i].choice {
				fmt.Println("Somedoby decline")
			}
		}
	default:
		fmt.Println("Somebody didn't accept the game ")
	}
}

func main() {
	p1 := Player{name: "1000-7?", id: 1, choice: false}
	p2 := Player{name: "all muted", id: 2, choice: true}
	lp := LeadPlayer{name: "mid or feed", id: 3, thisChoiceMatter: true}

	game := Game{
		name:    "Dota 2",
		players: []Player{},
		isStart: false,
	}
	p1.accept(&game)
	p2.accept(&game)
	lp.accept(&game)
	Start(&p1, &game)
}
