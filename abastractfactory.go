package assigmentgolangobservermainstrategygotour

import "fmt"

type IHealFactory interface {
	makeHeal() IHeal
}
type IHeal interface {
	pushMana() string
	pushHeal() string
}
type Heal struct {
	impact int
}

func (h *Heal) pushHeal() string {
	return fmt.Sprintf("healing you, +%B hp", h.impact)
}
func (h *Heal) pushMana() string {
	return fmt.Sprintf("pushing your mana, +%B mana", h.impact)
}

type Druid struct {
	additionalImpact int
	nickname         string
	Heal
}
type Priest struct {
	Heal
	additionalImpact int
	nickname         string
}
type DruidFactory struct {
}

func (d *DruidFactory) makeHeal() IHeal {
	return &Druid{
		nickname:         "Nature's Prophet",
		additionalImpact: 200,
		Heal: Heal{
			impact: 400,
		},
	}
}

type PriestFactory struct {
}

func (f *PriestFactory) makeHeal() IHeal {
	return &Priest{
		nickname:         "Nature's Prophet",
		additionalImpact: 300,
		Heal: Heal{
			impact: 700,
		},
	}
}
