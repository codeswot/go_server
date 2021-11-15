package genfunctions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/elmubarak/go_server/connection"
	studentModel "github.com/elmubarak/go_server/model"
	"github.com/elmubarak/go_server/studentDao"
	"github.com/julienschmidt/httprouter"
)

func CreateStudent(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	res.Header().Add("Content-Type", "application/json")
	newStudent := studentModel.Student{}
	json.NewDecoder(req.Body).Decode(&newStudent)
	insertQuery := fmt.Sprintf("insert into students(name,age) values('%v', '%v')", newStudent.Name, newStudent.Age)
	result, err := connection.Db.Exec(insertQuery)
	HandleErr(err)
	lid, err := result.LastInsertId()

	HandleErr(err)
	fmt.Println("insert result >", lid)
	newStudent.Id = lid
	json.NewEncoder(res).Encode(newStudent)

}

func UpdateStudent(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	res.Header().Add("Content-Type", "application/json")
	id, _ := strconv.Atoi(params[0].Value)
	student := studentModel.Student{}
	json.NewDecoder(req.Body).Decode(&student)
	rows, err := studentDao.UpdateOneFromDb(student.Name, student.Age, int64(id))
	HandleErr(err)
	students := []studentModel.Student{}
	for rows.Next() {
		var id int64
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		HandleErr(err)
		updatedStudent := studentModel.Student{id, name, age}
		students = append(students, updatedStudent)

	}

	err1 := json.NewEncoder(res).Encode(students)
	HandleErr(err1)
	json.NewEncoder(res).Encode(student)

}

func DeleteStudent(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	res.Header().Add("Content-Type", "application/json")
	id, _ := strconv.Atoi(params[0].Value)

	rows, err := studentDao.DeleteOneFromDb(int64(id))
	HandleErr(err)
	students := []studentModel.Student{}
	for rows.Next() {
		var id int64
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		HandleErr(err)
		deletedStudent := studentModel.Student{id, name, age}
		students = append(students, deletedStudent)
	}
	err1 := json.NewEncoder(res).Encode(students)
	HandleErr(err1)

}

func FetchById(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	res.Header().Add("Content-Type", "application/json")

	id, _ := strconv.Atoi(params[0].Value)

	rows, err := studentDao.DeleteOneFromDb(int64(id))
	HandleErr(err)
	students := []studentModel.Student{}
	for rows.Next() {
		var id int64
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		HandleErr(err)
		newStudent := studentModel.Student{id, name, age}
		students = append(students, newStudent)
	}
	err1 := json.NewEncoder(res).Encode(students)
	HandleErr(err1)

}

func FetchAll(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	res.Header().Add("Content-Type", "application/json")
	rows, err := studentDao.FetchAllFromDb()
	HandleErr(err)
	students := []studentModel.Student{}
	for rows.Next() {
		var id int64
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		HandleErr(err)
		newStudent := studentModel.Student{id, name, age}
		students = append(students, newStudent)
	}
	err1 := json.NewEncoder(res).Encode(students)
	HandleErr(err1)

}

func HandleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
