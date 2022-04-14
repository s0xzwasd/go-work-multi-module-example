package producer

import "fmt"

type Producer struct{}

func (p *Producer) Produce() {
	fmt.Println("Welcome to Go Workspaces!")
}
