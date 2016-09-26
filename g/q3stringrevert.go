package main
import "fmt"

func revert(s string) string{
    bytes := []byte(s[:])
    len := len(bytes)
    fmt.Println("len=",len)
    for i:=0;i<len/2;i++ {
        bytes[i],bytes[len-1-i] = bytes[len-1-i],bytes[i]
    }
    return string(bytes)
}
func revert2(s string) string{
    r := []rune(s)
    len := len(r)
    fmt.Println("len=",len)
    for i:=0;i<len/2;i++ {
        r[i],r[len-1-i] = r[len-1-i],r[i]
    }
    return string(r)
}

func main() {
    s:= "abcdefg"
    s1 := s
    s2 := &s    
    s3 := "zhong国人"
    
    fmt.Println("s=",s,",&s=",&s)
    fmt.Println("s1=",s1,",&s1=",&s1)
    fmt.Println("s2=",s2,",&s2=",&s2)

    fmt.Println("s=",s, "revert(s)=",revert(s))
    fmt.Println("s=",s3, "revert2(s)=",revert2(s3))
}
