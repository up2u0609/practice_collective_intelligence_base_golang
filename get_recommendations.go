package main
import(
  "fmt"
  "math"
  "sort"
  "pearson_correlation"
  "euclidean_distance"
)

type sortCriticsMapWrap struct {
  Keys_list []float64
  Values_list []float64
  fn func(float64, float64) bool
}

func (sw sortCriticsMapWrap) Len() int {
  return len(sw.Values_list)
}
func (sw sortCriticsMapWrap) Swap(i, j int) {
  vl := sw.Values_list
  kl := sw.Keys_list
  vl[i], vl[j] = vl[j], vl[i]
  kl[i], kl[j] = kl[j], kl[i]
}
func (sw sortCriticsMapWrap) Less(i, j int) bool {
  return sw.fn(sw.Values_list[i], sw.Values_list[j])
}

type LessFn func( float64 , float64 ) bool

func SortCriticMapWith(v map[string]float64, fn LessFn) {
  v_l := 0
  for k, _ := range v{
    v_l += 1
  }
  var vl [v_l]float64
  var kl [v_l]float64
  vi , ki := 0 , 0
  for k , v := range v{
    kl[ki] = k
    vl[vi] = v
  }
  sort.Sort(sortCriticsMapWrap{kl , vl , fn})
}

func get_recommendation( prefs map[string]map[string]float64 ,
                         per string, similarity func( prefs map[string]map[string]float64 , p1 string, p2 string)(sim float64 ) ){
  var totals map[string]float64;
  var sim_sums map[string]float64;
  var rankings map[string]float64;
  for critic , item_point := range prefs{
    if per == critic {
      continue
    }
    
    sim := similarity( prefs , per , critic )
    if sim <= 0{
      continue
    }

    for book , point := range prefs[critic]{
      if per_point , per_point_exist := prefs[per][book] ; !per_point_exist || per_point == 0 {
        totals[book] += point * sim ;
        sim_sums[book] += sim ;
      }
    }
  }
  for book , _ := range totals{
    rankings[book] = totals[book] / sim_sums[book]
  }
  SortCriticMapWith(rankings , func(a, b interface{}){ return a.(float64) < b.(float64) })
  
}
func main(){
  
  prefs := map[string]map[string]float64{
    "johnson": map[string]float64{
      "programming ruby": 5.0 ,
      "learning go": 5.0 },
    "britney": map[string]float64{
      "programming ruby": 1.0 ,
      "learning go": 1.0 ,
      "SpongeBob" : 0.0},
    "jingjing":map[string]float64{
      "programming ruby": 0.0 ,
      "learning go": 0.0 ,
      "SpongeBob" : 1.0 ,
      "huluwa" : 5.0}}

  fmt.Printf("%s 和 %s的兴趣相似度" , "johnson" , "britney")
  fmt.Printf("%v\n", sim_distance( prefs , "johnson" , "britney" ))
  fmt.Printf("%s 和 %s的兴趣相似度" , "johnson" , "jingjing")
  fmt.Printf("%v\n", sim_distance( prefs , "johnson" , "jingjing" ))
}