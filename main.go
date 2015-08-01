// main
package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const DIR = "/tmp/tmpFiles/"

func serveFiles(responceWriter http.ResponseWriter, request *http.Request) {
	method := request.Method
	fileName := strings.Split(request.URL.Path, "/")[2]
	//buff:=make([]byte,1024)
	switch method {
	case "POST":
		reader, _ := request.MultipartReader()
		file, e := os.Create(DIR + fileName)
		defer file.Close()
		if e != nil {
			log.Println("Can not create file ", e.Error())
			responceWriter.WriteHeader(500)
			return
		}

		for {
			part, e := reader.NextPart()

			if e == io.EOF {
				responceWriter.WriteHeader(201)
				return
			}
			fileName := part.FileName()
			log.Println("Read ", part)
			if fileName != "" {
				io.Copy(file, part)
			}

			//ioutil.ReadAll()

		}

	case "GET":
		if fileName != "" {
			name := (DIR + fileName)

			http.ServeFile(responceWriter, request, name)
		}
		files, _ := ioutil.ReadDir(DIR)
		responceWriter.Write([]byte(" { "))
		for _, file := range files {
			responceWriter.Write([]byte("\"files\": " + file.Name()))
		}
		responceWriter.Write([]byte(" } "))
	case "DELETE":
		if file, e := os.Open(DIR + fileName); os.IsExist(e) {
			e := os.Remove(file.Name())
			if e != nil {
				log.Println("Error while deleting ", e.Error())
				responceWriter.WriteHeader(500)
			}
			responceWriter.WriteHeader(200)
		}
	default:
		responceWriter.Write([]byte("Un supported method"))
		responceWriter.WriteHeader(500)
	}

}

func main() {
	if _, e := os.Open(DIR); os.IsNotExist(e) {
		os.Mkdir(DIR, os.ModePerm)
		log.Println("Making tmp directory")
	}
	http.HandleFunc("/files/", serveFiles)
	er := http.ListenAndServe(":8080", nil)
	if er != nil {
		log.Fatalln("Could not start the server")
	}
}
