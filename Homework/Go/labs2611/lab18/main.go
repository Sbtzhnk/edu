// todo: Заданы массивы
//
// Даны читатели книг
// var readers_books = [...] string {'id3', 'id5', 'id9', 'id8', 'id2', 'id1' }
//
// Даны читатели газет
// var readers_magazines = [...] string { 'id8', 'id2', 'id1', 'id4', 'id6', 'id7', 'id10'}
// Найти пользователей кто читает и книги и газеты
package main

import (
	"fmt"
)

func main() {
	var readers_books = [...]string{"id3", "id5", "id9", "id8", "id2", "id1"}
	var readers_magazines = [...]string{"id8", "id2", "id1", "id4", "id6", "id7", "id10"}
	var common_readers = []string{}

	for _, rbVal := range readers_books {
		isCommon := false
		for _, rmVal := range readers_magazines {
			if rbVal == rmVal {
				isCommon = true
			}
		}
		if isCommon == true {
			common_readers = append(common_readers, rbVal)
		}
	}
	fmt.Println(common_readers)
}
