package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type Book struct{
	Title string `json:"title"`
	Author string `json:"author"`
	Pages int `json:"pages"`
}

func main() {
	
	http.HandleFunc("/header/", reqHeader)
	http.HandleFunc("/version/", getVersion)
	http.HandleFunc("/book", getBook)

	http.HandleFunc("/",  Default)
	http.HandleFunc("/time/", timePrint)

	http.ListenAndServe(":8080", nil)
}

func Default(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello World... GO ...")
}

func timePrint(w http.ResponseWriter, r *http.Request){
	t := time.Now()
	timeStr := fmt.Sprintf("{\"time\": \"%s\"}", t)
	w.Write([]byte(timeStr))
}

func reqHeader(w http.ResponseWriter, req *http.Request){
	
	for name, headers := range req.Header {
		var hkey  strings.Builder
		var hval strings.Builder
		hkey.WriteString(name)
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
			hval.WriteString(h )
        }
		w.Header().Set(hkey.String(), hval.String())
		//fmt.Println(hkey.String())
    }
}

func getVersion(w http.ResponseWriter, req *http.Request){
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	
}



func getBook(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "Application-jsion")
	book := Book{
		Title: "茶花女",
		Author: "小仲马",
		Pages: 299,
	}

	json.NewEncoder(w).Encode(book)
}