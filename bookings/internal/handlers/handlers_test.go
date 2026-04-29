package handlers

type postData struct {
	key   string
	value string
}

var theTests []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOk},
}

func TestHandlers(t *testing.T)  {
	routes := getRoutes()
}