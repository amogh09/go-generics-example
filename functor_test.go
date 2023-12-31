package main

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSecsSum1(t *testing.T) {
	assert.Equal(t, 360, SecondsSum1([]int{1, 2, 3}))
}

func TestScoresSum1(t *testing.T) {
	assert.Equal(t, 240, ScoresSum1([]SubjectScore{
		{"math", 70, true},
		{"science", 80, true},
		{"language", 90, true}}),
	)
}

func TestSliceMap(t *testing.T) {
	assert.Equal(t,
		[]int{0, 2, 4},
		MapSlice(func(x int) int { return x * 2 }, []int{0, 1, 2}))
}

func TestSum(t *testing.T) {
	assert.Equal(t, 10, Sum([]int{1, 2, 3, 4}))
}

func TestSecsSum2(t *testing.T) {
	assert.Equal(t, 360, SecondsSum2([]int{1, 2, 3}))
}

func TestScoresSum2(t *testing.T) {
	assert.Equal(t, 240, ScoresSum2([]SubjectScore{
		{"math", 70, true},
		{"science", 80, true},
		{"language", 90, true}}),
	)
}

func TestRetry(t *testing.T) {
	t.Skip() // Remote to test
	contents, err := Retry(100, time.Second, func() ([]byte, error) {
		return ioutil.ReadFile("test.txt")
	})
	require.NoError(t, err)
	t.Log("contents: ", string(contents))
}

func TestFromPointer(t *testing.T) {
	var x *string = nil
	assert.Equal(t, "default", FromPointer(x, "default"))

	y := "hello"
	x = &y
	assert.Equal(t, "hello", FromPointer(x, "default"))
}

func TestToPointer(t *testing.T) {
	assert.Equal(t, 5, *ToPointer(5))
}

func TestGroupBy(t *testing.T) {
	subjectScores := []SubjectScore{
		{"math", 85, true},
		{"science", 90, true},
		{"math", 92, true},
	}
	grouped := groupBy(GetSubject, GetScore, subjectScores)
	assert.Equal(t, map[string][]int{"math": {85, 92}, "science": {90}}, grouped)
}

func TestScoresOf(t *testing.T) {
	subjectScores := []SubjectScore{
		{"math", 85, true},
		{"science", 90, true},
		{"math", 92, true},
	}
	assert.Equal(t, []int{85, 92}, scoresOf("math", subjectScores))
}

func TestTreeMap(t *testing.T) {
    tree := EmptyTreeMap[int, string]()
	assert.Nil(t, tree.Lookup(5))

	tree.Insert(5, "a")
	assert.Equal(t, "a", *tree.Lookup(5))

	tree.Insert(3, "b")
	tree.Insert(7, "c")
	assert.Equal(t, "a", *tree.Lookup(5))
	assert.Equal(t, "b", *tree.Lookup(3))
	assert.Equal(t, "c", *tree.Lookup(7))
}
