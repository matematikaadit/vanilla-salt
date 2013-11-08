package main

import "fmt"

func main() {
	var n, m, d int64
	fmt.Scanf("%d %d %d", &n, &m, &d)
	fmt.Println(ceildiv(n, d) * ceildiv(m, d))
}

func ceildiv(n, d int64) int64 {
	switch n % d {
	case 0:
		return n / d
	default:
		return (n / d) + 1
	}
}
