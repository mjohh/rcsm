package main
import "fmt"

func meanvalue(s []float64) float64{
    var sum float64
    for _,v := range s{
        sum += v
    }
    mean := sum/float64(len(s))
    return mean
}

func main(){
    s1 := []float64{1.1,2.2,3.3,4.4}
    fmt.Println(meanvalue(s1))
    s2 := []float64{1,2,3,4}
    fmt.Println(meanvalue(s2))
}
