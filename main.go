package main

import (
	"fmt"

	"github.com/LukeOrth/go-maze/maze"
)

func main() {
    maze := maze.NewMaze(10, 10, 3)
    data := maze.Json()
    fmt.Printf("%s\n", data)
    //server.Run()
}
