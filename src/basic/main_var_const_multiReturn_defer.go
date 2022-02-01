package main_var_const_multiReturn_defer

import (
	"fmt"
	"something"
	"strings"
)

// fmt = formatting

var name3 = "juwon3"

// type 선언은 typeScript와 비슷함
// func multiply(a int, b int) int {
func multiply(a, b int) int {
	return a * b
}

// multiple return function
func lenAndUpper(name string) (int, string) {
	// https://pkg.go.dev/search?q=package&m=
	// len => String.length()
	// strings.ToUpper => upperCase()
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

// naked return
// go compiler가 미리 length int, uppercase string을 리턴할 것이라는걸 알아듣는다
// return type에 변수를 미리 선언하고 return 하는것을
// naked return 이라고 한다.
func lenAndUpper2(name string) (length int, uppercase string) {
	// defer는 function이 값을 return 하고나면 실행하는 코드다
	// 약간 promise의 .then .catch같은??
	defer fmt.Println("lenAndUpeer2 함수가 끝나고 실행하는 구문")

	// 위에서 리턴 구문에 length int가  var length int 이걸 의미함 이미 정해져있음
	//  length :=
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main_var_const_multiReturn_defer() {
	// function export => function을 대문자로 작성
	// export default를 하지않는다
	fmt.Println("Hello world!")

	something.SayHello()

	// 변수와 상수
	// unType
	// const name  = "juwon"

	// 상수
	// const name string = "juwon"
	// name = "jjjjjjjj"
	// fmt.Print(name)

	// 변수, type
	var name string = "juwon"
	name = "jjjjjjj"
	fmt.Println(name)

	// 같음 type 바인딩
	// :=  => 축약형 선언이라고한다
	// 오로지 func 안에서만 가능하고 변수( variable )에서만 가능하다
	name2 := "juwon"
	fmt.Println(name2)
	fmt.Println(name3)

	// ./main.go:44:2: too many arguments to return
	// 변수의 type을 정해주어야함 typeScript처럼
	fmt.Println(multiply(2, 2))

	// totalLength, upperName := lenAndUpper("juwon")
	// _로 variable을 무시하고 컴파일 할 수 있다
	totalLength, _ := lenAndUpper("juwon")
	// fmt.Println(totalLength, upperName)
	fmt.Println(totalLength)

	repeatMe("juwon1", "juwon2", "juwon3", "juwon4", "juwon5")

	totalLength2, upperCase := lenAndUpper2("juwon2")
	fmt.Println(totalLength2, upperCase)

}
