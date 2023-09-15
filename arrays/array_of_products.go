package arrays

func ArrayOfProducts(array []int) []int {

	if array == nil { // edge case
		var emptyResult []int
		return emptyResult
	}

	track := NewProductTrack()

	for i := 0; i < len(array); i++ {
		if array[i] == 0 {
			track.incrementZeros()

			if track.allProductsWillBeZero() {
				break //don't need to loop anymore since all possible products will be 0
			}

			track.replaceMinZeroPosition(i)
			track.replaceMaxZeroPosition(i)
			continue // go to next since without need to update product
		}

		track.multiplyProductByNonZero(array[i])
	}

	result := make([]int, len(array))

	for i := 0; i < len(array); i++ {
		if track.allProductsWillBeZero() {
			// just fill everything with zeros
			for i < len(array) {
				result[i] = 0
				i++
			}
			break
		}

		if track.hasZerosAtTheSides(i) {
			result[i] = 0 // this position product will be zero since has zero as element
			continue
		}

		if array[i] == 0 {
			result[i] = track.product // will be the full product (zeroless) itself
			continue
		}

		result[i] = track.product / array[i] // will be the product of all non zero elements excluding the current one
	}

	return result
}

const positionInitialValue = -999999

type ProductTrack struct {
	product, zeros, minZeroPosition, maxZeroPosition int
}

func NewProductTrack() *ProductTrack {
	return &ProductTrack{
		product:         1,
		minZeroPosition: positionInitialValue,
		maxZeroPosition: positionInitialValue,
	}
}

func (p *ProductTrack) incrementZeros() {
	p.zeros++
}

func (p *ProductTrack) allProductsWillBeZero() bool {
	return p.zeros > 1
}

func (p *ProductTrack) replaceMinZeroPosition(i int) {
	if positionInitialValue == p.minZeroPosition || i < p.minZeroPosition {
		p.minZeroPosition = i
	}
}

func (p *ProductTrack) replaceMaxZeroPosition(i int) {
	if positionInitialValue == p.maxZeroPosition || i > p.maxZeroPosition {
		p.maxZeroPosition = i
	}
}

func (p *ProductTrack) multiplyProductByNonZero(value int) {
	if value == 0 {
		return
	}
	p.product = p.product * value
}

func (p *ProductTrack) hasZerosAtTheSides(i int) bool {
	return (positionInitialValue != p.minZeroPosition && p.minZeroPosition < i) ||
		(positionInitialValue != p.maxZeroPosition && p.maxZeroPosition > i)
}
