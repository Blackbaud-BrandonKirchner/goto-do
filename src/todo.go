package main

type Todo struct {
  Title string `json:"name"`
  Notes string `json:"notes"`
  Completed bool `json:"completed"`
}

type Todos []Todo
