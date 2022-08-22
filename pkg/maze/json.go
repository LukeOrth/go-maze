package maze

import (
	"encoding/json"
)

func (m Maze) MarshalJSON() ([]byte, error) {
    return json.Marshal(map[string]interface{}{
        "cells":    m.cells,
        "scale":    m.scale,
        "columns":  m.cols,
        "rows":     m.rows,
    })
}

func (c Cell) MarshalJSON() ([]byte, error) {
    return json.Marshal(map[string]interface{}{
        "border":   c.border,
    })
}
