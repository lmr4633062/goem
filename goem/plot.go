package goem

import (
	"fmt"
	"image/color"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
)

func (em EM) Plot() {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "EM Algorithm Plot"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	bs, err := plotter.NewBubbles(em.clusterTriples(), vg.Points(40), vg.Points(50))
	if err != nil {
		panic(err)
	}
	bs.Color = color.RGBA{R: 255, B: 255, A: 255}
	p.Add(bs)

	ss, err := plotter.NewScatter(em.dataTriples())
	if err != nil {
		panic(err)
	}
	bs.Color = color.Black //color.RGBA{R: 255, B: 255, A: 255}
	p.Add(ss)

	if err := p.Save(10*vg.Inch, 10*vg.Inch, "bubble.png"); err != nil {
		panic(err)
	}
}

func (em EM) dataTriples() plotter.XYs {
	data := make(plotter.XYs, len(em.data))
	for i, v := range em.data {
		data[i].X = v[0]
		data[i].Y = v[1]
	}
	return data
}

func (em EM) clusterTriples() plotter.XYZs {
	data := make(plotter.XYZs, em.k)
	for i, v := range em.mu {
		data[i].X = v[0]
		data[i].Y = v[1]
		data[i].Z = em.pi[i] * 10
		fmt.Println(em.pi[i])
	}

	return data
}
