// Write a series of tests for the Min and Max functions you wrote in the previous chapter.

package minmax

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