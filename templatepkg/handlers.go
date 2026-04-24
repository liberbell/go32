package main

import (
	"fmt"
	"log"
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
	if err != nil {
		_, _ = fmt.Fprint(w, "Error executing template:", err)
	}
}

func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println("Unexpected template error:", err.Error())
	}
}
