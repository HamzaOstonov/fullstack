package controllers

import (
	"fmt"
	"net/http"

	"github.com/victorsteven/fullstack/api/models"
	"github.com/victorsteven/fullstack/api/responses"
)

func (server *Server) GetTests(w http.ResponseWriter, r *http.Request) {

	test := models.Test{}

	tests, err := test.FindAllTests(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, tests)
}

func (server *Server) GetTestsAndVars(w http.ResponseWriter, r *http.Request) {

	test := models.Test{}

	tests, err := test.FindAllTests(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	item1 := models.Test{ID: 1, Title: "Hamza aka"}
	*tests = append(*tests, item1)

	fmt.Println(tests)
	responses.JSON(w, http.StatusOK, tests)
}
