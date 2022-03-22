package main

import (
	"fmt"

	"github.com/helelou/Oppg3/event"
	"github.com/helelou/Oppg3/state"
)

func main() {
	fmt.Println(state.ViewState())
	fmt.Println(event.PutIn("rev"))
	fmt.Println(event.CrossRiver("rev"))
	fmt.Println(event.TakeOut("rev"))
}
