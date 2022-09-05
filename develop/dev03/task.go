package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	f := ParceFlags()

	lines, err := ReadLines(f)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	sortedLines, err := Sort(lines, f)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	WriteLines(sortedLines)
}

func Sort(lines []string, flags *Flags) (res []string, err error) {
	var ls []*Line
	for _, l := range lines {
		line := NewLine(l, flags)
		ls = append(ls, line)
	}

	sort.Slice(ls, func(i, j int) bool {
		if ls[i].keyInt < ls[j].keyInt {
			return true
		} else if ls[i].keyInt == ls[j].keyInt {
			if ls[i].keyStr < ls[j].keyStr {
				return true
			} else if ls[i].keyStr == ls[j].keyStr {
				return ls[i].fullStr < ls[j].fullStr
			}
		}

		return false
	})

	ls = deleteRepetitions(ls, flags)
	reverseOrder(ls, flags)

	for _, l := range ls {
		res = append(res, l.line)
	}

	return
}

type Line struct {
	line    string
	keyInt  int
	keyStr  string
	fullStr string
}

func (l *Line) Print() {
	fmt.Printf("line: '%s' int: '%d'  str: '%s' full: '%s'\n", l.line, l.keyInt, l.keyStr, l.fullStr)
}

func NewLine(l string, f *Flags) *Line {
	res := &Line{line: l}

	fields := strings.Fields(l)
	res.fullStr = strings.Join(fields, "")

	colInt := 0
	colStr := 0
	if f.ColumnSort > 1 {
		colStr = f.ColumnSort - 1
		colInt = f.ColumnSort - 1
	}

	if f.NumberSort {
		colStr++
	}

	if colStr < len(fields) {
		res.keyStr = strings.Join(fields[colStr:], "")
	}

	if colInt < len(fields) && f.NumberSort {
		res.keyInt = parseKeyInt(fields[colInt], f)
		if res.keyInt == 0 {
			res.keyStr = ""
		}
	}

	return res
}

func parseKeyInt(str string, f *Flags) int {
	res, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return res
}

func deleteRepetitions(lines []*Line, f *Flags) (res []*Line) {
	if !f.UniqueStrings {
		res = lines
		return
	}

	var equal func(l, r *Line) bool
	if f.NumberSort {
		equal = func(l, r *Line) bool {
			return l.keyInt == r.keyInt
		}
	} else {
		equal = func(l, r *Line) bool {
			return l.keyStr == r.keyStr
		}
	}

	for i, line := range lines {
		if i == 0 {
			res = append(res, line)
			continue
		}

		if equal(res[len(res)-1], line) {
			continue
		}

		res = append(res, line)
	}

	return
}

func reverseOrder(lines []*Line, f *Flags) {
	if !f.ReverseSort {
		return
	}

	for i := 0; i <= (len(lines)-1)/2; i++ {
		lines[i], lines[len(lines)-1-i] = lines[len(lines)-1-i], lines[i]
	}
}
