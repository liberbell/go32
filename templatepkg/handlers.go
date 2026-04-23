package main

import (
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

var views = jet.NewSet{
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := make(jet.VarMap)
}
