package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../translater"
)

// Server is a struct
type Server struct {
	translater *translater.Translater
}

// OriginText is a struct
type OriginText struct {
	Text string `json:"text,omitempty"`
}

// TranslatedText is a struct
type TranslatedText struct {
	Origin     string `json:"origin,omitempty"`
	Translated string `json:"translated,omitempty"`
}

// Init is a function
func (server *Server) Init() {
	var translaterObject translater.Translater
	// イニシャライズ
	err := translaterObject.Init()
	if err != nil {
		fmt.Println("error")
		fmt.Println(err.Error())
		return
	}
	server.translater = &translaterObject
}

// PostTranslate is a function
func (server *Server) PostTranslate(w http.ResponseWriter, r *http.Request) {
	var originText OriginText
	json.NewDecoder(r.Body).Decode(&originText)
	result, err := server.translater.TranslateString(originText.Text)
	if err != nil {
		fmt.Println("error")
		fmt.Println(err.Error())
		return
	}
	json.NewEncoder(w).Encode(TranslatedText{
		Origin:     originText.Text,
		Translated: result,
	})
}
