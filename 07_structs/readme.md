A struct is a user-defined type that contains a collection of named fields/properties. It is used to group related data together to form a single unit. Any real-world entity that has a set of properties can be represented using a struct.

If you’re coming from an object-oriented background, you can think of a struct as a lightweight class that supports composition but not inheritance.

Defining a struct type
You can define a new struct type like this -

type Person struct {
	FirstName string
	LastName  string
	Age       int
}
The type keyword introduces a new type. It is followed by the name of the type (Person) and the keyword struct to indicate that we’re defining a struct. The struct contains a list of fields inside the curly braces. Each field has a name and a type.

Note that, you can collapse fields of the same type like this:

type Person struct {
	FirstName, LastName string
	Age       int
}
Declaring and Initializing a struct
Declaring a variable of a struct type
Just like other data types, you can declare a variable of a struct type like this -

// Declares a variable of type 'Person'
var p Person // All the struct fields are initialized with their zero value
The above code creates a variable of type Person that is by default set to zero. For a struct, zero means all the fields are set to their corresponding zero value. So the fields FirstName and LastName are set to "", and Age is set to 0.

Initializing a struct
You can initialize a variable of a struct type using a struct literal like so -

// Initialize a struct by supplying the value of all the struct fields.
var p = Person{"Rajeev", "Singh", 26}
Note that you need to pass the field values in the same order in which they are declared in the struct. Also, you can’t initialize only a subset of fields with the above syntax -

var p = Person{"Rajeev"} // Compiler Error: too few values in struct initializer
Naming fields while initializing a struct
Go also supports the name: value syntax for initializing a struct (the order of fields is irrelevant when using this syntax).

var p = Person{FirstName: "Rajeev", LastName: "Singh", Age: 25}
You can separate multiple fields by a new line for better readability (the trailing comma is mandatory in this case) -

var p = Person{
	FirstName: "John",
	LastName:  "Snow",
	Age:       45,
}
The name: value syntax allows you to initialize only a subset of fields. All the uninitialized fields are set to their corresponding zero value -

var p = Person{FirstName: "Alien"} // LastName: "", Age: 0
var p = Person{} // FirstName: "", LastName: "", Age: 0
Complete Example: Defining and Initializing struct types
package main

import (
	"fmt"
)

// Defining a struct type
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// Declaring a variable of a `struct` type
	var p Person // // All the struct fields are initialized with their zero value
	fmt.Println(p)

	// Declaring and initializing a struct using a struct literal
	p1 := Person{"Rajeev", "Singh", 26}
	fmt.Println("Person1: ", p1)

	// Naming fields while initializing a struct
	p2 := Person{
		FirstName: "John",
		LastName:  "Snow",
		Age:       45,
	}
	fmt.Println("Person2: ", p2)

	// Uninitialized fields are set to their corresponding zero-value
	p3 := Person{FirstName: "Robert"}
	fmt.Println("Person3: ", p3)
}
# Output
{  0}
Person1:  {Rajeev Singh 26}
Person2:  {John Snow 45}
Person3:  {Robert  0}
Accessing fields of a struct
You can access individual fields of a struct using the dot (.) operator -

package main

import (
	"fmt"
)

type Car struct {
	Name, Model, Color string
	WeightInKg         float64
}

func main() {
	c := Car{
		Name:       "Ferrari",
		Model:      "GTC4",
		Color:      "Red",
		WeightInKg: 1920,
	}

	// Accessing struct fields using the dot operator
	fmt.Println("Car Name: ", c.Name)
	fmt.Println("Car Color: ", c.Color)

	// Assigning a new value to a struct field
	c.Color = "Black"
	fmt.Println("Car: ", c)
}
# Output
Car Name:  Ferrari
Car Color:  Red
Car:  {Ferrari GTC4 Black 1920}
Pointer to a struct
You can get a pointer to a struct using the & operator -

package main

import (
	"fmt"
)

type Student struct {
	RollNumber int
	Name       string
}

func main() {
	// instance of student struct type
	s := Student{11, "Jack"}

	// Pointer to the student struct
	ps := &s
	fmt.Println(ps)

	// Accessing struct fields via pointer
	fmt.Println((*ps).Name)
	fmt.Println(ps.Name) // Same as above: No need to explicitly dereference the pointer

	ps.RollNumber = 31
	fmt.Println(ps)
}
# Output
&{11 Jack}
Jack
Jack
&{31 Jack}
As demonstrated in the above example, Go lets you directly access the fields of a struct through the pointer without explicit dereference.

Creating a struct and obtaining a pointer to it using the built-in new() function
You can also use the built-in new() function to create an instance of a struct. The new() function allocates enough memory to fit all the struct fields, sets each of them to their zero value and returns a pointer to the newly allocated struct -

package main

import "fmt"

type Employee struct {
	Id   int
	Name string
}

func main() {
	// You can also get a pointer to a struct using the built-in new() function
	// It allocates enough memory to fit a value of the given struct type, and returns a pointer to it
	pEmp := new(Employee)

	pEmp.Id = 1000
	pEmp.Name = "Sachin"

	fmt.Println(pEmp)
}
# Output
&{1000 Sachin}
Exported vs Unexported Structs and Struct Fields
Any struct type that starts with a capital letter is exported and accessible from outside packages. Similarly, any struct field that starts with a capital letter is exported.

On the contrary, all the names starting with a small letter are visible only inside the same package.

Let’s see an example. Consider the following package hierarchy of our Go program -

example
  main
    main.go
  model
    address.go
    customer.go  
customer.go

package model

type Customer struct {  // exported struct type
	Id int				// exported field
	Name string			// exported field
	addr address        // unexported field (only accessible inside package `model`)
	married bool  		// unexported field (only accessible inside package `model`)
}
address.go

package model

// Unexported struct (only accessible inside package `model`)
type address struct {
	houseNo, street, city, state, country string
	zipCode                               int
}
And, here is the driver program main.go -

package main

import (
	"fmt"
	"example/model"
)

func main() {
	c := model.Customer{
		Id: 1, 
		Name: "Rajeev Singh",
	}

	c.married = true	// Error: can not refer to unexported field or method

	a := model.address{} // Error: can not refer to unexported name

	fmt.Println("Programmer = ", c);
}
As you can see, the names address and married are unexported and not accessible from the main package.

Structs are value types
Structs are value types. When you assign one struct variable to another, a new copy of the struct is created and assigned. Similarly, when you pass a struct to another function, the function gets its own copy of the struct.

package main

import "fmt"

type Point struct {
	X float64
	Y float64
}

func main() {
	// Structs are value types.
	p1 := Point{10, 20}
	p2 := p1 // A copy of the struct `p1` is assigned to `p2`
	fmt.Println("p1 = ", p1)
	fmt.Println("p2 = ", p2)

	p2.X = 15
	fmt.Println("\nAfter modifying p2:")
	fmt.Println("p1 = ", p1)
	fmt.Println("p2 = ", p2)
}
# Output
p1 =  {10 20}
p2 =  {10 20}

After modifying p2:
p1 =  {10 20}
p2 =  {15 20}
Struct Equality
Two struct variables are equal if all their corresponding fields are equal -

package main

import "fmt"

type Point struct {
	X float64
	Y float64
}

func main() {
	// Two structs are equal if all their corresponding fields are equal.
	p1 := Point{3.4, 5.2}
	p2 := Point{3.4, 5.2}

	if p1 == p2 {
		fmt.Println("Point p1 and p2 are equal.")
	} else {
		fmt.Println("Point p1 and p2 are not equal.")
	}
}
# Output
Point p1 and p2 are equal.