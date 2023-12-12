package main
import "fmt"

func main(){
	firstName, _ := getNames()
	fmt.Println("Wlecome to Textio", firstName)
}

func getNames() (string, string){
	return "John", "Doe"
}