package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower float64 = 2

func main() {

	for {
		fmt.Println("__ Калькулятор индекса массы тела__")
		userKg, userHeight := getUserInput()
		IMT, err := calculateIMT(userKg, userHeight)
		if err != nil {
			panic("Не заданы параметры для расчёта")
		}
		outputResult(IMT)
		isRepeateCalculation := checkRepeatCalculation()
		if !isRepeateCalculation {
			break
		}
	}

}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.2f", IMT)
	fmt.Println(result)
	switch true {
	case IMT < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case IMT < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case IMT < 25:
		fmt.Println("У вас дефицит массы тела")
	case IMT < 30:
		fmt.Println("У вас избыточная масса тела")
	default:
		fmt.Println("У вас cтепень ожирения")
	}
}
func calculateIMT(userKg float64, userHeight float64) (float64, error) {
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("не указан вес или высота")
	}
	IMT := userKg / math.Pow(userHeight, IMTPower)
	return IMT, nil
}
func getUserInput() (float64, float64) {
	var userHeight, userKg float64
	fmt.Print("Введите свой рост в метрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	return userKg, userHeight
}
func checkRepeatCalculation() bool {
	var userChoise string
	fmt.Print("Вы хотите сделать ещё расчёт (y/n): ")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}
