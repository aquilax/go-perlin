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


func Normalize2 (v * [2]float64) {
  s := math.Sqrt(v[0] * v[0] + v[1] * v[1]);
  v[0] = v[0] / s;
  v[1] = v[1] / s;
}


func Normalize3 (v * [3]float64) {
  s := math.Sqrt (v[0] * v[0] + v[1] * v[1] + v[2] * v[2]);
  v[0] = v[0] / s;
  v[1] = v[1] / s;
  v[2] = v[2] / s;
}


func PerlinInit() {
  var i int

  for i = 0; i < B; i++ {
    p[i] = i
    g1[i] =  float64((rand.Int() % (B + B)) - B) / B

    for j := 0; j < 2; j++ {
      g2[i][j] =  float64((rand.Int() % (B + B)) - B) / B
    }

    Normalize2 (&g2[i]);

    for j := 0; j < 3; j++ {
      g3[i][j] = float64((rand.Int() % (B + B)) - B) / B
    }
    Normalize3 (&g3[i])
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

func SCurve(t float64) (float64){
  return t * t * (3. - 2. * t)
}

func Lerp(t, a, b float64) (float64) {
  return a + t * (b - a) ;
}

func Noise1(arg float64) (float64) {

  var vec[1]float64
  vec[0] = arg

  if (!start){
    start = true
    PerlinInit()
  }

  t := vec[0] + N
  bx0 := int(t) & BM
  bx1 := (bx0 + 1) & BM
  rx0 := t - float64(int(t))
  rx1 := rx0 - 1.

  sx := SCurve (rx0);
  u := rx0 * g1[p[bx0]];
  v := rx1 * g1[p[bx1]];

  return Lerp(sx, u, v);
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
    val := Noise1(p);
    sum += val / scale;
    scale *= alpha;
    p *= beta;
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

