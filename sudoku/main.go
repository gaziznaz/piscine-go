package main

import (
	"os"

	"github.com/01-edu/z01"
)

var (
	arr   = make([][]int, 9) // создание массива
	count = 0                // если отснется 0, то судоку не решается
)

func main() {
	arg := os.Args[1:] // cчитываем аргументы
	if !CheckInput(arg) {
		return
	}
	Converter(arg)
	Solve(0, 0)
	if count != 1 { // если не решится солв то каунт =  0
		PrintError()
	}
}

func CheckInput(arg []string) bool {
	if len(arg) != 9 {
		PrintError()
		return false
	}
	for _, line := range arg {
		if len(line) != 9 {
			PrintError()
			return false
		}
		for _, element := range line {
			if element != '.' && (element < '1' || element > '9') {
				PrintError()
				return false
			}
		}
	}
	return true
}

func Converter(arg []string) { // обьявление функции для конвертирования стринг в инт
	for i := 0; i < 9; i++ { // перебор строк
		var p []int
		for j := 0; j < 9; j++ { // перебор символов в строке
			for k := 49; k < 58; k++ { // подбор цифр по аски
				if rune(arg[i][j]) == rune(k) {
					p = append(p, k-48) // добавление в массив цифр
				}
			}
			if rune(arg[i][j]) == '.' { // определение точки как 0
				p = append(p, 0)
			}
		}
		arr[i] = p // добавление в массив арр
	}
}

func Solve(row int, col int) { // функция решения
	if row == 8 && col == 9 { // если проходит условия, судоку решена
		count++
		if !CheckRes() {
			return
		}
		for i := 0; i < 9; i++ { // распечатка
			for j := 0; j < 9; j++ {
				z01.PrintRune(rune(arr[i][j]) + 48) // перевод в символы
				if j != 8 {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
		return
	}
	if col == 9 { // если значение выходит за пределы массива, то переход на сл строку
		row++
		col = 0
	}
	if arr[row][col] > 0 { // проверка на наличие цифры
		Solve(row, col+1)
		return
	}
	for num := 1; num <= 9; num++ { // подбор цифр для пустой ячейки(ячейка = 0)
		if CheckValue(row, col, num) == true {
			arr[row][col] = num
			Solve(row, col+1)
		}
		arr[row][col] = 0 // если ни одно из значений не подходит, ставится 0 и возврат на ячейку назад
	}
}

func CheckValue(row int, col int, num int) bool { // функция проверки
	for i := 0; i <= 8; i++ { // проверка по горизонтали
		if arr[row][i] == num || arr[i][col] == num {
			return false
		}
	}

	var startRow int = row - row%3 // определение координаты внутри маленького квадрата
	var startCol int = col - col%3
	for i := 0; i < 3; i++ { // проверка по квадратам
		for j := 0; j < 3; j++ {
			if arr[i+startRow][j+startCol] == num { // проверка на наличие числа
				return false
			}
		}
	}
	return true
}

func CheckRes() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			for k := 0; k < 9; k++ {
				if j != k && (arr[i][j] == arr[i][k] || arr[j][i] == arr[k][i]) {
					PrintError()
					return false
				}
			}
		}
	}
	return true
}

func PrintError() {
	str := "Error"
	for _, letter := range str {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}
