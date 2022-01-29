package something

import "fmt"

// 함수의 앞이 소문자인 경우 private 접근제한자가 생김
func sayBye() {
	fmt.Println("Bye")
}

// 함수의 앞이 대문자인 경우 export가 자동으로됨
func SayHello() {
	fmt.Println("Hello")
}
