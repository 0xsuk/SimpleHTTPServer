package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"
)

func Usage() string {
	usage := `
  Usage: ` + os.Args[0] + ` <port number>
  `
	return usage
}

func HandleBase(w http.ResponseWriter, r *http.Request) {
	base_dir, _ := os.Getwd()
	http.FileServer(http.Dir(base_dir)).ServeHTTP(w, r)
	rv := reflect.ValueOf(w).Elem()
	status := fmt.Sprintf("%v", rv.FieldByName("status"))
	time := time.Now()
	fmt.Printf("%v %v %v at %d/%d/%d %d:%d from %v\n", r.Method, r.URL.Path, status, time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(), r.RemoteAddr)
}

func main() {
	PORT := "8080"
	if len(os.Args) > 1 {
		PORT = os.Args[1]
		if _, err := strconv.Atoi(PORT); err != nil {
			log.Fatalln(Usage())
		}
	}
	fmt.Println("Listening on port", PORT)
	http.HandleFunc("/", HandleBase)

	log.Fatalln(http.ListenAndServe(":"+PORT, nil))

}
