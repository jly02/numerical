package main

import (
	"math"
	"numerical/diff"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func f(x float64) float64 {
	return math.Pow(x, 3.)
}

func main() {
	x_0 := -3.
	x_1 := 3.
	dx := 0.01
	df_dx := diff.DiffQuot(f, x_0, x_1, dx)
	df_dx_sec := diff.SymDiffQuot(f, x_0, x_1, dx)

	p := plot.New()

	p.Title.Text = "Numerical Differentiation"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err := plotutil.AddLinePoints(p,
		"f", PlotFunc(x_0, x_1, dx),
		"df/dx", PlotSlice(df_dx, x_0, dx),
		"df/dx_sec", PlotSlice(df_dx_sec, x_0, dx))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

// PlotSlice returns a plot of a given slide of floats
func PlotSlice(vals []float64, x_0, dx float64) plotter.XYs {
	pts := make(plotter.XYs, len(vals))
	x_i := x_0
	for i := range pts {
		pts[i].X = x_i
		pts[i].Y = vals[i]
		x_i += dx
	}

	return pts
}

// PlotFunc returns a plot of a function in a given range
func PlotFunc(x_0, x_1, dx float64) plotter.XYs {
	sample_num := int(math.Abs(x_1-x_0) / dx)
	pts := make(plotter.XYs, sample_num)
	x_i := x_0
	for i := range pts {
		pts[i].X = x_i
		pts[i].Y = f(x_i)
		x_i += dx
	}

	return pts
}
