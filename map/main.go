package main

import "fmt"

type bookmarkMap = map[string]string

func main() {

	bookmarks := bookmarkMap{}
	fmt.Println("Приложение закладок")
	for {
		variant := getMenu()
		switch variant {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			addBookmark(bookmarks)
		case 3:
			deleteBookmark(bookmarks)
		case 4:
			return
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Выберете вариант")
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
	fmt.Scan(&variant)
	return variant
}

func printBookmarks(bookmarks bookmarkMap) {
	if len(bookmarks) == 0 {
		fmt.Println("Пока нет закладок")
	}
	for key, value := range bookmarks {
		fmt.Println(key, ":", value)
	}
}
func addBookmark(bookmarks bookmarkMap) {
	var name, address string
	fmt.Print("Введите название электронного ресурса: ")
	fmt.Scan(&name)
	fmt.Print("Введите адрес электронного ресурса: ")
	fmt.Scan(&address)
	bookmarks[name] = address
}
func deleteBookmark(bookmarks bookmarkMap) {
	var name string
	fmt.Print("Введите название электронного ресурса: ")
	fmt.Scan(&name)
	delete(bookmarks, name)
}
