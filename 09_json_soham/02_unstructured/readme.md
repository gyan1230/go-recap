link:
https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/

Unstructured data (Decoding JSON to maps)
If you have data whose structure or property names you are not certain of, you cannot use structs to unmarshal your data. Instead you can use maps. Consider some JSON of the form:

{
  "birds": {
    "pigeon":"likes to perch on rocks",
    "eagle":"bird of prey"
  },
  "animals": "none"
}
There is no struct we can build to represent the above data for all cases since the keys corresponding to the birds can change, which will change the structure.

To deal with this case we create a map of strings to empty interfaces:

birdJson := `{"birds":{"pigeon":"likes to perch on rocks", eagle":"bird of prey"},"animals":"none"}`
var result map[string]interface{}
json.Unmarshal([]byte(birdJson), &result)

// The object stored in the "birds" key is also stored as 
// a map[string]interface{} type, and its type is asserted from
// the interface{} type
birds := result["birds"].(map[string]interface{})

for key, value := range birds {
  // Each value is an interface{} type, that is type asserted as a string
  fmt.Println(key, value.(string))
}
Try it here

Each string corresponds to a JSON property, and its mapped interface{} type corresponds to the value, which can be of any type. The type is asserted from this interface{} type as is needed in the code. These maps can be iterated over, so a variable number of keys can be handled by a simple for loop.