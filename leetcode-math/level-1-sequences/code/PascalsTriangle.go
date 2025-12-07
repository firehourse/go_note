package code

/*
	# Pascal's Triangle

Given an integer numRows, return the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:
Example 1:

Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
Example 2:

Input: numRows = 1
Output: [[1]]
*/
func Generate(numRows int) [][]int {
	// 宣告一個二維陣列 視作x跟y軸
	result := make([][]int, numRows)
	result[0] = []int{1}
	if numRows == 1 {
		return result
	}
	result[1] = []int{1, 1}
	if numRows == 2 {
		return result
	}
	for i := 2; i <= numRows-1; i++ {
		// 直接初始化陣列
		result[i] = make([]int, i+1)
		// 內部循環
		for j := 0; j <= i; j++ {
			// 先處理頭尾
			if j == 0 {
				result[i][j] = 1
				continue
			}
			if j == i {
				result[i][j] = 1
				continue
			}
			// 3,1 = 2,0 + 2,1
			result[i][j] = result[i-1][j] + result[i-1][j-1]
		}
	}
	return result
}
