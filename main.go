package main
import (
  "net/http"
  "strings"
  "fmt"
  "os"
)
func sayHello(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  message = "Hello " + message
  fmt.Fprintf(w, "Hello World!")
}

func sayNothing(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  message = "Hello " + message
  fmt.Fprintf(w, "Nothing")
}

func getPort() string{
    if ( os.Getenv("APP_ENV") == "production") {
      return ":80"
    } else {
      return "127.0.0.1:8080"
    }
}

func main() {
  http.HandleFunc("/dont/", sayNothing)
  http.HandleFunc("/", sayHello) 
  http.Handle("/serve-static-files/", http.StripPrefix("/serve-static-files/", http.FileServer(http.Dir("static/"))))
  
  if err := http.ListenAndServe( getPort(), nil); err != nil {
    panic(err)
  }
}