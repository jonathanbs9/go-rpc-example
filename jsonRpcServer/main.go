package main

import (
	jsonparse "encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

type Args struct {
	ID string
}

// Book Struct holds Book JSON structure
type Book struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
}

type JSONServer struct{}

// Give Book detail is RPC implementation
func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error {
	var books []Book
	// Read JSON file and load data
	absPath, _ := filepath.Abs("jsonRpcServer/books.json")
	raw, reader := ioutil.ReadFile(absPath)
	if reader != nil {
		log.Println("Error del reader: ", reader)
		os.Exit(1)
	}

	// Unmarshal JSON raw data into books array
	marshaller := jsonparse.Unmarshal(raw, &books)
	if marshaller != nil {
		log.Println("Error del marshaler: ", reader)
		os.Exit(1)
	}

	// Iterateover each book to find the given book
	for _, book := range books {
		if book.ID == args.ID {
			// If book found, fill the reply with it
			*reply = book
			break
		}
	}
	return nil
}

func main() {
	// Create a new RPC Server
	s := rpc.NewServer()
	// register the type of data requested as JSON
	s.RegisterCodec(json.NewCodec(), "application/json")
	//Register the service by creating a new JSON Server
	s.RegisterService(new(JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":1234", r)

}
