package main

import "fmt"

// 메모리 주소를 보는법과
// 어떤것이 메모리 주소에 접근하도록하는 방법

func main() {
	// low level programming
	// go의 특징 중 하나 :
	// hight level programming 기법으로
	// low level programming이 가능하게 해줌
	a := 2
	// b := a
	b := &a
	// value 복사임
	a = 10
	// javascript의 깊은복사와 얕은복사를 의미하나?
	// fmt.Println(a, b)

	// 0xc0000140d0 0xc0000140d8 메모리 주소를 볼 수 있다
	fmt.Println(&a, &b)

	// 0xc0000b8010 0xc0000b8010 메모리 주소를 복사함
	fmt.Println(&a, b)

	// 여기서 b 는 a의 메모리 주소에 연결되어있어 a의 값을 참조한다.

	// 살펴보거나 훑어보는 느낌의 *
	fmt.Println(*b)

	// pointer를 사용하는 이유는 값 복사 대신에
	// 메모리에 저장된 Object( 객체 )를 서로 똑같이 가지고 싶도록
	// 원할 때 사용한다 (=> javascript의 [...Object] )

	// 결과적으로 이러한 메모리 주소 복사 방법의 코드가
	// 프로그램을 더 빠르게 동작할 수 있게하는 기반을 만들어준다.
	// 즉 계속 value 복사를 하게되면 느려짐 (=> 다른 메모리 주소에 같은 값 복사)

	// 깊은 복사의 상태에서 b를 가지고 a의 값 변경이 가능하다.
	*b = 20
	fmt.Println(a)

	// 이것들이 바로 go에서의 low level programming 기법이다.

	// 결론 & => address
	// * => 살펴보는것 ( see through )
	// *를 사용해서 주소를 살펴보면 주소에 담긴 값을 볼 수 있다.
	// 또한 주소를 갖고있는 변수는 주소에 담긴 값을 변경할 수 있다.

	// 이게 Go에서 a를 poiner b를 사용해 값을 변경할 수 있는 이유
}
