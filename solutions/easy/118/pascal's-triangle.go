package main

import "fmt"

/*
	Given an integer numRows, return the first numRows of Pascal's triangle.

	In Pascal's triangle, each number is the sum of the two numbers
	directly above it as shown:

	Example 1:
	Input: numRows = 5
	Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

	Example 2:
	Input: numRows = 1
	Output: [[1]]

	Constraints:
		1 <= numRows <= 30
*/

func generate(numRows int) [][]int {

	matris := make([][]int, numRows)
	matris[0] = append(matris[0], 1)

	if numRows == 1 {
		return matris[:1]
	}

	matris[1] = append(matris[1], 1, 1)		

	for i := 2; i < numRows; i++ {
		
		matris[i] = make([]int, i + 1)
		
		for j := 1; j < i; j++ {
			
			matris[i][0] = 1
			matris[i][i] = 1

			matris[i][j] = matris[i-1][j-1] + matris[i-1][j]

		}
	}

	return matris
}

func main()  {
	
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
	fmt.Println(generate(1))

}