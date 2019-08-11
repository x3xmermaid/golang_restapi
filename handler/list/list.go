package listController

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"../../database"
	"../../models"
	"github.com/gorilla/mux"
)

//ShowAllLists data
func ShowAllLists(w http.ResponseWriter, r *http.Request) {
	var List models.List
	var arrList []models.List
	var response models.ResponseList

	db := database.Connect()
	defer db.Close()

	SQL := "select * from list"

	queryID := r.URL.Query().Get("id")
	querySearch := r.URL.Query().Get("search")

	if queryID != "" {

		SQL = "select * from list where id = '" + queryID + "'"

	} else if querySearch != "" {

		SQL = "select * from list where name like '%" + querySearch + "%'"

	}

	rows, err := db.Query(SQL)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&List.ID, &List.Name, &List.CategoryID); err != nil {
			log.Fatal(err.Error())

		} else {
			arrList = append(arrList, List)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arrList

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

//InsertList insert data to table list
func InsertList(w http.ResponseWriter, r *http.Request) {
	// var arrList []models.List
	var response models.ResponseList

	db := database.Connect()
	defer db.Close()

	var List models.List
	decode := json.NewDecoder(r.Body)
	errBody := decode.Decode(&List)
	if errBody != nil {
		panic(errBody)
	}

	_, errBody = db.Exec("insert into List (id, Name, category) values (?, ?, ?)", List.ID, List.Name, List.CategoryID)

	// check error or not
	if errBody != nil {
		log.Print(errBody)
		w.WriteHeader(401)
		w.Write([]byte("Something when error"))
	} else {
		response.Status = 1
		response.Message = "Data Successfully added"
		log.Print("Data inserted to List")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

}

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

//UpdateList for update data in yabel list
func UpdateList(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseList

	db := database.Connect()
	defer db.Close()

	var List models.List
	decode := json.NewDecoder(r.Body)
	errBody := decode.Decode(&List)
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
		// execute update List
		_, errBody = db.Exec("UPDATE List set category = ?, Name = ? where id = ?", List.CategoryID, List.Name, ParamsID)

		// check error or not
		if errBody != nil {
			log.Print(errBody)
			w.WriteHeader(401)
			w.Write([]byte("Something when error"))
			log.Panic(errBody)
		} else {
			response.Status = 1
			response.Message = "Data Successfully updated"
			log.Print("Data updated to List")

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
	}

}

//DeleteList for delete data in Delete List
func DeleteList(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseList

	db := database.Connect()
	defer db.Close()

	ParamsID, err := getVarsID(r)

	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("ID doesnt inputed or something wrong with inputs"))
		log.Panic(err)
	} else {

		_, err = db.Exec("DELETE from List where id = ?", ParamsID)

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

func Test(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Parses the request body
	x := r.Form.Get("search")
	// querySearch := r.URL.Query()

	w.Write([]byte(x))
}
