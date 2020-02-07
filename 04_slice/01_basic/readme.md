1. Grow and shrink in runtime
2. Default value is nil
3. Can’t get /set into nil slice
4. Append elements to a slice return a new slice
Note : appending to a sliced slice overwrites original slice
5. Slice doesn’t store any elements , it creates a baking array and returns a slice (means header) that refer to the array
6. Nil slice doesn’t have baking array but it has slice header
7. When the capacity is full append allocates a new larger array.
8. Make initialize and returns a slice with given length and capacity