package main
import "fmt"
import "reflect"

func main(){
    s := "asSASA ddd dsjkdsjs dk"
    fmt.Println("s=",s,"len(s)=",len(s))
  
    s1 := "asSA汉字"
    fmt.Println("s1=",s1,"len(s1)=",len(s1))
 
    for i,v := range s1{
        fmt.Println(i,v,string(v))
    }
    // chang to rune slice, note the index val changed
    r := []rune(s1)
    for i,v := range r {
        fmt.Println(i, v, string(v))
    }

    //replace "abc" from pos 4
    s2 := s[:]
    // for s2's is string still, could not be edit
    //s2[3]='a'
    //s2[4]='b'
    //s2[5]='c'
    fmt.Println(s2)
    fmt.Println("&s=",&s, "&s2=", &s2)
    fmt.Println("type of s2=", reflect.TypeOf(s2))

    // string to slice
    ss := []byte(s)
    ss[3] = 'a'
    ss[4] = 'b'
    ss[5] = 'c'
    sss := string(ss)
    fmt.Println("sss=",sss)    


    array := [5]byte{'h','e','l','l','o'}
    array[3] = 'k'
    fmt.Println("array type is ", reflect.TypeOf(array))
    slice := array[:]
    fmt.Println("slice type is ", reflect.TypeOf(slice))
}

//conclusion:[]byte only for english char, 
//[]rune could cover char and chinese, ie, all character
