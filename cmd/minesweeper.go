package main

import (
	pkg "Jacob953-minesweeper/pkg"

	"fmt"
	"os"
)

const (
	FLAG   = true
	UNFALG = false
)

var (
	cmd  string
	x, y int
)

func main() {
	var rows, cols int
	fmt.Print("Enter [rows cols]: ")
	_, err := fmt.Fscanf(os.Stdin, "%d %d", &rows, &cols)
	if err != nil {
		fmt.Print("Wrong type with rows or cols:", err)
		return
	}
	// init a minesweeper game
	game, flagleft := pkg.NewGame(rows, cols, 0.3)
	for {
		fmt.Printf("Go [x y [f -flag|u -unflag|o -open]]: ")
		_, err := fmt.Fscanf(os.Stdin, "%d %d %s", &x, &y, &cmd)
		if err != nil {
			fmt.Println("Wrong action:", err)
			game.ShowGrid()
			continue
		}
		switch cmd {
		case "f":
			flagleft = game.Flag(x, y, FLAG)
		case "u":
			flagleft = game.Flag(x, y, UNFALG)
		case "o":
			// open the mask
			left, ok := game.Open(x, y)
			if !ok {
				game.ShowGrid()
				fmt.Println("GAME OVER!")
				return
			}
			if left == 0 {
				game.ShowGrid()
				fmt.Println("WIN!")
				return
			}
		default:
			fmt.Println("Wrong command: ", cmd, ". Please try again!")
		}
		fmt.Println("Left flags:", flagleft)
		game.ShowGrid()
	}
}
