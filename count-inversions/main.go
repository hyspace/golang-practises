package main

import(
  "./countInversions"
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func bruteForce(slice []int) int {
  length := len(slice)
  count := 0
  for i := 0; i < length; i++ {
    num := slice[i]
    for j := i + 1; j < length; j++ {
      if slice[j] < num {
        count++
      }
    }
  }
  return count
}

func test(){
  array1 := []int{6,10,13,5,8,2,11,15,4,25,21,5,73,1,46,3,6,31,5,7}
  array2 := make([]int, len(array1))
  copy(array2, array1)
  count := countInversions.Count(array1)
  fmt.Println(count)
  fmt.Println(bruteForce(array2))
}

func main(){
  file, err := os.Open("./IntegerArray.txt")
  if err != nil {
    panic(err)
  }
  defer file.Close()

  var slice []int

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    integer, err := strconv.Atoi(scanner.Text())
    if err != nil {
      panic("convert failed")
    }
    slice = append(slice, integer)
  }
  if scanner.Err() != nil {
    panic("error happened when reading file")
  }
  fmt.Printf("length of input: %d\n", len(slice))
  count := countInversions.Count(slice)
  // count := bruteForce(slice)
  fmt.Println(count)

}