package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic("ошибка")
	}

	input = strings.TrimSpace(input)

	if len(input) > 20 {
		panic("ошибка")
	}

	var result string
	switch {
	case strings.Contains(input, "-"):
		result, err = strV(input)
	case strings.Contains(input, "+"):
		result, err = strS(input)
	case strings.Contains(input, "*"):
		result, err = strU(input)
	case strings.Contains(input, "/"):
		result, err = strD(input)
	default:
		panic("ошибка")
	}

	if err != nil {
		panic("ошибка")
	}

	result = remove(result)
	if len(result) > 40 {
		fmt.Printf("Результат: \"%s\"\n", result[:40]+"...")
	}
	fmt.Printf("Результат: \"%s\"\n", result)
}

func remove(str string) string {
	return strings.ReplaceAll(str, `"`, "")
}

func strS(input string) (string, error) {
	parts := strings.Split(input, "+")
	if len(parts) != 2 {
		panic("ошибка")
	}
	str1 := strings.TrimSpace(parts[0])
	str2 := strings.TrimSpace(parts[1])

	return str1 + str2, nil
}

func strV(input string) (string, error) {
	index := strings.Index(input, "-")
	if index == -1 {
		return input, nil
	}
	result := strings.TrimSpace(input[:index])
	return result, nil
}

func strU(input string) (string, error) {
	parts := strings.Split(input, "*")
	if len(parts) != 2 {
		panic("ошибка")
	}
	str1 := strings.TrimSpace(parts[0])
	times, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		panic("ошибка")
	}
	result := strings.Repeat(str1, times)
	return result, nil
}

func strD(input string) (string, error) {
	parts := strings.Split(input, "/")
	if len(parts) != 2 {
		panic("ошибка")
	}
	str1 := strings.TrimSpace(parts[0])
	divisor, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		panic("ошибка")
	}
	if divisor <= 0 || divisor > len(str1) {
		panic("ошибка")
	}
	result := str1[:len(str1)/divisor]
	return result, nil
}
