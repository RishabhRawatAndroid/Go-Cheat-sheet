package main

import (
	"fmt"
	"reflect"
)

type User struct {
	firstname string
	lastname string
}

func (u User) fullName() string{
	u.firstname="Ris"
	u.lastname="ckbkjd"
	return u.firstname+" "+u.lastname
}


//GO LANG IS STATICALLY TYPE LANGUAGE
func main() {

	//for printing the console
	fmt.Println("This is console output") //no need of semicolon

	//DATA TYPE IN GO LANG
	//string
	//bool
	//int int8 int16 int32 int64
	//uint uint8 uint16 uint32 uint64
	//byte
	//rune
	//float32 float64
	//complex64 complex128

	//Now create a variable
	var name string ="Rishabh Rawat"
	fmt.Println(name)

	//We can use the type of variable and we also have an option for not using the type go lang infer the type automatically
	var age=24
	fmt.Println(age)

	//for checking the type of variable
	var check=false
	fmt.Printf("%T\n",check)  //first way
	fmt.Println(reflect.TypeOf(check)) //second way

	//for creating a shorthand variable
	tv:="Apple"
	fmt.Println(tv)

	//call the function
	var myname string=getname("Rishabh Rawat")
	fmt.Println(myname)

	//call a function which return two variable
	add,sub :=addsub(10,5)
	fmt.Println(add)
	fmt.Println(sub)

	//calling the Anonymous function //something wrong in this code
	//adding,subing:=adding_subtracting(addsub(10,50),10,50)

	//for making an array
	x:=[5]int{1,23,4,5,6}
	fmt.Println(x)

	//When you declare arrays, you always specify the length of the array, but when you declare and initialize
	//arrays, you can use the expression … instead of specifying the length, as shown here:
	z := [...] int { 5,10,15,20,25}
	fmt.Println(z)

	//go lang only support the for loop
	langs := [4]string{"Go", "Rust", "Scala","Julia"}
	for i := 0; i < len(langs); i++ {
		fmt.Printf("langs[%d]:%s \n", i, langs[i])
	}

	for k, v := range langs {
		fmt.Printf("langs[%d]:%s \n", k, v)
	}

	// Initialize values at specific index with array literal
	lang := [4]string{0: "Go", 3: "Julia"}
	fmt.Println("Value of langs:", langs)
	// Assign values to remaining positions
	lang[1] = "JAVA"
	lang[2] = "Scala"


	//SLICE  A slice is a data structure that
	//is built on top of Go’s array type, which allows you to store a dynamic length of elements of a single type
	//Here is the code block that creates a slice using the make function specifying the length and capacity:
	y:= make ([]int, 3,5)
	fmt.Println(y)// 3 is the length and 5 is the capacity after the capacity full then go make an new array with add 5 more capacity

	//now enlarge the size of the slice using copy and append function
	xy := []int{10, 20, 30}
	fmt.Printf("[Slice:xy] Length is %d Capacity is %d\n", len(x), cap(x))
	// Create a bigger slice
	yx := make([]int, 5, 10)
	copy(yx, xy)
	fmt.Printf("[Slice:yx] Length is %d Capacity is %d\n", len(yx), cap(yx))
	fmt.Println("Slice yx after copying:", yx)
	yx[3] = 40
	yx[4] = 50
	fmt.Println("Slice yx after adding elements:", yx)

	//Now using the map which is key value pair
	language:=map[string]string{
		"name":"Rishabh",
		"age":"22",
		"degree":"B-tech",
	}
	//m = make(map[string]int)
	fmt.Println(language)
	delete(language,"age") //for delete an element from the map
	fmt.Println(language)
	language["äge"]="22"
	fmt.Println(language)//adding new element in map
	for key, value := range language {
		fmt.Println("Key:", key, "Value:", value)
	}

	//Checking if a key exists in a map
	//value, ok := m[key]
	value,ok:=language["age"]
	if ok{
		fmt.Println(value)
	} else {
		fmt.Println("Not exit")
	}


	user:=User{}
	fmt.Println(user.fullName())











}

//function syntax
//func name(list of parameters) (list of return types)
//{
//function body
//}

//creating a function using the func keyword
func getname(name string) string{  //this function return the single variable
	return "Hi.."+name
}

//creating a function to return multiple type
func addsub(a,b int) (int,int){
	add:=a+b
	sub:=a-b
	return add,sub
}

//naming return type in go
func applePrice()(apple float32) {
	apple=784.256
	return
}

//Anonymous Functions
func adding_subtracting(res func(a,b int)(int,int),a int, b int)(add, sub int){
	add,sub=res(a,b)
	return
}