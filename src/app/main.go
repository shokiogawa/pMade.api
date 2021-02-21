package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"pMade.api/src/app/handler"
	"pMade.api/src/app/infrastructure"
	"pMade.api/src/app/usecase"
)

func main() {
	db := infrastructure.NweMysql()
	userPersistence := infrastructure.NewUserPersistence(db)
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)

	r := mux.NewRouter()

	//r.HandleFunc("/", ctrl.hello).Methods("GET")
	r.HandleFunc("/users", userHandler.Index).Methods("GET")
	//r.HandleFunc("/users/{id}", ctrl.getUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r))
}

//
//type Controller struct {
//	db *sqlx.DB
//}
//
////Contropllerのコンストラクタ
//func NewController (db *sqlx.DB) Controller{
//	ctrl := new(Controller)
//	ctrl.db = db
//	return *ctrl
//}
//
//func (ctrl Controller) hello(w http.ResponseWriter, r *http.Request){
//	fmt.Println("hello")
//}
//
//
//type UserDTO struct{
//	ID int `db:"user_id" json:"id"`
//	Name string `db:"name" json:"name"`
//	Email string `db:"email" json:"email"`
//}
//
//func (ctrl Controller)users (w http.ResponseWriter, r *http.Request){
//	var usersDto []UserDTO
//	query := `SELECT user_id, name, email FROM users`
//	err := ctrl.db.Select(&usersDto, query)
//	if err != nil{
//		fmt.Println(err)
//	}
//	json.NewEncoder(w).Encode(usersDto)
//
//
//}
//
//func (ctrl Controller) getUser (w http.ResponseWriter, r *http.Request){
//	params := mux.Vars(r)
//	var userDto []UserDTO
//	query := `SELECT user_id, name, email FROM users WHERE user_id = ?`
//	err := ctrl.db.Select(&userDto, query, params["id"])
//	if err != nil{
//		fmt.Println(err)
//	}
//	json.NewEncoder(w).Encode(userDto)
//}
