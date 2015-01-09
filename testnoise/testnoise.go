package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aquilax/go-perlin"
)

const (
	alpha = 2.
	beta  = 2.
	n     = 3
)

func show1d(seed int64) {
	p := perlin.NewPerlin(alpha, beta, n, seed)
	for x := 0.; x < 100; x++ {
		fmt.Printf("%0.0f\t%0.4f\n", x, p.Noise1D(x/10))
	}
}

func show2d(seed int64) {
	p := perlin.NewPerlin(alpha, beta, n, seed)
	for x := 0.; x < 10; x++ {
		for y := 0.; y < 10; y++ {
			fmt.Printf("%0.0f\t%0.0f\t%0.4f\n", x, y, p.Noise2D(x/10, y/10))
		}
	}
}

func show3d(seed int64) {
	p := perlin.NewPerlin(alpha, beta, n, seed)
	for x := 0.; x < 10; x++ {
		for y := 0.; y < 10; y++ {
			for z := 0.; z < 3; z++ {
				fmt.Printf("%0.0f\t%0.0f\t%0.0f\t%0.4f\n", x, y, z, p.Noise3D(x/10, y/10, z/10))
			}
		}
	}
}

func main() {
	d := "1"
	if len(os.Args) > 1 {
		d = os.Args[1]
	}
	seed := int64(time.Now().Nanosecond())
	switch d {
	case "2":
		show2d(seed)
	case "3":
		show3d(seed)
	default:
		show1d(seed)
	}
}
