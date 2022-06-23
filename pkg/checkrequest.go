package pkg

import (
	"fmt"
	"medcalcbot/set"
	"regexp"
	"strings"
)

func CheckRequest(f string) string {
	charSum := len(f) / 2                  // Сразу определяем идентификатор количества символов для кирилицы
	search := []string{}                   // Определяем промежуточный массив выводимых в ответе данных
	f = strings.ToLower(f)                 // Намеренно приводим все символы в lowercase
	re := regexp.MustCompile(set.RuSymbol) //проверяем на киррилические символы
	isRussian := re.MatchString(f)
	if isRussian {
		if charSum < 3 {
			req1 := "🚫 Введите более 2-х начальных букв из названия показателя\n💬 К примеру для поиска гемаглобина введите'гем' или 'гемо'!"
			search = append(search[:], req1)
		} else {
			search = append(search, f)
		}
	} else {
		req2 := "🚫 Проверьте вводимые символы.\n💬 Поиск возможен только с использованием русских букв!"
		search = append(search[:], req2)
	}
	x := fmt.Sprintf(strings.Join(search, ""))
	return x
}
