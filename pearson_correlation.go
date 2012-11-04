package pearson_correlation
import(
  "fmt"
  "math"
)
func Sim_pearson( prefs map[string]map[string]float64 ,
                   per1 string, 
                   per2 string)( pearson float64 ){
  var si [10]string
  si_index := 0
  for book , _ := range prefs[per1]{
    if _ , point_per2_exist := prefs[per2][book]; point_per2_exist {
      si[si_index] = book
      si_index += 1
    }
  }
  //pearson 越趋近于1，则两人相似度一致；趋近于-1则完全不一致；为0则为无法评价
  if len( si ) == 0{
    pearson = 0
    return
  }

  per1_points_sum := 0.0
  per1_points_sumq := 0.0
  per2_points_sum := 0.0
  per2_points_sumq := 0.0
  per_sum := 0.0

  for _ , book := range si{
    per1_points_sum += prefs[per1][book]
    per1_points_sumq += math.Pow(prefs[per1][book] , 2)
    per2_points_sum += prefs[per2][book]
    per2_points_sumq += math.Pow(prefs[per2][book] , 2)
    per_sum += prefs[per1][book] * prefs[per2][book]
  }
  fmt.Printf("%v\n" , [...]float64{ per1_points_sum , per1_points_sumq , per2_points_sum ,per2_points_sumq ,per_sum , float64(si_index)})

  // 计算皮尔逊差值
  num := per_sum - ( per1_points_sum * per2_points_sum / float64(si_index) )
  fmt.Printf("%v\n" , [...]float64{ num })

  den := math.Sqrt(
    (per1_points_sumq - math.Pow(per1_points_sum , 2)) *
    (per2_points_sumq - math.Pow(per2_points_sum , 2)) /
    math.Pow( float64(si_index) , 2))
  fmt.Printf("%v\n" , [...]float64{ den })
  if den == 0{
    pearson = 0
    return
  }
  pearson = num / den
  return
}
// 范例函数
// func main(){
  
//   prefs := map[string]map[string]float64{
//     "johnson": map[string]float64{
//       "programming ruby": 5.0 ,
//       "learning go": 3.0 },
//     "britney": map[string]float64{
//       "programming ruby": 5.0 ,
//       "learning go": 3.0 ,
//       "SpongeBob" : 0.0},
//     "jingjing":map[string]float64{
//       "programming ruby": 0.0 ,
//       "learning go": 0.0 ,
//       "SpongeBob" : 1.0 ,
//       "huluwa" : 5.0}}

//   fmt.Printf("%s 和 %s的兴趣相似度" , "johnson" , "britney")
//   fmt.Printf("%v\n", Sim_pearson( prefs , "johnson" , "britney" ))
//   fmt.Printf("%s 和 %s的兴趣相似度" , "johnson" , "jingjing")
//   fmt.Printf("%v\n", Sim_pearson( prefs , "johnson" , "jingjing" ))
// }