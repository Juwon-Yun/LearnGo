package main_if_switch

import "fmt"

func canIDrink(age int) bool {

	// koreanAge := age + 2

	//variable expression go의 변수 표현식
	// if koreanAge := age - 2; koreanAge < 20 {

	// 조건을 체크하기전에 variable( 변수 )를 만들 수 있다
	// if문에서만 사용하는 변수를 말한다
	if koreanAge := age - 2; koreanAge < 20 {
		return false
	}
	return true
}

func switchCanIDring(age int) bool {
	// switch koreanAge := age - 2; koreanAge {
	switch koreanAge := age - 2; {
	case koreanAge < 18:
		return false
	case koreanAge > 50:
		return false
	}
	return true
}

func mamain_if_switch() {
	fmt.Println(canIDrink(22))
	fmt.Println(switchCanIDring(19))
}
