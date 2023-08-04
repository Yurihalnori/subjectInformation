package model

type SearchCommonDBRequest struct {
	Title    string `json:"title"`
	Category string `json:"category" binding:"required"`
	Module   int    `json:"module" binding:"required"`
	Page     int    `json:"page" binding:"required"`
	Limit    int    `json:"limit" binding:"required"`
	Name     string `json:"name"`
	Order    string `json:"order"`
}

type SearchCommonDBPreview struct {
	Title     string
	Id        string
	TableName string
	Author    string
	Time      string
}

type SearchUniqueDBRequest struct {
	Title    string `json:"title" binding:"required"`
	Category string `json:"category" binding:"required"`
	Page     int    `json:"page" binding:"required"`
	Limit    int    `json:"limit" binding:"required"`
	Name     string `json:"name"`
	Order    string `json:"order"`
}
