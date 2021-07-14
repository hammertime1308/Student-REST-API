package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"github.com/hammertime1308/Student-REST-API/database"
	"github.com/hammertime1308/Student-REST-API/model"
)

var student *model.Student

func GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var (
		id      string
		name    string
		class   int
		subject string
		age     int
		deleted bool
	)
	if err := database.Session.Query("SELECT * FROM student_db WHERE id=?", params["uid"]).Consistency(gocql.One).Scan(&id, &age, &class, &deleted, &name, &subject); err != nil {
		json.NewEncoder(w).Encode(struct {
			Err string `json:"error"`
		}{err.Error()})
		return
	}
	if deleted {
		json.NewEncoder(w).Encode(struct {
			Student [0]model.Student `json:"student"`
		}{})
		return
	}
	student = &model.Student{
		Uid:     id,
		Name:    name,
		Class:   class,
		Subject: subject,
		Age:     age,
	}
	json.NewEncoder(w).Encode(struct {
		Student *model.Student `json:"student"`
	}{student})
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var (
		id      string
		name    string
		class   int
		subject string
		age     int
		deleted bool
	)
	if err := database.Session.Query("SELECT * FROM student_db WHERE id=?", params["uid"]).Consistency(gocql.One).Scan(&id, &age, &class, &deleted, &name, &subject); err != nil {
		json.NewEncoder(w).Encode(struct {
			Err string `json:"error"`
		}{err.Error()})
		return
	}
	if deleted {
		json.NewEncoder(w).Encode(struct {
			Err string `json:"error"`
		}{"User is already deleted"})
		return
	}
	database.Session.Query("UPDATE student_db SET deleted = true WHERE id=?", params["uid"]).Consistency(gocql.One).Scan(&id, &age, &class, &deleted, &name, &subject)

	json.NewEncoder(w).Encode(struct {
		Err string `json:"message"`
	}{"Deleted successfully"})
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student1 model.Student
	_ = json.NewDecoder(r.Body).Decode(&student1)

	student1.Uid = gocql.TimeUUID().String()

	if err := database.Session.Query("INSERT INTO student_db(id,age,class,deleted,name,subject) values(?,?,?,false,?,?)", student1.Uid, student1.Age, student1.Class, student1.Name, student1.Subject).Consistency(gocql.One).Exec(); err != nil {
		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
		}{err.Error()})
		return
	}
	json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
		Id      string `json:"id"`
	}{"Created successfully", student1.Uid})
}
