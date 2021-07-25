package perlin_test

import (
	"fmt"
	"math/rand"

	"github.com/aquilax/go-perlin"
)

const (
	alpha       = 2.
	beta        = 2.
	n           = 3
	seed  int64 = 100
)

func ExampleNewPerlinRandSource() {
	p := perlin.NewPerlinRandSource(alpha, beta, n, rand.NewSource(seed))
	for x := 0.; x < 3; x++ {
		fmt.Printf("%0.0f;%0.4f\n", x, p.Noise1D(x/10))
	}
	// Output:
	// 0;0.0000
	// 1;-0.0086
	// 2;-0.0017
}

func ExamplePerlin_Noise1D() {
	p := perlin.NewPerlin(alpha, beta, n, seed)
	for x := 0.; x < 3; x++ {
		fmt.Printf("%0.0f;%0.4f\n", x, p.Noise1D(x/10))
	}
	// Output:
	// 0;0.0000
	// 1;-0.0086
	// 2;-0.0017
}

func ExamplePerlin_Noise2D() {
	p := perlin.NewPerlin(alpha, beta, n, seed)
	for x := 0.; x < 2; x++ {
		for y := 0.; y < 2; y++ {
			fmt.Printf("%0.0f;%0.0f;%0.4f\n", x, y, p.Noise2D(x/10, y/10))
		}
	}
	// Output:
	// 0;0;0.0000
	// 0;1;-0.2002
	// 1;0;-0.3389
	// 1;1;-0.5045
}

func ExamplePerlin_Noise3D() {
	p := perlin.NewPerlin(alpha, beta, n, seed)
	for x := 0.; x < 2; x++ {
		for y := 0.; y < 2; y++ {
			for z := 0.; z < 2; z++ {
				fmt.Printf("%0.0f;%0.0f;%0.0f;%0.4f\n", x, y, z, p.Noise3D(x/10, y/10, z/10))
			}
		}
	}
	// Output:
	// 0;0;0;0.0000
	// 0;0;1;0.2616
	// 0;1;0;-0.0755
	// 0;1;1;0.2020
	// 1;0;0;-0.2138
	// 1;0;1;0.0616
	// 1;1;0;-0.2208
	// 1;1;1;0.0304
}
