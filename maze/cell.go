package maze

import (
    "encoding/json"
	"github.com/ajstarks/svgo"
)

type Cell struct {
    x       int
    y       int
    border  uint8
    visited bool
    current bool
}

func (c Cell) MarshalJSON() ([]byte, error) {
    return json.Marshal(c.border)
}

func (c *Cell) DrawBorder(canvas *svg.SVG, scale int) {
    x1, y1, x2, y2 := c.x * scale, c.y * scale, c.x * scale + scale, c.y * scale + scale
    s := "stroke:black;fill:none;"

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
    if c.x > n.x {          // left
        return 1
    } else if c.x < n.x {   // right
        return 4
    } else if c.y > n.y {   // top
        return 8
    } else {                // down
        return 2
    }
}

func (c *Cell) removeWall(n *Cell) {
    dir := uint8(c.direction(n))
    c.border = c.border & ^dir
    n.border = n.border & ^(dir >> 2 | dir << 2)
}
