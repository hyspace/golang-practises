package countInversions

func mergeAndCount(dest []int, left []int, right []int) int {
  length := len(dest)
  lenLeft := len(left)
  lenRight := len(right)

  tmp := make([]int, length)
  result := 0
  for i:=0, j:=0, k:=0; i < lenLeft && j < lenRight; ++k{
    if left[i] <= right[j]{
      tmp[k] = left[i]
      ++i
    } else {
      tmp[k] = right[j]
      result += lenLeft - i
      ++j
    }
  }
  if lenLeft

  copy(dest, tmp)
  return result
}

func Count(slice []int) int {
  length := len(slice)
  mid := length / 2

  left := slice[0:mid]
  right := slice[mid:length]

  return Count(left) + Count(right) + mergeAndCount(slice, left, right)
}