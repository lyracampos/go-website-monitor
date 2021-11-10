package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"website-monitor/domain/entities"
)

func List(w http.ResponseWriter, r *http.Request) {
	log.Println("websites - list")
}

func Get(w http.ResponseWriter, r *http.Request) {
	log.Println("websites - get")
}

func Add(w http.ResponseWriter, r *http.Request) {

	requestBody, _ := ioutil.ReadAll(r.Body)
	var webSite entities.WebSite
	json.Unmarshal(requestBody, &webSite)

	log.Println(webSite)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(webSite)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	log.Println("websites - edit")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("websites - delete")
}
