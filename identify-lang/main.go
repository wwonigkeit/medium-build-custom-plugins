package main

import (
	"github.com/abadojack/whatlanggo"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type Message struct {
	Text string `json:"text"`
}

type ReturnMessage struct {
	Language string `json:"language"`
	Script string `json:"script"`
	Confidence float64 `json:"confidence"`
}

const (
	DirektivActionIDHeader = "Direktiv-ActionID"
	DirektivErrorCodeHeader    = "Direktiv-ErrorCode"
	DirektivErrorMessageHeader = "Direktiv-ErrorMessage"
)

const code = "com.greeting-%s.error"

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", LanguageHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)

	go func() {
		<-sigs
		shutdown(srv)
	}()

	srv.ListenAndServe()
}

func LanguageHandler(w http.ResponseWriter, r *http.Request) {
	message := new(Message)
	aid := r.Header.Get(DirektivActionIDHeader)

	log(aid, "Reading Input")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithErr(w, fmt.Sprintf(code, "readdata"), err.Error())
		return
	}

	rdr := bytes.NewReader(data)
	dec := json.NewDecoder(rdr)

	dec.DisallowUnknownFields()

	log(aid, "Decoding Input")
	err = dec.Decode(message)
	if err != nil {
		respondWithErr(w, fmt.Sprintf(code, "decode"), err.Error())
		return
	}

	var output ReturnMessage
	info := whatlanggo.Detect(message.Text)
	output.Language = info.Lang.String()
	output.Script = whatlanggo.Scripts[info.Script]
	output.Confidence = info.Confidence

	marshalBytes, err := json.Marshal(output)
	if err != nil {
		respondWithErr(w, fmt.Sprintf(code, "marshal"), err.Error())
		return
	}

	log(aid, "Writing Output")
	respond(w, marshalBytes)
}

func shutdown(srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}

func log(aid, l string) {
	http.Post(fmt.Sprintf("http://localhost:8889/log?aid=%s", aid), "plain/text", strings.NewReader(l))
}

func respond(w http.ResponseWriter, data []byte) {
	w.Write(data)
}

func respondWithErr(w http.ResponseWriter, code, err string) {
	w.Header().Set(DirektivErrorCodeHeader, code)
	w.Header().Set(DirektivErrorMessageHeader, err)
}

