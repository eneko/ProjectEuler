package main

import "fmt"

// use a hashmap to avoid calculating the same sequence multiple times
var seq_map = make(map[int]int)

func nextElement(n int) int {
  if n&1 == 1 {
    return 3*n + 1
  } else {
    return n/2
  }
}

func Collatz(n int) (seq_len int) {
  seq_len = seq_map[n]
  if seq_len == 0 {
    seq_len = 1 + Collatz(nextElement(n))
    seq_map[n] = seq_len
  }
  return
}

func main() {
  seq_map[1] = 1
  max := 0
  max_i := 0
  for i := 1; i < 1000000; i++ {
    fmt.Printf("Computing sequence for %v: ", i)
    seq_len := Collatz(i)
    fmt.Printf("%v items\n", seq_len)
    if seq_len > max {
      max = seq_len
      max_i = i
    }
  }

  fmt.Printf("Longest sequence: f(%v) [%v items]\n", max_i, max)
}
