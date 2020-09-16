package service

//import (
//	"github.com/gorilla/mux"
//	"net/http"
//	"testing"
//)
//
//func ServiceMethods() *mux.Router {
//	r := mux.NewRouter()
//	r.HandleFunc("/massages", GetAllMassages).Methods("GET")
//	r.HandleFunc("/massages/{id}", GetMassage).Methods("GET")
//	r.HandleFunc("/massages", WriteMassage).Methods("POST")
//	r.HandleFunc("/massages/{id}", UpdateMassage).Methods("PUT")
//	r.HandleFunc("/massages/{id}", DeleteMassage).Methods("DELETE")
//	return r
//}
//
//func TestService(t *testing.T) {
//
//	test := []struct {
//		method       string
//		path         string
//		body         string
//		responseCode int
//		responseBody string
//	}{
//		{
//			"GET",
//			"/1",
//			"`{\"id\":\"1\"}`",
//			http.StatusBadRequest,
//			"not found\n",
//		},
//		{
//			"PUT",
//			"/1",
//			"`{\"id\":\"1\"}`",
//			http.StatusOK,
//			"`{\"id\":\"1\"}`",
//		},
//		{
//			"GET",
//			"",
//			"",
//			http.StatusOK,
//			"`{\"id\":\"1\"}`",
//		},
//		{
//			"DELETE",
//			"/1",
//			"",
//			http.StatusOK,
//			"`{\"id\":\"1\"}`",
//		},
//		{
//			"UPDATE",
//			"/1",
//			"`{\"In\":\"7777\"}`",
//			http.StatusOK,
//			"`{\"Isbn\":\"7777\"}`",
//		},
//		{
//			"POST",
//			"/1",
//			"`{\"Content\":\"Hello, this test works!\"}`",
//			http.StatusOK,
//			"`{\"Content\":\"Hello, this test works!\"}`",
//		},
//	}
//}
