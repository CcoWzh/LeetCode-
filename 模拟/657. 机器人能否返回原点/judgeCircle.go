package main

import "fmt"

func judgeCircle(moves string) bool {
	//不需要真的去模拟，只需要保证最后能回到(0,0)即可
	x, y := 0, 0
	for i := 0; i < len(moves); i++ {
		if moves[i] == 'R' {
			x++
		} else if moves[i] == 'L' {
			x--
		} else if moves[i] == 'U' {
			y++
		} else if moves[i] == 'D' {
			y--
		}
	}

	return x == 0 && y == 0
}

func main() {
	moves := "UD"
	fmt.Println(judgeCircle(moves))
}
