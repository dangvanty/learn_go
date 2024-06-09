package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "start server go neeee nooooo")
    })

    http.ListenAndServe(":3000", nil)
}
