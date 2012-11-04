package main
import(
  "fmt"
  "sort"
)
type sortCriticsMapWrap struct {
  Keys_list [3]string
  Values_list [3]float64
  fn func(interface{}, interface{}) bool
}

func (sw sortCriticsMapWrap) Len() int {
  return len(sw.Values_list)
}
func (sw sortCriticsMapWrap) Swap(i, j int) {
  sw.Values_list[i], sw.Values_list[j] = sw.Values_list[j], sw.Values_list[i]
  sw.Keys_list[i], sw.Keys_list[j] = sw.Keys_list[j], sw.Keys_list[i]
}
func (sw sortCriticsMapWrap) Less(i, j int) bool {
  return sw.fn(sw.Values_list[i], sw.Values_list[j])
}

type LessFn func( a interface{} , b interface{} ) bool

func SortCriticMapWith(v interface{}, fn LessFn)( scmw sortCriticsMapWrap ){
  map_obj := v.(map[string]float64)
  var vl [3]float64
  var kl [3]string
  sc_index := 0
  for k , val := range map_obj{
    kl[sc_index] = k
    vl[sc_index] = val
    sc_index += 1
  }
  scmw = sortCriticsMapWrap{kl , vl , fn}
  sort.Sort(scmw)
  return
}
func main() {
  m := map[string]float64{ "j": 1.0 , "b": 0.5 , "c" : 2 }
  sc_list := SortCriticMapWith(m , func(a, b interface{})bool{
    fmt.Printf("%v\n" , [...]float64{a.(float64) , b.(float64)})
    fmt.Printf("%v\n" , a.(float64) < b.(float64))
    return a.(float64) < b.(float64)
  })
  fmt.Printf("%v\n" , sc_list.Values_list)
  fmt.Printf("%v\n" , sc_list.Keys_list)
  s := []int{5, 2, 6, 3, 1, 4} // unsorted
  sort.Ints(s)
  fmt.Println(s)
}