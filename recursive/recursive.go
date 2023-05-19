package recursive

// GetMax: T(N) = 2T(N/2) + O(1) => O(N)
func GetMax(arr []int, left, right int) int {
	if left == right {
		return arr[left]
	}
	mid := (left + right) / 2
	leftMax := GetMax(arr, left, mid)     // 子过程
	rightMax := GetMax(arr, mid+1, right) // 子过程
	if leftMax > rightMax {
		return leftMax
	}
	return rightMax
}

/*
归并排序的扩展
Question：
 1. 小和问题，在一个数组中，每一个数左边比当前数小的数累加起来，叫做这个数组的小和
    例子：[1, 3, 4, 2, 5]
         1左边比1小的数，没有
         3左边比3小的数，1
         4左边比4小的数，1,3
         2左边比2小的数，1
         5左边比5小的数，1,3,4,2
        小和为 1+(1+3)+1+(1+3+4+2)
 2. 逆序对问题，在一个数组中，左边的数如果比右边的数大，则这两个数构成一个逆序对
*/

// todo
func SmallSum() {

}

// todo
func ReverseOrderPair() {

}

/*
快排Partition(分治思想)的应用
 1. 给定一个数组arr和一个数num，把小于等于num的数放在左边，大于num的数放在数组右边；额外空间复杂度O(1)，时间复杂度O(N)
 2. 荷兰国旗问题，给定一个数组arr和一个数num，把小于num的数放在左边，等于num的数放中间，大于num的数放在数组右边；额外空间复杂度O(1)，时间复杂度O(N)
*/
