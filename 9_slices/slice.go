package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic (dynamic array )
// most used construct in go
// + useful methods
func main(){
	// unintialized slice value are nil (like null in java )
	// var nums[] int // creates an empty slice nil
	// fmt.Println(nums)
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// so use this synatx for slice 

	// var nums = make([]int , 2 , 5) // creates a integer slice (dynamic array of 2 integer values with default values  and capacity 5 so itially only are telling preallocate 5 as capapcity as we know )

	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	// nums = append(nums, 5)
	// nums = append(nums, 6)
	// nums = append(nums, 7)
	// nums = append(nums, 8)
	// nums = append(nums, 9)

	// fmt.Println(nums)
	// fmt.Println(cap(nums)) // prints the capacity of the slice 
	// fmt.Println(len(nums))

	// fmt.Println()


	// var nums2 = make([]int , 0 , 5); // create a dynamic array (slice in go) having intial capacity to store 5 element after which capcacity doubling happend dynamically as you keep adding elements 
	// nums2 = append(nums2, 1)
	// nums2 = append(nums2, 2)

	// fmt.Println(nums2)
	// fmt.Println(len(nums2))

	// fmt.Println(cap(nums2))
	
	// fmt.Println()

	// nums4 := []int{}
	// nums4 = append(nums4, 1)
	// nums4 = append(nums4, 2)
	// nums4 = append(nums4, 3)
	// fmt.Println(nums4)
	// fmt.Println(len(nums4))
	// fmt.Println(cap(nums4))

	var nums5 = make([]int , 0 , 5);
	nums5 = append(nums5, 12)

	var num6 = make([]int , len(nums5))
	fmt.Println(nums5 , num6)

	// copy(dest , src) if dest length is 0 then copy wont have
	copy(num6 , nums5)
	
	fmt.Println(nums5 , num6)


	// slice operator

	var nums = []int{1 , 2 , 3}

	fmt.Println(nums[0:])

	fmt.Println(nums[1:])

	fmt.Println(nums[2:])
	fmt.Println(nums[:3])
	fmt.Println(nums[:2]) // after collan end value is (non-inclusive)


	// slices (package)

	var num48 = []int{1, 2 , 3}
	// var num49 = []int{1, 2 , 3}
	var num49 = []int{1, 2 , 4}
	fmt.Println(slices.Equal(num48 , num49))

	// 2d slice 

	var twodnums = [][]int{{1 ,2 , 3} , {2 , 3 , 4, 5}}
	fmt.Println(twodnums)
}