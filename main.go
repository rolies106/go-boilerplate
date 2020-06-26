package main

import (
	"fmt"
	"net/http"
	"os"

	"mortred/config"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	//info version service
	fmt.Printf("Service: %s\nVersion: %s\n\nApp Info:\nApplication running at 0.0.0.0:"+port+"\n\n", os.Getenv("APP_NAME"), os.Getenv("APP_VER"))

	err := http.ListenAndServe(":"+port, config.Setup()) //Launch the app
	if err != nil {
		fmt.Print(err)
	}
}
