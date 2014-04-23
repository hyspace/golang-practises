package main

import(
  "./quicksort"
  "fmt"
)

func main(){
  array := []int{6,10,13,5,8,2,11,15,4,25,21,5,73,1,46,3,6,31,5,7}
  fmt.Println(array)
  quicksort.Sort(array)
  fmt.Println(array)
}