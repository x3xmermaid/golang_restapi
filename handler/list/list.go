package listController

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../../database"
	"../../models"
	"github.com/gorilla/mux"
)

const (
	SQLSelect = "select list.id, list.name, category.name from list"
	SQLInsert = "insert into list"
	SQLDelete = "Delete from list"
	SQLUpdate = "Update list set"
)

//ShowAllLists data
func ShowAllLists(w http.ResponseWriter, r *http.Request) {
	var List models.List
	var arrList []models.List
	var response models.ResponseList

	db := database.Connect()
	defer db.Close()

	var SQL bytes.Buffer
	SQL.WriteString(SQLSelect)

	// queryID := r.URL.Query().Get("idy")
	Search := r.URL.Query().Get("search")

	SQL.WriteString(" inner join category on list.category=category.id")

	if Search != "" {
		SQL.WriteString(" where list.name like '%" + Search + "%'")
	}

	fmt.Println(SQL.String())
	rows, err := db.Query(SQL.String())

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&List.ID, &List.Name, &List.CategoryName); err != nil {
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
	Execute(w, r, 1)
}

func getID(r *http.Request) (id int, err error) {
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
	Execute(w, r, 2)
}

//DeleteList for delete data in Delete List
func DeleteList(w http.ResponseWriter, r *http.Request) {
	Execute(w, r, 3)
}

func Execute(w http.ResponseWriter, r *http.Request, queryID int) {
	var response models.ResponseList

	var SQL bytes.Buffer
	var s []string
	db := database.Connect()
	defer db.Close()

	var List models.List
	decode := json.NewDecoder(r.Body)
	errBody := decode.Decode(&List)
	if errBody != nil {
		panic(errBody)
	}

	if queryID == 1 {
		SQL.WriteString(SQLInsert)
		s = append(s, List.CategoryName)
		s = append(s, List.Name)
		SQL.WriteString(" set category = ?, Name = ?")
		response.Message = "Data Successfully Inserted"
	} else if queryID == 2 {
		SQL.WriteString(SQLUpdate)
		s = append(s, List.CategoryName)
		s = append(s, List.Name)
		SQL.WriteString(" category = ?, Name = ? where id = ?")
		fmt.Println(SQL.String())
		response.Message = "Data Successfully Updated"
	} else if queryID == 3 {
		SQL.WriteString(SQLDelete)
		SQL.WriteString(" where id = ?")
		response.Message = "Data Successfully Deleted"
	}

	ParamsID, err := getID(r)

	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("ID doesnt inputed or something wrong with inputs"))
		log.Panic(err)
	} else {
		// execute update List
		if queryID == 3 {
			_, errBody = db.Exec(SQL.String(), ParamsID)
		} else if queryID == 2 {
			_, errBody = db.Exec(SQL.String(), s[0], s[1], ParamsID)
		} else {
			_, errBody = db.Exec(SQL.String(), s[0], s[1])
		}

		// check error or not
		if errBody != nil {
			log.Print(errBody)
			w.WriteHeader(401)
			w.Write([]byte("Something when error"))
			log.Panic(errBody)
		} else {
			response.Status = 1
			// response.Message = "Data Successfully updated"

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
	}
}

func Test(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm() // Parses the request body
	// // x := r.Form.Get("search")
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// bodyString := string(body)
	// querySearch := r.URL.Query()
	// var data = string
	data := r.URL.Path
	w.Write([]byte(trimLeftChar(data)))

}

func trimLeftChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return s[:0]
}
