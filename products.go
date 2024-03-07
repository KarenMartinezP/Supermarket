package main

import (
	"encoding/json"
	"fmt"
	"os"
)


type Product struct {
	ID int
	Name string
	Quantity int 
	CodeValue string `json:"code_value"`
	IsPublished bool `json:"is_published"`
	Expiration string
	Price float64
}

var Products = []Product{}


func InsertFile(path string){
	file, err := os.Open(path)
    if err != nil {
        fmt.Println("Error al abrir el archivo:", err)
        return
    }

    defer file.Close()

	err = json.NewDecoder(file).Decode(&Products)
	if err != nil {
		fmt.Println("Error al decodificar el archivo JSON:", err)
		return
	}
	
}


