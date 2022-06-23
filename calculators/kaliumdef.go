package calculators

import "fmt"

func KaliumDef(ks float64, w float64) string {
	// Нормальное содержание

	deficiency := float64(w) * 0.2 * (5 - ks)
	volume7 := 1.0 * deficiency
	volume4 := 1.8 * deficiency
	maxdoz := 3 * float64(w)
	return fmt.Sprintf("Дефицит калия составляет %4.1f моль/л\nДля возмещения необходимо: 7,5 проц-го раствора: %4.1f мл \nили 4-х проц-го: %4.1f мл\n\nМаксимальная дозировка в сутки составляет: : %4.0f ммоль", deficiency, volume7, volume4, maxdoz)
}
