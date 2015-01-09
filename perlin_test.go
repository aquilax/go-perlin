package perlin

import (
	"math/rand"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rand.Seed(3.0)
	retCode := m.Run()
	os.Exit(retCode)
}

func TestPerlinNoise1D(t *testing.T) {
	expected := 0.0
	noise := Noise1D(10, 2, 2, 3)
	if noise != expected {
		t.Fail()
		t.Logf("Wrong node result: given: %f, expected: %f", noise, expected)
	}
}

func TestPerlinNoise2D(t *testing.T) {
	expected := 0.0
	noise := Noise2D(10, 10, 2, 2, 3)
	if noise != expected {
		t.Fail()
		t.Logf("Wrong node result: given: %f, expected: %f", noise, expected)
	}
}

func TestPerlinNoise3D(t *testing.T) {
	expected := 0.0
	noise := Noise3D(10, 10, 10, 2, 2, 3)
	if noise != expected {
		t.Fail()
		t.Logf("Wrong node result: given: %f, expected: %f", noise, expected)
	}
}
