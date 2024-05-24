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
	input, err := reader.ReadString('\n')
	if err != nil {
		panic("Ошибка чтения ввода")
	}
	input = strings.TrimSuffix(input, "\n")

	if strings.Contains(input, "-") {
		result := strV(input)
		fmt.Println(result)
	} else if strings.Contains(input, "+") {
		result := strS(input)
		fmt.Println(result)
	} else if strings.Contains(input, "*") {
		result := strU(input)
		fmt.Println(result)
	} else if strings.Contains(input, "/") {
		result := strD(input)
		fmt.Println(result)
	} else {
		panic("Оператор не распознан")
	}
}

func strS(input string) string {
	parts := strings.Split(input, "+")
	if len(parts) != 2 {
		panic("некорректный формат ввода")
	}
	str1 := strings.TrimSpace(parts[0])
	str2 := strings.TrimSpace(parts[1])
	return str1 + str2
}

func strV(input string) string {
	index := strings.Index(input, "-")
	if index == -1 {
		return input
	}
	result := strings.TrimSpace(input[:index])
	return result
}

func strU(input string) string {
	parts := strings.Split(input, "*")
	if len(parts) != 2 {
		panic("некорректный формат ввода")
	}
	str1 := strings.TrimSpace(parts[0])
	times, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		panic("второй аргумент должен быть числом")
	}
	result := strings.Repeat(str1, times)
	return result
}

func strD(input string) string {
	parts := strings.Split(input, "/")
	if len(parts) != 2 {
		panic("некорректный формат ввода")
	}
	str1 := strings.TrimSpace(parts[0])
	divisor, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		panic("второй аргумент должен быть числом")
	}
	if divisor <= 0 || divisor > len(str1) {
		panic("делитель должен быть положительным и не больше длины строки")
	}
	result := str1[:len(str1)/divisor]
	return result
}
