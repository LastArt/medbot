package calculators

import "fmt"

func Alco(gender string, weight, vol, str int) string {
	var someResult string

	var vidmark_man int = int(0.70 * 100)
	var vidmark_wmn int = int(0.60 * 100)
	var res float64
	if gender == "true" {
		res = float64(vol / (weight * vidmark_man))
		someResult = fmt.Sprintf("Количество алкоголя в крови - %.4f промилле", res)
	} else if gender == "false" {
		res = float64(vol / (weight * vidmark_wmn))
		someResult = fmt.Sprintf("Количество алкоголя в крови - %.4f промилле", res)
	} else {
		someResult = "🚫 Пол не определен\n💬 Введите пожалуйста '1'-Мужчина или '2'-Женщина !"
	}
	return someResult
}
