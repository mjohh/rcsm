package main
import (
 "fmt"
 //"strings"
 "errors"
 "strconv")

var s1 *stack
var s2 *stack
var s3 *stack

func init() {
    s1 = &stack{top:-1}
    s2 = &stack{top:-1}
    s3 = &stack{top:-1}
}

type stack struct {
    array [32]string
    top int
}

func push(s *stack, a string){
    s.top++
    s.array[s.top] = a
}

func pop(s *stack) string{
    if s.top<=-1{
        return ""
    }
    a := s.array[s.top]
    s.top--
    return a
}

func top(s *stack) string{
    if s.top<=-1 {
        return ""
    }
    return s.array[s.top]
}
/*
// we just support +,-,*,/
func pri(op string) int {
    if strings.EqualFold(op, ")" {
        return 0
    }
    if strings.EqualFold(op, "-") || strings.EqualFold(op, "+"){
        return 1
    }else if strings.EqualFold(op, "*") || strings.EqualFold(op, "/"){
        return 2
    }
    return -1
}
*/
func pri(op string) int {
    if op== ")" {
        return 0
    }
    if op== "-"||op== "+"{
        return 1
    }else if op=="*" || op== "/"{
        return 2
    }
    return -1
}



type Type int
const(
    Number Type=iota
    Operator
    Lbracket
    Rbracket
)

// for this func will be called cyclely, and for  more efficiency, give len of string,
// instead of calling len() within func.
//func parse(str string, cursor int, len int) (s string, t Type, nextcursor int, err error){
func parse(str string, cursor int, len int) ( string,  Type,  int,  error){
    ///fmt.Println("args into parse() are str,cursor,len=",str, cursor, len)
    i := cursor
    if str[i]=='(' {
        ///fmt.Println("in ( case")
        return "(", Lbracket,i+1,nil 
    }
    if str[i]==')'{
        ///fmt.Println("in ) case")
        return ")", Rbracket, i+1, nil
    }
    if str[i]=='+'||str[i]=='-'||str[i]=='*'||str[i]=='/'{
        ///fmt.Println("in op case")
        return str[i:i+1], Operator, i+1, nil
    }
    for i<len && str[i]>='0'&&str[i]<='9'{
       i++
    }
    if i>cursor {
        if i>=len {
            ///fmt.Println("i>=len, i,len=",i,len)
            return str[cursor:], Number, i, nil
        }
        ///fmt.Println("i<len, i,len=",i,len)
        return str[cursor:i], Number, i, nil
    }
    
    return "",Number,-1,errors.New("illegal string!")
}

func typename(t Type) string{
    switch t{
    case Number:
        return "Number"
    case Operator:
        return "Operator"
    case Lbracket:
        return "Lbracket"
    case Rbracket:
        return "Rbracket"
    default:
        return ""
    }
}

func parse_str(str string){
    l:= len(str)
    var cursor int
    var s string
    var t Type
    var err error
    fmt.Println(str)
    for cursor < l{
        s,t,cursor,err =parse(str,cursor,l)
        fmt.Println("after parse(),s,t,cursor=",s,typename(t),cursor,err)
    }
}

func parse_test(){
    str0 := "2+1"
    str1 := "2*1"
    str2 := "(2+1)"
    str3 := "(2+1)*3"
    str4 := "3*(2+1)+4*(4-2)"
    str5 := "(33*(2378+1)+4*(400-2))/55"

    parse_str(str0)
    parse_str(str1)
    parse_str(str2)
    parse_str(str3)
    parse_str(str4)
    parse_str(str5)
}

func cal(op, a, b  string) int{
    // todo: err handle
    c,_ := strconv.Atoi(a)
    d,_ := strconv.Atoi(b)
    
    switch op{
    case "+":
        return c+d 
    case "-":
        return c-d
    case "*":
        return c*d
    default: //"/"
        return c/d
    }
}

// 
func handlestr(s string, t Type){
    switch t{
    case Lbracket:
        push(s1, "(")
    case Rbracket:
        for top(s1)!="("{
            push(s2, pop(s1))
        }
        pop(s1)
    case Operator:
        for pri(top(s1))>=pri(s){
            push(s2, pop(s1))
        }
        push(s1, s)
    case Number:
        push(s2, s)
    default:
    }
}

// use global stack s1,s2 to gen reversed polish
// result stores in s2
func genrpolished(str string){
    l:= len(str)
    var cursor int
    var s string
    var t Type
    var err error
    //fmt.Println(str)
    for cursor < l{
        s,t,cursor,err=parse(str,cursor,l)
        if err==nil{
            handlestr(s, t)
        }else{
            fmt.Println("parse err:", err)
        }
        
        //fmt.Println("after parse(),s,t,cursor=",s,typename(t),cursor,err)
    }
    for top(s1)!=""{
        push(s2,pop(s1))
    }
}

func calrpolished() int{
    //fmt.Println("s2.top=",s2.top)
    for i:=0; i<=s2.top; i++{
        switch s2.array[i] {
        case "+":
            fallthrough
        case "-":
            fallthrough
        case "*":
            fallthrough
        case "/":
            //attention sequence
            a := pop(s3)
            b := pop(s3)
            r := cal(s2.array[i],b,a)    
            push(s3, strconv.Itoa(r))
            //fmt.Println("op case s3,r=", s3, r)
        default:
            // todo: err handle
            //c,_ := strconv.Atoi(s2.array[i])
            push(s3, s2.array[i])
            //fmt.Println("number case, s2.array[i]=", s2.array[i])
        }
    }
    r,_ := strconv.Atoi(pop(s3))
    //fmt.Println("r=",r)
    return r
}

func caculation(s string) int{
    genrpolished(s)
    return calrpolished()
}

func reset(){
    s1.top=-1
    s2.top=-1
    s3.top=-1
}
func genrpolished_testone(s string){
    genrpolished(s)
    fmt.Println(s)
    for top(s2) != ""{
        fmt.Println(pop(s2))
    }    
}
func genrpolished_test(){
    genrpolished_testone("1+4")   
    reset()
    genrpolished_testone("(1+4)")   
    reset()
    genrpolished_testone("(1+4)*3")   
    reset()
    genrpolished_testone("(1+4)*3+8")   
    reset()
    genrpolished_testone("(1+4)*3+8/(2+2)*5")   
    reset()
    genrpolished_testone("8/(2+2)*5")   
    reset()
}

func calrpolished_test(){
    genrpolished("1+4")   
    fmt.Println(calrpolished())
    reset()
    genrpolished("(1+4)")   
    fmt.Println(calrpolished())
    reset()
    genrpolished("(1+4)*3")   
    fmt.Println(calrpolished())
    reset()
    genrpolished("(1+4)*3+8")   
    fmt.Println(calrpolished())
    reset()
    genrpolished("(1+4)*3+8/(2+2)*5")   
    fmt.Println(calrpolished())
    reset()
    genrpolished("8/(2+2)*5")   
    fmt.Println(calrpolished())
    reset()
}
func main(){
    parse_test()
    genrpolished_test()
    calrpolished_test()
}
