package app

import (
	"encoding/xml"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab4/internal/controller"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab4/internal/model"
	"log"
	"net/http"
)

func RPCHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received POST request")

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var args model.MatrixArgs
	err := xml.NewDecoder(r.Body).Decode(&args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Обработка матрицы
	reply := controller.ProcessMatrix(args.Matrix)

	log.Println("Sending GET response with results")

	w.Header().Set("Content-Type", "text/xml")
	err = xml.NewEncoder(w).Encode(reply)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
