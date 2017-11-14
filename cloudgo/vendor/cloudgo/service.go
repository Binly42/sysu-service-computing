package cloudgo

import (
	"fmt"
	"model"
	"net/http"
	"strings"
	log "util/logger"
)

const (
	DefaultPort = "8080"
)

func init() {
}

func LoadAll() {
	model.Load()
}
func SaveAll() {
	if err := model.Save(); err != nil {
		log.Error(err)
	}
}

type Cloudgo struct {
	*Server
}

func New() *Cloudgo {
	mux := NewServeMux()
	mux.HandleFunc("/", sayhelloName)

	server := NewServer()
	server.SetHandler(mux)

	cloudgo := new(Cloudgo)
	cloudgo.Server = server

	return cloudgo
}

func (cloudgo *Cloudgo) Listen(addr string) error {
	if addr == "" {
		addr = DefaultPort
	}
	return cloudgo.Server.Listen(addr)
}

// detail handlers, etc ... ----------------------------------------------------------------

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	/* fmt.Println(r.Form)
		fmt.Println("path", r.URL.Path)
		fmt.Println("scheme", r.URL.Scheme)
		fmt.Println(r.Form["url_long"])
		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
	    }
	*/

	segments := strings.Split(r.URL.Path, "/")
	name := segments[len(segments)-1]
	fmt.Fprintf(w, "Hello %v!\n", name)
}
