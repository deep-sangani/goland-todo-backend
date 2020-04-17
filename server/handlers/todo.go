package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"../models"
)

type Todos struct{}

var todoModel models.Todo = models.Todo{}

// NewTodo creates an instance of todos
func NewTodo() *Todos {
	return &Todos{}
}
func (p *Todos) CreateTodo(rw http.ResponseWriter, r *http.Request) {
	// newTodo := models.Todo{
	// 	Task: "teste",
	// }

	// err := todoModel.FromJSON(r.Body)
	log.Println("model", todoModel)

	err := json.NewDecoder(r.Body).Decode(&todoModel)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		log.Println("error", err)
		return
	}

	todoModel.InsertOne()

	rw.WriteHeader(http.StatusCreated)
}

func (todo *Todos) GetTodos(rw http.ResponseWriter, r *http.Request) {
	todoModel, _ := todoModel.GetAll()
	err := json.NewEncoder(rw).Encode(todoModel)
	// err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unabble to marshal json", http.StatusInternalServerError)
	}
}
