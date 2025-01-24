package main

import "fmt"

var maze = []string{
	"xxxxxxxxxx x",
	"x        x x",
	"x        x x",
	"x xxxxxxxx x",
	"x          x",
	"x xxxxxxxxxx",
}

type Point struct {
	x, y int
}

var DIR = []Point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func isEqual(p1, p2 Point) bool {

	return p1.x == p2.x && p1.y == p2.y

}

func walk(maze []string, wall string, curr, end Point, seen [][]bool, path *[]Point) bool {
	if curr.x < 0 || curr.x >= len(maze[0]) || curr.y < 0 || curr.y >= len(maze) {
		return false
	}

	if string(maze[curr.y][curr.x]) == wall {
		return false
	}

	if isEqual(curr, end) {
		*path = append(*path, end)
		return true
	}

	if seen[curr.y][curr.x] {
		return false
	}
	seen[curr.y][curr.x] = true
	*path = append(*path, curr)

	for i := 0; i < 4; i++ {
		tmp := DIR[i]
		if walk(maze, wall, Point{x: curr.x + tmp.x, y: curr.y + tmp.y}, end, seen, path) {
			return true
		}

	}
	*path = (*path)[:len(*path)-1]

	return false

}

func solve(maze []string, wall string, start, end Point) []Point {
	seen := make([][]bool, len(maze), len(maze))
	path := make([]Point, 0, len(maze))

	for i := range maze {
		seen[i] = make([]bool, len(maze[i]))
	}
	walk(maze, wall, start, end, seen, &path)

	return path
}

func main() {
	fmt.Println(solve(maze, "x", Point{x: 10, y: 0}, Point{x: 1, y: 5}))

}
