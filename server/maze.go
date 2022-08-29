package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/LukeOrth/go-maze/maze"
	"github.com/gorilla/mux"
)

func JSONError(w http.ResponseWriter, err interface{}, code int) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.Header().Set("X-Content-Type-Options", "nosniff")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(err)
}

func validate(r *http.Request, params ...string) []string {
    errs := []string{}

    for _, param := range params {
       if !r.URL.Query().Has(param) {
           errs = append(errs, fmt.Sprintf("missing parameter: %s", param))
       }
    }
    return errs
}

func generateMazeGET(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    q := r.URL.Query()
    columns, _ := strconv.Atoi(q.Get("columns"))
    rows, _ := strconv.Atoi(q.Get("rows"))
    scale, _ := strconv.Atoi(q.Get("scale"))
    //missingParams := make([]string, 3)
    if q.Has("columns") {
        fmt.Println("HI MOM")
    }
    if columns == 0 {
        JSONError(w, "missing parameter: columns", 400)
        return
    }
    maze := maze.NewMaze(columns, rows, scale)
    json.NewEncoder(w).Encode(maze)
}

func mazeToPngPOST(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	var maze maze.Maze
    err := json.NewDecoder(r.Body).Decode(&maze)
    if err != nil {
        log.Panic(err)
    }
    maze.Png(w)
}

func mazeToSvgPOST(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "image/svg+xml")
    var maze maze.Maze
    err := json.NewDecoder(r.Body).Decode(&maze)
    if err != nil {
        log.Panic(err)
    }
    maze.Svg(w)
}

func Run() {
    r := mux.NewRouter()

    r.HandleFunc("/maze", generateMazeGET).Methods("GET")
    r.HandleFunc("/maze/png", mazeToPngPOST).Methods("POST")
    r.HandleFunc("/maze/svg", mazeToSvgPOST).Methods("POST")

    log.Fatal(http.ListenAndServe(":8000", r))
}
