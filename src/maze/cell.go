package maze

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/ajstarks/svgo"
)

type Cell struct {
    x       int
    y       int
    border  uint8
    visited bool
    current bool
}

func NewCell(x int, y int, border uint8) *Cell {
    cell := &Cell {
        x: x,
        y: y,
        border: border,
        visited: false,
        current: false,
    }
    return cell
}

func (c Cell) MarshalJSON() ([]byte, error) {
    return json.Marshal(map[string]interface{}{
        "border":   c.border,
        "x":        c.x,
        "y":        c.y,
    })
}

func (c *Cell) DrawPNG(img *image.RGBA, scale int) {
    // weight of cell wall in pixels
    weight := int(math.Ceil(float64(scale) / 4))

    // cell corners (including walls)
    x1 := c.x * (scale + weight)
    x2 := x1 + (2 * weight) + scale - 1
    y1 := c.y * (scale + weight)
    y2 := y1 + (2 * weight) + scale - 1

    black := color.RGBA{0, 0, 0, 255}
    white := color.RGBA{255, 255, 255, 255}

    for x := x1; x <= x2; x++ {
        for y := y1; y <= y2; y++ {
            // initialize all pixels to black
            img.Set(x, y, black)
            // set pixels to white (open) where needed
            if x > x1 && x < x2 {
                // cell body
                if y > y1 && y < y2 {
                    img.Set(x, y, white)
                }
                // top wall
                if y < y1 + weight && c.border & 8 == 0 {
                    img.Set(x, y, white)
                }
                // bottom wall
                if y > y2 - weight && c.border & 2 == 0 {
                    img.Set(x, y, white)
                }
            }
            if y > y1 && y < y2 {
                // left wall
                if x < x1 + weight && c.border & 1 == 0 {
                    img.Set(x, y, white)
                }
                // right wall
                if x > x2 - weight && c.border & 4 == 0 {
                    img.Set(x, y, white)
                }
            }
        }
    }
}

func (c *Cell) DrawSVG(canvas *svg.SVG, scale int) {
    // cell corners 
    x1 := c.x * (scale + 1)
    x2 := c.x * (scale + 1) + (scale + 1)
    y1 := c.y * (scale + 1)
    y2 := c.y * (scale + 1) + (scale + 1)
    s := fmt.Sprintf("stroke:black;stroke-width:%f;fill:none;", math.Ceil(float64((scale + 1)) / 4))

    // functions to draw cell walls
    borders := map[uint8]func() {
        1: func() {canvas.Line(x1, y1, x1, y2, s)},
        2: func() {canvas.Line(x1, y2, x2, y2, s)},
        3: func() {canvas.Polyline([]int{x1, x1, x2}, []int{y1, y2, y2}, s)},
        4: func() {canvas.Line(x2, y1, x2, y2, s)},
        5: func() {
            canvas.Line(x1, y1, x1, y2, s) 
            canvas.Line(x2, y1, x2, y2, s)
        },
        6: func() {canvas.Polyline([]int{x2, x2, x1}, []int{y1, y2, y2}, s)},
        7: func() {canvas.Polyline([]int{x1, x1, x2, x2}, []int{y1, y2, y2, y1}, s)},
        8: func() {canvas.Line(x1, y1, x2, y1, s)},
        9: func() {canvas.Polyline([]int{x1, x1, x2}, []int{y2, y1, y1}, s)},
        10: func() {
            canvas.Line(x1, y1, x2, y1, s)
            canvas.Line(x1, y2, x2, y2, s)
        },
        11: func() {canvas.Polyline([]int{x2, x1, x1, x2}, []int{y1, y1, y2, y2}, s)},
        12: func() {canvas.Polyline([]int{x1, x2, x2}, []int{y1, y1, y2}, s)},
        13: func() {canvas.Polyline([]int{x1, x1, x2, x2}, []int{y2, y1, y1, y2}, s)},
        14: func() {canvas.Polyline([]int{x1, x2, x2, x1}, []int{y1, y1, y2, y2}, s)},
        15: func() {canvas.Square(x1, y1, scale, s)},
    }

    if c.border > 0 {
        borders[c.border]()
    }
}

func (c *Cell) direction(n *Cell) uint {
    // direction for getting from Cell(c) to Cell(n)
    if c.x > n.x {          // left
        return 1
    } else if c.x < n.x {   // right
        return 4
    } else if c.y > n.y {   // top
        return 8
    } else {                // bottom
        return 2
    }
}

func (c *Cell) removeWall(n *Cell) {
    // remove wall between Cell(c) and Cell(n)
    dir := uint8(c.direction(n))
    c.border = c.border & ^dir
    n.border = n.border & ^(dir >> 2 | dir << 2)
}
