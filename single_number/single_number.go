package single_number

import "fmt"

func singleNumber(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			delete(m, num)
		} else {
			m[num] = true
		}
	}
	for key, _ := range m {
		return key
	}
	return 0
}

func singleNumberV2(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

func main() {
	// 位运算
	fmt.Println(1 & 2)
	fmt.Println(1 | 2)
	fmt.Println(1 | 1)
	fmt.Println(1 ^ 2)

}