package main
import (
	"fmt"
	"slices"
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

	//Make function to create slice
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

	//Copying slice
	var nums4=make([]int,0,5)
	nums4 =append(nums4,2)
	var nums5 =make([]int,len(nums4))

	//copy function
	copy(nums5,nums4) //copies elements from source to destination
	fmt.Println(nums4,nums5)

	//slice operators
	var nums7=[]int{1,2,3,4,5}
	fmt.Println(nums7[1:4]) //4th index is excluded and 1st index is included
	fmt.Println(nums7[1:])

	//slice
	var nums8=[]int{1,2,3}
	var nums9 = []int{1,2,3}
	fmt.Println(slices.Equal(nums8,nums9)) //true
	
	var numss = [][]int{{1,2},{3,4}}
	fmt.Println(numss)
}