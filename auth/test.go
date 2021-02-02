package main

import (
	"fmt"
)

func test() error {

	type point struct {
	    x, y int
	}
	p := point{1, 2}
	fmt.Printf("point: %v\n", p)

	fmt.Println("end. ")
  return nil
}
