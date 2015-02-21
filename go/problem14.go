package main

import "fmt"

func nextElement(n int) int {
  if n&1 == 1 {
    return 3*n + 1
  } else {
    return n/2
  }
}

func Collatz(n int) []int {
  var seq []int
  seq = append(seq, n)
  next := nextElement(n)
  for next > 1 {
    seq = append(seq, next)
    next = nextElement(next)
  }
  seq = append(seq, 1)
  return seq
}

func printSequence(seq []int) {
  for i:=0; i<len(seq)-1; i++ {
    fmt.Printf("%v", seq[i])
    fmt.Print(" -> ")
  }
  fmt.Printf("%v\n", seq[len(seq)-1])
}

func main() {
  // solution 1: brute force, down from 1m-1 to 2, find max length
  max := 0
  max_i := 0
  for i := 999999; i > 1; i-- {
    fmt.Printf("Computing sequence for %v: ", i)
    seq := Collatz(i)
    seq_len := len(seq)
    fmt.Printf("%v items\n", seq_len)
    if seq_len > max {
      max = seq_len
      max_i = i
    }
  }
  fmt.Printf("Longest sequence: f(%v) [%v items]\n", max_i, max)
  printSequence(Collatz(max_i))

  // solution 2: save sequences to avoid calculating same sequence multiple times
}
