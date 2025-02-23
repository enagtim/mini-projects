package main

import "fmt"

func main() {

	transactions := []float64{}
	for {
		transaction := getTransaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}
	balance := calculateBalanse(transactions)
	fmt.Println("Ваша сумма баланса:", balance)
}

func getTransaction() float64 {
	var transaction float64
	fmt.Print("Введите вашу транзакцию(n для выхода): ")
	fmt.Scan(&transaction)
	return transaction
}
func calculateBalanse(transactions []float64) float64 {
	var balance float64
	for _, value := range transactions {
		balance += value
	}
	return balance
}
