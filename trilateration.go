package main

import (
  "math"
)

type Cordinate struct {
  x, y, d float64
}

func (c1 Cordinate) trilaterate(c2, c3 Cordinate) (x, y float64) {

  return trilaterate(c1, c2, c3)
}

func trilaterate(c1, c2, c3 Cordinate) (x, y float64) {
  // https://confluence.slac.stanford.edu/display/IEPM/TULIP+Algorithm+Alternative+Trilateration+Method
  // Possiblility to divide by zero. Rearranging c1,c2,c3 may result in correct answer when that happens

  d1, d2, d3, i1, i2, i3, j1, j2, j3 := &c1.d, &c2.d, &c3.d, &c1.x, &c2.x, &c3.x, &c1.y, &c2.y, &c3.y

  x = ((((math.Pow(*d1, 2)-math.Pow(*d2, 2))+(math.Pow(*i2, 2)-math.Pow(*i1, 2))+(math.Pow(*j2, 2)-math.Pow(*j1, 2)))*(2**j3-2**j2) - ((math.Pow(*d2, 2)-math.Pow(*d3, 2))+(math.Pow(*i3, 2)-math.Pow(*i2, 2))+(math.Pow(*j3, 2)-math.Pow(*j2, 2)))*(2**j2-2**j1)) / ((2**i2-2**i3)*(2**j2-2**j1) - (2**i1-2**i2)*(2**j3-2**j2)))

  y = ((math.Pow(*d1, 2) - math.Pow(*d2, 2)) + (math.Pow(*i2, 2) - math.Pow(*i1, 2)) + (math.Pow(*j2, 2) - math.Pow(*j1, 2)) + x*(2**i1-2**i2)) / (2**j2 - 2**j1)

  return x, y
}
