package math

import "testing"

type testpair struct {
	values   []float64
	expected float64
}

var averageTests = []testpair{
	{[]float64{1, 2, 3, 4}, 2.5},
	{[]float64{1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
	{[]float64{}, 0},
}

var minTests = []testpair{
	{[]float64{1, 2, 3, 4}, 1},
	{[]float64{-100, -20, 0, 20, 100}, -100},
	{[]float64{1000, 2000, 3000, 4444}, 1000},
}

var maxTests = []testpair{
	{[]float64{1, 2, 3, 4}, 4},
	{[]float64{-100, -20, 0, 20, 100}, 100},
	{[]float64{1000, 2000, 3000, 4444}, 4444},
}

func TestAverage(t *testing.T) {
	for _, pair := range averageTests {
		v := Average(pair.values)
		if v != pair.expected {
			t.Error("For", pair.values, "expected", pair.expected, "got", v)
		}
	}
}

func TestMin(t *testing.T) {
	for _, pair := range minTests {
		v := Min(pair.values)
		if v != pair.expected {
			t.Error("For", pair.values, "expected", pair.expected, "got", v)
		}
	}
}

func TestMax(t *testing.T) {
	for _, pair := range maxTests {
		v := Max(pair.values)
		if v != pair.expected {
			t.Error("For", pair.values, "expected", pair.expected, "got", v)
		}
	}
}
