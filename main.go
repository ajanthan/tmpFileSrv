// main
package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"encoding/json"
	
)

const DIR = "/tmp/tmpFiles/"
type Artifact struct {
	Name string
	Size float64
	CreatedOn string	
}

func serveFiles(responceWriter http.ResponseWriter, request *http.Request) {
	method := request.Method
	//fileName := strings.Split(request.URL.Path, "/")[2]
	//buff:=make([]byte,1024)
	switch method {
	case "POST":
		reader, _ := request.MultipartReader()
		part, e := reader.NextPart()
		if e == io.EOF {
				responceWriter.WriteHeader(500)
				return
			}
		fileName := part.FileName()
		file, e := os.Create(DIR + fileName)
		defer file.Close()
		if e != nil {
			log.Println("Can not create file ", e.Error())
			responceWriter.WriteHeader(500)
			return
		}
		if fileName != "" {
				io.Copy(file, part)
			}
		responceWriter.WriteHeader(200)
			

		/*for {
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
			*/

			//ioutil.ReadAll()

		//}

	case "GET":
	fileName := strings.Split(request.URL.Path, "/")[2]
		if fileName != "" {
			name := (DIR + fileName)

			http.ServeFile(responceWriter, request, name)
		}
		files, _ := ioutil.ReadDir(DIR)
		//responceWriter.Write([]byte(" { "))
		artifacts :=make ([]Artifact,0)
		for _, file := range files {
			//responceWriter.Write([]byte("\"files\": " + "\""+file.Name()+"\""))
			artifact:=Artifact{Name:file.Name(),Size:float64(file.Size()/1000000.0),CreatedOn:file.ModTime().Local().String()}
			artifacts=append(artifacts,artifact)
		}
		b,_:=json.Marshal(artifacts)
		responceWriter.Write(b)
	case "DELETE":
	fileName := strings.Split(request.URL.Path, "/")[2]
	log.Println("Filename ",fileName)
		if file,_ := os.Open(DIR + fileName); file !=nil {
			e := os.Remove(file.Name())
			if e != nil {
				log.Println("Error while deleting ", e.Error())
				responceWriter.WriteHeader(500)
			}
			log.Println("Deleted file ",file.Name())
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
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("Port must be not empty")
	}
	http.HandleFunc("/files/", serveFiles)
	http.Handle("/public/",http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.Handle("/",http.RedirectHandler("/public/index.html",http.StatusFound))
	log.Println("starting the server on port:", port)
	er := http.ListenAndServe(":"+port, nil)
	if er != nil {
		log.Fatalln("Could not start the server")
	}
}
