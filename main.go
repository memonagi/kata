package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var m1 = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}
var m2 = map[string]int{
	"1":  1,
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
}

const (
	p1 = "Вывод ошибки, так как используются одновременно разные системы счисления."
	p2 = "Вывод ошибки, так как в римской системе нет отрицательных чисел."
	p3 = "Вывод ошибки, так как строка не является математической операцией."
	p4 = "Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	p5 = "Калькулятор умеет работать только с арабскими целыми числами или римскими числами от 1 до 10 включительно."
	p6 = "Вывод ошибки, так как в римской системе нет числа 0."
)

// 1. Ввод данных. На вход подается строка. Программа определяет параметры "число"-"оператор"-"число" (x, z, y)
// Если один из параметров пустой, то выдается ошибка p3.
func input(console string) (string, string, string) {
	var s []rune
	var x, z, y string
	s = []rune(console)
	for i := 0; i < len(s); i++ {
		if unicode.IsSpace(s[i]) == true {
			continue
		} else if z == "" && (unicode.IsDigit(s[i]) == true || unicode.IsLetter(s[i]) == true) {
			x += string(s[i])
		} else if x != "" && z == "" {
			z = string(s[i])
			continue
		} else if z != "" && (unicode.IsDigit(s[i]) == true || unicode.IsLetter(s[i]) == true) {
			y += string(s[i])
		}
	}
	if x == "" || z == "" || y == "" {
		panic(p3)
	} else {
		return x, z, y
	}
}

// 2. Функция определяет, какие данные поданы на вход - римские или арабские.
// Если поданы римское и арабское число в одной строке, выдает ошибку p1.
// Если поданы числа <= 0 или >10, выдает ошибку p5.
func whichMap(a, b string) (int, int, bool, bool) {
	var usl1, usl2 bool
	var a2, b2 int
	for i, e := range m1 {
		if strings.ToUpper(a) == i || strings.ToUpper(b) == i {
			if strings.ToUpper(a) == i {
				a2 = e
				usl1 = true
			}
			if strings.ToUpper(b) == i {
				b2 = e
				usl2 = true
			}
		}
	}
	for j, el := range m2 {
		if a == j {
			a2 = el
		}
		if b == j {
			b2 = el
		}
	}
	if a2 == 0 || b2 == 0 {
		panic(p5)
	} else if usl1 != usl2 {
		panic(p1)
	} else {
		return a2, b2, usl1, usl2
	}
}

// 3. Сумма. Функция определяет сумму параметров, если оператор z == "+".
func plus(a, b int, usl1, usl2 bool) string {
	var res string
	c := a + b
	if usl1 == true && usl2 == true {
		if c <= 10 {
			for i, e := range m1 {
				if c == e {
					res = i
				}
			}
		} else if c > 10 && c <= 20 {
			s1 := make([]string, 2)
			s1[0] = strings.Repeat("X", c/10)
			for i, e := range m1 {
				if c%10 == 0 {
					res = s1[0]
				} else if c%10 == e {
					s1[1] = i
					res = s1[0] + s1[1]
				}
			}
		}
	} else {
		res = strconv.Itoa(c)
	}
	return res
}

// 4. Вычитание. Функция вычитает один параметр из другого, если оператор z == "-".
// При работе с римскими цифрами выдает ошибку p2, если ответ меньше 0, и ошибку p6, если ответ равен 0.
func minus(a, b int, usl1, usl2 bool) string {
	var res string
	c := a - b
	if usl1 == true && usl2 == true {
		if c > 0 && c <= 10 {
			for i, e := range m1 {
				if c == e {
					res = i
				}
			}
		} else if c < 0 {
			panic(p2)
		} else if c == 0 {
			panic(p6)
		}
	} else {
		res = strconv.Itoa(c)
	}
	return res
}

// 5. Произведение. Функция определяет произведение параметров, если оператор z == "*".
func comp(a, b int, usl1, usl2 bool) string {
	var res string
	c := a * b
	if usl1 == true && usl2 == true {
		if c <= 10 {
			for i, e := range m1 {
				if c == e {
					res = i
				}
			}
		} else if c > 10 && c < 40 {
			s1 := make([]string, 2)
			s1[0] = strings.Repeat("X", c/10)
			for i, e := range m1 {
				if c%10 == 0 {
					res = s1[0]
				} else if c%10 == e {
					s1[1] = i
					res = s1[0] + s1[1]
				}
			}
		} else if c >= 40 && c < 50 {
			s2 := make([]string, 2)
			s2[0] = "XL"
			for i, e := range m1 {
				if c%10 == 0 {
					res = s2[0]
				} else if c%10 == e {
					s2[1] = i
					res = s2[0] + s2[1]
				}
			}
		} else if c >= 50 && c < 60 {
			s3 := make([]string, 2)
			s3[0] = "L"
			for i, e := range m1 {
				if c%10 == 0 {
					res = s3[0]
				} else if c%10 == e {
					s3[1] = i
					res = s3[0] + s3[1]
				}
			}
		} else if c >= 60 && c < 90 {
			s4 := make([]string, 2)
			s4[0] = "L" + strings.Repeat("X", c/10-5)
			for i, e := range m1 {
				if c%10 == 0 {
					res = s4[0]
				} else if c%10 == e {
					s4[1] = i
					res = s4[0] + s4[1]
				}
			}
		} else if c >= 90 && c < 100 {
			s5 := make([]string, 2)
			s5[0] = "XC"
			for i, e := range m1 {
				if c%10 == 0 {
					res = s5[0]
				} else if c%10 == e {
					s5[1] = i
					res = s5[0] + s5[1]
				}
			}
		} else if c == 100 {
			res = "C"
		}
	} else {
		res = strconv.Itoa(c)
	}
	return res
}

// 6. Деление. Функция делит один параметр на другой, если оператор z == "/".
// При работе с римскими цифрами выдает ошибку p6, если ответ равен 0.
func div(a, b int, usl1, usl2 bool) string {
	var res string
	c := a / b
	if usl1 == true && usl2 == true {
		if c > 0 && c <= 10 {
			for i, e := range m1 {
				if c == e {
					res = i
				}
			}
		} else if c == 0 {
			panic(p6)
		}
	} else {
		res = strconv.Itoa(c)
	}
	return res
}

// 7. Калькулятор. Выводит ответ в зависимости от введенных параметров.
// Если оператор z != "+", "-", "/", "*", выводит ошибку p4.
func calc(a2, b2 int, usl1, usl2 bool, z string) string {
	var result string
	if z != "+" && z != "-" && z != "*" && z != "/" && z != "" {
		panic(p4)
	} else {
		switch z {
		case "+":
			result = plus(a2, b2, usl1, usl2)
		case "-":
			result = minus(a2, b2, usl1, usl2)
		case "*":
			result = comp(a2, b2, usl1, usl2)
		case "/":
			result = div(a2, b2, usl1, usl2)
		}
	}
	return result
}

// 8. Старт работы программы.
func start(console string) string {
	x, z, y := input(console)
	x2, y2, usl1, usl2 := whichMap(x, y)
	result := calc(x2, y2, usl1, usl2, z)
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите формулу для расчета:")
		console, _ := reader.ReadString('\n')
		result := start(console)
		fmt.Println("Ответ: ", result)
	}
}
