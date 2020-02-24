Retrieving the value associated with a given key in a map
1. To retrieve a value from a map use map[key] -> value is present means you get a key, otherwise zero value of map's valur type.

Deleting a key from a map
1. You can delete a key from a map using the built-in delete() function. The syntax looks like this -

// Delete the `key` from the `map`
delete(map, key)

Maps are reference types
1. When you assign a map to a new variable, they both refer to the same underlying data structure. Therefore changes done by one variable will be visible to the other 

Iterating over a map
You can iterate over a map using range form of the for loop. It gives you the key, value pair in every iteration