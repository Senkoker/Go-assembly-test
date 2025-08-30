package Tests

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"unicode"
)

const (
	attempt    = 100
	wordLength = 10
)

var (
	dictionary = "qwertyuiopasdfghjklzxcvbnm"
)

func SumSlice(slice []int32) int32
func WordCount(slice []rune) int32
func IsSpace(char rune) bool {
	return unicode.IsSpace(char)
}

func GenerateSlice(lenght int) []int32 {
	output := make([]int32, 0, lenght)
	for i := 0; i < lenght; i++ {
		output = append(output, int32(rand.Int()))
	}
	return output
}

func GenerateSentence(wordCount int) []rune {
	buff := make([]byte, 0)
	for i := 0; i < wordCount; i++ {
		for k := 0; k < wordLength; k++ {
			buff = append(buff, dictionary[rand.Intn(len(dictionary))])
		}
		buff = append(buff, ' ')
	}
	return []rune(string(buff[:len(buff)-1]))
}

func TestSumSlice(t *testing.T) {
	for i := 0; i < attempt; i++ {
		length := 100
		input := GenerateSlice(length)
		actual := lo.Sum[int32](input)
		assert.Equal(t, SumSlice(input), actual, fmt.Sprintf("attempt %d", i))
	}
}

func TestCountWord(t *testing.T) {
	for i := 0; i < attempt; i++ {
		wordNumber := int32(rand.Intn(100))
		assert.Equal(t, WordCount(GenerateSentence(int(wordNumber))), wordNumber)
	}
}
