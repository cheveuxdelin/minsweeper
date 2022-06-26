package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/eiannone/keyboard"
	"github.com/gookit/color"
)

const (
	bomb         int          = -2
	undiscovered int          = -1
	blank        int          = 0
	arrowLeft    keyboard.Key = 65515
	arrowRight   keyboard.Key = 65514
	arrowUp      keyboard.Key = 65517
	arrowDown    keyboard.Key = 65516
	key_quit     rune         = 113
	key_bomb     rune         = 97
	key_safe     rune         = 115
)

var axis [9]int = [9]int{0, 1, 1, 0, -1, 1, -1, -1, 0}

var prints map[int]string = map[int]string{
	bomb:         "b",
	undiscovered: " ",
	0:            "0",
	1:            "1",
	2:            "2",
	3:            "3",
	4:            "4",
	5:            "5",
	6:            "6",
	7:            "7",
	8:            "8",
}

type MineSweeper struct {
	board      [][]int
	discovered [][]bool
	n          int
	m          int
	nBombs     int
	x          int
	y          int
}

func initialize(n int, m int, nBombs int) MineSweeper {
	var board [][]int = make([][]int, n)
	var discovered [][]bool = make([][]bool, n)

	for i := 0; i < n; i++ {
		board[i] = make([]int, m)
		discovered[i] = make([]bool, m)

		for j := 0; j < m; j++ {
			discovered[i][j] = false
		}
	}

	for i := 0; i < nBombs; i++ {
		var x int = rand.Intn(n)
		var y int = rand.Intn(m)
		for board[x][y] == bomb {
			x = rand.Intn(n)
			y = rand.Intn(m)
		}

		board[x][y] = bomb

		for k := 0; k < len(axis)-1; k++ {
			var xx = x + axis[k]
			var yy = y + axis[k+1]

			if xx >= 0 && yy >= 0 && xx < n && yy < m && board[xx][yy] != bomb {
				board[xx][yy] += 1
			}
		}

	}

	return MineSweeper{board, discovered, n, m, nBombs, 1, 1}
}

func (m *MineSweeper) print() {
	for i := 0; i < m.n; i++ {

		for j := 0; j < m.m; j++ {
			if (i == m.x || i == m.x+1) && j == m.y {
				color.Red.Print("+---")
			} else if (i == m.x || i == m.x+1) && j == m.y+1 {
				fmt.Print(color.Red.Sprint("+") + "---")
			} else {
				fmt.Print("+---")
			}
		}

		if (i == m.x || i == m.x+1) && m.y == m.m-1 {
			color.Red.Println("+")
		} else {
			fmt.Println("+")
		}

		for j := 0; j < m.m; j++ {
			var r string

			if m.discovered[i][j] {
				r = prints[m.board[i][j]]
			} else {
				r = prints[undiscovered]
			}

			if i == m.x && j == m.y {
				fmt.Print(color.Red.Sprint("|") + fmt.Sprintf(" %s ", r))
			} else if i == m.x && j == m.y+1 {
				fmt.Printf(color.Red.Sprint("|")+" %s ", r)
			} else {
				fmt.Printf("| %s ", r)
			}
		}

		if m.x == i && m.y == m.m-1 {
			color.Red.Println("|")
		} else {
			fmt.Println("|")
		}
	}

	for i := 0; i < m.m; i++ {
		if m.x == m.n-1 && i == m.y {
			color.Red.Print("+---")
		} else if m.x == m.n-1 && i == m.y+1 {
			fmt.Print(color.Red.Sprint("+") + "---")
		} else {
			fmt.Print("+---")
		}
	}

	if m.y == m.m-1 && m.x == m.n-1 {
		color.Red.Println("+")
	} else {
		fmt.Println("+")
	}

	fmt.Println()
}

func (m *MineSweeper) move(k keyboard.Key) {
	switch k {
	case arrowLeft:
		if m.y-1 >= 0 {
			m.y -= 1
		}
	case arrowRight:
		if m.y+1 < m.n {
			m.y += 1
		}
	case arrowUp:
		if m.x-1 >= 0 {
			m.x -= 1
		}
	case arrowDown:
		if m.x+1 < m.m {
			m.x += 1
		}
	}
}

func (m *MineSweeper) expand(x int, y int) {
	if x >= 0 && y >= 0 && x < m.n && y < m.m && !m.discovered[x][y] {
		m.discovered[x][y] = true
		if m.board[x][y] == blank {
			for i := 0; i < len(axis)-1; i++ {
				var xx int = x + axis[i]
				var yy int = y + axis[i+1]
				m.expand(xx, yy)
			}
		}
	}
}

func (m *MineSweeper) mark(x, y int) {

}

func (m *MineSweeper) play() {

	for {
		value, arrowKey, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		// clearing console
		fmt.Print("\033[H\033[2J")

		fmt.Println(value, arrowKey)

		if arrowKey != 0 {
			m.move(arrowKey)
		} else if value == key_bomb {
			fmt.Println("bomb")
		} else if value == key_safe {
			m.expand(m.x, m.y)
		} else if value == key_quit {
			os.Exit(0)
		}

		m.print()
	}
}

func main() {
	ms := initialize(10, 10, 50)
	ms.play()
}
