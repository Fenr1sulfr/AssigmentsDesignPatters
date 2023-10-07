package main

import "fmt"

type Attack interface {
	attack() string
}

type Staff struct {
	name   string
	damage int
	Attack Attack
}
type Fire struct {
	name   string
	damage int
	Attack Attack
}
type Ice struct {
	name   string
	damage int
	Attack Attack
}

func (s *Staff) attack() string {
	battlelog := fmt.Sprintf("%a is attacking with damage %b", s.name, s.damage)
	return battlelog
}

func (i *Ice) attack() string {
	battlelog := fmt.Sprintf("%a is attacking with damage %b", i.name, i.damage)
	return i.Attack.attack() + battlelog
}

func (f *Fire) attack() string {
	battlelog := fmt.Sprintf("%a is attacking with damage %b", f.name, f.damage)
	return f.Attack.attack() + battlelog
}
func decorator() {
	combination := &Staff{name: "Staff of Great Wizrd", damage: 190}
	firstSpell := &Fire{
		name: "FireWall", damage: 120,
		Attack: combination,
	}
	secondSpell := &Ice{
		name: "FrostBite", damage: 130,
		Attack: firstSpell,
	}
	fmt.Println(secondSpell.attack())
}
