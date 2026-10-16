package main

import "fmt"

/*
	Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle.

	In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

	Example 1:
	Input: rowIndex = 3
	Output: [1,3,3,1]

	Example 2:
	Input: rowIndex = 0
	Output: [1]
	
	Example 3:
	Input: rowIndex = 1
	Output: [1,1]	

	Constraints:

		0 <= rowIndex <= 33

*/


func getRow(rowIndex int) []int {

    matris := make([][]int, rowIndex + 1)

	matris[0] = append(matris[0], 1)

	if rowIndex == 0 {
		return matris[0]
	}

	matris[1] = append(matris[1], 1, 1)

	if rowIndex == 1 {
		return matris[1]
	}	

	for i := 2; i <= rowIndex; i++ {
		
		matris[i] = make([]int, i + 1)
		
		for j := 1; j < i; j++ {
			
			matris[i][0] = 1
			matris[i][i] = 1

			matris[i][j] = matris[i-1][j-1] + matris[i-1][j]

		}
	}

	return matris[rowIndex]
}

func main() {

	matris := [][]int{
		    {1},
		   {1,1},
		  {1,2,1},
		 {1,3,3,1},
		{1,4,6,4,1},
	}
	
	fmt.Println("expected :")
	fmt.Println(matris)
	fmt.Println("result :")
	fmt.Println(getRow(0))
}