package main

import (
	"fmt"
	"strconv"
)

type V2 struct {
	x int
	y int
}

func CreateV2(a, b int) V2 {
	vec := V2{a, b}
	return vec
}

func (vec V2) ToString() (str string) {
	str = "[" + strconv.Itoa(vec.x) + ", " + strconv.Itoa(vec.y) + "]"
	return
}

func (vec V2) GetX() int {
	return vec.x
}

func (vec V2) GetY() int {
	return vec.y
}

func AddV2(vec1 V2, vec2 V2) V2 {
	var vec V2
	vec.x = vec1.x + vec2.x
	vec.y = vec1.y + vec2.y

	return vec
}

func MultiplyV2(vec1 V2, scalar int) V2 {
	var vec V2

	vec.x = vec1.x * scalar
	vec.y = vec1.y * scalar

	return vec
}

func main() {
	vec := CreateV2(1, 2)
	vec2 := CreateV2(2, -2)
	fmt.Println(vec)

	fmt.Println(vec.ToString())
	fmt.Println(vec.GetX())
	fmt.Println(vec.GetY())
	fmt.Println(AddV2(vec, vec2))
	fmt.Println(MultiplyV2(vec, 5))

}
