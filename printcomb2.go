package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	firstrun := false
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '9'; b++ {
			for c := '0'; c <= '9'; c++ {
				for d := '0'; d <= '9'; d++ {
					if a > c || (b >= d && a == c) {
						continue
					}
					if firstrun {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					} else {
						firstrun = true
					}
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(' ')
					z01.PrintRune(c)
					z01.PrintRune(d)
				}
			}
		}
	}
	z01.PrintRune('\n')
}
