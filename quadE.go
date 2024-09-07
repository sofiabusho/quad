package piscine

import "github.com/01-edu/z01"

// corner topleft always there
// corner topright if x!=1
// corner bottomleft if y!=1
// corner bottomright if x!=1 && y!=1
func QuadE(x, y int) {
	if x > 0 && y > 0 {
		z01.PrintRune('A')
		for i := 2; i < x; i++ {
			z01.PrintRune('B')
		}
		if x != 1 {
			z01.PrintRune('C')
		}
		z01.PrintRune('\n')
		// first line ends here
		if y != 1 {
			for j := 2; j < y; j++ {
				z01.PrintRune('B')
				for i := 2; i < x; i++ {
					z01.PrintRune(' ')
				}
				if x != 1 {
					z01.PrintRune('B')
				}
				z01.PrintRune('\n')
			}
		}
		// middle line ends here
		if y != 1 {
			z01.PrintRune('C')
		}
		if x != 1 && y != 1 {
			for i := 2; i < x; i++ {
				z01.PrintRune('B')
			}

			z01.PrintRune('A')
			// z01.PrintRune('\n')
		}
		if y != 1 {
			z01.PrintRune('\n')
		}
	}
}
