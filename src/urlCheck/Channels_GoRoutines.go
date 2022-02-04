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

func Channels_GoRoutines() {

	// go routines 작동 => 병렬형 프로그래밍
	// go sexyCount("juwon1")
	// go sexyCount("juwon2")

	// gorouine 함수밖에 존재하지않으면 main함수는 go만 실행하고 바로 종료된다
	// 따라서 go만 실행하는 코드만 있다면 그냥 종료하는것처럼 보임

	// 기다려주지않는다

	// go sexyCount("juwon2")

	// time.Sleep(time.Second * 5)

	// 변수명 := make(chan returnType)
	// channel := make(chan bool)
	// people := [2]string{"juwon1", "juwon2"}
	// for _, person := range people {
	// go isSexy(person, channel)

	// 변수에 담는건 불가능
	// resule := go isSexy(person)
	// }
	// time.Sleep(time.Second * 10)

	// goroutine의 메세지를 channel로 통해 받은것을 result 변수에 저장함
	// 이 경우에는 goroutine 함수만 있어도 main이 종료되지 않는다
	// result := <-channel
	// fmt.Println(result)

	// 바로 받는것과 두개 이상으로 받는것이 가능함
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)

	// 2개의 고루틴이 채널을 통해 메세지를 보내는것을 알고있으면 2개만 받는다
	// 3개 이상의 고루틴이 채널을 통해 메세지를 보내면 예외처리하고 main을 종료함
	// fatal error : all goroutines are asleep - deadlock!

	channel2 := make(chan string)
	people2 := []string{"juwon3", "juwon4", "juwon5", "juwon6", "juwon7", "juwon8", "juwon9"}
	for _, person := range people2 {
		go isSexy2(person, channel2)

		// 변수에 담는건 불가능
		// resule := go isSexy(person)
	}
	// resultOne := <-channel2
	// resultTwo := <-channel2
	// goroutine이 2개인데 3개를 recive하면 에러
	// fatal error: all goroutines are asleep - deadlock!
	// resultThree := <-channel2
	// fmt.Println("Waiting for meesage")
	// fmt.Println("Received this message : ", resultOne)
	// fmt.Println("Received this message : ", resultTwo)
	// fmt.Println("Received this message : ", resultThree)

	for i := 0; i < len(people2); i++ {
		fmt.Println("Waiting for : ", i)
		// 메세지를 받는 것은 Blocking Operation이라 한다.
		// 메세지를 받기 전까지 main함수를 멈추는 것
		fmt.Println(<-channel2)
	}
	// 동시성( Concurrency )
	// fmt.Println(<-channel2)
	// fmt.Println(<-channel2)
	// fmt.Println(<-channel2)
	// fmt.Println(<-channel2)
	// fmt.Println(<-channel2)
	// fmt.Println(<-channel2)
	// fmt.Println(<-channel2)

	// main 함수가 끝나면 go routine은 무의미해진다.
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

func isSexy2(person string, c chan string) {
	time.Sleep(time.Second * 5)
	// channel을 이용해서 메세지를 보낸다
	c <- person + " is sexy"
	// return true
}
