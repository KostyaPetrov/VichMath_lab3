package main

import (
	"math"
)

func firstEquation(x float64) float64 {

	return (-1)*(math.Pow(x, 3)) - math.Pow(x, 2) + x + 3
}

func secondEquation(x float64) float64 {
	return 4*math.Pow(x, 3) - 5*math.Pow(x, 2) + 6*x - 7
}

func thirdEquation(x float64) float64 {
	return -2*math.Pow(x, 3) - 5*math.Pow(x, 2) + 7*x - 13
}

func f(x float64, number int) float64 {
	if number == 1 {
		return firstEquation(x)
	} else if number == 2 {
		return secondEquation(x)
	} else {
		return thirdEquation(x)
	}
}

func LeftRectangles(lowBorder float64, heightBorder float64, numParts int, numEquation int, acc float64) (float64, int) {

	sumLess := 0.0
	sumMore := 0.0
	j := 0
	for j != 1 {

		step := (heightBorder - lowBorder) / float64(numParts)

		for i := 0; i < numParts; i++ {
			sumLess += step * f(lowBorder+float64(i)*step, numEquation)
		}

		step = (heightBorder - lowBorder) / float64(numParts*2)
		for i := 0; i < numParts*2; i++ {
			sumMore += step * f(lowBorder+float64(i)*step, numEquation)
		}

		tochnost := math.Abs(sumLess - sumMore)
		if tochnost < acc {
			j = 1

		} else {
			numParts *= 2
			sumLess = 0.0
			sumMore = 0.0
		}

	}

	return sumMore, numParts
}

func RightRectangles(lowBorder float64, heightBorder float64, numParts int, numEquation int, acc float64) (float64, int) {
	sumLess := 0.0
	sumMore := 0.0
	j := 0
	for j != 1 {

		step := (heightBorder - lowBorder) / float64(numParts)

		for i := 1; i < numParts+1; i++ {
			sumLess += step * f(lowBorder+float64(i)*step, numEquation)
		}

		step = (heightBorder - lowBorder) / float64(numParts*2)
		for i := 1; i < numParts*2+1; i++ {
			sumMore += step * f(lowBorder+float64(i)*step, numEquation)
		}

		tochnost := math.Abs(sumLess - sumMore)
		if tochnost < acc {
			j = 1

		} else {
			numParts *= 2
			sumLess = 0.0
			sumMore = 0.0
		}

	}

	return sumMore, numParts
}

func MiddleRectangles(lowBorder float64, heightBorder float64, numParts int, numEquation int, acc float64) (float64, int) {
	sumLess := 0.0
	sumMore := 0.0
	j := 0
	for j != 1 {

		step := (heightBorder - lowBorder) / float64(numParts)

		for i := 0.5; i < float64(numParts)+0.5; i++ {
			sumLess += step * f(lowBorder+float64(i)*step, numEquation)
		}

		step = (heightBorder - lowBorder) / float64(numParts*2)
		for i := 0.5; i < float64(numParts)*2+0.5; i++ {
			sumMore += step * f(lowBorder+float64(i)*step, numEquation)
		}

		tochnost := math.Abs(sumLess - sumMore)
		if tochnost < acc {
			j = 1

		} else {
			numParts *= 2
			sumLess = 0.0
			sumMore = 0.0
		}

	}

	return sumMore, numParts
}

func Trapeze(lowBorder float64, heightBorder float64, numParts int, numEquation int, acc float64) (float64, int) {
	sumLess := f(lowBorder, numEquation) + f(heightBorder, numEquation)
	sumMore := f(lowBorder, numEquation) + f(heightBorder, numEquation)
	j := 0
	for j != 1 {

		step := (heightBorder - lowBorder) / float64(numParts)

		for i := 1; i < numParts; i++ {
			sumLess += 2 * f(lowBorder+float64(i)*step, numEquation)
		}
		sumLess *= step / 2

		step = (heightBorder - lowBorder) / float64(numParts*2)
		for i := 1; i < numParts*2; i++ {
			sumMore += 2 * f(lowBorder+float64(i)*step, numEquation)
		}
		sumMore *= step / 2

		tochnost := math.Abs(sumLess - sumMore)
		if tochnost < acc {
			j = 1

		} else {
			numParts *= 2
			sumLess = f(lowBorder, numEquation) + f(heightBorder, numEquation)
			sumMore = sumLess
		}

	}

	return sumMore, numParts
}

func Simpson(lowBorder float64, heightBorder float64, numParts int, numEquation int, acc float64) (float64, int) {
	sumLess := f(lowBorder, numEquation) + f(heightBorder, numEquation)
	sumMore := sumLess
	j := 0
	k := 0.0
	for j != 1 {

		step := (heightBorder - lowBorder) / float64(numParts)

		for i := 1; i < numParts; i++ {
			k = float64(2 + 2*(i%2))
			sumLess += k * f(lowBorder+float64(i)*step, numEquation)
		}
		sumLess *= step / 3

		step = (heightBorder - lowBorder) / float64(numParts*2)
		for i := 1; i < numParts*2; i++ {
			k = float64(2 + 2*(i%2))
			sumMore += k * f(lowBorder+float64(i)*step, numEquation)
		}
		sumMore *= step / 3

		tochnost := math.Abs(sumLess - sumMore)
		if tochnost < acc {
			j = 1

		} else {
			numParts *= 2
			sumLess = f(lowBorder, numEquation) + f(heightBorder, numEquation)
			sumMore = sumLess
		}

	}

	return sumMore, numParts
}
