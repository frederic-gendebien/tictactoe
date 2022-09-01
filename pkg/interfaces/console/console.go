package console

import (
	"bufio"
	"fmt"
	"os"
	"tictactoe/pkg/domain"
	"unicode"
)

func NewConsole() Console {
	return Console{}
}

type Console struct {
}

func (c Console) ShowGrid(grid *domain.Grid) {
	printColumns := func() {
		fmt.Println("    a   b   c")
	}

	fmt.Println()
	printColumns()
	for row, line := range grid.Matrix {
		fmt.Print(row + 1)
		for _, symbole := range line {
			fmt.Print(" | ", string(symbole))
		}
		fmt.Print(" | ")
		fmt.Println(row + 1)
	}
	printColumns()
}

func (c Console) AskCoordinatesTo(player domain.Player) (domain.Coordinates, error) {
	normalizeColumn := func(col rune, _ int, err error) (int, error) {
		if err != nil {
			return -1, err
		}

		switch unicode.ToLower(col) {
		case 'a':
			return 0, nil
		case 'b':
			return 1, nil
		case 'c':
			return 2, nil
		default:
			return -1, fmt.Errorf("wrong coordinates, dumbass")
		}
	}

	normalizeRow := func(row rune, _ int, err error) (int, error) {
		if err != nil {
			return -1, err
		}

		switch unicode.ToLower(row) {
		case '1':
			return 0, nil
		case '2':
			return 1, nil
		case '3':
			return 2, nil
		default:
			return -1, fmt.Errorf("wrong coordinates, dumbass")
		}
	}

	fmt.Println()
	fmt.Printf("player '%s' enter coordinates, e.g b2 : ", string(player.Symbol))
	reader := bufio.NewReader(os.Stdin)
	col, err := normalizeColumn(reader.ReadRune())
	if err != nil {
		return domain.Coordinates{}, err
	}
	row, err := normalizeRow(reader.ReadRune())
	if err != nil {
		return domain.Coordinates{}, err
	}

	return domain.Coordinates{
		Row: row,
		Col: col,
	}, nil
}

func (c Console) Congratulate(player domain.Player) {
	fmt.Println()
	fmt.Printf("congratulations player '%s' you win!", string(player.Symbol))
	fmt.Println()
}

func (c Console) Slap(player domain.Player, err error) {
	fmt.Println()
	fmt.Printf("player '%s' %s", string(player.Symbol), err.Error())
	fmt.Println()
}

func (c Console) NoWinner() {
	fmt.Println()
	fmt.Println("nobody wins, losers!")
	fmt.Println()
}
