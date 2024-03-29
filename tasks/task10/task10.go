package main

import (
	"fmt"
)

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

Идея решения следущая - поскольку расположение внутри групп не важно, то дробную часть смело можем откидывать, так как она вообще никакой роли играть не будет
Поскольку группы идут с шагом в 10 градусов, то саму группы можно вычислить отбросив последний разряд в числе, а потом умножив получившееся число на 10
*/
func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	temperatureGroups := make(map[int][]float64)

	for _, temp := range temperatures {
		group := int(temp/10) * 10
		temperatureGroups[group] = append(temperatureGroups[group], temp)
	}

	for group, temps := range temperatureGroups {
		fmt.Printf("%d: %v\n", group, temps)
	}
}
