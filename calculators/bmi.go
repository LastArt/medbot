package calculators

import (
	"fmt"
	"math"
)

func Bmi(h, w float64) string {

	height := float64(h) / 100
	weight := float64(w)
	index := weight / math.Pow(height, 2)
	var category string
	if index <= 16 {
		category = fmt.Sprintf("Выраженный дефицит массы! \nИМТ =%.2f", index)
	} else if index > 16.01 && index <= 18.5 {
		category = fmt.Sprintf("Недостаточная (дефицит) масса тела! \nИМТ =%.2f", index)
	} else if index > 18.51 && index <= 25 {
		category = fmt.Sprintf("Норма! \nИМТ =%.2f", index)
	} else if index > 25.01 && index <= 30 {
		category = fmt.Sprintf("Избыточная масса тела (предожирение)! \nИМТ =%.2f", index)
	} else if index > 30.01 && index <= 35 {
		category = fmt.Sprintf("Ожирение первой степени! \nИМТ =%.2f", index)
	} else if index > 35.01 && index <= 40 {
		category = fmt.Sprintf("Ожирение второй степени! \nИМТ =%.2f", index)
	} else if index > 40.01 {
		category = fmt.Sprintf("Ожирение третьей степени (морбидное)! \nИМТ =%8.2f", index)
	}

	return category
}
