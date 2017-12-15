package util

import (
	"fmt"
	"sort"
)

func AarraySort(){
	bind := []int{1,3,4,23,6}
	//sort.Ints(bind) //升序
	sort.Sort(sort.Reverse(sort.IntSlice(bind))) //降序
	fmt.Println(bind)


	names := []string{"bigticket","allen","silang","meizi"}
	sort.Strings(names)
	fmt.Println(names)

	//sort.Sort()
}

func  BubbleSort(data  sort.Interface){
    r := data.Len()-1
    for i := 0; i < r ; i++{
        for j := r; j > i; j--{
            if data.Less(j, j-1){
                data.Swap(j, j-1)
            }
        }
    }
}
func InsertSort(data sort.Interface){
    r := data.Len()-1
    for i := 1; i <= r; i++{
        for j := i; j > 0 && data.Less(j, j-1); j--{
            data.Swap(j, j-1)
        }
    }
}
func  SelectSort(data sort.Interface){
    r := data.Len()-1
    for i := 0; i < r; i++{
        min := i
        for j:= i+1; j <= r; j++ {
            if data.Less(j, min) {  min = j }
        }
        data.Swap(i, min)
    }
}