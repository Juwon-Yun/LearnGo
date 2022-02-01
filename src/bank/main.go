package main

import (
	"bank/dict"
	"fmt"
)

func main() {
	// account
	// account := accounts.NewAccout("juwon")
	// &{juwon 0} 복사본이 아닌 Object라는 뜻의 &
	// 이제는 account.owner, account.balance로 접근하지 못한다
	// account.Deposit(10)
	// err := account.Whithdraw(20)
	// if err != nil {
	// fmt.Println(err)
	// log를 PrintLn을 하고 프로그램을 종료함
	// log.Fatal(err)
	// }
	// fmt.Println(account.Balance(), account.Onwer())
	// fmt.Println(account.String())

	// dict

	// type 에도 struct 뿐만아니라 method를 추가할 수 있다
	// dictionary := dict.Dictionary{"first": "First word"}
	// definition, err := dictionary.Search("first")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	// dictionary["hello"] = "hello"
	// fmt.Println(dictionary)
	// go에서는 map의 key 여부값을 알수있는 기능이있다.

	// Add하고 나서 Search하고 또 Add
	// dictionary := dict.Dictionary{}
	// word := "hello"
	// definition := "Greeting"
	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// hello, _ := dictionary.Search(word)
	// fmt.Println("found", word, "definition ", hello)

	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)

	// }

	// update
	// dictionary := dict.Dictionary{}
	// baseWord := "hello"
	// dictionary.Add(baseWord, "First")

	// err := dictionary.Update(baseWord, "Second")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// word, _ := dictionary.Search(baseWord)
	// fmt.Println(word)

	dictionary := dict.Dictionary{}
	baseword := "Bye"
	dictionary.Add(baseword, "First")
	dictionary.Search(baseword)
	dictionary.Delete(baseword)
	word, err := dictionary.Search(baseword)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}

}
