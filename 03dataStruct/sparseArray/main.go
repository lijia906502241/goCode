package main

import "fmt"

/*
	稀疏数组
*/
type ValNode struct {
	row int
	col int
	val int
}

func main() {
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
	//转成稀疏数组
	var sparseArr []ValNode
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	//输出稀疏数组,或将此数组 存盘
	fmt.Println("当前的稀疏数组是...")
	for i, valNode := range sparseArr {
		fmt.Printf("%d:%d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}
	//恢复原始的数组
	var chessMap2 [11][11]int
	for i, valNode := range sparseArr {
		if i != 0 {
			chessMap2[valNode.row][valNode.col] = valNode.val
		}
	}
	fmt.Println("h恢复后的数据")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
