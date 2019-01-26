package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type rowStruct struct {
	ID      int    `json:"id"`
	PetID   int    `json:"petID"`
	Class   string `json:"class"`
	Family  string `json:"family"`
	Species string `json:"species"`
	ImgURL  string `json:"imgURL"`
}

func dbQuery(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	defer func() {
		StatCount("goserver.http.get.recommends.request", 1)
		StatTime("goserver.http.get.recommends.response_time", time.Since(t1))
	}()
	w.Header().Set("Content-Type", "application/json")
	searchQuery := strings.TrimPrefix(r.URL.Path, "/api/recommends/")
	v, err := strconv.Atoi(searchQuery)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	if v < 0 {
		w.WriteHeader(400)
		return
	}

	rows, err := db.Query("SELECT * FROM pets WHERE family = (SELECT family FROM pets WHERE pet_id = $1) LIMIT 20", v)
	if err != nil {
		panic(err)
	}
	finalData := make([]rowStruct, 0)

	for rows.Next() {
		var row rowStruct
		err := rows.Scan(&row.ID, &row.PetID, &row.Class, &row.Family, &row.Species, &row.ImgURL)
		if err != nil {
			fmt.Println("Error with db response", err)
		}
		finalData = append(finalData, row)
	}

	if len(finalData) == 0 {
		w.WriteHeader(204)
		return
	}

	petInfo, err := json.Marshal(finalData)
	if err != nil {
		panic(err)
	}
	w.Write(petInfo)
	return

}

func main() {
	db.SetMaxOpenConns(30)
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	http.HandleFunc("/api/recommends/", dbQuery)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
