package main

import (
	"bufio"
	"fmt"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab1/internal/pkg/arguments"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab1/internal/pkg/factorial"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab1/internal/pkg/fibonacci"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab1/internal/pkg/leap_years"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab1/internal/pkg/prime_nums"
	"os"
	"strconv"
)

func getUserInput(prompt string) int {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	value, err := strconv.Atoi(input[:len(input)-1])
	if err != nil {
		fmt.Println("Ошибка ввода. Используется значение по умолчанию.")
		return -1
	}
	return value
}

func main() {
	for {
		fmt.Println("Выберите функцию для выполнения:")
		fmt.Println("1 - Вывести аргументы командной строки")
		fmt.Println("2 - Найти високосные годы в заданном диапазоне")
		fmt.Println("3 - Вывести последовательность Фибоначчи")
		fmt.Println("4 - Вычислить факториал числа")
		fmt.Println("5 - Найти простые числа до заданного числа")
		fmt.Println("0 - Выйти")
		fmt.Print("Введите номер: ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(input[:len(input)-1])
		if err != nil {
			fmt.Println("Ошибка ввода. Попробуйте снова.")
			continue
		}

		switch choice {
		case 1:
			arguments.GiveArguments()
		case 2:
			start := getUserInput("Введите начальный год: ")
			end := getUserInput("Введите конечный год: ")
			if start != -1 && end != -1 {
				leap_years.GiveLeapYears(start, end)
			}
		case 3:
			n := getUserInput("Введите количество чисел Фибоначчи: ")
			if n != -1 {
				fibonacci.GiveFibonacci(n)
			}
		case 4:
			n := getUserInput("Введите число для факториала: ")
			if n != -1 {
				factorial.Factorial(n)
			}
		case 5:
			n := getUserInput("Введите предел для поиска простых чисел: ")
			if n != -1 {
				prime_nums.SieveOfEratosthenes(n)
			}
		case 0:
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Некорректный ввод, попробуйте снова.")
		}
		fmt.Println()
	}
}
