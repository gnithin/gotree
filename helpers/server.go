package helpers

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func showFile(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "Hello World!")
}

type handlerStruct struct{}

var mux map[string]func(http.ResponseWriter, *http.Request)

func CreateServer() {
	// Create all the muxing right here
	mux = make(map[string]func(http.ResponseWriter, *http.Request))

	// This enlists that path handling
	mux["/"] = showFile
	addr := ":8000"

	server := http.Server{
		Addr:    addr,
		Handler: &handlerStruct{},
	}
	fmt.Println("Starting server at - ", addr)
	server.ListenAndServe()
}

func fileHandler(reqUrl string) (string, error) {
	// check the file
	fullPath, pathErr := filepath.Abs(reqUrl[1:])
	if pathErr != nil {
		fmt.Println("The filepath does not exist")
		fmt.Println(pathErr)
		return "", pathErr
	} else {
		fmt.Println("Getting the new file")
		fileContents, err := getFile(fullPath)
		if err != nil {
			fmt.Println("Error fetching the file")
			fmt.Println(err)
			return "", err
		} else {
			return string(fileContents), nil
		}
	}
}

func (h *handlerStruct) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("*************************************")
	reqUrl := req.URL.String()
	fmt.Println("Hitting - ", reqUrl)
	f, isExists := mux[reqUrl]
	if !isExists {
		fmt.Println("The url does not exist - ", reqUrl)
		respStr, err := fileHandler(reqUrl)
		if err == nil {
			io.WriteString(resp, respStr)
		}
	} else {
		// Calling the function
		f(resp, req)
	}
	fmt.Println("*************************************")
}

func getFile(filename string) ([]byte, error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return body, nil
}
