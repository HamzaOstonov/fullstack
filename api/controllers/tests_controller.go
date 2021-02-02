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

func (server *Server) GetTestsVars(w http.ResponseWriter, r *http.Request) {

	testvar := models.Testvar{}

	testvars, err := testvar.FindAllTestvars(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, testvars)
}

func (server *Server) GetTest(w http.ResponseWriter, r *http.Request) {

	general := models.General{}

	test := models.Test{}
	tests, err := test.FindAllTests(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Println(tests)

	quests := []*models.Question{}
	item1 := models.Question{ID: 1, Text: "Hamza aka", Type: "mc"}

	anss := []*models.Answer{}
	item2 := models.Answer{VariantText: "Javob1"}
	anss = append(anss, &item2)
	item2 = models.Answer{VariantText: "Javob2"}
	anss = append(anss, &item2)
	item1.Answers = anss

	ss := []*string{}
	var numberString string
	numberString = "hamza1"
	ss = append(ss, &numberString)
	numberString = "hamza2"
	ss = append(ss, &numberString)
	item1.Javoblar = ss

	s1 := []string{}
	s1 = append(s1, "javob 3")
	s1 = append(s1, "javob 4")
	s1 = append(s1, "javob 5")
	item1.Javob = s1

	quests = append(quests, &item1)
	fmt.Println(&quests)

	item1 = models.Question{ID: 2, Text: "Hamza aka", Type: "mc"}
	item2 = models.Answer{VariantText: "Javob11"}
	anss = append(anss, &item2)
	item2 = models.Answer{VariantText: "Javob22"}
	anss = append(anss, &item2)
	item1.Answers = anss
	quests = append(quests, &item1)
	general.Questions = quests

	testvar := models.Testvar{}
	testvars, err := testvar.FindAllTestvars(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Println(testvars)

	responses.JSON(w, http.StatusOK, general)
}
