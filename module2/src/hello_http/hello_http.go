package main

import (

	"fmt"
	"io"
	"log"
	"net/http"
	"os"

)


func Log(handler http.HandlerFunc) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf( "%s %s %d %s\n", r.RemoteAddr, r.Method, http.StatusOK, r.URL)
        handler(w, r)
    })
}

func main() {
	
	http.HandleFunc("/header", Log(reqHeader))
	http.HandleFunc("/version", Log(getVersion))

	http.HandleFunc("/",  Log(Default))

	if err := http.ListenAndServe(":8080", nil);err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Default(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello World... GO ...")
}


func reqHeader(w http.ResponseWriter, req *http.Request){
	for name, headers := range req.Header {
        for _, h := range headers {
            io.WriteString(w, fmt.Sprintf( "%v: %v\n", name, h))		
        }
    }
}

func getVersion(w http.ResponseWriter, req *http.Request){
	version := os.Getenv("VERSION")
	w.Header().Add("VERSION", version)
	io.WriteString(w, fmt.Sprintf("Version: %s\n", version))

}
