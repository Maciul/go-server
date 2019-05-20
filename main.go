package main
import (
  "net/http"
  "strings"
  "fmt"
  "os"
  "log"
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

func main() {
  port := "8080"
  
  if (os.Getenv("PORT") != "" ) {
    port = os.Getenv("PORT")
  }

  if port == "" {
        log.Fatal("$PORT must be set")
  }

  http.HandleFunc("/dont/", sayNothing)
  http.HandleFunc("/", sayHello) 
  http.Handle("/serve-static-files/", http.StripPrefix("/serve-static-files/", http.FileServer(http.Dir("static/"))))
  
  if err := http.ListenAndServe( ":" + port, nil); err != nil {
    panic(err)
  }
}