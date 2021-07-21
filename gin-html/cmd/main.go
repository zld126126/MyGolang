package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/httpfs/union"
	"github.com/shurcooL/vfsgen"
)

func main() {
	// use union for combin resource
	var Assets = union.New(map[string]http.FileSystem{
		"/static": http.Dir("../static"),
	})
	var fs http.FileSystem = Assets
	err := vfsgen.Generate(fs, vfsgen.Options{
		PackageName:  "main",
		Filename:     "../dist.go",
		VariableName: "Dir",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
