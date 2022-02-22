package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/library/models"
	"github.com/library/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {

	database, client, ctx, cancel, err := models.Connect(utils.DbURL)
	if err != nil {
		http.Error(w, "Error while connecting to database", http.StatusInternalServerError)
		return
	}

	defer models.Close(client, ctx, cancel)

	cursor, err := database.Collection(utils.Books).Find(ctx, bson.M{})
	if err != nil {
		log.Println("collections could not be connected.")
		http.Error(w, "Error while connecting to database", http.StatusInternalServerError)
		return
	}

	var books []models.Book
	err = cursor.All(ctx, &books)
	if err != nil {
		log.Println("collections could not find data: ",err)
		http.Error(w, "Error while finding books", http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func GetBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database, client, ctx, cancel, err := models.Connect(utils.DbURL)
	if err != nil {
		log.Println("error: could not connect to Database")
		http.Error(w, "Error while connecting to database", http.StatusInternalServerError)
		return
	}

	defer models.Close(client, ctx, cancel)

	idVar := vars["id"]
	id, err := strconv.Atoi(idVar)
	if err != nil {
		log.Println("error: non int value")
		http.Error(w, "Error: non int ID provided", http.StatusBadRequest)
		return
	}

	var book models.Book
	err = database.Collection(utils.Books).FindOne(ctx, bson.M{"_id": id}).Decode(&book)
	if err != nil {
		log.Println("error: could not find ID")
		http.Error(w, "Error while connecting to database", http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	var book models.Book
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&book)
	if err != nil {
		log.Println("error while decoding: ", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	database, client, ctx, cancel, err := models.Connect(utils.DbURL)
	if err != nil {
		log.Println("error: could not connect to Database")
		http.Error(w, "Error while connecting to database", http.StatusInternalServerError)
		return
	}

	defer models.Close(client, ctx, cancel)

	_, err = database.Collection(utils.Books).InsertOne(ctx, book)
	if err != nil && strings.Contains(err.Error(), "E11000") {
		log.Println("error: could not connect to Database. Err: ", err)
		http.Error(w, "Error record with same ID present", http.StatusConflict)
		return
	}
	if err != nil {
		log.Println("error: could not connect to Database. Err: ", err)
		http.Error(w, "Error while connecting to database", http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
