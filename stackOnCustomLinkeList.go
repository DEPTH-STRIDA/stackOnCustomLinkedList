package main

import "fmt"

//Создание структы "стэк", которая содержит другую структуру.
type stack struct {
	list LinkedList
}

//Создание методов для работы со "стэком"

//Добавление в конец стэка элемента.
func (stack *stack) push(number int) {
	stack.list.AppEnd(number)
}

//Получение из конца элемента с его удалением из стэка.
func (stack *stack) pop() int {
	n := stack.list.getElement(0)
	stack.list.DeletFirst()
	return n
}
func main() {
	//Из-за того что стэк основан на списке хранящем ссылки, то вывод делаю
	//с помощью метода привязанного к структуре списка.
	var myStack stack
	fmt.Println(`Стэк в начале: `, myStack)

	myStack.push(1)
	fmt.Print(`Стэк после добавления одного элемента `)
	myStack.list.Print()
	fmt.Println()

	myStack.push(2)
	myStack.push(3)
	fmt.Print(`Стэк после добавления двух элементов `)
	myStack.list.Print()
	fmt.Println()

	myStack.pop()
	fmt.Print(`Стэк после удаления одного элемента `)
	myStack.list.Print()
	fmt.Println()

}
