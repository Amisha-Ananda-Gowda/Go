package main
import (
	"fmt"
	"strings"
)
func main(){
	s:="I am a string"
	fmt.Printf("Ends with string? %v\n",strings.HasSuffix(s,"string"))
}