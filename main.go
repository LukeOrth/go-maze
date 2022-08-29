package main

import "github.com/LukeOrth/go-maze/server"

func main() {
    /*
    f, _ := os.Open("maze.json")
    defer f.Close()
    
    s := bufio.NewScanner(f)
    
    var b []byte
    for s.Scan() {
        b = s.Bytes()
    }

    var m maze.Maze
    err := json.Unmarshal(b, &m)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Printf("%+v\n", m)
    */

    server.Run()
    
    /*
    f, _ := os.Create("test.png")
    defer f.Close()

    maze := maze.NewMaze(10, 10, 5)
    maze.Png(f)
    */
}
