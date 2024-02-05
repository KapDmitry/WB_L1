package main

import "fmt"

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
Я сделал решение для самой сложной вариации задачи, где могут быть лишние пробелы перед началом строки, после строки и между словами
Воспользуемся методом двух указателей.
Первоначально, оба указателя указывают на 0-й элемент строки
Далее, двигаем один из указателей(j) до тех пор пока не наткнемся на букву
Далее, двигаем указатель на (j) до тех пор, пока не наткнемся на пробел, в это время *
меняем между собой значения под i и j. В случае, когда строка начинается с пробелов, получается, что i указывало бы на первый пробел, а j на первую букву(на момент начала шага *)
Таким образом, мы бы на место начальных пробелов ставили бы букву, а пробелы как-бы переносили за слово.
Потом делаем i++ (один пробел в норме у нас всегда должен быть)
и повторяем алгоритм
Следом, срезаем строку(отрезая все лишние пробелы, которые окажутся перенесенными за последнее слово)
Переворачиваем каждое слово в строке
Переворачиваем строку
"___hello__world)" i = 0  j = 1                      // <>- - обозначает пробел
"___hello__world)" i = 0 j = 2
"___hello__world)" i = 0 j = 3
"h___ello__world)" i = 1 j = 4
"he___llo__world)" i = 2 j = 5
"hel___lo__world)" i = 3 j = 6
"hell___o__world)" i = 4 j = 7
"hello_____world)" i = 5 j = 7
i = 6
и снова повторяем
*/

func reverseS(str []rune, i, j int) {
	for i < j {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}
}
func reverseWords(s string) string {
	i := 0
	j := 0
	end := 0
	str := []rune(s)

	for j < len(str) {
		for j < len(str) && str[j] == ' ' {
			j++
		}
		for j < len(str) && str[j] != ' ' {
			str[i], str[j] = str[j], str[i]
			j++
			end = i
			i++
		}
		i++
	}
	str = str[:end+1]

	i = 0
	j = 0
	for j <= len(str) {
		if j != len(str) && str[j] != ' ' {
			j++
		} else {
			reverseS(str, i, j-1)
			j++
			i = j
		}
	}
	reverseS(str, 0, len(str)-1)

	return string(str)
}

func main() {
	s := "  sun     dog snow     "
	s2 := reverseWords(s)
	fmt.Println(s2)
}
