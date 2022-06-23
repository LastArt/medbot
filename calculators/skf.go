package calculators

import (
	"fmt"
	"math"
)

func Skf(gender, creatinine, age, height, weight float64) string {
	var someResult string
	// math.pow(число, степень)
	switch gender {
	case 1:
		// По Кокрофту-Голту
		skfCrocf := ((140 - age) * weight) / (0.81 * creatinine)
		// По формуле MDRD
		skfMdrd := 186 * math.Pow((creatinine*0.0113096584483149), -1.154) * math.Pow(age, -0.203)
		someResult = fmt.Sprintf("У мужчины:\nпо Крокрофту-Голту - %.2f \nпо формуле MDRD - %.2f", skfCrocf, skfMdrd)
	case 2:
		// По Кокрофту-Голту женщины
		skfCrocf := (((140 - age) * weight) / (0.81 * creatinine)) * 0.81
		skfMdrd := (186 * math.Pow((creatinine*0.0113096584483149), -1.154) * math.Pow(age, -0.203)) * 0.742
		someResult = fmt.Sprintf("У женщины:\nпо Крокрофту-Голту - %.1f \nпо формуле MDRD - %.2f", skfCrocf, skfMdrd)
	default:
		someResult = "🚫 Пол не определен\n💬 Введите пожалуйста '1'-Мужчина или '2'-Женщина !"
	}

	return someResult
}
