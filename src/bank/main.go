package main

import (
	"bank/accounts"
	"fmt"
	"log"
)

func main() {
	account := accounts.NewAccout("juwon")
	// &{juwon 0} 복사본이 아닌 Object라는 뜻의 &
	// 이제는 account.owner, account.balance로 접근하지 못한다
	account.Deposit(10)
	err := account.Whithdraw(20)
	if err != nil {
		fmt.Println(err)

		// PrintLn을 하고 프로그램을 종료함
		log.Fatal(err)
	}
	fmt.Println(account.Balance())
}
