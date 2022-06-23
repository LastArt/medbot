package calculators

import (
	"fmt"
	"math"
)

func RomaIndex(wmnAge, he, ca125 float64) string {
	var someResult string

	// формула индекса РОМА
	//  Расчет прогностического индекса (ПИ, Predictive Index, PI):
	//  1. Пременопауза: ПИ = -12,0 + 2,38 х LN [HЕ4] + 0,0626 х LN [CA125].
	//  2. Постменопауза: ПИ = -8,09 + 1,04 х LN [HЕ4] + 0,732 х LN [CA125].
	// 	Расчет ROMA:
	//	ROMA (%) = exp (ПИ) / [1 + exp (ПИ) ] х 100.
	var pIndex float64
	var roma float64
	var risk string
	switch wmnAge {
	case 1: // Пременопауза
		pIndex = -12.0 + 2.38*math.Log(he) + 0.0626*math.Log(ca125)
		roma = math.Exp(pIndex) / (1 + math.Exp(pIndex)) * 100
		if roma >= 12.5 {
			risk = "Высокий риск"
		} else {
			risk = ""
		}
	case 2: // Постменопауза
		pIndex = -8.09 + 1.04*math.Log(he) + 0.732*math.Log(ca125)
		roma = math.Exp(pIndex) / (1 + math.Exp(pIndex)) * 100
		if roma >= 14.4 {
			risk = "Высокий риск"
		} else {
			risk = ""
		}
	default:
		someResult = "🚫 Не определена возрастная категория\n💬 Введите пожалуйста '1'- Пременопауза или '2'- Постменопауза !"
	}

	someResult = fmt.Sprintf("\nПри показателях:  HE-4 - %.2f пмоль/л и СА-125 - %.2f Ед/мл \nИндекса РОМА = %.2f \n%s", he, ca125, roma, risk)

	return someResult
}
