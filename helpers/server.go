package helpers

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
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
	fmt.Println("Getting the new file")
	fileContents, err := getFile(reqUrl)
	if err != nil {
		fmt.Println("Error fetching the file")
		fmt.Println(err)
		return "", err
	} else {
		return string(fileContents), nil
	}

}

func (h *handlerStruct) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("*************************************")
	reqUrl := req.URL.Path
	fmt.Println("Hitting - ", req.URL.String())
	f, isExists := mux[reqUrl]
	if !isExists {
		fmt.Println("The url does not exist, so fetching from file - ", reqUrl)
		reqUrl, pathErr := filepath.Abs(reqUrl[1:])
		if pathErr != nil {
			fmt.Println("The filepath does not exist")
			fmt.Println(pathErr)
			return
		}
		respStr, err := fileHandler(reqUrl)
		if err == nil {
			// Check if it has params
			filePath := req.URL.Query().Get("filePath")
			if len(filePath) != 0 {
				// Host the template
				filePath, escErr := url.QueryUnescape(filePath)
				if escErr != nil {
					fmt.Println("Error when unescaping url")
				} else {
					t, templateErr := template.ParseFiles(reqUrl)
					fmt.Println("Starting to template!!")
					if templateErr != nil {
						fmt.Println("Tempalte error")
						fmt.Println(templateErr)
					} else {
						type DummyStruct struct {
							filePath string
						}
						dummy := DummyStruct{filePath: filePath}
						fmt.Println("Dummt - ", dummy.filePath)
						t.Execute(resp, dummy)
					}
				}
			} else {
				io.WriteString(resp, respStr)
			}
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
