package maze

import (
	"fmt"

	"github.com/ajstarks/svgo"
)

type Cell struct {
    x       int
    y       int
    border  uint8
    visited bool
    current bool
}

func (c *Cell) removeWall(n *Cell) {
    xDelta := c.x - n.x
    yDelta := c.y - n.y

    // X
    if xDelta == 1 {
        fmt.Println("removing: cur - left, nex - right")
        c.border = c.border & 14
        n.border = n.border & 11
    } else if xDelta == -1 {
        fmt.Println("removing: cur - right, nex - left")
        c.border = c.border & 11
        n.border = n.border & 14
    }
    // Y
    if yDelta == 1 {
        fmt.Println("removing: cur - top, nex - bottom")
        c.border = c.border & 7
        n.border = n.border & 13
    } else if yDelta == -1 {
        fmt.Println("removing: cur - bottom, nex - top")
        c.border = c.border & 13
        n.border = n.border & 7
    }
}


func (c *Cell) DrawBorder(canvas *svg.SVG, scale int) {
    x1, y1, x2, y2 := c.x * scale, c.y * scale, c.x * scale + scale, c.y * scale + scale
    s := "stroke:black;fill:none"

    borders := map[uint8]func() {
        1: func() {canvas.Line(x1, y1, x1, y2, s)},
        2: func() {canvas.Line(x1, y2, x2, y2, s)},
        3: func() {canvas.Polyline([]int{x1, x1, x2}, []int{y1, y2, y2}, s)},
        4: func() {canvas.Line(x2, y1, x2, y2, s)},
        5: func() {
            canvas.Line(x1, y1, x2, y1, s) 
            canvas.Line(x1, y2, x2, y2, s)
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
        15: func() {canvas.Square(c.x * scale, c.y * scale, scale, "stroke:none;fill:none")},
    }
    borders[c.border]()
}
