package models

type Input struct {
	TestCase1      TestData `json:"1"`
	TestCase2      TestData `json:"2"`
	TestCase3      TestData `json:"3"`
	TestCase4      TestData `json:"4"`
	TestCase5      TestData `json:"5"`
	TestCase6      TestData `json:"6"`
	TestCase7      TestData `json:"7"`
	TestCase8      TestData `json:"8"`
	TestCase9      TestData `json:"9"`
	TestCase10     TestData `json:"10"`
	TotalTestCases int      `json:"Total Test Cases"`
	//TestCases      map[string]TestData
}

type TestData struct {
	TotalCars   int   `json:"N"`
	SpeedOfCars []int `json:"S"`
}
