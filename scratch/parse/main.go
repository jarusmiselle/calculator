package main

import "fmt"

func main() {
	b()
}

func b() {
	daddy := "Seanny 38"

	for i, r := range daddy {
		fmt.Printf("index: %d, rune: %c, ASCII: %d\n", i, r, r)
	}
}

func a() {
	ex := "stuff"
	outer(inner(len(ex), 20))
}

func inner(a int, b int) int {
	fmt.Printf("running inner(%d)\n", a)
	return a + b
}

func outer(b int) {
	fmt.Printf("running outer(%d)\n", b)
}
