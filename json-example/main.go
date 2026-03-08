package main

type Person struct {
	FirstName string `json: "first_name")`
	LastName  string `json: "last_name"`
	HairColor string `json: "hair_color"`
	HasDog    bool   `json: "has_dog"`
}

func main() {

	myJson := `
	[
		{
			"first_name": "Eric",
			"last_name": "Clapton",
			"hair_color": "Gray",
			"has_dog": false
		},
		{
			"first_name": "Bob",
			"last_name": "Bary",
			"hair_color": "Black",
			"has_dog": true
		}
	]`
}
