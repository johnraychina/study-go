package main

import (
	"fmt"

	"github.com/mjibson/go-dsp/fft"
)

func main() {
	fmt.Println(fft.FFTReal([]float64{1, 2, 3}))
}

//import (
//	"fmt"
//	"math"
//	"math/cmplx"
//	"github.com/mjibson/go-dsp/fft"
//)

//func main() {
//	numSamples := 8
//
//	// Equation 3-10.
//	x := func(n int) float64 {
//		wave0 := math.Sin(2.0 * math.Pi * float64(n) / 8.0)
//		wave1 := 0.5 * math.Sin(2*math.Pi*float64(n)/4.0+3.0*math.Pi/4.0)
//		return wave0 + wave1
//	}
//
//	// Discretize our function by sampling at 8 points.
//	a := make([]float64, numSamples)
//	for i := 0; i < numSamples; i++ {
//		a[i] = x(i)
//	}
//
//	X := fft.FFTReal(a)
//
//	// Print the magnitude and phase at each frequency.
//	for i := 0; i < numSamples; i++ {
//		r, θ := cmplx.Polar(X[i])
//		θ *= 360.0 / (2 * math.Pi)
//		if dsputils.Float64Equal(r, 0) {
//			θ = 0 // (When the magnitude is close to 0, the angle is meaningless)
//		}
//		fmt.Printf("X(%d) = %.1f ∠ %.1f°\n", i, r, θ)
//	}
//}
