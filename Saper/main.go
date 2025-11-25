package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Cell struct {
	mine     bool
	revealed bool
	count    int
}

type Board struct {
	width  int
	height int
	cells  [][]Cell
}

func NewBoard(w, h, mines int) *Board {
	b := &Board{
		width:  w,
		height: h,
		cells:  make([][]Cell, h),
	}

	for i := range b.cells {
		b.cells[i] = make([]Cell, w)
	}

	b.placeMines(mines)
	b.calculateCounts()
	return b
}

func (b *Board) placeMines(mines int) {
	rand.Seed(time.Now().UnixNano())
	placed := 0
	for placed < mines {
		x := rand.Intn(b.width)
		y := rand.Intn(b.height)

		if !b.cells[y][x].mine {
			b.cells[y][x].mine = true
			placed++
		}
	}
}

func (b *Board) calculateCounts() {
	dirs := [8][2]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			if b.cells[y][x].mine {
				continue
			}

			count := 0
			for _, d := range dirs {
				nx := x + d[0]
				ny := y + d[1]

				if nx >= 0 && nx < b.width && ny >= 0 && ny < b.height {
					if b.cells[ny][nx].mine {
						count++
					}
				}
			}
			b.cells[y][x].count = count
		}
	}
}

func (b *Board) Reveal(x, y int) bool {
	if x < 0 || x >= b.width || y < 0 || y >= b.height {
		return true
	}

	c := &b.cells[y][x]

	if c.revealed {
		return true
	}

	c.revealed = true

	if c.mine {
		return false
	}

	// авто-раскрытие пустых областей
	if c.count == 0 {
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				if dx != 0 || dy != 0 {
					b.Reveal(x+dx, y+dy)
				}
			}
		}
	}

	return true
}

func (b *Board) Print(showMines bool) {
	fmt.Print("   ")
	for x := 0; x < b.width; x++ {
		fmt.Printf("%d ", x)
	}
	fmt.Println()

	for y := 0; y < b.height; y++ {
		fmt.Printf("%2d ", y)
		for x := 0; x < b.width; x++ {
			c := b.cells[y][x]
			if !c.revealed {
				if showMines && c.mine {
					fmt.Print("* ")
				} else {
					fmt.Print("# ")
				}
				continue
			}

			if c.mine {
				fmt.Print("* ")
			} else if c.count == 0 {
				fmt.Print(". ")
			} else {
				fmt.Printf("%d ", c.count)
			}
		}
		fmt.Println()
	}
}

func main() {
	var w, h, mines int
	w, h, mines = 10, 10, 10

	board := NewBoard(w, h, mines)

	for {
		board.Print(false)

		var x, y int
		fmt.Print("Введите x y: ")
		_, err := fmt.Scan(&x, &y)
		if err != nil {
			fmt.Println("Ошибка ввода.")
			continue
		}

		if !board.Reveal(x, y) {
			fmt.Println("💥 Вы взорвались!")
			board.Print(true)
			return
		}

		// Победа
		if checkWin(board) {
			fmt.Println("🎉 Победа!")
			board.Print(true)
			return
		}
	}
}

func checkWin(b *Board) bool {
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			c := b.cells[y][x]
			if !c.mine && !c.revealed {
				return false
			}
		}
	}
	return true
}
