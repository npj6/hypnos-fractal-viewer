package fractal

import (
	"math"
	"math/cmplx"
	"npj1610/hypnos-fractal-viewer/types"
)

type Mandelbrot struct {
	types.ScreenInt

	colorizer ColorizerMandelbrot
	initial   complex128
	maxVal    float64
	limit     int
}

func (mb Mandelbrot) Limit() int {
	return mb.limit
}

func (mb Mandelbrot) MaxVal() float64 {
	return mb.maxVal
}

func (mb Mandelbrot) CalcScreen(positions types.CoordinatesComplex) types.ScreenInt {
	topLeft := positions.TopLeft()
	size := positions.BottomRight() - positions.TopLeft()
	rightstep := complex(real(size)/float64(mb.Width()), 0)
	downstep := complex(0, imag(size)/float64(mb.Height()))

	mb.colorizer.PreCalc(&mb)

	*mb.Screen() = make([][][]int, mb.Height())
	for row := range *mb.Screen() {
		(*mb.Screen())[row] = make([][]int, mb.Width())

		for point := range (*mb.Screen())[row] {
			complexPoint := topLeft + complex(float64(point), 0)*rightstep + complex(float64(row), 0)*downstep

			(*mb.Screen())[row][point] = make([]int, 1)
			(*mb.Screen())[row][point][0] = mb.calcPoint(complexPoint)

			mb.colorizer.ForPoint(&mb, row, point, complexPoint)
		}

		mb.colorizer.ForRow(&mb, row, float64(row)*real(downstep))
	}

	mb.colorizer.PostCalc(&mb)

	return mb.ScreenInt
}

func (mb Mandelbrot) calcPoint(c complex128) int {
	p := mb.initial
	for counter := 0; counter < mb.Limit(); counter++ {
		if mb.MaxVal() < math.Pow(real(p), 2)+math.Pow(imag(p), 2) {
			return counter
		}
		p = cmplx.Pow(p, complex128(2)) + c
	}
	return 0
}
