package main

import (
	"sort"
	"time"

	"golang.org/x/exp/constraints"
)

// Converts minutes to seconds.
func MinutesToSeconds(minutes int) int { return minutes * 60 }

// Sums up the seconds in list of minutes.
func SecondsSum1(minutesArr []int) int {
	res := 0
	for _, m := range minutesArr {
		res += MinutesToSeconds(m)
	}
	return res
}

// Represents score of a student in a particular subject.
type SubjectScore struct {
	subject   string
	score     int
	important bool
}

func GetSubject(s SubjectScore) string {
	return s.subject
}

func GetScore(s SubjectScore) int {
	return s.score
}

func IsImportant(s SubjectScore) bool {
	return s.important
}

// Returns the sum of all scores of a list of subject scores.
func ScoresSum1(subjectScores []SubjectScore) int {
	res := 0
	for _, s := range subjectScores {
		res += GetScore(s)
	}
	return res
}

// Returns the sum of seconds of a list of minutes.
func SecondsSum2(minutesArr []int) int {
	return Sum(MapSlice(MinutesToSeconds, minutesArr))
}

// Returns the sum of all scores of a list of subject scores.
func ScoresSum2(subjectScores []SubjectScore) int {
	return Sum(MapSlice(GetScore, subjectScores))
}

// Maps each element of a slice using a given mapper function.
// func MapSlice(f func(interface{}) interface{}, xs []interface{}) []interface{} {
// 	res := make([]interface{}, 0, len(xs))
// 	for _, x := range xs {
// 		res = append(res, f(x))
// 	}
// 	return res
// }

// func secsSum2(minutesArr []int) int {
//     secs := mapSliceI(minsToSecs, minutesArr)
// }

// Maps each element of a slice using a given mapper function.
func MapSlice[A any, B any](f func(A) B, xs []A) []B {
	res := make([]B, 0, len(xs))
	for _, x := range xs {
		res = append(res, f(x))
	}
	return res
}

// Sums up slice's elements
func Sum[T constraints.Integer | constraints.Float](xs []T) T {
	if len(xs) == 0 {
		return 0
	}

	res := xs[0]
	for _, x := range xs[1:] {
		res += x
	}
	return res
}

// Returns sum of scores of important subjects.
func ImportantSubjectScores1(subjectScores []SubjectScore) int {
	res := 0
	for _, ss := range subjectScores {
		if IsImportant(ss) {
			res += GetScore(ss)
		}
	}
	return res
}

// Returns sum of scores of important subjects.
func ImportantSubjectScores2(subjectScores []SubjectScore) int {
	return Sum(MapSlice(GetScore, FilterSlice(IsImportant, subjectScores)))
}

// Filters a slice.
func FilterSlice[A any](f func(A) bool, xs []A) []A {
	res := []A{}
	for _, x := range xs {
		if f(x) {
			res = append(res, x)
		}
	}
	return res
}

// Retry a function maxAttempts times or until success.
func Retry[T any](maxAttempts int, interval time.Duration, f func() (T, error)) (T, error) {
	var result T
	var err error
	for i := 0; i < maxAttempts; i++ {
		result, err = f()
		if err != nil {
			time.Sleep(interval)
			continue
		}
		break
	}
	return result, err
}

// Safely dereference a pointer
func FromPointer[T any](p *T, def T) T {
	if p == nil {
		return def
	}
	return *p
}

// Returns a pointer to a value
func ToPointer[T any](x T) *T {
	v := x
	return &v
}

// BEGIN extra material
func groupBy[A any, K comparable, V any](keyFn func(A) K, valFn func(A) V, xs []A) map[K][]V {
	res := map[K][]V{}
	for _, x := range xs {
		key := keyFn(x)
		value := []V{valFn(x)}
		if existing, ok := res[key]; ok {
			value = append(existing, valFn(x))
		}
		res[key] = value
	}
	return res
}

func scoresOf(subject string, subjectScores []SubjectScore) []int {
	return groupBy(GetSubject, GetScore, subjectScores)[subject]
}

func sortSlice[T constraints.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

func average[T constraints.Integer | constraints.Float](xs []T) float64 {
	return float64(Sum(xs)) / float64(len(xs))
}

func mapMap[A comparable, B any, C any](f func(B) C, kvs map[A]B) map[A]C {
	res := map[A]C{}
	for k, v := range kvs {
		res[k] = f(v)
	}
	return res
}

// ------- BEGIN Tree Example ----------
type TreeMap[K constraints.Ordered, V any] struct {
	root *BinaryTreeNode[K, V]
}

func EmptyTreeMap[K constraints.Ordered, V any]() *TreeMap[K, V] {
	return &TreeMap[K, V]{nil}
}

type BinaryTreeNode[K constraints.Ordered, V any] struct {
	left, right *BinaryTreeNode[K, V]
	key         K
	value       V
}

func (t *TreeMap[K, V]) Lookup(key K) *V {
	return t.root.Lookup(key)
}

func (t *BinaryTreeNode[K, V]) Lookup(key K) *V {
	if t == nil {
		return nil
	} else if key < t.key {
		return t.left.Lookup(key)
	} else if key > t.key {
		return t.right.Lookup(key)
	} else {
		return &t.value
	}
}

func (t *TreeMap[K, V]) Insert(key K, value V) {
	t.root = t.root.Insert(key, value)
}

func (t *BinaryTreeNode[K, V]) Insert(key K, value V) *BinaryTreeNode[K, V] {
	if t == nil {
		t = &BinaryTreeNode[K, V]{nil, nil, key, value}
		return t
	} else if key < t.key {
		t.left = t.left.Insert(key, value)
		return t
	} else if key > t.key {
		t.right = t.right.Insert(key, value)
		return t
	} else {
		t.value = value
		return t
	}
}
