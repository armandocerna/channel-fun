package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	//n := 3
	//
	http.HandleFunc("/", helloTest)
	fmt.Println("starting http server")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}

func goToSleep(in <-chan string, out chan<- string) {
	exec.Command("sleep", <-in).Run()
	output, err := exec.Command("ps", "aux").Output()
	if err != nil {
		panic("Error running command")
	}
	out <- string(output)
}

func helloTest(w http.ResponseWriter, r *http.Request) {
	in := make(chan string)
	out := make(chan string)

	go goToSleep(in, out)
	in <- r.URL.Query()["int"][0]
	fmt.Println(<-out)

}
