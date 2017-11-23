package cloudgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"model"
	"net/http"
	"path"
	"strings"
	"time"
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
	mux.HandleFunc("/api/test", apiTestHandler())

	mux.HandleFunc("/post-phone-info", showTablePhoneInfo)

	mux.HandleFunc("/unknown/", sayDeveloping)

	mux.HandleFunc("/say/", sayhelloName)

	mux.HandleFunc("/", showFormToPostPhoneInfo)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./asset/"))))

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

	segments := strings.Split(r.URL.Path, "/")
	name := segments[len(segments)-1]
	fmt.Fprintf(w, "Hello %v!\n", name)
}

func sayDeveloping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)

	fmt.Fprintf(w, "Developing!\n")
	fmt.Fprintf(w, "Now NotImplemented!\n")
}

func apiTestHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := struct {
			ID      string `json:"id"`
			Content string `json:"content"`
		}{ID: "9527", Content: "Hello from Go!\n"}

		// json.NewEncoder(w).Encode(res)
		j, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rand.Seed(time.Now().UnixNano())
		prettyPrint := rand.Float32() < 0.5
		if prettyPrint {
			var out bytes.Buffer
			json.Indent(&out, j, "", "\t")
			j = out.Bytes()
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	}
}

var tmplFilepath = path.Join("asset", "template", "phone-info-table.html")
var tableTemplatePhoneInfo = template.Must(template.ParseFiles(tmplFilepath))

type T struct {
	Items []Item
}

type Item map[string]interface{}

var args T

func showTablePhoneInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.Form) == 0 {
		return
	}

	fmt.Println(r.Form)
	w.Header().Set("Content-Type", "text/html")

	var item = make(Item)
	for k, v := range r.Form {
		item[k] = strings.Join(v, "")
	}
	args.Items = append(args.Items, item)

	err := tableTemplatePhoneInfo.Execute(w, args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err)
	}
}

func showFormToPostPhoneInfo(w http.ResponseWriter, r *http.Request) {
	f := path.Join("asset", "post-phone-info.html")
	res, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(res)
}
