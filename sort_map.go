package sort_map
import (
  "sort"
  "fmt"
)

type sortWrap struct {
  sl [interface{}]interface{}
  fn func([]interface{}, []interface{}) bool
}

func (sw sortWrap) Len() int {
  return len(sw.sl)
}
func (sw sortWrap) Swap(i, j interface{}) {
  s := sw.sl
  s[i], s[j] = s[j], s[i]
}
func (sw sortWrap) Less(i, j interface{}) bool {
  return sw.fn(sw.sl[i], sw.sl[j])
}

type LessFn func(interface{}, interface{}) bool

func SortWith(v [interface{}]interface{}, fn LessFn) {
  sort.Sort(sortWrap{v, fn})
}