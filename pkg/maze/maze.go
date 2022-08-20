package maze

import (
    "log"
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
    // file operations
    fname := "maze.svg"
    outf, err := os.Create(fname)
    if err != nil {
        log.Panicf("ERROR: unable to create '%s' file\n", fname)
    }
    defer outf.Close()
    
    width := m.cols * m.cellSize
    height := m.rows * m.cellSize

    canvas := svg.New(outf)
    canvas.Start(width, height)
    canvas.Rect(0, 0, width, height, canvas.RGB(255, 255, 255))

    for _, c := range m.cells {
        c.DrawBorder(canvas, m.cellSize) 
    }
    canvas.End()
}

/*
func (m *Maze) Print() {
    // file operations
    fname := "output.log"
    outf, err := os.Create(fname)
    if err != nil {
        log.Panicf("ERROR: unable to creat the '%s' file\n", fname)
    }
    defer outf.Close()

    for _, c := range m.cells {
        err := os.WriteFile(outf, []byte(c.border), 0600)
        err := binary.Write(outf, binary.LittleEndian, c.border)
        if err != nil {
            log.Panicf("ERROR: could not write data to '%s' file\n", fname)
        }
    }
}
*/
