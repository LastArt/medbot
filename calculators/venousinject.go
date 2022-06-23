package calculators

import (
	"fmt"
)

func VenInject(vol int, ti float64) (string, string) {

	res := float64(vol*20) / ti // время в минутах обязательно
	res2 := res / 60
	return fmt.Sprintf("Количество капель в минуту: %4.0f", res), fmt.Sprintf("\nКоличество капель в секунду: %4.1f", res2)
}
