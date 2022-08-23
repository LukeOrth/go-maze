package maze

import (
	"encoding/json"
)

func (m Maze) MarshalJSON() ([]byte, error) {
    return json.Marshal(map[string]interface{}{
        "cells":    m.cells,
        "columns":  m.cols,
        "rows":     m.rows,
        "moves":    m.moves,
        "scale":    m.scale,
    })
}

func (c Cell) MarshalJSON() ([]byte, error) {
    return json.Marshal(c.border)
}

type moves []uint8

func (u moves) MarshalJSON() ([]byte, error) {
    return json.Marshal(u)
}
