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

func showGraph(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("Show graph")
	// Check if it has params
	getFilePath := req.URL.Query().Get("filePath")
	if len(getFilePath) != 0 {
		filePath_, escErr := url.QueryUnescape(getFilePath)
		getFilePath = filePath_
		if escErr != nil {
			panic("Error when unescaping url")
		}
	} else {
		fmt.Println("Could not read get parameter, setting default")
		// Setting a default filePath
		//getFilePath = "assets/data/sample.json"
		getFilePath = "assets/data/autogen.json"
	}

	// Using the template
	t, templateErr := template.ParseFiles("assets/draw_graph.html")
	fmt.Println("Starting to template!!")
	if templateErr != nil {
		fmt.Println("Tempalte error")
		fmt.Println(templateErr)
	} else {
		type DummyStruct struct {
			FilePath string
		}
		dummy := DummyStruct{FilePath: getFilePath}
		fmt.Println("Dummt - ", dummy.FilePath)
		t.Execute(resp, dummy)
	}
}

type handlerStruct struct{}

var mux map[string]func(http.ResponseWriter, *http.Request)

func CreateServer() {
	// Create all the muxing right here
	mux = make(map[string]func(http.ResponseWriter, *http.Request))

	// This enlists that path handling
	mux["/"] = showFile
	mux["/graph"] = showGraph
	addr := ":8000"

	server := http.Server{
		Addr:    addr,
		Handler: &handlerStruct{},
	}
	fmt.Println("Starting server at - ", addr)
	server.ListenAndServe()
}

func fileHandler(reqUrl string) (string, error) {
	// Slicing because, the / at the start messes with getting the
	// absolute path
	reqUrl, pathErr := filepath.Abs(reqUrl)
	if pathErr != nil {
		fmt.Println("The filepath does not exist")
		fmt.Println(pathErr)
		return "", pathErr
	}

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

/*
This basic logic of this method is to find if any key matches the muxed
value. If not, then the current file directory is searched
*/
func (h *handlerStruct) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("*************************************")
	// This gets the url without the query params
	reqUrl := req.URL.Path
	fmt.Println("Hitting - ", req.URL.String())

	// Check if the url is already mapped
	f, isExists := mux[reqUrl]
	if !isExists {
		// If the path is not muxed, then find the path in the current
		// directory path.
		fmt.Println("The url does not exist, so fetching from file - ", reqUrl)

		// Get the response string
		respStr, err := fileHandler(reqUrl[1:])
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
