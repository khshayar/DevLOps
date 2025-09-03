package helper

import "math/rand"


func Rand(n int) int {
	value:= rand.Intn(n)
	return value

}