package main

import "fmt"

// 35.搜索插入位置
func searchInsert(nums []int, target int) int {
    var len int = len(nums)
    var left int = 0
    right := len-1
    for left <= right{
        middle := left + (right-left)/2
        if target < nums[middle]{
            right = middle - 1
        } else if target > nums[middle]{
            left = middle + 1
        } else{
            return middle
        }
    }
    return right+1
    
}


func main() {
	nums := []int{1, 3, 5, 7}
	target := 6

	// 调用函数并获取结果
	insertionPoint := searchInsert(nums, target)

	fmt.Println("====================================")
	fmt.Printf("函数计算出的插入点是: %d\n", insertionPoint)
	fmt.Println("====================================")
	fmt.Println("\n解读：对于数组 [1, 3, 5, 7]，如果要插入 6，应该放在 5 和 7 之间，也就是索引为 3 的位置。")
}
