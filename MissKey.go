package main

import "fmt"

func main() {
	inputMap := map[string]string{
		"Яблоко":   "красное",
		"Банан":    "желтый",
		"Вишня":    "красная",
		"Груша":    "зеленая",
		"Клубника": "красная",
	}

	keys := []string{"Банан", "Груша", "Киви", "Ананас", "Апельсин"}

	missingKeys := []string{}
	for _, key := range keys {
		if _, exists := inputMap[key]; !exists {
			missingKeys = append(missingKeys, key)
		}
	}

	fmt.Println("Отсутствующие фрукты:", missingKeys)
}
