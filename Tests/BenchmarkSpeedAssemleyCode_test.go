package Tests

import (
	"fmt"
	"testing"
)

const wordCount = 1000

func WordCountDefault(slice []rune) int32 {
	var words int32 = 0
	inWord := false
	for i, char := range slice {
		space := IsSpace(char)
		if space && inWord {
			words++
		} else if !space && i == len(slice)-1 {
			words++
		}
		if space {
			inWord = false
		} else {
			inWord = true
		}

	}
	return words
}

func BenchmarkSpeedAssemblyCode(b *testing.B) {
	sizes := []int{100, 1000, 10000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Benchmark size %d", size), func(b *testing.B) {
			sentence := GenerateSentence(size)

			b.Run("AssembleyCode", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					WordCount(sentence)
				}

			})

			b.Run("DefaultRealisation", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					WordCountDefault(sentence)
				}
			})
		})
	}

}
