package main

import (
	"fmt"
	//"html"
	"log"
	"os/exec"
        //"reflect"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "30b8d47eae26a3b0bf989cbbb199d29eca19dc133a9ed208c455874c")
	// 	//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/images/{todoId}", TodoShow)

	

	log.Fatal(http.ListenAndServe(":8080", router))

}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "30b8d47eae26a3b0bf989cbbb199d29eca19dc133a9ed208c455874c")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	todoId := vars["todoId"]
	app := "docker"

	arg0 := "inspect"

	// arg1 := os.Args[1]

	cmd := exec.Command(app, arg0, todoId)
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

    fmt.Fprintln(w,  string(stdout[:]))
}

func Cmd(cmd string, shell bool) []byte {
	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			panic("some error found")
		}
		return out
	} else {
		out, err := exec.Command(cmd).Output()
		if err != nil {
			panic("some error found")
		}
		return out
	
	}
}
