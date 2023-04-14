package service

import (
	"math"
	"strconv"

	"vincs/models"
)

func FindMinDisparity(input models.Input) map[string]int64 {
	var result = make(map[string]int64)

	var T = input.TotalTestCases
	//calculate minimum possible disparity for each test case
	for i := 1; i <= T; i++ {
		var N int
		var S []int
		testCase := testCaseNumber(i, input)
		N = testCase.TotalCars
		S = testCase.SpeedOfCars

		var greatestNumber int = 0

		for i := 0; i < N; i++ {
			if S[i] > greatestNumber {
				greatestNumber = S[i]
			}
		}

		//initialising minimum number with the biggest array element and maximum number with zero
		var minS, maxS int = greatestNumber, 0

		//initialising dp array
		dp := make([][][]int64, N+1)
		for n := 0; n < N+1; n++ {
			dp[n] = make([][]int64, greatestNumber+1)
			for m := 0; m < greatestNumber+1; m++ {
				dp[n][m] = make([]int64, greatestNumber+1)
			}
		}

		//compute minimum total disparity
		result[strconv.Itoa(i)] = solve(N, S, minS, maxS, dp)
	}
	return result
}

// helper private functions
func solve(n int, S []int, minS int, maxS int, dp [][][]int64) int64 {
	if n == 0 {
		return 0
	}
	if dp[n][minS][maxS] != 0 {
		return dp[n][minS][maxS]
	}

	var totalDisparity int64 = math.MaxInt64 //initializing disparity

	for i := 0; i < n; i++ {
		ele := S[i]
		minimumSpeed := minS
		maximumSpeed := maxS

		if ele < minimumSpeed {
			minimumSpeed = ele
		}
		if ele > maximumSpeed {
			maximumSpeed = ele
		}

		S = deleteAtIndex(S, i)
		d := int64(maximumSpeed-minimumSpeed) + solve(n-1, S, minimumSpeed, maximumSpeed, dp)
		if d < totalDisparity {
			totalDisparity = d
		}
		S = insertAtIndex(S, i, ele) //backtrack
	}

	dp[n][minS][maxS] = totalDisparity
	return totalDisparity
}

func deleteAtIndex(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func insertAtIndex(arr []int, index int, value int) []int {
	arr = append(arr[:index+1], arr[index:]...)
	arr[index] = value
	return arr
}

func testCaseNumber(i int, input models.Input) models.TestData {
	var testcase models.TestData
	switch i {
	case 1:
		testcase = input.TestCase1
		break
	case 2:
		testcase = input.TestCase2
		break
	case 3:
		testcase = input.TestCase3
		break
	case 4:
		testcase = input.TestCase4
		break
	case 5:
		testcase = input.TestCase5
		break
	case 6:
		testcase = input.TestCase6
		break
	case 7:
		testcase = input.TestCase7
		break
	case 8:
		testcase = input.TestCase8
		break
	case 9:
		testcase = input.TestCase9
		break
	case 10:
		testcase = input.TestCase10
		break
	}
	return testcase
}
