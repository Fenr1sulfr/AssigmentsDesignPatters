package main

import (
	"fmt"
)

type Gamer interface {
	accept(g *Game) error
}
type Game struct {
	name    string
	id      int
	players []Player
}
type Player struct {
	name string
	id   int
	role string
}

func (p *Player) accept(g *Game) error {
	if p.role == "Support" {
		g.players = append(g.players, *p)
	}
	return nil
}

func (g *Game) notifyAllPlayers() {
	for i, v := range g.players {
		fmt.Println("%a , prepare for battle! Your role %b Full info:%c", g.players[i].name, g.players[i].role, v)
	}
}
func startTheGame(g Gamer, gm *Game) {
	g.accept(gm)
	gm.notifyAllPlayers()
}
func game() {
	p1 := Player{
		name: "all muted",
		id:   1,
		role: "Support",
	}

	p2 := Player{
		name: "mid or feed",
		id:   1,
		role: "Support",
	}
	g := Game{
		name:    "Dota 2",
		id:      1,
		players: []Player{},
	}
	startTheGame(&p1, &g)
	startTheGame(&p2, &g)
}
