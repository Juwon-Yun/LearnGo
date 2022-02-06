package main

import "fmt"

type person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *person {
	person := person{name: name, age: age}
	return &person
}

// 메소드 정의 (메소드 변수명 메소드명) 함수명(){...}
func (p person) introduce() string {
	return fmt.Sprint("name : ", p.name, " Age : ", p.getAge())
}

func (p *person) setName(name string) {
	p.name = name
}

func (p person) setAge(age int) {
	p.age = age
}

func (p person) getAge() int {
	return p.age
}

func main() {
	person1 := NewPerson("juwon1", 1)
	fmt.Println(person1.introduce())

	var person2 = person{}
	person2.setAge(2)
	person2.setName("juwon2")
	fmt.Println(person2)

	person3 := person{name: "juwon3", age: 3}
	person3.setAge(4)
	fmt.Println(person3)
}
