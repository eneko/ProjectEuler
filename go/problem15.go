package main

import "fmt"

type Point struct {
  x, y int
}
var route_map = make(map[Point]int)

func routes(x, y int) int {
  if x == 0 && y == 0 {
    return 1
  }

  // Check if we have already calculated this node
  paths := route_map[Point{x,y}]
  if paths > 0 {
    return paths
  }

  // Find all paths down to 0,0
  if x > 0 {
    paths += routes(x-1, y)
  }
  if y > 0 {
    paths += routes(x, y-1)
  }
  route_map[Point{x,y}] = paths
  return paths
}

func main() {
  side := 2
  fmt.Printf("Found %v routes for an %vx%v grid\n", routes(side,side), side, side)
  side = 20
  fmt.Printf("Found %v routes for an %vx%v grid\n", routes(side,side), side, side)
}
