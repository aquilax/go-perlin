package main

import (
	"fmt"
	"os"

	"github.com/aquilax/go-perlin"
)

const (
	alpha = 2.
	beta  = 2.
	n     = 3
)

func show1d() {
	for x := 0.; x < 100; x++ {
		fmt.Printf("%0.0f\t%0.4f\n", x, perlin.Noise1D(x/10, alpha, beta, 3))
	}
}

func show2d() {
	for x := 0.; x < 10; x++ {
		for y := 0.; y < 10; y++ {
			fmt.Printf("%0.0f\t%0.0f\t%0.4f\n", x, y, perlin.Noise2D(x/10, y/10, alpha, beta, 3))
		}
	}
}

func show3d() {
	for x := 0.; x < 10; x++ {
		for y := 0.; y < 10; y++ {
			for z := 0.; z < 3; z++ {
				fmt.Printf("%0.0f\t%0.0f\t%0.0f\t%0.4f\n", x, y, z, perlin.Noise3D(x/10, y/10, z/10, alpha, beta, 3))
			}
		}
	}
}

func main() {
	d := "1"
	if len(os.Args) > 1 {
		d = os.Args[1]
	}
	switch d {
	case "2":
		show2d()
	case "3":
		show3d()
	default:
		show1d()
	}
}
