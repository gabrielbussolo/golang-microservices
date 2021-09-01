package utils

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestBubbleSortWorstCase(t *testing.T) {
	//inicialization
	els := []int{9, 8, 7, 6, 5}

	//execution
	sort := BubbleSort(els)

	//Validation
	assert.NotNil(t, sort)
	assert.EqualValues(t, 5, len(sort))
	assert.EqualValues(t, sort, []int{5, 6, 7, 8, 9})
}

func TestBubbleSortBestCase(t *testing.T) {
	//inicialization
	els := []int{5, 6, 7, 8, 9}

	//execution
	sort := BubbleSort(els)

	//Validation
	assert.NotNil(t, sort)
	assert.EqualValues(t, 5, len(sort))
	assert.EqualValues(t, sort, []int{5, 6, 7, 8, 9})
}

func BenchmarkBubbleSort10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(getElements(10))
	}
}

func BenchmarkSort10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Ints(getElements(10))
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(getElements(1000))
	}
}

func BenchmarkSort1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Ints(getElements(1000))
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(getElements(100000))
	}
}

func BenchmarkSort100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Ints(getElements(100000))
	}
}

func getElements(size int) []int {
	result := make([]int, size)
	i := 0
	for j := size - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElements(t *testing.T) {
	elements := getElements(5)

	assert.EqualValues(t, elements, []int{4, 3, 2, 1, 0})
	assert.EqualValues(t, 5, len(elements))
}
