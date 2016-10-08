package main
import "fmt"
import "os"
import "bufio"
import "io"

func main(){
    argc := len(os.Args)
    if argc<2 || argc >3{
        return
    }
    flag := ""
    if argc==3 {
        flag = os.Args[1]
    }
    fname := os.Args[argc-1]
    f,err := os.Open(fname)
    defer f.Close()
    if err!=nil {
        return
    }
    
    rd:=bufio.NewReader(f)
    cnt := 0
    for {
        line,err := rd.ReadBytes('\n')
        if flag == "-n" {
            cnt++
            fmt.Printf("%d ",cnt)
        }
        os.Stdout.Write(line)
        if  err != nil{
            return
        }
        if err == io.EOF{
            return
        }
    }
    
    
}
