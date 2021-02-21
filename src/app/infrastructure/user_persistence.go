package infrastructure

import (
	"context"
	"fmt"
	"pMade.api/src/app/domain/model"
	"pMade.api/src/app/domain/repository"
)

type UserPersistence struct {
	Mysql
}

func (up *UserPersistence) GetUser(ctx context.Context) (users []*model.User, err error) {
	query := `SELECT user_id, name, email FROM users`
	//マルチ
	rows, err := up.Mysql.DB.Query(query)
	fmt.Println(rows.Next())
	defer rows.Close()
	if err != nil {
		fmt.Println("persistenceのGetUserでのエラー")
	}

	//nextというループがあるのでuserストラクトを指定して、一つ一つストラクトに入れてappendでスライスに入れる。
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		//一個一個エラーが出ているが、下のerr = rows.Err()でまとめることもできる。
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, &user)
	}
	err = rows.Err()
	return users, nil

}

func NewUserPersistence(mysql *Mysql) repository.UserRepository {
	userPer := UserPersistence{
		*mysql,
	}
	return &userPer
}
