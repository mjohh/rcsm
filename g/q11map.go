package main
import "fmt"

type mapfunction func(a int)int

func maping(f mapfunction, a []int) []int{
    len := len(a)
    r := make([]int, len)
    for i,v := range a {
        r[i] = f(v)
    }
    return r
}

func double(a int) int{
    return a*2
}


//extend for string op
type mapfunction2 func(a interface{}) interface{}
func maping2(f mapfunction2, a []interface{}) []interface{}{
    len := len(a)
    r := make([]interface{}, len)
    for i,v := range a {
        r[i] = f(v)
    }
    return r
}

func double2(a interface{}) interface{}{
    if v,ok := a.(int); ok{
        return v*2
    }
    return nil
}

func addprefix(a interface{}) interface{}{
    if v,ok := a.(string); ok{
        return "bingo-"+v
    }
    return nil
}

func main(){
    a := []int{1,3,5,7,9,10}
    fmt.Println(a)
    aa := maping(double, a)
    fmt.Println(aa)

    a2 := []interface{}{1,3,89,2,4,0}
    aa2 := maping2(double2, a2)
    fmt.Println(aa2)

    a3 := []interface{}{"i","like","sports","and","singing","you?"}
    aa3 := maping2(addprefix, a3)
    fmt.Println(aa3)
}
