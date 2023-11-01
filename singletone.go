package main

import "fmt"

type progressBar struct {
	hours int
}

func (p *progressBar) getProgress() {
	fmt.Println(p.hours)
}

var instanceOfProgress *progressBar

func getInstance() *progressBar {
	if instanceOfProgress == nil {
		fmt.Println("Your progress starts here!")
		instanceOfProgress = &progressBar{}
	} else {
		fmt.Println("You already have started your story")
		instanceOfProgress.getProgress()
	}
	return instanceOfProgress
}
func main() {

	progress := getInstance()
	progress.hours = 5
	progress.getProgress()
	progress1 := getInstance()
	progress1.getProgress()

}
