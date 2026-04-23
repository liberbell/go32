package main

import "github.com/CloudyKit/jet/v6"

var views = jet.NewSet{
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
}
