package perlin

import (
	"testing"
)

const (
	seed = 123
)

func Test_PerlinNoise1D(t *testing.T) {
	expected := 0.0
	p := NewPerlin(2, 2, 3, seed)
	noise := p.Noise1D(10)
	if noise != expected {
		t.Fail()
		t.Logf("Wrong node result: given: %f, expected: %f", noise, expected)
	}
}

func Test_PerlinNoise2D(t *testing.T) {
	expected := 0.0
	p := NewPerlin(2, 2, 3, seed)
	noise := p.Noise2D(10, 10)
	if noise != expected {
		t.Fail()
		t.Logf("Wrong node result: given: %f, expected: %f", noise, expected)
	}
}

func Test_PerlinNoise3D(t *testing.T) {
	expected := 0.0
	p := NewPerlin(2, 2, 3, seed)
	noise := p.Noise3D(10, 10, 10)
	if noise != expected {
		t.Fail()
		t.Logf("Wrong node result: given: %f, expected: %f", noise, expected)
	}
}

func Benchmark_PerlinNoise1D(b *testing.B) {
	p := NewPerlin(2, 2, 3, seed)
	for n := 0; n < b.N; n++ {
		p.Noise1D(10)
	}
}

func Benchmark_PerlinNoise2D(b *testing.B) {
	p := NewPerlin(2, 2, 3, seed)
	for n := 0; n < b.N; n++ {
		p.Noise2D(10, 10)
	}
}
func Benchmark_PerlinNoise3D(b *testing.B) {
	p := NewPerlin(2, 2, 3, seed)
	for n := 0; n < b.N; n++ {
		p.Noise3D(10, 10, 10)
	}
}
