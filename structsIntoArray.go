package main

import "fmt"

//Мы создаем новый кастомный тип, который строится на структуре "элемент", которая содержит значение элемента и ссылку на следующий.
type Element struct {
	number int
	next   *Element
}

//Структура "связанный список" содержит только голову.
type LinkedList struct {
	head *Element
}

//newElement - новый элемент.
//oldElement - уже существовавший элемент.

func (list *LinkedList) AppFirst(number int) {
	//newElement - этот элемент будет помещен в начало. & указывает о том, что мы обращаемся к адрессу памяти, где хранится Element. Т.е. newElement содержит ссылку.
	//Из-за этого вывести содержимое элементов встроенными фунциями печати не получится. Далее я напишу метод println, который выведет все содержимое списка подобно слайсу (расширяемому массиву)
	newElement := &Element{number, nil}
	//Если головы нет, новый элемент станет головой.
	if list.head == nil {
		list.head = newElement
	} else {
		//Если голова есть, её место займет новый элемент и изменится ссылка на следющий элемент. Т.е. следующий элемент не просто пустота (nil), а предыдющий элемент.
		//Далее почти весь код строится по похожему принципу.
		oldElement := list.head
		list.head = newElement
		newElement.next = oldElement
	}
}
func (list *LinkedList) AppEnd(number int) {
	newElement := &Element{number, nil}

	if list.head == nil {
		list.head = newElement
	} else {
		oldElement := list.head
		for oldElement.next != nil {
			oldElement = oldElement.next
		}
		oldElement.next = newElement
	}
}
func (list *LinkedList) InsertAt(number int, index int) {
	if index == 0 {
		list.AppFirst(number)
	} else if index == list.Len() {
		list.AppEnd(number)
	} else {
		newElement := &Element{number, nil}
		oldElement := list.head

		//Чтобы дойти до элемента, находящимся под индексом, нужно просмотреть каждый элемент до и считать их.
		for i := 0; i < index-1; i++ {
			oldElement = oldElement.next
		}
		newElement.next = oldElement.next
		oldElement.next = newElement
	}
}
func (list *LinkedList) DeletFirst() {
	if list.Len() == 1 {
		list.head = nil
		return
	}
	newElement := list.head
	list.head = newElement.next
}
func (list *LinkedList) DeleteEnd() {
	if list.Len() == 1 {
		list.head = nil
		return
	}
	newElement := list.head
	var oldElement *Element
	//Чтобы удалить последний элемент, необходимо прошагать по списку до него.
	for newElement.next != nil {
		oldElement = newElement
		newElement = newElement.next
	}
	oldElement.next = nil

}

func (list *LinkedList) DeleteElement(index int) {
	if index == 0 {
		list.DeletFirst()
	} else if index == list.Len()-1 {
		list.DeleteEnd()
	} else {
		newElement := list.head
		for i := 0; i < index-1; i++ {
			newElement = newElement.next
		}
		oldElement := newElement.next
		newElement.next = oldElement.next
	}
}

//Чтобы посчитать длину, необходимо пройтись по всем элементам.
func (list *LinkedList) Len() int {
	now := list.head
	lenght := 0
	for now != nil {
		lenght++
		next := now.next
		now = next
	}
	return lenght
}

//Чтобы реверсировать список, я создаю новый список и переношу из старого все элементы поэтапно в начало.
func (list *LinkedList) ReverseList() {
	var copyList LinkedList
	now := list.head
	next := now.next
	for now != nil {
		copyList.AppFirst(now.number)
		next = now.next
		now = next

	}
	list = &copyList
}
func (list *LinkedList) getElement(index int) int {
	if index == 0 {
		return list.head.number
	}
	now := list.head
	next := now.next
	for i := 0; i <= index; i++ {
		if i == index {
			return now.number
		}
		next = now.next
		now = next
	}
	return 0
}

//Вывод списка обременен некоторыми обстоятельствами. Т.к. "связыный список" содержит ссылки, при простом вызове метода Println
//вместо содержимого нам предоставят ссылки на "элементы", а  метод Printline напечатает содержимое.
func (list *LinkedList) Print() {
	if list.Len() == 0 {
		fmt.Println(`[]`)
		return
	}
	now := list.head
	next := now.next
	fmt.Print(`[`)
	count := 0
	for now != nil {
		count++
		next = now.next
		if now != nil {
			fmt.Print(now.number)
		}
		if count != list.Len() {
			fmt.Print(` `)
		}
		now = next
	}
	fmt.Print("]")
}
