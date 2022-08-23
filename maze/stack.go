package maze

import (
    "errors"
    "sync"
)

type Stack struct {
    lock sync.Mutex
    cell []Cell
}

func NewStack() *Stack {
    return &Stack{ sync.Mutex{}, make([]Cell, 0) }
}

func (s *Stack) Push(v *Cell) {
    s.lock.Lock()
    defer s.lock.Unlock()

    s.cell = append(s.cell, *v)
}

func (s *Stack) Pop() (*Cell, error) {
    s.lock.Lock()
    defer s.lock.Unlock()


    l := len(s.cell)
    if l == 0 {
        return nil, errors.New("Empty Stack")
    }

    res := s.cell[l-1]
    s.cell = s.cell[:l-1]
    return &res, nil
}
