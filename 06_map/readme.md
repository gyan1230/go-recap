1. Map is an unordered collections of key-value pairs. Keys are unique, values may not be.
2. Map is used for fast lookups, retrieval and deletion of data based on keys.

Declaring a Map
1. var identifier map[key type] value type
    eg: var m map[string]int
2. zero value of a map is nil
3. A nil map has no keys, any attempt to add keys to a nil map will result in runtime error.

package main
import "fmt"

func main() {
	var m map[string]int
	fmt.Println(m)
	if m == nil {
		fmt.Println("m is nil")
	}

	// Attempting to add keys to a nil map will result in a runtime error
	// m["one hundred"] = 100
}

Initializing a Map
1. initializing a map using make function -> not nil map
    eg: var m = make(map[string]int)
2. make function returns an initialized and ready to use map, means can add keys
3. Initialize by using map literal
    eg: var m = map[string]int{
	                "one": 1,
	                "two": 2,
                	"three": 3,
                    }
4. Create an empty map using map literal by leaving the curly braces empty 
    eg: var m = map[string]int{}   //functionally identical to using the make() function
5. If you try to add a key that already exists in the map, then it will simply be overridden by the new value.