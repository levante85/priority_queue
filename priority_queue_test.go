package pqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func less(i, j int, values []interface{}) bool {
	l, _ := values[i].(int)
	r, _ := values[j].(int)
	return l < r
}

func TestInsert(t *testing.T) {
	pq := New(less)
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range num {
		pq.Insert(num[n])
	}

	assert.Equal(t, pq.Size(), len(num), "Inserting went wrong")
}

func TestInserttop(t *testing.T) {
	pq := New(less)
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range num {
		pq.Insert(num[n])
	}

	assert.Equal(t, pq.Size(), len(num), "Inserting went wrong")
	top := pq.Top().(int)
	assert.Equal(t, top, 9, "Inserting went wrong")

}

func TestPopTopEdge3Elements(t *testing.T) {
	pq := New(less)
	num := []int{1, 2, 3}
	for n := range num {
		pq.Insert(num[n])
	}

	_ = pq.PopTop()
}

func TestInsertPopTop(t *testing.T) {
	pq := New(less)
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range num {
		pq.Insert(num[n])
	}

	assert.Equal(t, pq.Size(), len(num), "Inserting went wrong")
	assert.Equal(t, pq.PopTop().(int), 9, "Pop not working")
	assert.Equal(t, pq.Size(), 8, "PopTop went wrong")

}

func TestInsertPopTopMulti(t *testing.T) {
	pq := New(less)
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range num {
		pq.Insert(num[n])
	}

	assert.Equal(t, pq.Size(), len(num), "Inserting went wrong")
	assert.Equal(t, pq.PopTop().(int), 9, "Pop not working")
	assert.Equal(t, pq.Size(), 8, "PopTop went wrong")
	assert.Equal(t, pq.PopTop().(int), 8, "Pop not working")
	assert.Equal(t, pq.Size(), 7, "PopTop went wrong")

}
