package calculators

import "fmt"

func MidAD(sisad float64, diastad float64) string {
	midleAd := (sisad / 3) + ((diastad / 3) * 2)
	var desc string
	if midleAd > 70 || midleAd < 110 {
		desc = "Норма"
	} else if midleAd < 70 {
		desc = "Пониженное артериальное давление. Возможна гипоксия и ишемия тканей"
	} else {
		desc = "Повышенное артериальное давление"
	}
	someResult := fmt.Sprintf("Среднее артериальное давление равно - %.2f мм.рт.ст\nКомментарий: %s", midleAd, desc)

	return someResult
}
