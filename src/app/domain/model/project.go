package model

type Project struct {
	Id           uint32 `db:"project_id"`
	Name         string `db:"name"`
	Description  string `db:"description"`
	CategoryId   uint32 `db:"category_id"`
	Thumbnail    string `db:"thumbnail"`
	MemberNumber uint32 `db:"member_number"`
	AuthorId     uint32 `db:"author_id"`
}
