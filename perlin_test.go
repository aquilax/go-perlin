package perlin

import "testing"

func TestPerlinNoise1D(t *testing.T) {
	noise := PerlinNoise1D(10, 2, 2, 1)
	if noise < -1 || noise > 1 {
		t.Fail();
		t.Log("Unexpected noise")
	}
}

func TestPerlinNoise2D(t *testing.T) {
	noise := PerlinNoise2D(10, 10, 2, 2, 1)
	if noise < -1 || noise > 1 {
		t.Fail();
		t.Log("Unexpected noise")
	}
}

func TestPerlinNoise3D(t *testing.T) {
	noise := PerlinNoise3D(10, 10, 10, 2, 2, 1)
	if noise < -1 || noise > 1 {
		t.Fail();
		t.Log("Unexpected noise")
	}
}
