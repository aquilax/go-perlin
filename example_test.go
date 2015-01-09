package perlin_test

import (
	"fmt"

	"github.com/aquilax/go-perlin"
)

const (
	alpha       = 2.
	beta        = 2.
	n           = 3
	seed  int64 = 100
)

func ExamplePerlin_Noise1D() {
	p := perlin.NewPerlin(alpha, beta, n, seed)
	for x := 0.; x < 3; x++ {
		fmt.Printf("%0.0f\t%0.4f\n", x, p.Noise1D(x/10))
	}
	// Output:
	// 0	0.0000
	// 1	-0.1670
	// 2	-0.2011
}

func ExamplePerlin_Noise2D() {
	p := perlin.NewPerlin(alpha, beta, n, seed)
	for x := 0.; x < 2; x++ {
		for y := 0.; y < 2; y++ {
			fmt.Printf("%0.0f\t%0.0f\t%0.4f\n", x, y, p.Noise2D(x/10, y/10))
		}
	}
	// Output:
	// 0	0	0.0000
	// 0	1	-0.0060
	// 1	0	0.1230
	// 1	1	0.0625
}

func ExamplePerlin_Noise3D() {
	p := perlin.NewPerlin(alpha, beta, n, seed)
	for x := 0.; x < 2; x++ {
		for y := 0.; y < 2; y++ {
			for z := 0.; z < 2; z++ {
				fmt.Printf("%0.0f\t%0.0f\t%0.0f\t%0.4f\n", x, y, z, p.Noise3D(x/10, y/10, z/10))
			}
		}
	}
	// Output:
	// 0	0	0	0.0000
	// 0	0	1	0.1881
	// 0	1	0	0.1384
	// 0	1	1	0.2507
	// 1	0	0	-0.0416
	// 1	0	1	0.1232
	// 1	1	0	0.0536
	// 1	1	1	0.1826
}
