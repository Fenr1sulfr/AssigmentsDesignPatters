package gotour

import (
	"fmt"
)

type Mailer interface {
	mailer(message string) error
}
type Gmail struct {
	price   int
	name    string
	profile string
}
type Yandex struct {
	price  int
	adress string
	key    string
}

func (g *Gmail) mailer(message string) error {
	fmt.Println("Sending b% to Gmail %a ...", message, g.profile)
	return nil
}
func (y *Yandex) mailer(message string) error {
	fmt.Printf("Sending %b to Yandex by %a key", message, y.key)
	return nil
}

type Sending struct {
	Mailer Mailer
}

func (s *Sending) sendMessage(message string) {
	s.Mailer.mailer(message)
}

func strategy() {
	y1 := Yandex{
		price:  10,
		adress: "zombie123@yandex.ru",
		key:    "lal24l12l41s012",
	}
	g1 := Gmail{
		price:   20,
		name:    "Temirlan T.T",
		profile: "zxcghoul123@gmail.com",
	}

	y1.mailer("123")
	g1.mailer("231")
}
