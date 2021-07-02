package main

import "fmt"

func main() {
	arr1 := [4]int{1,2,3,4}
	r:= len(arr1)
	l:= 0
	ans := 3
	for l<r {
		if arr1[l]==ans {
			break
		}
		var mid int
		mid = (r-l)>>1+l
		if arr1[mid]<=ans{
			l=mid
		}
		if arr1[mid]>ans {
			r=mid
		}
	}
	fmt.Println(l)
}
