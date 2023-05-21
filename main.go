package main

import (
	"fmt"
)

func main() {
	i := 0
	repeat := ""
	for i != 1 {

		numEquation := choiceEquation()

		lowLimitIntegration, heightLimitIntegration := choiceLimitIntegration()

		accuracy := choiceAccuracy()

		method := choiceMethod()

		n := 4

		answ, lines := 0.0, 0

		if method == 1 {
			answ, lines = LeftRectangles(float64(lowLimitIntegration), float64(heightLimitIntegration), n, numEquation, accuracy)
		} else if method == 3 {
			answ, lines = RightRectangles(float64(lowLimitIntegration), float64(heightLimitIntegration), n, numEquation, accuracy)
		} else if method == 2 {
			answ, lines = MiddleRectangles(float64(lowLimitIntegration), float64(heightLimitIntegration), n, numEquation, accuracy)
		} else if method == 4 {
			answ, lines = Trapeze(float64(lowLimitIntegration), float64(heightLimitIntegration), n, numEquation, accuracy)
		} else {
			answ, lines = Simpson(float64(lowLimitIntegration), float64(heightLimitIntegration), n, numEquation, accuracy)
		}

		fmt.Print("Ответ: ")
		fmt.Println(answ)

		fmt.Println("Для получения ответа с заданной точностью понадобилось разбить отрезок на ", lines, " части(ей)")

		fmt.Println("Повторим? (y/n)")
		_, _ = fmt.Scan(&repeat)
		if repeat != "y" {
			i = 1
		}
	}
}

func choiceEquation() int {

	fmt.Println("Выберите нужную функцию:\n" +
		"1: -x^3 - x^2 + x + 3\n" +
		"2: 4x^3 - 5x^2 + 6x - 7\n" +
		"3: -2x^3 - 5x^2 +7x -13")

	numEquation := 0
	i := 0
	for i != 1 {
		_, err := fmt.Scanf("%d\n", &numEquation)
		if err != nil {
			fmt.Println("Вы должны ввести цифру 1, 2 или 3")

		} else {

			if numEquation == 1 || numEquation == 2 || numEquation == 3 {
				i = 1
			} else {
				fmt.Println("Вы должны выбрать уравнение и ввести соответственно цифру 1, 2 или 3")
			}
		}
	}

	if numEquation == 1 {
		fmt.Printf("Вы выбрали уравнение -x^3 - x^2 + x + 3\n")
		return 1
	} else if numEquation == 2 {
		fmt.Printf("Вы выбрали уравнение 4x^3 - 5x^2 + 6x - 7\n")
		return 2
	} else {
		fmt.Printf("Вы выбрали уравнение -2x^3 - 5x^2 +7x -13\n")
		return 3
	}

}

func choiceLimitIntegration() (int, int) {

	fmt.Print("Введите пределы интегрирования через пробел: ")

	lowLimit := 0
	heightLimit := 0
	i := 0
	for i != 1 {
		_, err := fmt.Scan(&lowLimit, &heightLimit)
		if err != nil {
			fmt.Println("Вы должны ввести два целых числа в поряден возростпния через пробел")
			fmt.Print("Введите еще раз: ")
		} else {
			if lowLimit > heightLimit {
				fmt.Println("Числа должны быть введены в порядке возрастания: ")
			} else {
				i = 1
			}
		}
	}

	return lowLimit, heightLimit

}

func choiceAccuracy() float64 {
	fmt.Print("Введите точность вычисления: ")

	acc := 0.0

	i := 0
	for i != 1 {
		_, err := fmt.Scan(&acc)
		if err != nil {
			fmt.Println("Вы должны ввести дробное число, равное необходимой точности вычисления")
			fmt.Print("Введите еще раз: ")
		} else {
			i = 1
		}
	}

	return acc
}

func choiceMethod() int {
	fmt.Println("Выберите нужный метод:\n" +
		"1: Метод левых прямоугольников\n" +
		"2: Метод средних прямоугольников\n" +
		"3: Метод правых прямоугольников\n" +
		"4: Метод трапеций\n" +
		"5: Метод Симпсона")

	method := 0

	i := 0
	for i != 1 {
		_, err := fmt.Scan(&method)
		if err != nil {
			fmt.Println("Вы должны ввести число соответствующее номеру выбранного метода")
			fmt.Print("Введите еще раз: ")
		} else {
			if method == 1 || method == 2 || method == 3 || method == 4 || method == 5 {
				i = 1
			} else {
				fmt.Println("Вы должны ввести число 1, 2, 3, 4 или 5")
				fmt.Print("Введите еще раз: ")
			}

		}
	}
	return method
}

//#############################################################################################################
