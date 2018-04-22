package pqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func debug(s []interface{}, t *testing.T) {
	t.Logf("Content %+v\n", s)
}

func compare(i, j int, values []interface{}) bool {
	left := values[i].(int)
	right := values[i].(int)
	return left < right
}

func TestInsert(t *testing.T) {
	pq := New(compare)
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range num {
		pq.Insert(num[n])
	}

	assert.Equal(t, pq.Count(), len(num), "Inserting went wrong")
	debug(pq.slice, t)
}

func TestInsertMax(t *testing.T) {
	pq := New(compare)
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range num {
		pq.Insert(num[n])
	}

	assert.Equal(t, pq.Count(), len(num), "Inserting went wrong")
	max := pq.Max().(int)
	debug(pq.slice, t)
	assert.Equal(t, max, 9, "Inserting went wrong")

}

func TestInsertPopMax(t *testing.T) {
	pq := New(compare)
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range num {
		pq.Insert(num[n])
	}

	if pq.Count() != len(num) {
		t.Fatal("Inserting went wrong!")
	}

	if max := pq.PopMax(); max != 9 || pq.Count() != 8 {
		t.Fatalf("PopMax not working, max %v, size: %v\n", max, pq.Size())
	}
}

func TestInsertPopMaxMulti(t *testing.T) {
	pq := New(compare)
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range num {
		pq.Insert(num[n])
	}

	if pq.Count() != len(num) {
		t.Fatal("Inserting went wrong!")
	}

	if max := pq.PopMax(); max != 9 || pq.Count() != 8 {
		t.Fatalf("PopMax not working, max %v, size: %v\n", max, pq.Size())
	}
	if max := pq.PopMax(); max != 8 || pq.Count() != 7 {
		t.Fatalf("PopMax not working, max %v, size: %v\n", max, pq.Size())
	}

	debug(pq.slice, t)
}
