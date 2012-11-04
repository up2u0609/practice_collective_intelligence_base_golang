package main
import(
  "fmt"
  "math"
)
func sim_distance( prefs map[string]map[string]float64 ,
                   per1 string, 
                   per2 string)( distance float64 ){
  for book , point_per1 := range prefs[per1]{
    if point_per2 , point_per2_exist := prefs[per2][book]; point_per2_exist {
      distance += math.Pow( point_per1 - point_per2 , 2 )
    }
  }
  return
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

  fmt.Printf("%s 和 %s的兴趣差异" , "johnson" , "britney")
  fmt.Printf("%v\n", sim_distance( prefs , "johnson" , "britney" ))
  fmt.Printf("%s 和 %s的兴趣差异" , "johnson" , "jingjing")
  fmt.Printf("%v\n", sim_distance( prefs , "johnson" , "jingjing" ))
}