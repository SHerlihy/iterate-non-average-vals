package main

import (
	"fmt"
	"testing"
)

type MatchingPairsTestData struct {
    Size int
    Expects [][2]int
    LongIndicies []int
    ShortIndicies []int
}
type CalculateCompletePairData struct {
    Size int
    Expects ArrangedFrag
    BinToFrags BinaryToFragments
    LongIndicies []int
    ShortIndicies []int
}

func TestRearrangeArray(t *testing.T){
    averageSlice := []int{1,2,3,4,5,6}
    nonAverageSlice := RearrangeArray(averageSlice)

    t.Log(nonAverageSlice)

    if len(nonAverageSlice) != len(averageSlice){
        t.Errorf("Length incorrect: %v", nonAverageSlice)
    }
}

func TestGenerateFragment(t *testing.T){
    averageSlice := []int{1,2,3,4,5,6}
    frag, indicies := GenerateFragment(averageSlice)

    t.Logf("\nfrag: %v", frag)
    t.Logf("\nindicies: %v", indicies)
}

func TestCalculateCompletePair(t *testing.T){
    ogFrag := [6]int{1,2,3,4,5,6}

    lIndicies := []int{45}
    sIndicies := []int{18}

    lUsableFrag := UsableFragment{-1,8,ArrangedFrag{1,3,4,6}}
    sUsableFrag := UsableFragment{-1,8,ArrangedFrag{2,5}}

    lFrags := []UsableFragment{
        lUsableFrag,
    }
    sFrags := []UsableFragment{
        sUsableFrag,
    }

    binToFrags := make(BinaryToFragments, 2)

    binToFrags[lIndicies[0]] = lFrags
    binToFrags[sIndicies[0]] = sFrags

    completePair := CalculateCompletePair(6, binToFrags, lIndicies, sIndicies)

    if len(completePair) != len(ogFrag) {
        t.Errorf("\nReturned len: %v, Expected len: %v", len(completePair), len(ogFrag))
    }

    for _, val := range completePair {
        inExpected := false
        for _, expVal := range ogFrag {
            if expVal == val {
                inExpected = true
                break
            }
        }
        if inExpected == false {
            t.Errorf("\nVal in returned slice not in expected: %v", completePair)
        }
    }
}

func TestFindMatchingPairs(t *testing.T){
    sizeThreePairs := [][2]int{
        [2]int{6,1},
    }

    sizeFourPairs := [][2]int{
        [2]int{9,6},
        [2]int{7,8},
    }

    succeedsData := []MatchingPairsTestData{
        MatchingPairsTestData{3, sizeThreePairs, []int{6}, []int{1}},
        MatchingPairsTestData{4, sizeFourPairs, []int{9,7}, []int{6,8}},
    }

    for _,data := range succeedsData {
    expects := data.Expects
    matches := FindMatchingPairs(data.Size, data.LongIndicies, data.ShortIndicies)

    for _,match := range matches {
        inExpects := false
        for _,expect := range expects{
            if match == expect{
                inExpects = true
                break
            }
        }
        
        if !inExpects {
            t.Errorf("Error on FindMatchingPairs, returned: %v", matches)
        }
    }
    }
}

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
