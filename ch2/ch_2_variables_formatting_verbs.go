package main 
import "fmt"

func main(){
	const name = "Saul Goodman"
	const OpenRate = 30.5
	msg := fmt.Sprintf("Hi %s, your open rate is %f percent", name, OpenRate)
	fmt.Println(msg)
}