package main

import (
	"fmt"
	"testing"
)

func TestBinaryConversions(t *testing.T) {
	smallNums := []int{0, 1, 2, 3}
	smallBins := [][]bool{
		[]bool{false},
		[]bool{true},
		[]bool{true, false},
		[]bool{true, true},
	}

	bigNums := []int{8, 12, 17}
	bigBins := [][]bool{
		[]bool{true, false, false, false},
		[]bool{true, true, false, false},
		[]bool{true, false, false, false, true},
	}

	testNums(t, smallNums, smallBins)
	testNums(t, bigNums, bigBins)

	testBins(t, smallBins, smallNums)
	testBins(t, bigBins, bigNums)

}

func testNums(t *testing.T, nums []int, binSlices [][]bool) {
	for ni, num := range nums {
		title := fmt.Sprintf("Num to Bin Slice: %v", num)
		t.Run(title, func(t *testing.T) {
			binSliceExpect := binSlices[ni]
			binSliceRet := IntToBinSlice(num)

			for i, binRet := range binSliceRet {
				if binSliceExpect[i] != binRet {
					t.Errorf("\nBin slice of %v is unexpected. Expected: %v , Returned: %v", binSliceRet, binSliceExpect, binSliceRet)
					break
				}
			}
		})
	}
}

func testBins(t *testing.T, binSlices [][]bool, nums []int) {
	for bi, binSlice := range binSlices {
		title := fmt.Sprintf("Bin Slice to Num: %v", binSlice)
		t.Run(title, func(t *testing.T) {
			numExpected := nums[bi]
			numRet := BinSliceToInt(binSlice)

			if numRet != numExpected {
				t.Errorf("\nNum of %v is unexpected. Expected: %v , Returned: %v", numRet, numExpected, numRet)
			}
		})

	}
}
