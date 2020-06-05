package main


// 顺时针打印二维数组


func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	tR := 0
	tC := 0
	dR := len(matrix) - 1
	dC := len(matrix[0]) - 1
	for tR <= dR && tC <= dC {
		printEdge(matrix, tR, tC, dR, dC, res)
	}
	return res
}



func printEdge(matrix [][]int, tR, tC, dR, dC int, res []int) {

	if tR == dR {
		// 表示子矩阵只有一行
		for i := tC; i <= dC; i++ {
			res = append(res, matrix[tR][i])
		}

	} else if tC == dC {
		// 表示子矩阵只有一列
		for i := tR; i <= dR; i++ {
			res = append(res, matrix[i][tC])
		}
	} else {
		// 一般情况
		curC := tC
		curR := tR

		for curC != dC {
			res = append(res, matrix[tR][curC])
			curC++
		}
		for curR != dR {
			res = append(res, matrix[curR][dC])
			curR++
		}
		for curC != tC {
			res = append(res, matrix[dR][curC])
			curC--
		}
		for curR != tR {
			res = append(res, matrix[curR][tC])
			curR--
		}

	}
}