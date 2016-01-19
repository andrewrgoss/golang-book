// We copied the average function from chapter 7 to our new package. Create Min and Max functions which find the minimum and maximum values in a slice of float64s.

package main

import (
    "fmt"
)

// In Go if something starts with a capital letter that means other packages (and programs) are able to see it. 

// Average finds the average of a series of numbers within an array
func Average(arr []float64) float64 {
  total := float64(0)
  for _, x := range arr {
    total += x
  }
  return total / float64(len(arr))
}

// Min finds the lowest number out of an array slice
func Min(s []float64) float64 {
    min := s[0]  
    for i := 1; i < len(s); i++ {
        if s[i] < min {
            min = s[i]
        }
    }
    return min
}

// Max finds the highest number out of an array slice
func Max(s []float64) float64 {
    max := s[0]  
    for i := 1; i < len(s); i++ {
        if s[i] > max {
            max = s[i]
        }
    }
    return max
}

func main() {  
    arr := []float64{1,2,3,4,5,6,7}
    s := arr[1:6]
    
    fmt.Println("Starting array=", arr)
    fmt.Println("Array average=", Average(arr))
    fmt.Println("Slice s from array=", s)
    fmt.Println("Minimum value of slice s =", Min(s))
    fmt.Println("Maximum value of slice s =", Max(s))
}