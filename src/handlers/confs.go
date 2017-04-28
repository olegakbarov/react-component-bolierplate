package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/olegakbarov/io.confs.api/src/db"
)

type Envelope struct {
	Result string      `json:"result"`
	Data   interface{} `json:"data"`
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func sendRespose(w http.ResponseWriter, data []byte) {
	// TODO: if ENV == develop
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(data)
}

func GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	recs, err := db.Read()

	log.Printf("DB rows: %v", recs)

	if err != nil {
		log.Fatal("Error quering the db- " + err.Error())
		w.WriteHeader(500)
		return
	}

	res := Envelope{
		Result: "OK",
		Data:   recs,
	}

	data, err := json.Marshal(res)

	if err != nil {
		log.Fatal("Failed marshaling json" + err.Error())
		w.WriteHeader(500)
		return
	}

	sendRespose(w, data)
}

func GetOne(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	rec, err := db.ReadOne(id)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(404)
			return
		}

		log.Fatal("Failed reading row" + err.Error())
		w.WriteHeader(500)
		return
	}

	res := Envelope{
		Result: "OK",
		Data:   rec,
	}

	data, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(500)
		return
	}

	sendRespose(w, data)
}

func DeleteOne(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	fmt.Printf("%s\n", id)

	if _, err := db.Remove(id); err != nil {
		log.Fatal("Failed deleting record" + err.Error())
		w.WriteHeader(500)
	}

	w.WriteHeader(204)
}

func Add(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var rec db.Conf

	err := decoder.Decode(&rec)
	fmt.Printf("%s\n", &rec)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	if _, err := db.Insert(rec); err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(201)
}

func Edit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var rec db.Conf

	err := decoder.Decode(&rec)
	fmt.Printf("%s\n", &rec)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	if _, err := db.Insert(rec); err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(201)
}

func NotImplemented(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Not Implemented"))
}
