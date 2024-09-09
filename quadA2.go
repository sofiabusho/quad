package piscine

import "fmt"

func Drawrec(x, y int, TLC, TRC, TS, BLC, BRC, S string) {
	if x > 0 && y > 0 {
		fmt.Print(TLC)
		if x != 1 {
			for i := 2; i < x; i++ {
				fmt.Print(TS)
			}

			fmt.Print(TRC)
		}
		fmt.Print('\n')
		// first line ends here
		if y != 1 {
			for j := 2; j < y; j++ {
				fmt.Print(S)
				if x != 1 {
					for i := 2; i < x; i++ {
						fmt.Print(' ')
					}

					fmt.Print(S)
				}
				fmt.Print('\n')
			}
			// middle line ends here
			fmt.Print(BLC)

			if x != 1 {
				for i := 2; i < x; i++ {
					fmt.Print(TS)
				}
				fmt.Print(BRC)

			}
			fmt.Print('\n')
		}

	}
}

func QuadA2(x, y int) {
	width := x
	height := y
	Drawrec(width, height, "o", "o", "-", "o", "o", "|")
}
