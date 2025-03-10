package main
import ("fmt")

/*
	2 cara mendeklarasikan variable dalam go
	pakai var
	atau :=
*/ 

// deklarasi variable dengan nilai awal
// func main(){
// 	var student1 string = "John" // type is string
// 	var student2 = "Jane" //type is inferred
// 	x := 2 // type is inferred

// 	fmt.Println(student1)
// 	fmt.Println(student2)
// 	fmt.Println(x)	
// }

// deklarasi variable tanpa nilai awal
// func main(){
// 	var a string
// 	var b int
// 	var c bool

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// }

// penugasan nilai setelah deklarasi
func main()  {
	var student1 string
	student1 = "John"
	fmt.Println(student1)
}