package main

func orangesRotting(grid [][]int) int {
	queue := [][2]int{}
	n, m := len(grid), len(grid[0])
	directions := [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	x, y := 0, 0
	flag, step := false, 0
	fresh := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}
	for len(queue) > 0 {
		flag = false
		for _, temp := range queue {
			queue = queue[1:]
			for _, dir := range directions {
				x = temp[0] + dir[0]
				y = temp[1] + dir[1]
				if x >= 0 && x < n && y >= 0 && y < m && grid[x][y] == 1 {
					grid[x][y] = 2
					queue = append(queue, [2]int{x, y})
					fresh--
					flag = true
				}
			}
			//fmt.Println(grid)
		}
		if flag {
			step++
		}
	}
	if fresh == 0 {
		return step
	}
	return -1
}
