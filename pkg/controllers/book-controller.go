package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alim7007/go-bookstore/pkg/models"
	"github.com/alim7007/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(resp http.ResponseWriter, req*http.Request){
	newBooks:=models.GetAllBooks()
	res,_:=json.Marshal(newBooks)
	resp.Header().Set("Content-type","pkglication/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(res)
}

func GetBookById(resp http.ResponseWriter, req*http.Request){
	vars:=mux.Vars(req)
	bookId:=vars["bookId"]
	ID,err:=strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _:=json.Marshal(bookDetails)
	resp.Header().Set("Content-type", "pkglication/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(res)
}

func CreateBook(resp http.ResponseWriter, req*http.Request){
	newBook := &models.Book{}
	utils.ParseBody(req,newBook)
	fmt.Println(newBook)
	b:=newBook.CreateBook()
	res,_:=json.Marshal(b)
	resp.WriteHeader(http.StatusOK)
	resp.Write(res)
}

func DeleteBook(resp http.ResponseWriter, req*http.Request){
	vars:=mux.Vars(req)
	bookId:=vars["bookId"]
	ID,err:=strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	book:=models.DeleteBook(ID)	
	res,_:=json.Marshal(book)
	resp.Header().Set("Content-type", "pkglication/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(res)
}

func UpdateBook(resp http.ResponseWriter, req*http.Request){
	changeBook := &models.Book{}
	utils.ParseBody(req,changeBook)
	vars:=mux.Vars(req)
	bookId:=vars["bookId"]
	ID,err:=strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)
	db.Model(&bookDetails).Updates(changeBook)
	// /////////
	// if changeBook.Name != ""{
	// 	bookDetails.Name = changeBook.Name
	// }
	// if changeBook.Author != ""{
	// 	bookDetails.Author = changeBook.Author
	// }
	// if changeBook.Publication != ""{
	// 	bookDetails.Publication = changeBook.Publication
	// }
	// db.Save(&bookDetails)
	// //////////
	res, _ := json.Marshal(bookDetails)
	resp.Header().Set("Content-type", "pkglication/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(res) 
}
