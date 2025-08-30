package SimdTests

import (
	"fmt"
	"testing"
)

const (
	benchmarkLength = 10000
)

type BenchmarkData struct {
	slice    []int32
	targets  []int32
	notFound []int32
}

func generateBenchmarkData() *BenchmarkData {
	slice := make([]int32, benchmarkLength)
	for i := range slice {
		slice[i] = int32(i)
	}

	targets := []int32{
		slice[0],
		slice[benchmarkLength/2],
		slice[benchmarkLength-1],
		slice[benchmarkLength/4],
		slice[3*benchmarkLength/4],
	}

	notFound := []int32{
		-1, -100, -1000, 99999, 100000,
	}

	return &BenchmarkData{
		slice:    slice,
		targets:  targets,
		notFound: notFound,
	}
}

func BenchmarkSimdSearch(b *testing.B) {
	data := generateBenchmarkData()
	b.Run("Found_Elements", func(b *testing.B) {
		b.Run("Go_Default", func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				target := data.targets[i%len(data.targets)]
				FindElem(data.slice, target)
			}
		})

		b.Run("SIMD_XMM_128", func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				target := data.targets[i%len(data.targets)]
				FindElemSimdYMM(data.slice, target)
			}
		})

		b.Run("SIMD_YMM_256", func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				target := data.targets[i%len(data.targets)]
				FindElemSimdYMM(data.slice, target)
			}
		})
	})

	b.Run("NotFound_Elements", func(b *testing.B) {
		b.Run("Go_Default", func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				target := data.notFound[i%len(data.notFound)]
				FindElem(data.slice, target)
			}
		})

		b.Run("SIMD_XMM_128", func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				target := data.notFound[i%len(data.notFound)]
				FindElemSimd(data.slice, target)
			}
		})

		b.Run("SIMD_YMM_256", func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				target := data.notFound[i%len(data.notFound)]
				FindElemSimdYMM(data.slice, target)
			}
		})
	})
}

func BenchmarkSimdScalability(b *testing.B) {
	sizes := []int{100, 8000, 80000, 8000000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			slice := make([]int32, size)
			for i := range slice {
				slice[i] = int32(i)
			}

			target := slice[size-1] // средний элемент

			b.Run("Go_Default", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					FindElem(slice, target)
				}
			})

			b.Run("SIMD_XMM_128", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					FindElemSimdYMM(slice, target)
				}
			})

			b.Run("SIMD_YMM_256", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					FindElemSimdYMM(slice, target)
				}
			})
		})
	}
}
