package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pMade.api/src/app/usecase"
)

type UserHandler interface {
	Index(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

func (uh userHandler) Index(w http.ResponseWriter, r *http.Request) {

	//userデータをjson形式に変換
	type userField struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	//userのリストを返す
	type response struct {
		Users []userField `json:"users"`
	}

	ctx := r.Context()

	//ユースケース呼び出し
	users, err := uh.userUseCase.GetUser(ctx)
	if err != nil {
		fmt.Println(err)
	}

	res := new(response)
	for _, user := range users {
		var uf userField
		uf = userField{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		}
		res.Users = append(res.Users, uf)
	}
	json.NewEncoder(w).Encode(res)

}
