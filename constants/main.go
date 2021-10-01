package main

import (
	"fmt"
	"math"
)

const s string = "constant"
const t = 33

func main() {
	fmt.Println(s)
	fmt.Println(t)

	const n = 50000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
