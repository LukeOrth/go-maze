package maze

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/ajstarks/svgo"
)

type Maze struct {
    cells       []*Cell
    cellSize    int
    cols        int
    rows        int
}

func NewMaze(cols int, rows int, cellSize int) *Maze {
    maze := &Maze{ cells: make([]*Cell, cols * rows), cellSize: cellSize, cols: cols, rows: rows  }

    for y := 0; y < rows; y++ {
        for x := 0; x < cols; x++ {
            maze.cells[cols * y + x] = &Cell{ x: x, y: y, border: 15, visited: false, current: false }
        }
    }

    maze.checkNeighbors(0, 0, 0, NewStack())

    return maze
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
           
            c.removeWall(randNeighbor)
            seen.Push(c)
            m.checkNeighbors(randNeighbor.x, randNeighbor.y, count + 1, seen)
            return randNeighbor
        }
    }

    if len(seen.cell) > 0 {
        c, _ := seen.Pop()
        c.current = true
        m.checkNeighbors(c.x, c.y, count + 1, seen)
    }

    return nil
}

func (m *Maze) MazeToSvg() {
    width := m.cols * m.cellSize
    height := m.rows * m.cellSize

    out, _ := os.Create(fmt.Sprintf("test.svg"))
    canvas := svg.New(out)
    canvas.Start(width, height)
    canvas.Rect(0, 0, width, height, canvas.RGB(255, 255, 255))

    for _, c := range m.cells {
        c.DrawBorder(canvas, m.cellSize) 
    }
    canvas.End()
}

func (m *Maze) Print() {
    for i, c := range m.cells {
        if i % m.cols == 0 {
            fmt.Println()
        }
        fmt.Printf("%d ", c.border)
    }
}
