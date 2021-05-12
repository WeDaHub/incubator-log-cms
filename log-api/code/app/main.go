package main

import (
	"api/code/irisapp"
)

func main() {
	irisapp.InitApp()
	irisapp.Start("/api/v1", "0.0.0.0:8848")
}
