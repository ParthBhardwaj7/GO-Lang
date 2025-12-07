package main
import (
	"fmt"
)	
func main() {
	//slice is dynamic in size
	//most used data structure in go lang
	//slice is built on array
	//slice is reference type
	//zeroed value of slice is nil
	var nums []int
	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(nums==nil)

	var nums1 =make([]int,0,5) //length 5
	//capacity -> Maximum numbers of elements can fit
	fmt.Println(nums1)
	fmt.Println(cap(nums1))
	nums1 = append(nums1,1)
	fmt.Println(nums1)
	fmt.Println(cap(nums1)) // Capacity double ho jata hai jab size exceed hota hai
	nums1 =append(nums1,2,3,4)
	fmt.Println(nums1)
	fmt.Println(cap(nums1))

	nums2 := []int{}
	fmt.Println(nums2)
	fmt.Println(len(nums2))
	fmt.Println(cap(nums2))

	
}