// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
package main

import "fmt"

// Human является базовым типом
type Human struct {
	Name string
}

func (human *Human) Work() string {

	return "I'm working"
}

// Механизм встраивания является аналогом наследования традиционных объектно-ориентированных языков
// Указав анонимное поле у структуры, появляется возможность использования полей и методов встроенного типа
type Action struct {
	Name string
	*Human
}

func (action *Action) Do() string {

	return fmt.Sprintf("%s is %s", action.Human.Name, action.Name)
}

func main() {

	human := &Human{
		Name: "CoolName",
	}

	sleepAction := &Action{
		Name:  "sleeping",
		Human: human,
	}

	fmt.Println(human.Work()) // вызов метода структуры Human

	fmt.Println(sleepAction.Human.Work()) // явный вызов метода встроенной структуры Human

	fmt.Println(sleepAction.Work()) // неявный вызов возможен в случаях, когда нет конфликта имен
	// в ином случае обязательно использовать явный вызов, иначе по умолчанию вызовается метод, который "ближе"

	fmt.Println(sleepAction.Do()) // вызов методы структуры Action

	// аналогичная логика и с полями структур
	fmt.Println(human.Name) // вывод поля структуры Human

	fmt.Println(sleepAction.Human.Name) // явный вывод, выведется имя структуры Human

	fmt.Println(sleepAction.Name) // неявный вывод не сработает, так как названия совпадают, выведется имя структуры Action
}
