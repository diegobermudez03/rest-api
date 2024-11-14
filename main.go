package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Employee represents an employee with basic attributes
type Employee struct {
	ID      int    `json:"id"`
	Salary 	float32	`json:"salary"`
	Name    string `json:"name"`
	Role    string `json:"role"`
	Tech    string `json:"tech"`
	Email   string `json:"email"`
}


var employees = []Employee{
	{ID: 1, Name: "Juan Diego Bermudez", Salary: 14000000, Role: "Bakcned engineer", Tech: "AWS", Email: "juand.bermudez@javeriana.edu.co"},
	{ID: 2, Name: "Nicolas Quiles", Salary: 1425000, Role: "DevOps Engineer", Tech: "Azure", Email: "nic@gmail.com"},
	{ID: 3, Name: "Zion", Salary: 740000, Role: "DevOps Engineer", Tech: "AWS", Email: "liam@gmail.com"},
	{ID: 4, Name: "Liam Charlie", Salary: 30228995, Role: "Frontend developer", Tech: "React", Email: "zion@gmail.com"},
	{ID: 5, Name: "Miley Estefanny", Salary: 14100520, Role: "Networking engineer", Tech: "IP", Email: "miley@gmail.com"},
	{ID: 6, Name: "Johan Ceballos", Salary: 4100000, Role: "Solutions Architect", Tech: "AWS", Email: "johan@gmail.com"},
	{ID: 7, Name: "Jhon moore", Salary: 9800025, Role: "DevOps Engineer", Tech: "Jira", Email: "jhon@gmail.com"},
	{ID: 8, Name: "Steve Anita Smith", Salary: 136999, Role: "Solutions Architect", Tech: "Jira", Email: "steve@gmail.com"},
	{ID: 9, Name: "Francine Smith", Salary: 7458520, Role: "Frontend developer", Tech: "Next js", Email: "francine@gmail.com"},
	{ID: 10, Name: "Luke", Salary: 4570000, Role: "Networking engineer", Tech: "AWS", Email: "luke@gmail.com"},
}


func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	for _, employee := range employees {
		if employee.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(employee)
			return
		}
	}

	http.Error(w, "Employee not found", http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/employee/{id}", GetEmployeeByID).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
