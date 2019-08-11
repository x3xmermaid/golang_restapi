package categoryController

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"../../database"
	"../../models"
	"github.com/gorilla/mux"
)

// showAllItems showing all field table of item
func ShowAllCategory(w http.ResponseWriter, r *http.Request) {
	var Category models.Category
	var arrCategory []models.Category
	var response models.ResponseCategory

	db := database.Connect()
	defer db.Close()

	SQL := "select * from Category"

	queryID := r.URL.Query().Get("id")
	querySearch := r.URL.Query().Get("search")

	if queryID != "" {

		SQL = "select * from Category where id = '" + queryID + "'"

	} else if querySearch != "" {

		SQL = "select * from Category where name like '%" + querySearch + "%'"

	}

	rows, err := db.Query(SQL)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&Category.ID, &Category.Name); err != nil {
			log.Fatal(err.Error())

		} else {
			arrCategory = append(arrCategory, Category)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arrCategory

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

// InsertItemData insert data into item tables
func InsertCategory(w http.ResponseWriter, r *http.Request) {
	// var arrItem []models.Item
	var response models.ResponseCategory

	db := database.Connect()
	defer db.Close()

	var Category models.Category
	decode := json.NewDecoder(r.Body)
	errBody := decode.Decode(&Category)
	if errBody != nil {
		panic(errBody)
	}

	_, errBody = db.Exec("insert into category (id, name) values ('', ?)", Category.Name)

	// check error or not
	if errBody != nil {
		log.Print(errBody)
		w.WriteHeader(401)
		w.Write([]byte("Something when error"))
	} else {
		response.Status = 1
		response.Message = "Data Successfully added"
		log.Print("Data inserted to Item")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

}

// getVarsID get an params from URLs
func getVarsID(r *http.Request) (id int, err error) {
	vars := mux.Vars(r)
	if val, ok := vars["id"]; ok {
		convertedVal, err := strconv.Atoi(val)
		if err != nil {
			return id, err
		}
		id = convertedVal
	}
	return
}

// updateItemData update an item data
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseCategory

	db := database.Connect()
	defer db.Close()

	var Category models.Category
	decode := json.NewDecoder(r.Body)
	errBody := decode.Decode(&Category)
	if errBody != nil {
		panic(errBody)
	}

	// get the params from URL
	ParamsID, errParams := getVarsID(r)

	if errParams != nil {
		w.WriteHeader(401)
		w.Write([]byte("ID not inserted or something wrong with inputs"))
		log.Panic(errParams)
	} else {
		// execute update item
		_, errBody = db.Exec("UPDATE category set name = ? where id = ?", Category.Name, ParamsID)

		// check error or not
		if errBody != nil {
			log.Print(errBody)
			w.WriteHeader(401)
			w.Write([]byte("Something when error"))
			log.Panic(errBody)
		} else {
			response.Status = 1
			response.Message = "Data Successfully updated"
			log.Print("Data updated to Item")

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
	}

}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseCategory

	db := database.Connect()
	defer db.Close()

	ParamsID, err := getVarsID(r)

	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("ID doesnt inputed or something wrong with inputs"))
		log.Panic(err)
	} else {

		_, err = db.Exec("DELETE from category where id = ?", ParamsID)

		if err != nil {

			w.WriteHeader(500)
			w.Write([]byte("failed to delete data"))
			log.Panic(err)

		} else {

			response.Status = 1
			response.Message = "Success Delete Data"
			log.Print("Delete data to database")

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		}
	}

}
