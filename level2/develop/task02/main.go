// Создать Go-функцию, осуществляющую примитивную распаковку
// строки, содержащую повторяющиеся символы/руны, например:
// ● "a4bc2d5e" => "aaaabccddddde"
// ● "abcd" => "abcd"
// ● "45" => "" (некорректная строка)
// ● "" => ""
// Дополнительно
// Реализовать поддержку escape-последовательностей.
// Например:
// ● qwe\4\5 => qwe45 (*)
// ● qwe\45 => qwe44444 (*)
// ● qwe\\5 => qwe\\\\\ (*)
// В случае если была передана некорректная строка, функция
// должна возвращать ошибку. Написать unit-тесты.

package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	test := []string{`a4bc2d5e`, `abcd`, `45`, ``, `ゴ\\2ウ\1ケ3ツ`, `qwe\4\5`, `qwe\45`, `qwe\\3`}
	expected := []string{"aaaabccddddde", "abcd", "", "", `ゴ\\ウ1ケケケツ`, "qwe45", "qwe44444", `qwe\\\`}
	expectedErrors := []error{nil, nil, fmt.Errorf("(некорректная строка)"), nil, nil, nil, nil, nil}

	for i, s := range test {
		res, err := StringConvertation(&s)
		fmt.Println("Res:", expected[i], "-----", res)
		fmt.Println("Err:", expectedErrors[i], "-----", err)
	}
}

func StringConvertation(str *string) (string, error) {
	runes := make([]rune, len(*str), 2*len(*str))
	CurrIsString := true
	PrevIsString := false
	var buff rune
	var err error

	for _, r := range []rune(*str) {
		v, errstr := strconv.Atoi(string(r))
		if errstr != nil {
			if CurrIsString == PrevIsString && string(buff) != "\\" {
				runes = append(runes, buff)
			}
			if string(buff) == "\\" && string(r) == "\\" {
				CurrIsString = false
			}
			PrevIsString = true
		} else {
			if string(buff) != "\\" {
				if PrevIsString == false {
					err = errors.New("(некорректная строка)")
					return "", err
				}
				for j := 0; j < v; j++ {
					runes = append(runes, buff)
				}
				PrevIsString = false
			} else if CurrIsString == false {
				for j := 0; j < v; j++ {
					runes = append(runes, buff)
				}
				PrevIsString = false
			}
			CurrIsString = true
		}
		buff = r
	}

	if PrevIsString == true {
		runes = append(runes, buff)
	}

	return string(runes), err
}
