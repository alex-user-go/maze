package maze

import (
	"errors"
)

type Maze struct {
	matrix  [][]int
	visited map[coordinate]bool
	allowed []int
}

func NewMaze(matrix [][]int, allowed []int) (*Maze, error) {
	if ok := matrixValid(matrix); !ok {
		return nil, ErrInvalidMatrix
	}
	mz := Maze{
		matrix:  matrix,
		allowed: allowed,
		visited: map[coordinate]bool{},
	}
	return &mz, nil
}

func matrixValid(m [][]int) bool {
	if len(m) == 0 {
		return false
	}
	for _, v := range m {
		if len(v) != len(m[0]) {
			return false
		}
	}
	return true
}

var (
	ErrInvalidInput  = errors.New("starting point isn't valid")
	ErrNoExit        = errors.New("can't find any exit")
	ErrInvalidMatrix = errors.New("matrix should be rectangular")
)

type vertex struct {
	coordinate
	distance int
}

type coordinate struct {
	x, y int
}

func (mz *Maze) isOpen(c coordinate) bool {
	return contains(mz.allowed, mz.matrix[c.y][c.x])
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (mz *Maze) isValid(c coordinate) bool {
	//check if valid
	if exist := c.x >= 0 && c.y >= 0 && c.x < len(mz.matrix[0]) && c.y < len(mz.matrix); !exist {
		return false
	}
	//check if we have been there
	if _, beenThere := mz.visited[c]; beenThere {
		return false
	}
	//check if it is open to go
	return mz.isOpen(c)
}

func (mz *Maze) isEnd(c coordinate) bool {
	return (c.x == 0 || c.y == 0 || c.x == len(mz.matrix[0])-1 || c.y == len(mz.matrix)-1)
}

func (mz Maze) getNeighbours(v vertex) []vertex {
	neighbours := make([]vertex, 0, 4)
	//to the left
	left := coordinate{
		x: v.x - 1,
		y: v.y,
	}
	if mz.isValid(left) {
		mz.visited[left] = true
		neighbours = append(neighbours, vertex{coordinate: left, distance: v.distance + 1})
	}
	//to the right
	right := coordinate{
		x: v.x + 1,
		y: v.y,
	}
	if mz.isValid(right) {
		mz.visited[right] = true
		neighbours = append(neighbours, vertex{coordinate: right, distance: v.distance + 1})
	}
	//top
	top := coordinate{
		x: v.x,
		y: v.y - 1,
	}
	if mz.isValid(top) {
		mz.visited[top] = true
		neighbours = append(neighbours, vertex{coordinate: top, distance: v.distance + 1})
	}
	//bottom
	btm := coordinate{
		x: v.x,
		y: v.y + 1,
	}
	if mz.isValid(btm) {
		mz.visited[btm] = true
		neighbours = append(neighbours, vertex{coordinate: btm, distance: v.distance + 1})
	}
	return neighbours
}

func (mz *Maze) FindLenExit(startx, starty int) (int, error) {
	start := coordinate{
		x: startx - 1,
		y: starty - 1,
	}
	if !mz.isValid(start) {
		return -1, ErrInvalidInput
	}
	//mark as visited
	mz.visited[start] = true
	return mz.findPath(start)

}

func (mz *Maze) findPath(start coordinate) (int, error) {
	pathLen := -1
	startVer := vertex{
		coordinate: start,
		distance:   0,
	}
	queue := []vertex{startVer}
	for len(queue) > 0 {
		current := queue[0]
		if mz.isEnd(current.coordinate) {
			//found it!
			pathLen = current.distance
			break
		}
		//get valid
		neighbours := mz.getNeighbours(current)
		queue = append(queue, neighbours...)
		//nowhere to go
		if len(queue) == 1 {
			return 0, ErrNoExit
		}
		//remove the current
		queue = queue[1:]
	}
	return pathLen, nil

}
