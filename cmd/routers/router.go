package routers

import (
	"api/cmd/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type LinkRequest struct {
	Link string `json:"link"`
}

// CreateLink godoc
// @Summary Create a new link
// @Description Create a new shortened link or perform some operation with the provided link
// @Tags links
// @Accept  json
// @Produce  json
// @Param link query string true "Link to be created"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Missing link parameter"
// @Router /link_create [post]
func CreateLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Получаем параметр link из URL
	link := r.URL.Query().Get("link")
	if link == "" {
		http.Error(w, "Missing link parameter", http.StatusBadRequest)
		return
	}

	db := database.InitDB("./example.db")
	defer db.Close()

	err := database.AddLink(db, link)
	if err != nil {
		log.Fatal("Error adding link:", err)
	}

	// Логируем или используем полученный линк
	log.Println("Received link:", link)

	output_code, err := database.GetCodeFromDB(db, link)

	if err != nil {
		log.Fatal("Error getting code:", err)
	}
	// Возвращаем ответ
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "success", "link": "` + "http://localhost:8000/" + output_code + `"}`))
}

// GetUser godoc
// @Summary Get user info
// @Description Get user information by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Router /{link} [get]
func GetLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)  // Get the variables from the request
	link := vars["link"] // Extract the 'link' parameter
	fmt.Println(w, "Link: %s", link)

	output, err := database.GetLinkFromDB(database.InitDB("./example.db"), link)
	if err != nil {
		log.Fatal("Error getting link:", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "success", "link": "` + output + `"}`))
}
