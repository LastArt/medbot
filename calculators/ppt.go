package calculators

import (
	"fmt"
	"math"
)

func Ppt(w, h int) ([]string, string) {
	weight := float64(w)
	height := float64(h)
	var ppt []string
	ppt = append(ppt, "ПЛОЩАДЬ ПОВЕРХНОСТИ ТЕЛА ПО\n")
	//fmt.Println("ПЛОЩАДЬ ПОВЕРХНОСТИ ТЕЛА ПО РАЗНЫМ ФОРМУЛАМ")
	// Формула Дюбуа
	res := 0.007184 * math.Pow(weight, 0.425) * math.Pow(height, 0.725)
	ppt = append(ppt, fmt.Sprintf("Формуле Дюбуа:  %.2f\n", res))
	//Формула Дюбуа и Дюбуа (модификация)
	res = math.Pow(weight, 0.425) * math.Pow(height, 0.725) / 139.2
	ppt = append(ppt, fmt.Sprintf("Формуле Дюбуа (модификация):  %.2f\n", res))
	//Формула Мостеллера
	res = math.Sqrt(weight * height / 3600)
	ppt = append(ppt, fmt.Sprintf("Формуле Мостеллера:  %.2f\n", res))
	//Формула Хейкока
	res = 0.024265 * math.Pow(weight, 0.5378) * math.Pow(height, 0.3964)
	ppt = append(ppt, fmt.Sprintf("Формуле Хейкока:  %.2f\n", res))
	//Формула Гехана и Джорджа
	res = 0.0235 * math.Pow(weight, 0.51456) * math.Pow(height, 0.42246)
	ppt = append(ppt, fmt.Sprintf("Формуле Гехана и Джорджа:  %.2f\n", res))
	//Формула Бойда
	weightGr := weight * 1000
	deg := 0.7285 - (0.0188 * math.Log10(weight))
	res = 0.0003207 * math.Pow(weightGr, deg) * math.Pow(height, 0.3)
	ppt = append(ppt, fmt.Sprintf("Формуле Бойда:  %.2f\n", res))
	//Формула Фудзимото
	res = 0.008883 * math.Pow(weight, 0.663) * math.Pow(height, 0.444)
	ppt = append(ppt, fmt.Sprintf("Формуле Фудзимото:  %.2f\n", res))
	//Формула Такахира
	res = 0.007241 * math.Pow(weight, 0.725) * math.Pow(height, 0.425)
	ppt = append(ppt, fmt.Sprintf("Формуле Такахира:  %.2f\n", res))
	normDescription := "\nНормальное значение ППТ для взрослых составляет - 1.73 м\n" + "Значения ППТ (м2)\n" + "Новорожденный - 0.25\nРебенок 2 года - 0.5\nРебенок 9 лет - 1.07\nРебенок 10 лет - 1.14\nРебенок 12-13 лет - 1.33\nДля мужчин - 1.9\nДля женщин - 1.6"
	return ppt, normDescription
}
