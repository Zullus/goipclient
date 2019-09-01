package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
)

//Saida [é a saída de dados para o app]
type Saida struct {
	Datahora string `json:"datahora"`
	Ip       string `json:"ip"`
}

//IndexPage eh a IndexPage
func IndexPage(w http.ResponseWriter, r *http.Request) {

	//Exibe data
	t := dataAtual()

	fmt.Fprintf(w, t+"\n\n")

	// get client ip address
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	// print out the ip address
	fmt.Fprintf(w, ip+"\n\n")

	// sometimes, the user acccess the web server via a proxy or load balancer.
	// The above IP address will be the IP address of the proxy or load balancer and not the user's machine.

	// let's get the request HTTP header "X-Forwarded-For (XFF)"
	// if the value returned is not null, then this is the real IP address of the user.
	fmt.Fprintf(w, "X-Forwarded-For :"+r.Header.Get("X-FORWARDED-FOR")+"\n\n")

	dadossaida := Saida{t, ip}

	j, err := json.Marshal(dadossaida)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(w, "dados de saída: ", string(j))
}

func dataAtual() string {

	t := time.Now()

	s := t.Format("2006-01-02 15:04:05")

	return s

}

func main() {

	fmt.Println("Iniciando servidor Web")

	http.HandleFunc("/", IndexPage)
	http.ListenAndServe(":8080", nil)
}
