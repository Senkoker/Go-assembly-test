package SimdTests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

const (
	attempts = 100
	length   = 1000
)

type SimdTest struct {
	slice  []int32
	target int32
	result bool
}

func GenerateSlice(lenght, failedTestCases int) ([]int32, []int32) {
	existNumber := make(map[int32]struct{}, lenght)
	output := make([]int32, 0, lenght)
	notInSliceNumbers := make([]int32, 0, failedTestCases)
	for i := 0; i < lenght; i++ {
		value := int32(rand.Int())
		output = append(output, int32(rand.Int()))
		existNumber[value] = struct{}{}
	}
	j := 0
	for j < failedTestCases {
		notEqualValue := int32(rand.Int())
		if _, ok := existNumber[notEqualValue]; !ok {
			notInSliceNumbers = append(notInSliceNumbers, notEqualValue)
			j++
		}
	}
	return output, notInSliceNumbers
}

func FindElem(slice []int32, target int32) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return true
		}
	}
	return false
}

func FindElemSimd(slice []int32, target int32) bool

func FindElemSimdYMM(slice []int32, target int32) bool

func TestSimd(t *testing.T) {
	happyTestCases := make([]SimdTest, 0, attempts)
	failedTestCases := make([]SimdTest, 0, attempts)
	slice, valuesNotInSlice := GenerateSlice(length, attempts)
	for i := 0; i < attempts; i++ {
		happyTestCases = append(happyTestCases, SimdTest{
			slice:  slice,
			target: slice[rand.Int63n(length-1)],
			result: true,
		})
		failedTestCases = append(failedTestCases, SimdTest{
			slice:  slice,
			target: valuesNotInSlice[rand.Int63n(attempts-1)],
			result: false,
		})
	}
	t.Run("SIMD XMM-128bits register", func(t *testing.T) {
		for _, testCase := range happyTestCases {
			assert.Equal(t, FindElemSimd(testCase.slice, testCase.target), testCase.result,
				fmt.Sprintf("testcase %v", testCase.target), "happyTest")
		}
		for _, testCase := range failedTestCases {
			assert.Equal(t, FindElemSimd(testCase.slice, testCase.target), testCase.result,
				fmt.Sprintf("testcase %v", testCase.target), "failedTest")
		}
	})
	t.Run("SIMD YMM-256bits register", func(t *testing.T) {
		for _, testCase := range happyTestCases {
			require.Equal(t, FindElemSimdYMM(testCase.slice, testCase.target), testCase.result, fmt.Sprintf("testcase %v", testCase))
		}
		for _, testCase := range failedTestCases {
			require.Equal(t, FindElemSimdYMM(testCase.slice, testCase.target), testCase.result, fmt.Sprintf("testcase %v", testCase))
		}
	})
}
