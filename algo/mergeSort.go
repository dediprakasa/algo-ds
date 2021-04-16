package main

import (
  "fmt"
)

func merge(l []int, r []int) []int {
  totalLength := len(l) + len(r)
  mergedArr := make([]int, totalLength)
  i := 0
  j := 0
  k := 0
  fmt.Println(totalLength, ">>>")
  
  for i < len(l) && j < len(r) {
    fmt.Println(i, j, k)
    if l[i] < r[j] {
      mergedArr[k] = l[i]
      i++
    } else {
      mergedArr[k] = r[j]
      j++
    }
    k++
  } 
  
  for i < len(l) {
    mergedArr[k] = l[i]
    i++
    k++
  }

  for j < len(r) {
    mergedArr[k] = r[j]
    j++
    k++
  }

  return mergedArr
}

func main() {
  leftArr := [2]int{2, 4}
  rightArr := [3]int{1, 1, 6}
  mergedArr := merge(leftArr[:], rightArr[:])
  fmt.Println(mergedArr)
}