package main

import (
	"fmt"
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

var views = jet.NewSet{
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := make(jet.VarMap)

	data.Set("user_id", 1)
	data.Set("first", "bob")
	dow := []string{
		"Sunday",
		"Monday",
		"Thuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Sarturday",
	}

	data.Set("dow", dow)

	err := renderPage(w, "home.jet", data)
	if err != nil {
		_, _ = fmt.Fprint(w, "Error executing template:", err)
	}
}

func SendData(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "send.jet", nil)
}
