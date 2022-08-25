package maze

import (
	"encoding/json"
	"io"
	"log"
	"math"
	"math/rand"
	"time"
    
    "github.com/ajstarks/svgo"
)

type Maze struct {
    cells   []*Cell
    cols    int
    rows    int
    moves   []uint
    scale   int
}

type moves []uint8

func (m Maze) MarshalJSON() ([]byte, error) {
    return json.Marshal(map[string]interface{}{
        "cells":    m.cells,
        "columns":  m.cols,
        "rows":     m.rows,
        "moves":    m.moves,
        "scale":    m.scale,
    })
}

func (u moves) MarshalJSON() ([]byte, error) {
    return json.Marshal(u)
}

func NewMaze(cols int, rows int, scale int) *Maze {
    maze := &Maze{ 
        cells: make([]*Cell, cols * rows), 
        cols: cols, 
        rows: rows, 
        moves: make([]uint, 0, int(math.Pow(float64(cols), 2) + math.Pow(float64(rows), 2))),
        scale: scale + 1, 
    }

    for y := 0; y < rows; y++ {
        for x := 0; x < cols; x++ {
            maze.cells[cols * y + x] = &Cell{
                x: x, 
                y: y, 
                border: 15, 
                visited: false, 
                current: false,
            }
        }
    }
    maze.cellAt(0, 0).border = uint8(7)
    maze.cellAt(cols - 1, rows - 1).border = uint8(13)

    // start at cell(0, 0)
    maze.checkNeighbors(0, 0, 0, NewStack())

    return maze
}

func (m *Maze) Svg(w io.Writer) {
    width := m.cols * m.scale
    height := m.rows * m.scale

    canvas := svg.New(w)
    canvas.Start(width, height)
    canvas.Rect(0, 0, width, height, canvas.RGB(255, 255, 255))

    for _, c := range m.cells {
        c.DrawBorder(canvas, m.scale) 
    }

    canvas.End()
}

func (m *Maze) Json() []byte {
    data, err := json.Marshal(m)

    if err != nil {
        log.Panicln("ERROR: unable to marshal JSON")
    }
    
    return data
}

func (m *Maze) cellAt(x int, y int) *Cell {
    if x < 0 || y < 0 || x > m.cols - 1 || y > m.rows - 1 {
        return nil
    }
    
    return m.cells[m.cols * y + x]
}

func (m *Maze) checkNeighbors(x int, y int, count int, seen *Stack) *Cell {
    c := m.cellAt(x, y)
    c.current = false

    neighbors := []*Cell{ 
        m.cellAt(x, y - 1),
        m.cellAt(x + 1, y), 
        m.cellAt(x, y + 1), 
        m.cellAt(x - 1, y),
    }

    rand.Seed(time.Now().UnixNano())
    random := rand.Intn(4)

    for i := range neighbors {
        randNeighbor := neighbors[(random + i) % 4]
        if randNeighbor != nil && !randNeighbor.visited {
            randNeighbor.visited = true
            randNeighbor.current = true
          
            m.moves = append(m.moves, c.direction(randNeighbor))
            c.removeWall(randNeighbor)
            seen.Push(c)
            m.checkNeighbors(randNeighbor.x, randNeighbor.y, count + 1, seen)
            return randNeighbor
        }
    }

    if len(seen.cell) > 0 {
        n, _ := seen.Pop()
        m.moves = append(m.moves, c.direction(n))
        n.current = true
        m.checkNeighbors(n.x, n.y, count + 1, seen)
    }

    return nil
}
