package server

import (
	"encoding/json"
    "errors"
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

func getParms(r *http.Request, params ...string) ([]interface{}, error) {
    values := make([]interface{}, len(params))
    var errString string

    for _, p := range params {
       if !r.URL.Query().Has(p) {
           errString = errString + "missing parameter: %s"
       }
       values = append(values, r.URL.Query().Get(p))
    }
    if len(errString) > 0 {
        return nil, errors.New(errString)
    }
    return values, nil
}

func generateMazeGET(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    params, err := getParms(r, "columns", "rows", "scale")
    if err != nil {
        JSONError(w, err, 400)
    }
    columns, _ := strconv.Atoi(params[0].(string))
    rows, _ := strconv.Atoi(params[1].(string))
    scale, _ := strconv.Atoi(params[2].(string))
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
