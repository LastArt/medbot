package calculators

import (
	"fmt"
)

func Bov(gender, height, weight, age, fiz float64) string {

	var someResult string
	var prefix string = "Для поддержания веса организму требуется:\n\n"
	if gender == 1 {
		switch fiz {
		case 1:
			res_haris := 66.5 + (13.75 * weight) + (5.003 * height) - (6.775 * age)
			res_maffin := 10*weight + 6.25*height - 5*age + 5
			res_venuto := 66 + (13.7 * weight) + (5 * height) - (6.8 * age)

			someResult = fmt.Sprintf("%s \nПо формуле Харриса-Бенедикта (разработана более полувека назад) - %.0f ккал в день\nПо формуле Маффина-Джеора (выведена в 1990г) - %.0f ккал в день\nПо формуле Тома Венуто (популярна среди спортсменов) - %.0f ккал в день", prefix, res_haris, res_maffin, res_venuto)
		case 2:
			res_haris := (66.5 + (13.75 * weight) + (5.003 * height) - (6.775 * age)) * 1.2
			res_maffin := (10*weight + 6.25*height - 5*age + 5) * 1.2
			res_venuto := (66 + (13.7 * weight) + (5 * height) - (6.8 * age)) * 1.2
			someResult = fmt.Sprintf("%s \nПо формуле Харриса-Бенедикта (разработана более полувека назад) - %.0f ккал в день\nПо формуле Маффина-Джеора (выведена в 1990г) - %.0f ккал в день\nПо формуле Тома Венуто (популярна среди спортсменов) - %.0f ккал в день", prefix, res_haris, res_maffin, res_venuto)
		case 3:
			res_haris := (66.5 + (13.75 * weight) + (5.003 * height) - (6.775 * age)) * 1.375
			res_maffin := (10*weight + 6.25*height - 5*age + 5) * 1.375
			res_venuto := (66 + (13.7 * weight) + (5 * height) - (6.8 * age)) * 1.375
			someResult = fmt.Sprintf("%s \nПо формуле Харриса-Бенедикта (разработана более полувека назад) - %.0f ккал в день\nПо формуле Маффина-Джеора (выведена в 1990г) - %.0f ккал в день\nПо формуле Тома Венуто (популярна среди спортсменов) - %.0f ккал в день", prefix, res_haris, res_maffin, res_venuto)
		case 4:
			res_haris := (66.5 + (13.75 * weight) + (5.003 * height) - (6.775 * age)) * 1.4625
			res_maffin := (10*weight + 6.25*height - 5*age + 5) * 1.4625
			res_venuto := (66 + (13.7 * weight) + (5 * height) - (6.8 * age)) * 1.4625
			someResult = fmt.Sprintf("%s \nПо формуле Харриса-Бенедикта (разработана более полувека назад) - %.0f ккал в день\nПо формуле Маффина-Джеора (выведена в 1990г) - %.0f ккал в день\nПо формуле Тома Венуто (популярна среди спортсменов) - %.0f ккал в день", prefix, res_haris, res_maffin, res_venuto)
		case 5:
			res_haris := (66.5 + (13.75 * weight) + (5.003 * height) - (6.775 * age)) * 1.550
			res_maffin := (10*weight + 6.25*height - 5*age + 5) * 1.550
			res_venuto := (66 + (13.7 * weight) + (5 * height) - (6.8 * age)) * 1.550
			someResult = fmt.Sprintf("%s \nПо формуле Харриса-Бенедикта (разработана более полувека назад) - %.0f ккал в день\nПо формуле Маффина-Джеора (выведена в 1990г) - %.0f ккал в день\nПо формуле Тома Венуто (популярна среди спортсменов) - %.0f ккал в день", prefix, res_haris, res_maffin, res_venuto)
		case 6:
			res_haris := (66.5 + (13.75 * weight) + (5.003 * height) - (6.775 * age)) * 1.6375
			res_maffin := (10*weight + 6.25*height - 5*age + 5) * 1.6375
			res_venuto := (66 + (13.7 * weight) + (5 * height) - (6.8 * age)) * 1.6375
			someResult = fmt.Sprintf("%s \nПо формуле Харриса-Бенедикта (разработана более полувека назад) - %.0f ккал в день\nПо формуле Маффина-Джеора (выведена в 1990г) - %.0f ккал в день\nПо формуле Тома Венуто (популярна среди спортсменов) - %.0f ккал в день", prefix, res_haris, res_maffin, res_venuto)
		case 7:
			res_haris := (66.5 + (13.75 * weight) + (5.003 * height) - (6.775 * age)) * 1.725
			res_maffin := (10*weight + 6.25*height - 5*age + 5) * 1.725
			res_venuto := (66 + (13.7 * weight) + (5 * height) - (6.8 * age)) * 1.725
			someResult = fmt.Sprintf("%s \nПо формуле Харриса-Бенедикта (разработана более полувека назад) - %.0f ккал в день\nПо формуле Маффина-Джеора (выведена в 1990г) - %.0f ккал в день\nПо формуле Тома Венуто (популярна среди спортсменов) - %.0f ккал в день", prefix, res_haris, res_maffin, res_venuto)
		case 8:
			res_haris := (66.5 + (13.75 * weight) + (5.003 * height) - (6.775 * age)) * 1.9
			res_maffin := (10*weight + 6.25*height - 5*age + 5) * 1.9
			res_venuto := (66 + (13.7 * weight) + (5 * height) - (6.8 * age)) * 1.9
			someResult = fmt.Sprintf("%s \nПо формуле Харриса-Бенедикта (разработана более полувека назад) - %.0f ккал в день\nПо формуле Маффина-Джеора (выведена в 1990г) - %.0f ккал в день\nПо формуле Тома Венуто (популярна среди спортсменов) - %.0f ккал в день", prefix, res_haris, res_maffin, res_venuto)
		}
	} else if gender == 2 {
		switch fiz {
		case 1:
			res_haris := 655.1 + (9.563 * weight) + (1.85 * height) - (4.775 * age)
			res_maffin := 10*weight + 6.25*height - 5*age - 161
			res_venuto := 665 + (9.6 * weight) + (1.8 * height) - (4.7 * age)
			someResult = fmt.Sprintf("%s \nПо формуле Харриса-Бенедикта (разработана более полувека назад) - %.0f ккал в день\nПо формуле Маффина-Джеора (выведена в 1990г) - %.0f ккал в день\nПо формуле Тома Венуто (популярна среди спортсменов) - %.0f ккал в день", prefix, res_haris, res_maffin, res_venuto)
		case 2:
			res_haris := (655.1 + (9.563 * weight) + (1.85 * height) - (4.775 * age)) * 1.2
			res_maffin := (10*weight + 6.25*height - 5*age - 161) * 1.2
			res_venuto := (665 + (9.6 * weight) + (1.8 * height) - (4.7 * age)) * 1.2
			someResult = fmt.Sprintf("%s \nПо формуле Харриса-Бенедикта (разработана более полувека назад) - %.0f ккал в день\nПо формуле Маффина-Джеора (выведена в 1990г) - %.0f ккал в день\nПо формуле Тома Венуто (популярна среди спортсменов) - %.0f ккал в день", prefix, res_haris, res_maffin, res_venuto)
		case 3:
			res_haris := (655.1 + (9.563 * weight) + (1.85 * height) - (4.775 * age)) * 1.375
			res_maffin := (10*weight + 6.25*height - 5*age - 161) * 1.375
			res_venuto := (665 + (9.6 * weight) + (1.8 * height) - (4.7 * age)) * 1.375
			someResult = fmt.Sprintf("%s \nПо формуле Харриса-Бенедикта (разработана более полувека назад) - %.0f ккал в день\nПо формуле Маффина-Джеора (выведена в 1990г) - %.0f ккал в день\nПо формуле Тома Венуто (популярна среди спортсменов) - %.0f ккал в день", prefix, res_haris, res_maffin, res_venuto)
		case 4:
			res_haris := (655.1 + (9.563 * weight) + (1.85 * height) - (4.775 * age)) * 1.4625
			res_maffin := (10*weight + 6.25*height - 5*age - 161) * 1.4625
			res_venuto := (665 + (9.6 * weight) + (1.8 * height) - (4.7 * age)) * 1.4625
			someResult = fmt.Sprintf("%s \nПо формуле Харриса-Бенедикта (разработана более полувека назад) - %.0f ккал в день\nПо формуле Маффина-Джеора (выведена в 1990г) - %.0f ккал в день\nПо формуле Тома Венуто (популярна среди спортсменов) - %.0f ккал в день", prefix, res_haris, res_maffin, res_venuto)
		case 5:
			res_haris := (655.1 + (9.563 * weight) + (1.85 * height) - (4.775 * age)) * 1.550
			res_maffin := (10*weight + 6.25*height - 5*age - 161) * 1.550
			res_venuto := (665 + (9.6 * weight) + (1.8 * height) - (4.7 * age)) * 1.550
			someResult = fmt.Sprintf("%s \nПо формуле Харриса-Бенедикта (разработана более полувека назад) - %.0f ккал в день\nПо формуле Маффина-Джеора (выведена в 1990г) - %.0f ккал в день\nПо формуле Тома Венуто (популярна среди спортсменов) - %.0f ккал в день", prefix, res_haris, res_maffin, res_venuto)
		case 6:
			res_haris := (655.1 + (9.563 * weight) + (1.85 * height) - (4.775 * age)) * 1.6375
			res_maffin := (10*weight + 6.25*height - 5*age - 161) * 1.6375
			res_venuto := (665 + (9.6 * weight) + (1.8 * height) - (4.7 * age)) * 1.6375
			someResult = fmt.Sprintf("%s \nПо формуле Харриса-Бенедикта (разработана более полувека назад) - %.0f ккал в день\nПо формуле Маффина-Джеора (выведена в 1990г) - %.0f ккал в день\nПо формуле Тома Венуто (популярна среди спортсменов) - %.0f ккал в день", prefix, res_haris, res_maffin, res_venuto)
		case 7:
			res_haris := (655.1 + (9.563 * weight) + (1.85 * height) - (4.775 * age)) * 1.725
			res_maffin := (10*weight + 6.25*height - 5*age - 161) * 1.725
			res_venuto := (665 + (9.6 * weight) + (1.8 * height) - (4.7 * age)) * 1.725
			someResult = fmt.Sprintf("%s \nПо формуле Харриса-Бенедикта (разработана более полувека назад) - %.0f ккал в день\nПо формуле Маффина-Джеора (выведена в 1990г) - %.0f ккал в день\nПо формуле Тома Венуто (популярна среди спортсменов) - %.0f ккал в день", prefix, res_haris, res_maffin, res_venuto)
		case 8:
			res_haris := (655.1 + (9.563 * weight) + (1.85 * height) - (4.775 * age)) * 1.9
			res_maffin := (10*weight + 6.25*height - 5*age - 161) * 1.9
			res_venuto := (665 + (9.6 * weight) + (1.8 * height) - (4.7 * age)) * 1.9
			someResult = fmt.Sprintf("%s \nПо формуле Харриса-Бенедикта (разработана более полувека назад) - %.0f ккал в день\nПо формуле Маффина-Джеора (выведена в 1990г) - %.0f ккал в день\nПо формуле Тома Венуто (популярна среди спортсменов) - %.0f ккал в день", prefix, res_haris, res_maffin, res_venuto)
		}
	} else {
		someResult = "🚫 Пол не определен\n💬 Введите пожалуйста '1'-Мужчина или '2'-Женщина !"
	}

	return someResult
}
