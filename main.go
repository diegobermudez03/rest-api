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
	Name    string `json:"name"`
	Role    string `json:"role"`
	Tech    string `json:"tech"`
	Email   string `json:"email"`
}


var employees = []Employee{
	{ID: 1, Name: "Juan Diego Bermudez", Role: "Bakcned engineer", Tech: "AWS", Email: "juand.bermudez@javeriana.edu.co"},
	{ID: 2, Name: "Nicolas Quiles", Role: "DevOps Engineer", Tech: "Azure", Email: "nic@gmail.com"},
	{ID: 3, Name: "Zion", Role: "DevOps Engineer", Tech: "AWS", Email: "liam@gmail.com"},
	{ID: 4, Name: "Liam Charlie", Role: "Frontend developer", Tech: "React", Email: "zion@gmail.com"},
	{ID: 5, Name: "Miley Estefanny", Role: "Networking engineer", Tech: "IP", Email: "miley@gmail.com"},
	{ID: 6, Name: "Johan Ceballos", Role: "Solutions Architect", Tech: "AWS", Email: "johan@gmail.com"},
	{ID: 7, Name: "Jhon moore", Role: "DevOps Engineer", Tech: "Jira", Email: "jhon@gmail.com"},
	{ID: 8, Name: "Steve Anita Smith", Role: "Solutions Architect", Tech: "Jira", Email: "steve@gmail.com"},
	{ID: 9, Name: "Francine Smith", Role: "Frontend developer", Tech: "Next js", Email: "francine@gmail.com"},
	{ID: 10, Name: "Luke", Role: "Networking engineer", Tech: "AWS", Email: "luke@gmail.com"},
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
