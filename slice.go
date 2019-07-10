package main
import "fmt"
//import "strings"
import "strconv"
import "sort"
func main(){


sli := make([]int , 0,3)
for{
fmt.Printf("type the no to be added:")
var x string
fmt.Scan(&x)
if x == "X"{
break
}
var t int
t,_=strconv.Atoi(x)
sli = append(sli , t)
sort.Ints(sli)

fmt.Printf("%d",sli)





}

}