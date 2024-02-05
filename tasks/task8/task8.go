package main

import (
	"fmt"
)

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0

1. Генерируем маску с помощью битового сдвига 1-цы на i битов
2. Если хотим установить бит в 1, то все просто - делаем дизъюнкцию заданного числа и маски, так как 1 V 0 = 1, 1 V 1 = 1, то нужный бит примет значение 1,
а биты, идущие до i-го не поменяют своих значений, так как будут в дизъюнкции с 0
3. Если хотим установить бит в 0, то все сложнее. Просто применить конъюнкцию будет недостаточно а) биты до I-го могут поменять свое значение, так как в маске идут 0-ли, 1 & 0 дает 0
б) сам i-й бит может не поменяться, если будет 1-ой, так как 1 & 1 = 1
Поэтому используем операцию И-НЕ(&^), где будет конъюнкция исходного числа с "сопряженной" маской, у которой вместо 1-цы будет 0, а остальные i-1 битов будут 1-цы.
Это гарантирует, что при конъюнкции поменяется лишь нужный i-й бит, а остальные окажутся прежними.
*/

func setBit(value int64, position int, bitValue int) int64 {
	if position < 0 || position > 63 {
		fmt.Println("Invalid bit position.")
		return value
	}

	mask := int64(1) << uint(position)

	if bitValue == 1 {
		value |= mask
	} else if bitValue == 0 {
		value &^= mask
	} else {
		fmt.Println("Invalid bit value. Use 0 or 1.")
	}

	return value
}

func main() {
	var num int64
	fmt.Println("Введите число")
	fmt.Scan(&num)
	fmt.Println("")
	fmt.Println("ВВедите номер бита")
	var numBit int
	fmt.Scan(&numBit)
	fmt.Printf("До изменения: %d\n", num)

	num = setBit(num, numBit, 1)
	fmt.Printf("После установки бита в 1: %d\n", num)

	num = setBit(num, numBit, 0)
	fmt.Printf("После установки бита в 0: %d\n", num)
}
