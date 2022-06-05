package usefulutilpackages

import (
	"fmt"
	"sort"
)

func sortingForPrimitiveDataType() {

	//Using the SORT package for the primitive data types
	str := []string{"apple", "oranges", "water melon", "mangoes"}

	sort.Strings(str)
	fmt.Println("Sorted string using sort package ==>", str)

	//In reverse order
	sort.Sort(sort.Reverse(sort.StringSlice(str)))
	fmt.Println("Reverse Sorted string using sort package ==>", str)

	//Interger data type
	intArray := []int{8, 9, 10, 2, 11, 99, 1}

	sort.Ints(intArray)
	fmt.Println("Sorted int using sort package ==>", intArray)

	sort.Sort(sort.Reverse(sort.IntSlice(intArray)))

	//Float data type

	floatArray := []float64{99.0, 2.0, 66, 11.12, 82.0, 0.0}

	sort.Float64s(floatArray)
	fmt.Println("Sorted float64 using sort package ==>", floatArray)

}

//The differnce between sort.Slice() vs sort.SliceStable() is that the latter once will keep the original order of
//equal elements while the former one may not
func sortingUsingComparatorFunction() {
	//Custom sort
	arr := [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}

	// Using sort.Slice
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})

	//Using sort.SliceStable
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})

}
