package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	rotate(s[:], 3)
	fmt.Println(s) // "[3 4 5 0 1 2]"
}

func rotate(s []int, pos int) {
	if pos < 0 {
		rotateLeft(s, -pos)
	} else {
		rotateRight(s, pos)
	}
}

func rotateLeft(s []int, pos int) {
	maxIndex := len(s) - 1
	step := pos % len(s)
	if needOtherRotate(step, s) {
		newPos := len(s) % step
		rotateRight(s, newPos)
		return
	}
	max := step
	cnt := 0
	for i := 0; i < max; i++ {
		temp := s[i]
		j := i
		for {
			ind := j + max

			if ind > maxIndex {
				ind -= maxIndex + 1
			}

			if ind == i {
				break
			}

			s[j] = s[ind]
			j = ind
			cnt++
		}
		s[j] = temp
		cnt++
		if cnt == len(s) {
			break
		}
	}
}

func needOtherRotate(step int, s []int) bool {
	return step > int(float64(len(s))/2.0+0.5)
}

func rotateRight(s []int, pos int) {
	maxIndex := len(s) - 1
	step := pos % len(s)
	if needOtherRotate(step, s) {
		newPos := len(s) % step
		rotateLeft(s, newPos)
		return
	}
	max := step

	cnt := 0
	for i := 0; i < max; i++ {
		temp := s[maxIndex-i]
		j := maxIndex - i
		for {
			ind := j - max

			if ind < 0 {
				ind += maxIndex + 1
			}

			if ind == maxIndex-i {
				break
			}

			s[j] = s[ind]
			j = ind
			cnt++
		}
		s[j] = temp
		cnt++
		if cnt == len(s) {
			break
		}
	}
}
