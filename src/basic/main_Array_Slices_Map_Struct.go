package basic

import "fmt"

func main_Array_Slice() {
	// names := [5]string{"juwon1", "juwon2", "juwon3"}
	// 길이제한 없애기
	names := []string{"juwon1", "juwon2", "juwon3"}

	// go 에서는 이런 식으로는 goroutine을 발생시킴
	// names[3] = "alalla"
	// names[4] = "alalla"

	// 크리와 범위를 정했기 때문에 out of bounds 에러 발생함
	// names[5] = "alalla"

	// 길이제한을 없앤 후에 요소 추가는 append() 한다
	// unshift도 있나?

	// 새로운 slice를 return 한다.
	names = append(names, "juwon5")

	// 없다
	// names = unshift(names, "unshift juwon")

	fmt.Println(names)

}

// struct => object와 class가 합쳐진 개녕

func main_Map() {
	// charactor 형태의 map
	// 		 map[key type]value type{key:value}
	juwon := map[string]string{"name": "juwon", "age": "12"}
	fmt.Println(juwon)

	// map도 for range로 loop 동작이 가능하다
	for _, value := range juwon {
		fmt.Println(value)
	}
}

type person struct {
	name    string
	age     int
	favFood []string
}

func main_Array_Slices_Map_Struct() {
	// 이와 같은 것들을 선언하고싶으면?
	// {
	// 	"name" : "juwon",
	// 	"age" : 30,
	// 	"favFood" : ["string"]
	// }
	favFood := []string{"kimchi", "ramen"}
	favFood2 := []string{"cake", "bread"}

	// struct를 만드는 방법 2개
	juwon := person{"juwon", 30, favFood}
	// key 값을 하나라도 명시했으면 다른 요소들도 필수로 명시해줘야한다.
	juwon2 := person{name: "juwon2", age: 55, favFood: favFood2}
	fmt.Println(juwon)
	fmt.Println(juwon2)

	// go에서는 생성자 메소드가 없다
	// javascript constructor
	// python __init__
	// java new Name();

	// struct는 없다
	// 그래서 스스로 생성자를 실행 시켜야함

	// 아직 설명하지 않은것들 channel, go routine
}
