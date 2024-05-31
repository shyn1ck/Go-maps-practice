package main

import (
	"fmt"
)

func main() {
	mapData := make(map[string]string)
	var n int
	fmt.Println("Введите количество пар ключ-значение:")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var key, value string
		fmt.Printf("Введите ключ %d: ", i+1)
		fmt.Scan(&key)
		fmt.Printf("Введите значение для ключа %s: ", key)
		fmt.Scan(&value)
		mapData[key] = value
	}
	var valueToRemove string
	fmt.Println("Введите значение, которое нужно удалить:")
	fmt.Scan(&valueToRemove)
	updatedMap := make(map[string]string)

	for key, value := range mapData {
		if value != valueToRemove {
			updatedMap[key] = value
		}
	}

	fmt.Println("Обновленная карта:")
	for key, value := range updatedMap {
		fmt.Printf("%s: %s\n", key, value)
	}
}
