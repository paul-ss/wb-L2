package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fmt.Println(FindAnagrams([]string{"пятак", "пятка", "тяпка", "пятка", "листок", "слиток", "столик"}))
}

func wordHash(word string) string {
	arr := []rune(strings.ToLower(word))
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	return string(arr)
}

func FindAnagrams(arr []string) map[string][]string {
	hashMap := make(map[string]string)
	anagramMap := make(map[string][]string)

	for _, w := range arr {
		w = strings.ToLower(w)
		hash := wordHash(w)

		firstWord, ok := hashMap[hash]
		if ok && firstWord != w {
			anagrams := anagramMap[firstWord]
			anagrams = append(anagrams, w)
			anagramMap[firstWord] = anagrams
		} else {
			hashMap[hash] = w
		}
	}

	for key, val := range anagramMap {
		sort.Slice(val, func(i, j int) bool {
			return val[i] < val[j]
		})

		var res []string
		for i, word := range val {
			if i == 0 {
				res = append(res, word)
				continue
			}

			if word == res[len(res)-1] {
				continue
			}

			res = append(res, word)
		}

		if len(res) != len(val) {
			anagramMap[key] = res
		}
	}

	return anagramMap
}
