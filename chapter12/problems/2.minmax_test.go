// Write a series of tests for the Min and Max functions you wrote in the previous chapter.

package minmax

import (
    "testing"
)

type testpair struct {
  values []float64
  average float64
}

var mintests = []testpair{
  { []float64{1,2,3,4,5}, 1},
  { []float64{19,10,11,15,46,47}, 10},
}

var maxtests = []testpair{
  { []float64{1,2,3,4,5}, 5},
  { []float64{19,10,11,15,46,47}, 47},
}


func TestMin(t *testing.T) {
  // loop through each pair
  for _, pair := range mintests {
    v := Min(pair.values)
    if v != pair.average {
      t.Error(
        "For", pair.values,
        "expected", pair.average,
        "got", v,
      )
    }
  }
}

func TestMax(t *testing.T) {
  // loop through each pair
  for _, pair := range maxtests {
    v := Max(pair.values)
    if v != pair.average {
      t.Error(
        "For", pair.values,
        "expected", pair.average,
        "got", v,
      )
    }
  }
}