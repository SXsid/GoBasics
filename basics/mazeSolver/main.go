package main

import (
	"errors"
	"fmt"
)

type point struct {
	x int
	y int
}

// we can walk or not true if each dead end
// recusrion function
func walk(maze []string, currPoint point, target string, wall string, path []point, seen [][]bool) (newPath []point, found bool) {
	//base cases

	//1) out of bound
	if currPoint.x < 0 || currPoint.x >= len(maze) || currPoint.y < 0 || currPoint.y >= len(maze[0]) {
		return path, false
	}

	//2 if alredy seen or wall
	x, y := currPoint.x, currPoint.y

	if seen[x][y] || string(maze[x][y]) == wall {
		return path, false
	}

	// 3) if we found the target
	if string(maze[x][y]) == target {
		newPath = append(path, currPoint)
		return newPath, true
	}

	// mark as seen
	seen[x][y] = true

	// add current point to path
	newPath = append(path, currPoint)

	// try all directions
	dirs := []point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range dirs {
		next := point{x + dir.x, y + dir.y}
		if newPath, ok := walk(maze, next, target, wall, newPath, seen); ok {
			//it will kill the recusion
			return newPath, true
		}
	}

	// if no direction works, remove the current point
	if len(newPath) > 0 {
		newPath = newPath[:len(newPath)-1]
	}

	//if all the parth formt the stating have been
	return newPath, false
}
func solve(maze []string, startPoint point, target string, wall string) (path []point, err error) {
	path = make([]point, 0)
	seen := make([][]bool, len(maze))
	for i, _ := range seen {
		seen[i] = make([]bool, len(maze[0]))
	}

	newPath, ok := walk(maze, startPoint, target, wall, path, seen)
	if !ok {
		return []point{}, errors.New("can't find teh ending")
	}
	return newPath, nil
}

func main() {
	maze := []string{
		"S#.....#",
		".#.###.#",
		"...#....",
		".#####.#",
		".....E.#",
	}

	startPoint := point{
		x: 0,
		y: 0,
	}

	target := "E"
	wall := "#"

	path, err := solve(maze, startPoint, target, wall)
	if err != nil {
		fmt.Println("cant' solve the maze or no exit point avaiblabe")
	}

	for i, point := range path {
		fmt.Printf("%d => (%v),\n", i+1, point)
	}

}
