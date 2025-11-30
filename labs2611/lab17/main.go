// todo: Заданы массивы
// все пользователи
// var all_users = [...]string {'id3', 'id5', 'id9', 'id8', 'id2', 'id1', 'id4', 'id6', 'id7', 'id10'}
// пользователи не в сети
// var offline_users = [...]string  {'id3', 'id9', 'id7', 'id2', 'id4', 'id6'}
// Вычислить пользователей online
package main

import (
	"fmt"
)

func main() {
	var all_users = [...]string{"id3", "id5", "id9", "id8", "id2", "id1", "id4", "id6", "id7", "id10"}
	var offline_users = [...]string{"id3", "id9", "id7", "id2", "id4", "id6"}
	var online_users []string
	for _, allVal := range all_users {
		isOnline := true
		for _, offVal := range offline_users {
			if allVal == offVal {
				isOnline = false
			}
		}
		if isOnline {
			online_users = append(online_users, allVal)
		}
	}
	fmt.Println(online_users)
}
