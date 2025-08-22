package main

import (
	"fmt"
	"github.com/samber/lo"
	"math/rand"
)

func WordCount(slice []rune) int32
func SumSlice(slice []int32) int
func GenerateSlice(lenght int) []int32 {
	output := make([]int32, 0, lenght)
	for i := 0; i < lenght; i++ {
		output = append(output, int32(rand.Int()))
	}
	return output
}

func main() {
	input := GenerateSlice(100)
	fmt.Println(int(lo.Sum[int32](input)))
	fmt.Println(SumSlice(input))
}
