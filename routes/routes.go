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
	insertQuery := fmt.Sprintf("insert into students(fName, lName, email, gender, age) values('%v', '%v', '%v','%v','%v')", newStudent.Fname, newStudent.Lname, newStudent.Email, newStudent.Gender, newStudent.Age)
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
	rows, err := studentDao.UpdateOneFromDb(int64(id), student.Fname, student.Lname, student.Email, student.Gender, student.Age)
	HandleErr(err)
	students := []studentModel.Student{}
	for rows.Next() {
		var id int64
		var fName string
		var lName string
		var email string
		var gender string
		var age int
		err := rows.Scan(&id, &fName, &lName, &email, &gender, &age)
		HandleErr(err)
		updatedStudent := studentModel.Student{id, fName, lName, email, gender, age}
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
		var fName string
		var lName string
		var email string
		var gender string
		var age int
		err := rows.Scan(&id, &fName, &lName, &email, &gender, &age)
		HandleErr(err)
		deletedStudent := studentModel.Student{id, fName, lName, email, gender, age}
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
		var fName string
		var lName string
		var email string
		var gender string
		var age int
		err := rows.Scan(&id, &fName, &lName, &email, &gender, &age)
		HandleErr(err)
		newStudent := studentModel.Student{id, fName, lName, email, gender, age}
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
		var fName string
		var lName string
		var email string
		var gender string
		var age int
		err := rows.Scan(&id, &fName, &lName, &email, &gender, &age)
		HandleErr(err)
		newStudent := studentModel.Student{id, fName, lName, email, gender, age}
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
