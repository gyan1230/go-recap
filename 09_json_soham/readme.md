Structured data (Decoding JSON to Structs)
Since this is much easier, let’s deal with it first. This is the sort of data where you know the structure beforehand. For example, let’s say you have a bird object, where each bird has a species field and a description field :

{
  "species": "pigeon",
  "decription": "likes to perch on rocks"
}
To work with this kind of data, create a struct that mirrors the data you want to parse. In our case, we will create a bird struct which has a Species and Description attribute:

type Bird struct {
  Species string
  Description string
}
And unmarshal it as follows:

birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
var bird Bird	
json.Unmarshal([]byte(birdJson), &bird)
fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)
//Species: pigeon, Description: likes to perch on rocks
Try it here

By convention, Go uses the same title cased attribute names as are present in the case insensitive JSON properties. So the Species attribute in our Bird struct will map to the species, or Species or sPeCiEs JSON property.

JSON Arrays
So what happens when you have an array of birds?

[
  {
    "species": "pigeon",
    "decription": "likes to perch on rocks"
  },
  {
    "species":"eagle",
    "description":"bird of prey"
  }
]
Since each element of the array is actually a Bird, you can actually unmarshal this, by just creating an array of birds :

birdJson := `[{"species":"pigeon","decription":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
var birds []Bird
json.Unmarshal([]byte(birdJson), &birds)
fmt.Printf("Birds : %+v", birds)
//Birds : [{Species:pigeon Description:} {Species:eagle Description:bird of prey}]
Try it here

Embedded objects
Now, consider the case when you have a property called Dimensions, that measures the Height and Length of the bird in question:

{
  "species": "pigeon",
  "decription": "likes to perch on rocks"
  "dimensions": {
    "height": 24,
    "width": 10
  }
}
As with our previous examples, we need to mirror the structure of the object in question in our Go code. To add an embedded dimensions object, lets create a dimensions struct :

type Dimensions struct {
  Height int
  Width int
}
Now, the Bird struct will include a Dimensions field:

type Bird struct {
  Species string
  Description string
  Dimensions Dimensions
}
And can be unmarshaled using the same method as before:

birdJson := `{"species":"pigeon","description":"likes to perch on rocks", "dimensions":{"height":24,"width":10}}`
var birds Bird
json.Unmarshal([]byte(birdJson), &birds)
fmt.Printf(bird)
// {pigeon likes to perch on rocks {24 10}}
Try it here

Primitives
We mostly deal with complex objects or arrays when working with JSON, but it’s easy to forget that data like 3, 3.1412 and "birds" are also perfectly valid JSON strings. We can unmarshal these values to their corresponding data type in Go:

numberJson := "3"
floatJson := "3.1412"
stringJson := `"bird"`

var n int
var pi float64
var str string

json.Unmarshal([]byte(numberJson), &n)
fmt.Println(n)
// 3

json.Unmarshal([]byte(floatJson), &pi)
fmt.Println(pi)
// 3.1412

json.Unmarshal([]byte(stringJson), &str)
fmt.Println(str)
// bird
Try it out

Custom attribute names
I mentioned earlier that Go uses convention to ascertain the attribute name it should map a property to. Many times though, you want a different attribute name than the one provided in your JSON data.

{
  "birdType": "pigeon",
  "what it does": "likes to perch on rocks"
}
In the JSON data above, I would prefer birdType to remain as the Species attribute in my Go code. It is also not possible for me to provide a suitable attribute name for a key like "what it does".

To solve this, we make use of struct field tags:

type Bird struct {
  Species string `json:"birdType"`
  Description string `json:"what it does"`
}
Now, we can explicitly tell our code which JSON property to map to which attribute.

birdJson := `{"birdType": "pigeon","what it does": "likes to perch on rocks"}`
var bird Bird
json.Unmarshal([]byte(birdJson), &bird)
fmt.Println(bird)
// {pigeon likes to perch on rocks}
Try it here