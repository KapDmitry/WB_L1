package main

import (
	"fmt"
	"math/rand"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

1. Использование глобальной переменной - всегда плохо и может привести к негативным последствиям, кто-то может поменять ее, могут быть конфликты имен и так далее
2. Несмотря на то, что работает с большой строкой - лучше всегда делать проверку на то, что длина оригинальной строки позволяет сделать срез в заданных пределах, иначе может быть выход за
границы массива и паника
3. Со строками надо работать крайне аккуратно: Строки в GO могут состоять из многобайтовых символов, поэтому для строки str := "こんにちは" slice := str[0:6]  будет содержать только "こん"
В строках с различными спец символами, например "😊🌍", получение слайса приведет к поломке символов (так один символ может занимать 3 байта, а мы возьмем [0:2])

Исходя из этого имеет смысл убрать глобальную переменну, в функции HugeString либо сразу возвращать массив рун, либо преобразовывать строку к такому массиву
Проверять его длину и создавать новую строку на его основе
*/

func createHugeString(size int) string {
	alphabet := []rune("ABCDEFGHIJKLMNOPQRST😊🌍UVWXYZabcdefghこんにちはijklmnopqrstuvwxyzäöüß")

	result := make([]rune, size)
	var i int
	for i < size {
		index := rand.Intn(len(alphabet))
		result[i] = alphabet[index]
		i++
	}
	return string(result)
}

func someFunc(justString *string) {
	v := createHugeString(1 << 10)
	runes := []rune(v)

	if len(runes) >= 100 {
		*justString = string(runes[:100])
	} else { // или возвращать ошибку
		*justString = string(runes)
	}
}

func main() {
	var justString string
	someFunc(&justString)
	fmt.Println(justString)
}
