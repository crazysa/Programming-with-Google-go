package main
import "fmt"
//import "strings"
//import "strconv"
//import "sort"
//import "Json"
import "encoding/json"
func main(){


fmt.Printf("Type the name followed by address")
var x , y string
fmt.Scan(&x,&y)
idmap := map[string] string{
	x:y}
barr , _ := json.Marshal(idmap)

fmt.Printf("%s",barr)
}