package model

type Msg struct {
	Module         string `form:"module" json:"module" binding:"binary"`
	Category       string `form:"category" json:"category" binding:"len=9"`
	Title          string `form:"title" json:"title" binding:"omitempty"`
	Text           string `form:"text" json:"text" binding:"omitempty"`
	Region         string `form:"region" json:"region" binding:"omitempty"`
	Name           string `form:"name" json:"name" binding:"omitempty"`
	Classification string `form:"classification" json:"classification" binding:"omitempty"` //项目类别 学位点类别 导师类型
	Sponsor        string `form:"sponsor" json:"" binding:"omitempty"`
	ApprovalNumber string `form:"approvalNumber" json:"approvalNumber" binding:"omitempty"`
	CreateDate     string `form:"createDate" json:"createDate" binding:"omitempty"`
	Superintendent string `form:"superintendent" json:"superintendent" binding:"omitempty"`
	Organization   string `form:"organization" json:"organization" binding:"omitempty"`
	Nation         string `form:"nation" json:"nation" binding:"omitempty"`
	Province       string `form:"province" json:"province" binding:"omitempty"`
	City           string `form:"city" json:"city" binding:"omitempty"`
	University     string `form:"university" json:"university" binding:"omitempty"`
	College        string `form:"college" json:"college" binding:"omitempty"`
	Author         string `form:"author" json:"author" binding:"omitempty"`
	Direction      string `form:"direction" json:"direction" binding:"omitempty"`
	Contact        string `form:"contact" json:"contact" binding:"omitempty"`
	Tutor          string `form:"tutor" json:"tutor" binding:"omitempty"`
	Year           string `form:"year" json:"year" binding:"omitempty"`
	Technique      string `form:"technique" json:"technique" binding:"omitempty"`
	KeyWord        string `form:"keyWord" json:"keyWord" binding:"omitempty"`
	Periodical     string `form:"periodical" json:"periodical" binding:"omitempty"`
	Digest         string `form:"digest" json:"digest" binding:"omitempty"`
	Data           string `form:"data" json:"data" binding:"omitempty"`
	Language       string `form:"language" json:"language" binding:"omitempty"`
	Publisher      string `form:"publisher" json:"publisher" binding:"omitempty"`
	Click          int    `form:"click" json:"click" binding:"omitempty"`
	Download       int    `form:"download" json:"download" binding:"omitempty"`
}

type Form struct {
	Total int   `form:"total" json:"total" binding:"required"`
	List  []Msg `form:"list" json:"list" binding:"required"`
}
