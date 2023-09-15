package greedy

import (
	"slices"
)

func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	slices.Sort(redShirtHeights)
	slices.Sort(blueShirtHeights)

	i := 0
	nthRedShirt := redShirtHeights[i]
	nthBlueShirt := blueShirtHeights[i]

	nthStudentSameHeight := nthRedShirt == nthBlueShirt
	if nthStudentSameHeight {
		return false
	}

	redRowShouldBeInFront := nthRedShirt < nthBlueShirt
	blueRowShouldBeInFront := !redRowShouldBeInFront

	for i := 1; i < len(redShirtHeights); i++ {
		nthRedShirt = redShirtHeights[i]
		nthBlueShirt = blueShirtHeights[i]

		nthBlueShirtIsInFront := nthRedShirt >= nthBlueShirt
		if redRowShouldBeInFront && nthBlueShirtIsInFront {
			return false
		}

		nthRedShirtIsInFront := nthBlueShirt >= nthRedShirt
		if blueRowShouldBeInFront && nthRedShirtIsInFront {
			return false
		}
	}

	return true
}
