package handler

import (
	. "UsersTest/internal/config"
	"UsersTest/internal/database"
	"UsersTest/internal/parser"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Start() {
	Log.Infof("Server is starting on port: %d", GetInstance().Server.Port)
	router := mux.NewRouter()
	router.HandleFunc("/users", getAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getFromId).Methods("GET")

	http.Handle("/", router)
	err := http.ListenAndServe(":"+strconv.Itoa(GetInstance().Server.Port), nil)
	if err != nil {
		Log.Fatal(err)
	}
}

func getFromId(w http.ResponseWriter, r *http.Request) {
	Log.WithField("Addr", r.RemoteAddr).Info("Get request get user from ID")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		Log.Error(err)
		return
	}
	var user = &parser.User{}
	switch str := strings.ToLower(GetInstance().Server.DataType); {
	case str == "inmemory" || str == "im":
		user = parser.GetUsers()[id]
	case str == "database" || str == "db":
		user, err = database.GetFromId(id)
	default:
		Log.Fatal("Wrong DataType")
	}
	if err != nil {
		Log.Error(err)
	}

	if user == nil {
		_, err = w.Write(GiveResponse(401, "Wrong User ID"))
		if err != nil {
			Log.Error(err)
			return
		}
		return
	}

	b, err := json.Marshal(user)
	_, err = w.Write(b)
	if err != nil {
		Log.Error(err)
		return
	}

	_, err = w.Write(GiveResponse(200, "OK"))
	if err != nil {
		Log.Error(err)
	}
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	Log.WithField("Addr", r.RemoteAddr).Info("Get request get all users")
	w.Header().Set("Content-Type", "application/json")

	b, err := json.Marshal(sortMap())
	_, err = w.Write(b)
	if err != nil {
		Log.Error(err)
		return
	}

	_, err = w.Write(GiveResponse(200, "OK"))
	if err != nil {
		Log.Error(err)
	}
}

func sortMap() []parser.ShortInfo {
	var keys []int
	var users map[int]*parser.User

	switch str := strings.ToLower(GetInstance().Server.DataType); {
	case str == "inmemory" || str == "im":
		users = parser.GetUsers()
	case str == "database" || str == "db":
		return database.GetAllUsers()
	default:
		Log.Fatal("Wrong DataType")
	}

	for key := range users {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	var info []parser.ShortInfo

	for i := range keys {
		var si parser.ShortInfo
		si.Id = users[i].Id
		si.UserName = users[i].UserName
		info = append(info, si)
	}

	return info
}

func GiveResponse(code int, message string) []byte {
	resp := &Response{code, message}
	ret, _ := json.Marshal(resp)
	return ret
}
