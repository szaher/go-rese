package main

import (
	"fmt"
	"srest"
	"flag"
	"net/http"
	"time"
	"encoding/json"
	"text/tabwriter"
	"os"
)

type Todo struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
	Due time.Time `json:"due"`
}

var Todos []Todo

func main() {
	fmt.Println("Welcome!")
	server := flag.Bool("server", false, "Enable this to work as server")
	action := flag.String("action", "list", "list, new, show")
	flag.Parse()
	if *server {
		fmt.Printf("Started as server %t ", *server)
		srest.Srest()
	}
	fmt.Println("Started as client!")
	fmt.Println(*action)
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	clnt := &http.Client{Transport:tr}
	if *action == "list"{
		resp, _ := clnt.Get("http://localhost:8080/todos")
		defer resp.Body.Close()

		err := json.NewDecoder(resp.Body).Decode(&Todos)
		if err != nil {
			panic(err)
		}
		// pretty print it in table!
		w := tabwriter.NewWriter(os.Stdout, 0, 20, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
		fmt.Fprintln(w, "ID \tName \tCompleted \tDue " )
		for _, t  := range Todos{
			row := fmt.Sprintf("%d \t %s \t %t \t %s ", t.Id, t.Name, t.Completed, t.Due)
			fmt.Fprintln(w, row)
		}
		// fmt.Println(Todos)
		w.Flush()

	}




}
