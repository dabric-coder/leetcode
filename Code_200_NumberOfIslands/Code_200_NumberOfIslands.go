package main

import "fmt"

func numIslands(grid [][]byte) int {
	if grid == nil || grid[0] == nil {
		return 0
	}
	res := 0
	R := len(grid)
	C := len(grid[0])
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if grid[i][j] == 1 {
				res++
				infect(grid, i, j, R, C)
			}
		}
	}
	return res
}


func infect(grid [][]byte, i, j, R, C int) {
	if i < 0 || i >= R || j < 0 || j >= C || grid[i][j] != 1 {
		return
	}
	grid[i][j] = 2
	infect(grid, i-1, j, R, C)
	infect(grid, i+1, j, R, C)
	infect(grid, i, j-1, R, C)
	infect(grid, i, j+1, R, C)
}


func main() {
	m1 := [][]byte{  { 0, 0, 0, 0, 0, 0, 0, 0, 0 },
		{ 0, 1, 1, 1, 0, 1, 1, 1, 0 },
		{ 0, 1, 1, 1, 0, 0, 0, 1, 0 },
		{ 0, 1, 1, 0, 0, 0, 0, 0, 0 },
		{ 0, 0, 0, 0, 0, 1, 1, 0, 0 },
		{ 0, 0, 0, 0, 1, 1, 1, 0, 0 },
		{ 0, 0, 0, 0, 0, 0, 0, 0, 0 }, }
	fmt.Println(numIslands(m1))

	m2 := [][]byte{  { 1, 1, 1, 1, 0},
		{ 1, 1, 0, 1, 0 },
		{ 1, 1, 0, 0, 0 },
		{ 0, 0, 0, 0, 0 },}

	fmt.Println(numIslands(m2))
}