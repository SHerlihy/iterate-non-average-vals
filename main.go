package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type ArrangedFrag []int

type UsableFragment struct {
	LeftInvalid  int
	RightInvalid int
	Frag         ArrangedFrag
}

type BinaryToFragments map[int][]UsableFragment

type SizedUsedIndicies map[int][]int

func RearrangeArray(nums []int) []int {
	var binToFrags BinaryToFragments
	var sizeToIndicies SizedUsedIndicies

	maxSize := len(nums)
	rand.Seed(time.Now().UnixNano())

	arrangedVals := iterateForArrangedValues(nums, maxSize, sizeToIndicies, binToFrags)

	return arrangedVals
}

func iterateForArrangedValues(nums []int, maxSize int, sizeToIndicies SizedUsedIndicies, binToFrags BinaryToFragments) ArrangedFrag {
	for i := 0; i < len(nums); i++ {
		genFrag, genIndicies := GenerateFragment(nums)

		size := len(genFrag)
		if size < 3 {
			continue
		}

		sizedIndicies, ok := sizeToIndicies[size]

		if !ok {
			sizeToIndicies[size] = []int{genIndicies}
		} else {
			sizeToIndicies[size] = append(sizedIndicies, genIndicies)
		}

		usableFrag := generateUsableFragment(genFrag)

		fragments, ok := binToFrags[genIndicies]

		if !ok {
			binToFrags[genIndicies] = []UsableFragment{usableFrag}
		} else {
			binToFrags[genIndicies] = append(fragments, usableFrag)
		}

	}

	matchedFragment := matchTwoParts(maxSize, sizeToIndicies, binToFrags)

	if len(matchedFragment) < len(nums) {
		iterateForArrangedValues(nums, maxSize, sizeToIndicies, binToFrags)
	}

	return matchedFragment
}

func matchTwoParts(maxSize int, sizeToIndicies SizedUsedIndicies, binToFrags BinaryToFragments) ArrangedFrag {
	matchedFragment := []int{}

	for i := maxSize; i > maxSize/2; i-- {
		mostIndicies, ok := sizeToIndicies[i]

		if !ok {
			continue
		}

		fewerIndicies, ok := sizeToIndicies[maxSize-i]

		if !ok {
			continue
		}

		arrangedVals := CalculateCompletePair(maxSize, binToFrags, mostIndicies, fewerIndicies)

		if len(arrangedVals) == maxSize {
			return arrangedVals
		}
	}

	return matchedFragment
}

func CalculateCompletePair(maxSize int, binToFrags BinaryToFragments, indicies1 []int, indicies2 []int) ArrangedFrag {
	matchingPairs := FindMatchingPairs(maxSize, indicies1, indicies2)

	for _, pair := range matchingPairs {
		lIndicies := pair[0]
		sIndicies := pair[1]

		longFrags, ok := binToFrags[lIndicies]

		if !ok {
			continue
		}

		shortFrags, ok := binToFrags[sIndicies]

		if !ok {
			continue
		}

		for _, longFrag := range longFrags {
			for _, shortFrag := range shortFrags {
				stitchFragment := stitchFragments(longFrag, shortFrag)

				if len(stitchFragment) == maxSize {
					return stitchFragment
				}
			}
		}
	}

	return []int{}
}

func stitchFragments(uFrag1 UsableFragment, uFrag2 UsableFragment) ArrangedFrag {
	frag1 := uFrag1.Frag
	frag2 := uFrag2.Frag

	leftInvalid1 := uFrag1.LeftInvalid
	leftInvalid2 := uFrag2.LeftInvalid

	rightInvalid1 := uFrag1.RightInvalid
	rightInvalid2 := uFrag2.RightInvalid

	if frag1[0] == rightInvalid2 && frag2[0] == rightInvalid1 {
		return frag1
	}

	f1End := len(frag1) - 1
	f2End := len(frag2) - 1

	if frag1[f1End] == leftInvalid2 && frag2[f2End] == leftInvalid1 {
		return frag1
	}

	return append(frag1[:], frag2...)
}

func generateUsableFragment(fragment ArrangedFrag) UsableFragment {
	size := len(fragment)
	leftInvalid := (fragment[0] * 2) - fragment[1]
	rightInvalid := (fragment[size-1] * 2) - fragment[size-2]

	usableFrag := UsableFragment{leftInvalid, rightInvalid, fragment}

	return usableFrag
}

func GenerateFragment(nums []int) (ArrangedFrag, int) {
	numsLen := len(nums)

	arrangedFrag := []int{}
	accessedIndicies := 0

    binSlice := make([]bool, numsLen)

	for i, _ := range nums {
		randIdx := rand.Intn(numsLen)

        for {
            if binSlice[randIdx] == false {
                break
            }
            randIdx = rand.Intn(numsLen)
        }


		if i < 2 {
			binSlice[randIdx] = true
			arrangedFrag = append(arrangedFrag, nums[randIdx])
            continue
		}

		var prev int
		var sub int
		var next int

		prev = arrangedFrag[i-2]
		sub = arrangedFrag[i-1]
		next = nums[randIdx]

		avg := (prev + next) / 2

		if sub == avg {
			break
		}

		binSlice[randIdx] = true
		arrangedFrag = append(arrangedFrag, nums[randIdx])
	}

	accessedIndicies = BinSliceToInt(binSlice)

	return arrangedFrag, accessedIndicies
}

func FindMatchingPairs(maxSize int, indicies1 []int, indicies2 []int) [][2]int {
	binSlice := make([]bool, maxSize)
	for i := range binSlice {
		binSlice[i] = true
	}

	maxInt := BinSliceToInt(binSlice)

	completeMatch := int(maxInt)

	var matchingIndicies [][2]int

	for _, lIndicies := range indicies1 {
		for _, sIndicies := range indicies2 {
			lsIndicies := lIndicies ^ sIndicies
			remainingIndicies := completeMatch & lsIndicies

			if remainingIndicies != completeMatch {
				continue
			}

			matchingIndicies = append(matchingIndicies, [2]int{lIndicies, sIndicies})
		}
	}

	return matchingIndicies
}

func BinSliceToInt(binSlice []bool) int {
	ret := 0
	finalIdx := len(binSlice) - 1

	for i, bin := range binSlice {
		if bin == false {
			continue
		}

		binIdx := float64(finalIdx - i)
		ret = ret + int(math.Pow(2, binIdx))
	}

	return ret
}

func IntToBinSlice(num int) []bool {
	if num == 0 {
		return []bool{false}
	}

	fltN := float64(num)
	exponent := int(math.Floor(math.Log(fltN) / math.Log(2)))
	fmt.Print(exponent)

	binSlice := make([]bool, exponent+1)

	for i := exponent; i != 0; i-- {
		deduct := math.Pow(float64(2), float64(i))

		remainder := int(fltN - deduct)

		if remainder >= 0 {
			fltN = float64(remainder)
			binSlice[exponent-i] = true
		}
	}

	deduct := float64(1)

	remainder := int(fltN - deduct)

	if remainder >= 0 {
		fltN = float64(remainder)
		binSlice[exponent] = true
	}

	return binSlice
}
