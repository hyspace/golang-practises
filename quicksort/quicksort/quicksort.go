package quicksort

import(
  "math/rand"
  "time"
)

func swap(i *int, j *int) {
  temp := *i
  *i = *j
  *j = temp
}

func sort(slice []int) {
  length := len(slice)

  r := rand.Intn(length)
  swap(&slice[r], &slice[0])
  i := 0
  for j := 1; j < length; j++ {
    if slice[j] < slice[0]{
      swap(&slice[j], &slice[i + 1])
      i++
    }
  }
  swap(&slice[0], &slice[i])
  if i > 1{
    sort(slice[0 : i])
  }
  if length - (i + 1) > 1{
    sort(slice[i + 1 : length])
  }
}

func Sort(slice []int) {
  rand.Seed(time.Now().UnixNano())
  sort(slice)
}