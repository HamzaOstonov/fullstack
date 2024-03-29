package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/victorsteven/fullstack/api/models"
	"github.com/victorsteven/fullstack/api/responses"
	"github.com/victorsteven/fullstack/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
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
	fmt.Println((*tests)[0].Title)

	quests := []models.Question{}

	for i := 0; i < len(*tests); i++ {

		item1 := models.Question{Text: (*tests)[i].Title, Type: "mc"}

		ss := []string{}
		var numberString string
		numberString = "Javob ha"
		ss = append(ss, numberString)
		numberString = "javob yuq"
		ss = append(ss, numberString)
		item1.Answers = ss

		//err := db.Debug().Model(&User{}).Where("id = ?", posts[i].AuthorID).Take(&posts[i].Author).Error
		//	if err != nil {
		//		return &[]Post{}, err
		//	}

		quests = append(quests, item1)

	}

	//item1 := models.Question{ID: 1, Text: "Hamza akani 1-savoli", Type: "mc"}
	//item1 = models.Question{ID: 2, Text: "Hamza aka ni 2-savoli", Type: "mc"}
	//ss = []string{}
	//numberString = "Javob ha 2-savol"
	//ss = append(ss, numberString)
	//numberString = "javob yuq 2-savol"
	//ss = append(ss, numberString)
	//item1.Answers = ss
	//quests = append(quests, item1)

	general.Questions = quests

	testvar := models.Testvar{}
	testvars, err := testvar.Find2(server.DB, 2)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Println(testvars)

	responses.JSON(w, http.StatusOK, general)
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (server *Server) GetKattaTest(w http.ResponseWriter, r *http.Request) {

	//general := models.General{}
	kattatest := models.Kattatest{}
	savollar := []models.Savol{}
	savol := models.Savol{}
	javoblar := []models.Javob{}
	javob := models.Javob{}

	test := models.Test{}
	//-----------
	savol.Savolnum = "001"
	savol.Savoltext = "tugri javobni topingda"

	javob.Javobnum = "1"
	javob.Javobtext = "Javob 1"
	javoblar = append(javoblar, javob)

	javob.Javobnum = "2"
	javob.Javobtext = "Javob 2"
	javoblar = append(javoblar, javob)

	savol.Javoblar = javoblar

	savollar = append(savollar, savol)
	//-------------
	savol.Savolnum = "002"
	savol.Savoltext = "no0tugri javobni topingda"
	javoblar = []models.Javob{}

	javob.Javobnum = "3"
	javob.Javobtext = "javob3"
	javoblar = append(javoblar, javob)

	javob.Javobnum = "4"
	javob.Javobtext = "javob4"
	javoblar = append(javoblar, javob)

	savol.Javoblar = javoblar

	savollar = append(savollar, savol)

	//savollarni bazada olamiz
	tests, err := test.FindAllTests(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	for i := 0; i < len(*tests); i++ {
		savol = models.Savol{}

		savol.Savolnum = strconv.FormatUint((*tests)[i].ID, 10)
		savol.Savoltext = (*tests)[i].Title

		hashedTugrijavob, err := Hash(strconv.FormatUint((*tests)[i].Answer_code, 10))
		if err != nil {
			return
		}
		savol.Tugrijavob = string(hashedTugrijavob)

		javoblar = []models.Javob{}

		testvar := models.Testvar{}
		testvars, err := testvar.Find3(server.DB, (*tests)[i].ID)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		fmt.Println(len(*testvars))
		for j := 0; j < len(*testvars); j++ {
			javob.Javobnum = strconv.FormatUint((*testvars)[j].ID, 10)
			javob.Javobtext = (*testvars)[j].VariantText
			javoblar = append(javoblar, javob)

			fmt.Println("javob.Javobnum =" + strconv.FormatUint((*testvars)[j].ID, 10) + ", text= " + (*testvars)[j].VariantText)
		}
		savol.Javoblar = javoblar
		savollar = append(savollar, savol)
	}

	kattatest.Mavzu = "Geometriya"
	kattatest.Savollar = savollar

	responses.JSON(w, http.StatusOK, kattatest)
}

func (server *Server) CreateTest(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post := models.Savol{}

	err = json.Unmarshal(body, &post)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	fmt.Println(post)

	test := models.Test{}
	test.Title = post.Savoltext

	/*post.Prepare()
	err = post.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != post.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}*/
	postCreated, err := test.SaveTest(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	fmt.Println(postCreated)

	testvar := models.Testvar{}

	for j := 0; j < len(post.Javoblar); j++ {
		testvar.VariantText = post.Javoblar[j].Javobtext
		testvar.IdxID = postCreated.ID
		testvar.ID = 0

		ansCreated, err := testvar.SaveTestvar(server.DB)
		if err != nil {
			formattedError := formaterror.FormatError(err.Error())
			responses.ERROR(w, http.StatusInternalServerError, formattedError)
			return
		}
		fmt.Println(ansCreated)
		if post.Javoblar[j].Javobsign == "1" {
			fmt.Println(int(ansCreated.ID))

			postCreated, err = test.UpdateTestAnswer(server.DB, postCreated.ID, ansCreated.ID)
			if err != nil {
				formattedError := formaterror.FormatError(err.Error())
				responses.ERROR(w, http.StatusInternalServerError, formattedError)
				return
			}
			fmt.Println(postCreated)
		}
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.URL.Path, postCreated.Title))
	responses.JSON(w, http.StatusCreated, postCreated)
}
