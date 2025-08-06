package main

import (
	"fmt"

	"github.com/Kuguchev/go-first-fl-codestyle/internal/game"
)

func main() {
	err := game.NewGame().Start()
	if err != nil {
		fmt.Println(err)
	}
}
