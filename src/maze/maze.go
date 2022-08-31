package maze

import (
	"encoding/json"
	"image"
	"image/png"
	"io"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/ajstarks/svgo"
)

type moves []uint

type Maze struct {
    cells   []*Cell
    cols    int
    rows    int
    moves   []uint
    scale   int
}

func NewMaze(cols int, rows int, scale int) *Maze {
    maze := initMaze(cols, rows, scale)
    maze.generateMaze(0, 0, NewStack())

    return maze
}

func (m *Maze) Png(w io.Writer) {
    // dimensions
    width := m.cols * (2 * m.scale) + m.scale
    height := m.rows * (2 * m.scale) + m.scale

    // setup image
    img := image.NewRGBA(image.Rect(0, 0, width, height))

    // draw cells
    for _, c := range m.cells {
        c.DrawPNG(img, m.scale)
    }
    png.Encode(w, img)
}

func (m *Maze) Svg(w io.Writer) {
    // dimensions
    width := m.cols * (m.scale + 1)
    height := m.rows * (m.scale + 1)

    // setup canvas
    canvas := svg.New(w)
    canvas.Start(width, height)
    canvas.Rect(0, 0, width, height, canvas.RGB(255, 255, 255))

    // draw cells
    for _, c := range m.cells {
        c.DrawSVG(canvas, m.scale) 
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

func (m Maze) MarshalJSON() ([]byte, error) {
    return json.Marshal(map[string]interface{}{
        "cells":    m.cells,
        "columns":  m.cols,
        "rows":     m.rows,
        "moves":    m.moves,
        "scale":    m.scale,
    })
}

func (m *Maze) UnmarshalJSON(b []byte) error {
    // temporary struct for capturing the received JSON data
    temp := &struct {
        Cells   []uint8 `json:"cells"`
        Cols    int     `json:"columns"`
        Rows    int     `json:"rows"`
        Scale   int     `json:"scale"`
    }{} 

    if err := json.Unmarshal(b, &temp); err != nil {
        return err
    }

    // initialize values
    m.cells = make([]*Cell, temp.Cols * temp.Rows)
    m.cols = temp.Cols
    m.rows = temp.Rows
    m.scale = temp.Scale

    // create cells
    for i, border := range temp.Cells {
        x := i % temp.Cols
        y := i / temp.Cols % temp.Rows
        m.cells[i] = NewCell(x, y, border)
    }
    return nil
}

func (u moves) MarshalJSON() ([]byte, error) {
    return json.Marshal(u)
}

func initMaze(cols int, rows int, scale int) *Maze {
    // instantiate maze
    maze := &Maze{ 
        cells: make([]*Cell, cols * rows), 
        cols: cols, 
        rows: rows, 
        moves: make([]uint, 0, int(math.Pow(float64(cols), 2) + math.Pow(float64(rows), 2))),
        scale: scale,
    }

    // initalize cells
    for y := 0; y < rows; y++ {
        for x := 0; x < cols; x++ {
            maze.cells[cols * y + x] = NewCell(x, y, 15)
        }
    }

    // make entrance at top left cell
    maze.cellAt(0, 0).border = uint8(7)
    // make bottom right cell the exit
    maze.cellAt(cols - 1, rows - 1).border = uint8(13)

    return maze
}

func (m *Maze) generateMaze(x int, y int, seen *Stack) {
    // mark initial cell as visited and push to stack
    c := m.cellAt(x, y)
    c.visited = true
    seen.Push(c)

    // while stack is not empty
    for len(seen.cell) > 0 {
        // pop cell from stack
        c, _ := seen.Pop()
        c = m.cellAt(c.x, c.y)

        // get any unvisited neighbors
        neighbors := m.unvisitedNeighbors(c)
        if len(neighbors) > 0 {
            // push cell to stack
            seen.Push(c)

            // get random unvisited neighbor
            rand.Seed(time.Now().UnixNano())
            random := rand.Intn(len(neighbors))
            randNeighbor := neighbors[(random) % len(neighbors)]

            // remove wall between cells
            c.removeWall(randNeighbor)

            // mark neighbor as visited and push to stack
            randNeighbor.visited = true
            seen.Push(randNeighbor)

            // track the move
            m.moves = append(m.moves, c.direction(randNeighbor))
        }
    }
}

func (m *Maze) cellAt(x int, y int) *Cell {
    // return nil if index is invalid
    if x < 0 || y < 0 || x > m.cols - 1 || y > m.rows - 1 {
        return nil
    }
    return m.cells[m.cols * y + x]
}

func (m *Maze) unvisitedNeighbors(c *Cell) []*Cell {
    // get all neighbors
    neighbors := []*Cell {
        m.cellAt(c.x, c.y - 1), // top
        m.cellAt(c.x + 1, c.y), // right
        m.cellAt(c.x, c.y + 1), // bottom
        m.cellAt(c.x - 1, c.y), // left
    }

    // return neighbors that have a valid index and aren't visited
    unvisited := make([]*Cell, 0, 4)
    for _, n := range neighbors {
        if n != nil && !n.visited {
            unvisited = append(unvisited, n)
        }
    }
    return unvisited
}
