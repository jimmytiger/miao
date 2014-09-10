package main

import (
	"fmt"
	"github.com/jimmytiger/miao/layout"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	readCookie, err := r.Cookie("haha")
	if err == nil && readCookie.Value != "" {
		fmt.Fprintf(w, "Loged in")
		log.Println("Loged in")
	} else {
		cookie := http.Cookie{Name: "haha", Value: "good", Path: "/"}
		http.SetCookie(w, &cookie)

	}
	fmt.Fprintf(w, r.URL.Path)
	fmt.Fprintf(w, "Hi, there, I love %s!", r.URL.Path[1:])
}

func main() {
	p := &layout.Page{"fds", []byte("fdfads")}
	p.Save()
	log.Println("Start...")
	http.Handle("/css/", http.FileServer(http.Dir("static")))
	http.Handle("/js/", http.FileServer(http.Dir("static")))

	http.HandleFunc("/", handler)
	er := http.ListenAndServe(":8081", nil)
	if er != nil {
		fmt.Println(er)
	}
}
