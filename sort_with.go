package sort_with
import (
  "sort"
  "fmt"
)

type sortWrap struct {
  sl []interface{}
  fn func(interface{}, interface{}) bool
}

func (sw sortWrap) Len() int {
  return len(sw.sl)
}
func (sw sortWrap) Swap(i, j int) {
  s := sw.sl
  s[i], s[j] = s[j], s[i]
}
func (sw sortWrap) Less(i, j int) bool {
  return sw.fn(sw.sl[i], sw.sl[j])
}

type LessFn func(interface{}, interface{}) bool

func SortWith(v []interface{}, fn LessFn) {
  sort.Sort(sortWrap{v, fn})
}