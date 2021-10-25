package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/urfave/negroni"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type ServerConfig struct {
	Scheme        string `default:"http"`
	ListenAddress string `default:":8080"`
	PrivateKey    string `default:""`
	Certificate   string `default:""`
}

type httpErr struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type Message struct {
	Test string `json:"test"`
}

func NewWSServer(config ServerConfig) error {

	router := newRouter()
	n := negroni.Classic()

	n.UseHandler(router)
	if config.Scheme == "https" {
		return http.ListenAndServeTLS(config.ListenAddress, config.Certificate, config.PrivateKey, n)
	}

	return http.ListenAndServe(config.ListenAddress, n)
}

func handleErr(w http.ResponseWriter, err error, status int) {
	fmt.Println(err.Error())
	msg, err := json.Marshal(&httpErr{
		Msg:  err.Error(),
		Code: status,
	})
	if err != nil {
		msg = []byte(err.Error())
	}

	http.Error(w, string(msg), status)
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		handleErr(w, err, http.StatusInternalServerError)
		return
	}
	defer func() {
		err = c.Close()
	}()

	ms := Message{
		"Openned",
	}


	ms_json, _ := json.Marshal(ms)

	err = c.WriteMessage(websocket.TextMessage, []byte(ms_json))

	if err != nil {
		handleErr(w, err, http.StatusInternalServerError)
	}
	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			handleErr(w, err, http.StatusInternalServerError)
			break
		}
		if mt != websocket.TextMessage {
			handleErr(w, errors.New("only text message are supported;"), http.StatusNotImplemented)
			break
		}
		var v Message
		err = json.Unmarshal(msg, &v)
		if err != nil {
			handleErr(w, err, http.StatusInternalServerError)
			break
		}

		fmt.Println(v.Test)

		err = c.WriteMessage(mt, []byte(msg))
		if err != nil {
			handleErr(w, err, http.StatusInternalServerError)
			break
		}


	}
}

// NewRouter is the constructor for all my routes
func newRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.
		Methods("GET").
		Path("/ws").
		Name("Communication Channel").
		HandlerFunc(serveWs)

	router.
		Methods("GET").
		PathPrefix("/").
		Name("Static").
		Handler(http.FileServer(http.Dir("./htdocs")))
	return router
}
