package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования)


*/

type Human struct {
	ID   string
	Rank int
	Name string
}

func (h *Human) PrintSomething() {
	fmt.Println("Меня можно вызвать из любых пакетов")
}

func (h *Human) printSomething() {
	fmt.Println("меня можно вызвать только из моего пакета, из других нельзя")
}

type Action struct {
	Human
	ID string // ID Action'а будет затенять ID Human'а
}

func (a *Action) PrintSomething() {
	fmt.Println("Я затеняю метод PrintSomething у Human'а")
}

func main() {
	hum := Human{
		ID:   "aaa",
		Rank: 1,
		Name: "Vasya",
	}
	hum.printSomething()
	hum.PrintSomething()
	a := Action{
		Human: hum,
		ID:    "opa",
	}
	a.PrintSomething() // метод Action'а
	a.printSomething() // метод Human'а
	fmt.Printf("ID Action %s\n", a.ID)
	fmt.Printf("ID Human %s\n", a.Human.ID) // явно указываем чье поле хоти получить
	a.Human.PrintSomething()                // явно вызываем метод от Human'а, если хотим его метод
}
