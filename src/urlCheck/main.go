package main

import (
	"fmt"
	"time"
)

// main 함수는 결과를 저장하는 곳이다

// channel은 goroutine과 main 함수 사이에 정보를 전달하기 위한 방법이다
// goroutine과 goroutine 끼리 소통도 가능하다

// channel은 파이프같다
// 메세지를 보내고 받고 하는 기능

func main() {

	// go routines 작동 => 병렬형 프로그래밍
	// go sexyCount("juwon1")
	// go sexyCount("juwon2")

	// gorouine 함수밖에 존재하지않으면 main함수는 go만 실행하고 바로 종료된다
	// 따라서 go만 실행하는 코드만 있다면 그냥 종료하는것처럼 보임

	// 기다려주지않는다

	// go sexyCount("juwon2")

	// time.Sleep(time.Second * 5)

	// 변수명 := make(chan returnType)
	channel := make(chan bool)
	people := [2]string{"juwon1", "juwon2"}
	for _, person := range people {
		go isSexy(person, channel)

		// 변수에 담는건 불가능
		// resule := go isSexy(person)
	}
	// time.Sleep(time.Second * 10)

	// goroutine의 메세지를 channel로 통해 받은것을 result 변수에 저장함
	// 이 경우에는 goroutine 함수만 있어도 main이 종료되지 않는다
	// result := <-channel
	// fmt.Println(result)

	// 바로 받는것과 두개 이상으로 받는것이 가능함
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	// 2개의 고루틴이 채널을 통해 메세지를 보내는것을 알고있으면 2개만 받는다
	// 3개 이상의 고루틴이 채널을 통해 메세지를 보내면 예외처리하고 main을 종료함
	// fatal error : all goroutines are asleep - deadlock!
}
func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

// func isSexy(person string ) bool {
func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	// channel을 이용해서 메세지를 보낸다
	c <- true
	// return true
}
