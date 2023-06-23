package model

// type ProjectForm struct {
// 	Total int       `form:"total" json:"total" binding:"numeric"`
// 	List  []Project `form:"list" json:"list" binding:"required"`
// }

// type InstituteForm struct {
// 	Total int         `form:"total" json:"total" binding:"numeric"`
// 	List  []Institute `form:"list" json:"list" binding:"required"`
// }

// type TutorForm struct {
// 	Total int     `form:"total" json:"total" binding:"numeric"`
// 	List  []Tutor `form:"list" json:"list" binding:"required"`
// }

// type BookForm struct {
// 	Total int    `form:"total" json:"total" binding:"numeric"`
// 	List  []Book `form:"list" json:"list" binding:"required"`
// }

// type DissertationForm struct {
// 	Total int            `form:"total" json:"total" binding:"numeric"`
// 	List  []Dissertation `form:"list" json:"list" binding:"required"`
// }

// type ArticleForm struct {
// 	Total int       `form:"total" json:"total" binding:"numeric"`
// 	List  []Article `form:"list" json:"list" binding:"required"`
// }

type NewsForm struct {
	Total    int    `json:"total"`
	NewsList []News `json:"list"`
}
