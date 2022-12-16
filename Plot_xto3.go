package main

import (
	"math"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
)

func main() {
	// Create a new plot
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	// Set the plot title
	p.Title.Text = "x^3"

	// Create a slice of x-values
	var xs []float64
	for i := -2.0; i <= 2.0; i += 0.1 {
		xs = append(xs, i)
	}

	// Create a slice of y-values by evaluating the function for each x-value
	var ys []float64
	for _, x := range xs {
		ys = append(ys, math.Pow(x, 3))
	}

	// Create a scatter plotter with the x and y values
	scatter, err := plotter.NewScatter(plotter.XYs{
		X: xs,
		Y: ys,
	})
	if err != nil {
		panic(err)
	}

	// Add the scatter plotter to the plot
	p.Add(scatter)

	// Save the plot to a PNG file
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "x3.png"); err != nil {
		panic(err)
	}
}
