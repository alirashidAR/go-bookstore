package controllers

import(
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"fmt"
	"github.com/aliRashidAR/go-bookstore/pkg/utils"
	"github.com/aliRashidAR/go-bookstore/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r* http.Request){
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetBookById(w http.ResponseWriter, r * http.Request){
	vars := mux.Vars(r)
	bookId = vars["bookId"]
	
}