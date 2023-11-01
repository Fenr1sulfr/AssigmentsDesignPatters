package main

import "fmt"

type Command interface {
	change()
}
type basedPhrases interface {
	say1()
	say2()
	say3()
	say4()
}
type Life struct {
	command Command
}

func (l *Life) doSmth() {
	l.command.change()
}

type say1Command struct {
	basedPhrases basedPhrases
}

func (s *say1Command) change() {
	s.basedPhrases.say1()
}

type say2Command struct {
	basedPhrases basedPhrases
}

func (s *say2Command) change() {
	s.basedPhrases.say2()
}

type say3Command struct {
	basedPhrases basedPhrases
}

func (s *say3Command) change() {
	s.basedPhrases.say3()
}

type say4Command struct {
	basedPhrases basedPhrases
}

func (s *say4Command) change() {
	s.basedPhrases.say2()
}

type human struct {
	phrase1 string
	phrase2 string
	phrase3 string
	phrase4 string
}

func (h *human) say1() {
	fmt.Println("Я хороший программист")
}
func (h *human) say2() {
	fmt.Println("Я ужасный программист")
}
func (h *human) say3() {
	fmt.Println("Я не очень хороший программист")
}
func (h *human) say4() {
	fmt.Println("Я ухожу на завод")
}
