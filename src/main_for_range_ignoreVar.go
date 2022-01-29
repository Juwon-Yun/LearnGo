package main_for_range_ignoreVar

import "fmt"

func superAdd(numbers ...int) int {
	total := 0
	// range는 무조건 for 구문 안에서만 loop동작을 할 수 있게 해준다

	// for number := range numbers {
	for index, number := range numbers {
		// for each, for in과 같다
		fmt.Println(index, number)
		// 0~5 까지 출력됨 => range keyword가 index를 부여함
	}
	fmt.Println()
	// 아래와 같은 코드로 작성할 수 있다.
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i], i)
	}
	fmt.Println()

	for _, number := range numbers {
		fmt.Println(number)
		total += number
	}
	fmt.Println()
	return total
}

func main_for_range_ignoreVar() {

	result := superAdd(1, 2, 3, 4, 5, 6)

	fmt.Println(result)

}
