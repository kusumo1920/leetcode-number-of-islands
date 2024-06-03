package main

import "fmt"

func main() {
	input := [][]byte{
		{'1', '1', '1', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(input)
	output := numIslands(input)
	fmt.Println(output)
}

func numIslands(grid [][]byte) int {
	if grid == nil {
		return 0
	}

	nr := len(grid)
	nc := len(grid[0])
	result := 0

	var dfs func([][]byte, int, int)
	dfs = func(grid [][]byte, r, c int) {
		nr := len(grid)
		nc := len(grid[0])

		if r < 0 || c < 0 || r >= nr || c >= nc || grid[r][c] == '0' {
			return
		}

		grid[r][c] = '0'
		dfs(grid, r-1, c)
		dfs(grid, r+1, c)
		dfs(grid, r, c-1)
		dfs(grid, r, c+1)
	}

	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				result++
				dfs(grid, r, c)
			}
		}
	}

	return result
}
