package main

import "golang.org/x/exp/constraints"

// Converts minutes to seconds.
func minsToSecs(mins int) int { return mins * 60 }

// Returns the sum of seconds represented by an array of minutes.
func secsSum1(minutesArr []int) int {
	res := 0
	for _, m := range minutesArr {
		res += minsToSecs(m)
	}
	return res
}

// func secsSum2(minutesArr []int) int {
//     secs := fmapI(minsToSecs, minutesArr)
// }

// Returns the sum of seconds of a list of minutes.
func secsSum2(minutesArr []int) int {
	return sum(mapSlice(minsToSecs, minutesArr))
}

// Represents score of a student in a particular subject.
type subjectScore struct {
	subject   string
	score     int
	important bool
}

func GetSubject(s subjectScore) string {
	return s.subject
}

func GetScore(s subjectScore) int {
	return s.score
}

func IsImportant(s subjectScore) bool {
	return s.important
}

// Returns the sum of all scores of a list of subject scores.
func scoresSum1(subjectScores []subjectScore) int {
	res := 0
	for _, s := range subjectScores {
		res += GetScore(s)
	}
	return res
}

// Returns the sum of all scores of a list of subject scores.
func scoresSum2(subjectScores []subjectScore) int {
	return sum(mapSlice(GetScore, subjectScores))
}

// Returns sum of scores of important subjects.
func importantSubjectScores1(subjectScores []subjectScore, subject string) int {
	res := 0
	for _, ss := range subjectScores {
		if IsImportant(ss) {
			res += GetScore(ss)
		}
	}
	return res
}

// Returns sum of scores of important subjects.
func importantSubjectScores2(subjectScores []subjectScore, subject string) int {
	return sum(mapSlice(GetScore, filterSlice(IsImportant, subjectScores)))
}

func sum[T constraints.Integer | constraints.Float](xs []T) T {
	if len(xs) == 0 {
		return 0
	}

	res := xs[0]
	for _, x := range xs[1:] {
		res += x
	}
	return res
}

func mapSliceI(f func(interface{}) interface{}, xs []interface{}) []interface{} {
	res := make([]interface{}, 0, len(xs))
	for _, x := range xs {
		res = append(res, f(x))
	}
	return res
}

// Maps each element of a slice using a given mapper function.
func mapSlice[A any, B any](f func(A) B, xs []A) []B {
	res := make([]B, 0, len(xs))
	for _, x := range xs {
		res = append(res, f(x))
	}
	return res
}

// Filters a slice.
func filterSlice[A any](f func(A) bool, xs []A) []A {
	res := []A{}
	for _, x := range xs {
		if f(x) {
			res = append(res, x)
		}
	}
	return res
}

// ------- BEGIN Tree Example ----------

type BinaryTree[K constraints.Ordered, V any] struct {
	left, right *BinaryTree[K, V]
	key         K
	value       *V
}

func (t *BinaryTree[K, V]) Lookup(key K) *V {
	if t == nil {
		return nil
	} else if key < t.key {
		return t.left.Lookup(key)
	} else if key > t.key {
		return t.right.Lookup(key)
	} else {
		return t.value
	}
}

func (t *BinaryTree[K, V]) Insert(key K, value *V) *BinaryTree[K, V] {
	if t == nil {
		return &BinaryTree[K, V]{nil, nil, key, value}
	} else if key < t.key {
		return t.left.Insert(key, value)
	} else if key > t.key {
		return t.right.Insert(key, value)
	} else {
		return &BinaryTree[K, V]{t.left, t.right, key, value}
	}
}
