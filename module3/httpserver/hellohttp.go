package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"io/ioutil"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Host	string `yaml:"host"`
	Port	int64  `yaml:"port"`
	Version string `yaml:"version"`
}


func Log(handler http.HandlerFunc) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf( "--> %s %s from %s, code: %d \n",r.Method, r.URL.Path, r.RemoteAddr, http.StatusOK)
        handler(w, r)
    })
}

func healthz(w http.ResponseWriter, r *http.Request) {
	// healthz test with code 200
	w.WriteHeader(200)
}

// get version
func version(w http.ResponseWriter, req *http.Request){
     	var setting Config
   	config, err := ioutil.ReadFile("./httpserverconfig.yaml")
   	if err != nil {
     		fmt.Print(err)
   	}
   	yaml.Unmarshal(config, &setting)
   	httpserver_version := string(setting.Version)
	w.Header().Set("VERSION", httpserver_version)
        w.WriteHeader(200)
}

// response request header 
func header(w http.ResponseWriter, req *http.Request){
	for name, headers := range req.Header {
        for _, h := range headers {
            //io.WriteString(w, fmt.Sprintf( "%v: %v\n", name, h))
		w.Header().Add(name, h)
        }
    }
}

func Default(w http.ResponseWriter, req *http.Request){
      fmt.Fprintln(w, "Hello world!")

}

func main() {

   mux := http.NewServeMux()

   // default request 
   mux.HandleFunc("/", Log(Default))

   // healthz test
   mux.HandleFunc("/healthz",  healthz)

   // get version
   mux.HandleFunc("/version", Log(version))

   // response request header 
   mux.HandleFunc("/header", Log(header))

   server := &http.Server{
      Addr:         ":8080",
      Handler:      mux,
   }
   // bind and start httpserver 
   go server.ListenAndServe()

   // elegant exit httpserver 
   listenSignal(context.Background(), server)
}

func listenSignal(ctx context.Context, httpSrv *http.Server) {
   sigs := make(chan os.Signal, 1)
   signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

   select {
   case <-sigs:
      log.Println("notify sigs")
      httpSrv.Shutdown(ctx)
      log.Println("http shutdown")
   }
}

