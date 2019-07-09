package main
import "fmt"
import "strings"
func main(){
fmt.Printf("Type a String here:")
var in string
fmt.Scan(&in)
//k:=1
if strings.HasPrefix(in , "i") || strings.HasPrefix(in , "I"){
 if strings.HasSuffix(in , "n")|| strings.HasSuffix(in , "N"){
 if strings.Contains(in ,"a") || strings.Contains(in , "A"){
 
	fmt.Printf("Found")
 
 }else{
 fmt.Printf("Not Found")
 }
 
 
 }else{fmt.Printf("Not Found")}
 

}else{fmt.Printf("Not Found")}



}