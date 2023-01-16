package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a bytye slice containing
// "Hello from Snippetbox" as the response body.

func home(w http.ResponseWriter, r *http.Request)  {
    // don't catch all /*
    if r.URL.Path != "/" {
        http.NotFound(w,r)
        return
    }

    w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request)  {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }

    fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request)  {
    // check method != POST
    if r.Method != "POST" {
        w.Header().Set("Allow", http.MethodPost)
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
    
    w.Write([]byte("Creat a new snippet..."))
}

func main()  {
    // handle requests
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet/view", snippetView)
    mux.HandleFunc("/snippet/create", snippetCreate)

    // Use the http.ListenAndServe() function to start a new web server. We pass in.
    // two parameters: the TCP network address to listen on (in this case ":4000")
    // and the servemux we just created. If the http.ListenAndServe() returns an error
    // we use the log.Fatal() function to log the error message and exit. Note
    // that any error returned by http.ListenAndServer() is always non-nil.

    log.Println("Starting server on :4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
