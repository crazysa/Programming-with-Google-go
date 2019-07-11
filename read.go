/*package main
import "fmt"
//import "strings"
//import "strconv"
//import "sort"
//import "Json"
//import "encoding/json"
import "io/ioutil"
func main(){

type Person struct {

fname string
lname string
}

sli := make([]Person ,0 ,  10)


var i int
dat , e := ioutil.ReadFile("test.txt")
for i = 0 ; i < len(dat)/2;i++{

sli.fname = append(sli.fname , dat [i*2])
sli.lname=append(sli.lname,dat [i*2+1])




}
fmt.Printf("%s",sli)

}*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Name Struct
type Name struct {
	fname string
	lname string
}

func main() {
	//inputReader := bufio.NewReader(os.Stdin)
	names := make([]Name, 0, 3)

		

	fileName := strings.Trim("test.txt", "\n")

	file, _:= os.Open(fileName)


	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		
		var aName Name
		aName.fname, aName.lname = s[0], s[1]
		names = append(names, aName)

	}

	file.Close()

	for _, v := range names {
		fmt.Println(v.fname, v.lname)
	}

}