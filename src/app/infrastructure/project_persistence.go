package infrastructure

import (
	"fmt"
	"pMade.api/src/app/domain/model"
	"pMade.api/src/app/domain/repository"
)

type ProjectPersistence struct {
	Mysql
}

func (pp *ProjectPersistence) GetProject() (projects []*model.Project, err error) {
	query := `SELECT project_id, name, description, category_id, thumbnail, member_number FROM projects`
	//マルチ
	rows, err := pp.Mysql.DB.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	//行セットに対して繰り返し処理を行う
	for rows.Next() {
		project := model.Project{}
		err := rows.Scan(&project.Id, &project.Name, &project.Description, &project.CategoryId, &project.AuthorId, &project.MemberNumber, &project.Thumbnail)
		if err != nil {
			fmt.Println(err)
		}
		projects = append(projects, &project)
	}
	err = rows.Err()
	return projects, nil

}

func NewProjectPersistence(mysql *Mysql) repository.ProjectRepository {
	pp := ProjectPersistence{
		*mysql,
	}
	return &pp
}
