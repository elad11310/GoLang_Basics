package main // the main package is the entry point of our app,every go app constructed from packages.

import (
	"fmt"     // format package , generally for input\output ,strings,error messages.
	"strconv" // string convert to print the "int" value and not the ASCI
)

// when declaring outside the "main" func , in the package level, we need full declaration
var o int = 22

//naming convention: in Go if we start with a capital letter it will be export from the package
//if we start with lower case letter it will be internal to the package , means from another package (not the main)
// we cannot access the variable unless it will start with capital letter
// when declaring alot of variables:

var (
	name    string = "Elad"
	counter int    = 0
)

//numeric const
// iota its int and first variable gets 0 and later 1 and 2... we dont have to explicity do b=iota
// iota its scooped - means in another scoop the values will be from 0 again
const (
	a = iota
	b
	c
)

const (
	isAdmin          = 1 << iota // 00000001
	isHeadQuarters               // 00000010
	canSeeFinancials             // 00000100
	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

// in Go we dont have inheritance , we have another model called composition which is similar to it.
// Bird indeed has name and origin but it's type is still Bird.

type Animal struct {
	name string 
	origin string
}

type Bird struct {
	Animal
	speed float32
	canFly bool
}

func main() {

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("IsAdmin? %v\n", isAdmin&roles == isAdmin)

	fmt.Println("Hello Elad2")
	fmt.Println(a) //0
	fmt.Println(b) //1
	fmt.Println(c) //2
	//declaring a variable
	// if we are not ready to init te variable yet we will use this syntax.
	var i int
	i = 42
	//second way
	var k int = 28
	//third way
	j := 32

	// when we try to convert int to string, it is printing the ASCI value of the int
	// if we want to print the int value in this case "42" we will use strconv package
	var str string
	//str =string(i)
	str = strconv.Itoa(i) //
	fmt.Println(str)

	fmt.Printf("%v , %T\n", k, k) // printing value and type of variable
	fmt.Println(name)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	// every variable in Go automatically is initialized with 0 , boolean with false
	var y int
	var b bool
	fmt.Println(y)
	fmt.Println(b)

	// we have some types of int in Go
	//int8 -128 -> 127 , int16 -32,768 -> 32,767 , int32 -2billion -> +2 billion , int64
	//also we have uint8 0-> 255 , uint16 0->65,535 , uint32 0-> 4 billion

	s := "this is a string"
	s2 := "continue the string"
	fmt.Printf("%v , %T\n", s, s)
	// because string is byte in Go , we get the asci value so need to cast
	fmt.Printf("%v , %T\n", string(s[2]), s)

	//example of concat of strings
	fmt.Printf("%v , %T\n", s+s2, s+s2)

	//if we want to convert each char in the string to its ASci value or utf (encoding)
	// result in uint8 which is a type alias for bytes - holds Asci value in integer 0-255

	b12 := []byte(s)
	fmt.Printf("%v , %T\n", b12, b12)

	//rune in Go its like char , alias for int32

	var r rune = 'a'
	fmt.Printf("%v , %T\n", r, r)

	// const - declared same as variable , if we want to export it we will start with big letter just like variable
	// we cannot assign const to anything that we know its value in run time
	// const cannot be shadowed
	const myConst int = 43
	fmt.Printf("%v , %T\n", myConst, myConst)

	// Arrays and Slices :

	grades := [3]int{99, 94, 92}
	//var grades [3]int
	//grades:= [...]int {99,94,92}
	fmt.Printf("Grades  %v\n", grades)
	fmt.Printf("Number of Grades:  %v\n", len(grades))

	// matrix

	var identityMatrix [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Printf("Matrix : %v\n", identityMatrix)

	// in go when we assign 1 array to another it wont point to the same array it will make a deep copy
	// and create a new one
	// they have fixed size that have to be known on compile size and cannot be change.

	a := [...]int{1, 2, 3}
	c := a
	c[1] = 4

	// if we want to point to the same array and avoid deep copy:
	d := &a
	d[2] = 6
	fmt.Println(d)
	fmt.Println(a)
	fmt.Println(c)

	//Slices
	// everything we can do in arr we can do in slices
	// in slices if we assign one to another it points to same value
	// slices dont need to have fixed size, we can add and delete elements

	aa := []int{1, 2, 3, 5, 6, 7, 8, 9, 10, 11}
	bb := aa
	bb[1] = 4
	fmt.Printf("%v , %T\n", aa, aa)
	fmt.Println(len(aa))
	fmt.Println(cap(aa))

	cc := aa[:]   // slice of all elements.
	dd := aa[3:]  // slice from 4th elemetns to end
	ee := aa[:6]  // slice first 6 elements
	ff := aa[3:6] // slice the 4th 5th and 6th elements
	fmt.Println(cc)
	fmt.Println(dd)
	fmt.Println(ee)
	fmt.Println(ff)

	// another way to make slices with function "make"

	g := make([]int, 3) // type , length , capacity like len unless we specify 3th argument
	fmt.Println(g)      // [0,0,0]
	fmt.Println(len(g))
	fmt.Println(cap(g))
	// add argument
	g = append(g, 1) // we can append many arguments
	fmt.Println(g)   // [0,0,0,1]
	fmt.Println(len(g))
	fmt.Println(cap(g))

	//when we add an element go making new array and copy all the elements and resize it to bigger
	// capacity , so we better know the capacity head to avoid resizeing and copying which takes a lot of time

	//if we want to append a whole slice to another :

	g = append(g, []int{2, 3}...) // we can append many arguments
	fmt.Println(g)                // [0,0,0,1]
	fmt.Println(len(g))
	fmt.Println(cap(g))

	// we can treat slices as stacks , append its like push .
	// pop like:

	g1 := g[1:] // like taking out the first
	//g2 := g[:1] // like taking out the last
	//
	fmt.Println(g1)
	g3 := append(g1[:2], g1[3:]...) // deleting in the middle, make sure we dont have another reference to this slice
	// it will make unwanted behaviour
	fmt.Println(g3)

	// maps and structs :

	statePopulations := map[string]int{
		"California": 1000000000,
		"Texas":      2000000000,
		"Florida":    3000000000,
		"New-York":   5000000000,
		"Ohio":       4000000000,
	}

	fmt.Println(statePopulations)

	//another way to init map is make function
	// good to use in a loop when we dont know from the first place the values of the map.

	statePopulations2 := make(map[string]int)
	statePopulations2 = map[string]int{
		"California": 1000000000,
		"Texas":      2000000000,
		"Florida":    3000000000,
		"New-York":   5000000000,
		"Ohio":       4000000000,
	}
	fmt.Println(statePopulations2)
	fmt.Println(statePopulations2["Ohio"])

	//adding value
	statePopulations2["Gerorgia"] = 600000000000

	fmt.Println(statePopulations2)

	//in array and slices the order is provided as we insert the elements.
	// in a map there is no guarantee it will be in the order we insert the elements.
	// also in maps like slices if we assign one to another it assigned by reference
	// in Go Slices maps cannot be tested for equality , so the key of the map cannot consist of slices as a key.

	//delete from map

	delete(statePopulations2, "Georgia")

	// if we are in situation that we dont sure if the key exists ( because it returns 0 and not error)
	// we will do :

	pop, ok := statePopulations2["Gerogia"]
	//,ok :=statePopulations2["Gerogia"]
	fmt.Println(ok)
	fmt.Println(pop, ok)

	// in structs like arr if we assign one to another it makes deep copy to a new one
	aDoctor := Doctor{ // Doctor declared in the package environment
		number:    3,
		actorName: "Elad",
		companions: []string{
			"Shani",
			"Omer",
			"Yossi",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)

	//we can also declare anonymous strut,cannot use it in another place because it doesnt have a var connected to it
	// few times we should use this.
	bDoctor := struct{ name string }{name: "Eladdd"}
	fmt.Println(bDoctor)


	//Tags in struct


}
