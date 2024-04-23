package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:(PASSWORD)@tcp(127.0.0.1:3306)/songs?charset=utf8mb4&parseTime=True&loc=Local"

type Songs struct {
	gorm.Model
	ArtistName string `json:"artistname"`
	SongtName  string `json:"songname"`
	Genre      string `json:"genre"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot Connect To Database")
	}
	DB.AutoMigrate(&Songs{})
}

func getSongs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var songs []Songs
	DB.Find(&songs)
	json.NewEncoder(w).Encode(songs)

}
func createSongs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var songs Songs
	json.NewDecoder(r.Body).Decode(&songs)
	DB.Create(&songs)
	json.NewEncoder(w).Encode(songs)

}

func getSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var songs Songs
	params := mux.Vars(r)
	DB.First(&songs, params["id"])
	json.NewEncoder(w).Encode(songs)

}

func deleteSongs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var songs Songs
	params := mux.Vars(r)
	DB.Delete(&songs, params["id"])
	json.NewEncoder(w).Encode(songs)
	json.NewEncoder(w).Encode("User Deleted Succesfully")
}

func updateSongs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var songs Songs
	params := mux.Vars(r)
	DB.First(&songs, params["id"])
	json.NewDecoder(r.Body).Decode(&songs)
	DB.Save(&songs)
	json.NewEncoder(w).Encode(songs)

}
