/*
* Addapted from git://git.gnome.org/gegl/operations/common/perlin/perlin.c
*  Coherent noise function over 1, 2 or 3 dimensions
* (copyright Ken Perlin)
*/

package main

import (
  "fmt"
  "math"
  "rand"
)

const (
  B = 0x100
  N = 0x1000
  BM = 0xff
)

var (
  start bool = false

  p [B + B + 2]int

  g3[B + B + 2][3]float64
  g2[B + B + 2][2]float64
  g1[B + B + 2]float64

)


func normalize2 (v * [2]float64) {
  s := math.Sqrt(v[0] * v[0] + v[1] * v[1]);
  v[0] = v[0] / s;
  v[1] = v[1] / s;
}


func normalize3 (v * [3]float64) {
  s := math.Sqrt (v[0] * v[0] + v[1] * v[1] + v[2] * v[2]);
  v[0] = v[0] / s;
  v[1] = v[1] / s;
  v[2] = v[2] / s;
}


func perlinInit() {
  var i int

  for i = 0; i < B; i++ {
    p[i] = i
    g1[i] =  float64((rand.Int() % (B + B)) - B) / B

    for j := 0; j < 2; j++ {
      g2[i][j] =  float64((rand.Int() % (B + B)) - B) / B
    }

    normalize2 (&g2[i]);

    for j := 0; j < 3; j++ {
      g3[i][j] = float64((rand.Int() % (B + B)) - B) / B
    }
    normalize3 (&g3[i])
  }


  for i = i; i > 0; i-- {
    k := p[i]
    j := rand.Int() % B 
    p[i] = p[j];
    p[j] = k;
  }

  for i := 0; i < B + 2; i++ {
    p[B + i] = p[i];
    g1[B + i] = g1[i];
    for j := 0; j < 2; j++ {
      g2[B + i][j] = g2[i][j];
    }
    for j := 0; j < 3; j++ {
      g3[B + i][j] = g3[i][j];
    }
  }
}

func sCurve(t float64) (float64){
  return t * t * (3. - 2. * t)
}

func lerp(t, a, b float64) (float64) {
  return a + t * (b - a) ;
}

func noise1(arg float64) (float64) {

  var vec[1]float64
  vec[0] = arg

  if (!start){
    start = true
    perlinInit()
  }

  t := vec[0] + N
  bx0 := int(t) & BM
  bx1 := (bx0 + 1) & BM
  rx0 := t - float64(int(t))
  rx1 := rx0 - 1.

  sx := sCurve (rx0);
  u := rx0 * g1[p[bx0]];
  v := rx1 * g1[p[bx1]];

  return lerp(sx, u, v);
}

func at2(rx, ry float64, q[2]float64) (float64) {
  return rx * q[0] + ry * q[1] 
}

func noise2 (vec[2]float64) (float64) {

  if (!start){
    start = true;
    perlinInit ();
  }

  t := vec[0] + N
  bx0 := int(t) & BM
  bx1 := (bx0 + 1) & BM
  rx0 := t - float64(int(t))
  rx1 := rx0 - 1.

  t = vec[1] + N
  by0 := int(t) & BM
  by1 := (by0 + 1) & BM
  ry0 := t - float64(int(t))
  ry1 := ry0 - 1.

  i := p[bx0];
  j := p[bx1];

  b00 := p[i + by0];
  b10 := p[j + by0];
  b01 := p[i + by1];
  b11 := p[j + by1];

  sx := sCurve (rx0);
  sy := sCurve (ry0);

  q := g2[b00];
  u := at2 (rx0, ry0, q);
  q = g2[b10];
  v := at2 (rx1, ry0, q);
  a := lerp (sx, u, v);

  q = g2[b01];
  u = at2 (rx0, ry1, q);
  q = g2[b11];
  v = at2 (rx1, ry1, q);
  b := lerp (sx, u, v);

  return lerp (sy, a, b);
}



/*
   In what follows "alpha" is the weight when the sum is formed.
   Typically it is 2, As this approaches 1 the function is noisier.
   "beta" is the harmonic scaling/spacing, typically 2.
*/

func PerlinNoise1D (x, alpha, beta float64, n int) (float64) {
  var scale float64 = 1;
  var sum float64 = 0;
  p := x;

  for i := 0; i < n; i++ {
    val := noise1(p);
    sum += val / scale;
    scale *= alpha;
    p *= beta;
  }
  return sum;
}

func PerlinNoise2D (x, y, alpha, beta float64, n int) (float64) {
  var scale float64 = 1;
  var sum float64 = 0;
  var p[2]float64;

  p[0] = x;
  p[1] = y;

  for i := 0; i < n; i++ {
    val := noise2(p);
    sum += val / scale;
    scale *= alpha;
    p[0] *= beta;
    p[1] *= beta;
  }
  return sum;
}

func main (){
  var alpha float64 = 2;
  var beta float64 = 2;

  for x := 1; x < 100; x++ {
    fmt.Printf("%d;%0.8f\n",x, PerlinNoise1D(float64(x)/10, alpha, beta, 3))
  }
}

