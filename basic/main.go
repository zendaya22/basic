package main

import "fmt"

// bubble sort

// func bubbeleSort(){
// 	for i := 0; i < len(num) - 1 ; i++{
// 		for j := 0; j < len(num) - 1; j++{
// 			if num[i] > num[i + 1] {
// 				num[i], num[i + 1] = num[i + 1], num[i]
// 			}
// 		}
// 	}
// }

// insertionsort
// func selectionSort(num []int){
// 	for i := 0; i < len(num); i++{

// 		for k := i; k > 0; k--{
// 			log.Println(num[k - 1], num[k])
// 			if num[k - 1] > num[k] {
// 				temp := num[k]
// 				num[k] = num[k - 1]
// 				num[k - 1] = temp
// 			}
// 		}
// 	}
// }
// func linearSearch(value int, nums ...int) (found bool, index int) {

// 	for i, num := range nums {
// 		if num == value {
// 			found = true
// 			index = i
// 			return
// 		}

// 	}

// 	index = -1
// 	return

// }

// func BinarySearch(find int, nums ...int) (found bool) {
	// 	if len(nums) <= 1 {
	// 		if nums[0] == find {
	// 			found = true
	// 		}
	// 	}
	
	// 	for {
	// 		middleIndex := len(nums) / 2
	
	// 		if nums[middleIndex] == find {
	// 			return true
	// 		} else if find < nums[middleIndex] {
	
	// 			nums = nums[:middleIndex]
	// 		} else if find > nums[middleIndex] {
	// 			nums = nums[middleIndex:]
	// 		}
	
	// 		if middleIndex == 0 {
	// 			break
	// 		}
	
	// 	}
	// 	return false
	// }
	

func main() {

	num := []int{1, 2, 3, 4, 5, 6}

}

// 