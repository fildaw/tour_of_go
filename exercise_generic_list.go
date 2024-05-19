package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (list *List[T]) Append(element T) {
	list_copy_ptr := list
	for list_copy_ptr.next != nil {
		list_copy_ptr = list_copy_ptr.next
	}
	new_el := List[T]{nil, element}
	list_copy_ptr.next = &new_el
}

func NewList[T any](first_element T) (List[T]) {
	new_list := List[T]{nil, first_element}
	return new_list
}

func (list List[T]) String() string {
	output := "["
	list_p := &list
	for list_p != nil {
		output += fmt.Sprintf("%v", list_p.val)
		if list_p.next != nil {
			output += ", "
		}
		list_p = list_p.next
	}
	output += "]"
	return output
}

func main() {
	a := NewList("x")
	a.Append("d")
	a.Append("d")
	fmt.Println(a)
}
