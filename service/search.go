package service

import (
	"strconv"
	"subjectInformation/model"
)

type SearchService struct {
}

func orderString(name string, ord string) (sql string) {
	order := " ORDER BY "
	if name == "" || ord == "" || name == "relativity" {
		order += "relativity DESC"
	} else {
		switch name {
		case "time":
			switch ord {
			case "increase":
				{
					order += "date asc"
				}
			case "decrease":
				{
					order += "date desc"
				}
			}
		case "citation":
			switch ord {
			case "increase":
				{
					order += "citation asc"
				}
			case "decrease":
				{
					order += "citation desc"
				}
			}
		}
	}
	return order
}

func (SearchService) SearchInCommonDB(form model.SearchCommonDBRequest) (res []model.SearchCommonDBPreview, err error) {
	JoinArticles := " INNER JOIN categories ON categories.foreign_key = articles.id AND categories.tablee = 'articles' AND ( "
	JoinDissertations := " INNER JOIN categories ON categories.foreign_key = dissertations.id AND categories.tablee = 'dissertations' AND ( "
	JoinBooks := " INNER JOIN categories ON categories.foreign_key = books.id AND categories.tablee = 'books' AND ( "
	JoinProjects := " INNER JOIN categories ON categories.foreign_key =projects.id AND categories.tablee = 'projects' AND ( "
	for key, value := range form.Category {
		if value == 49 { // ascii 1 = 49
			JoinArticles += "categories.category" + strconv.Itoa(key+1) + " ='1'" + " OR "
			JoinDissertations += "categories.category" + strconv.Itoa(key+1) + " ='1'" + " OR "
			JoinBooks += "categories.category" + strconv.Itoa(key+1) + " ='1'" + " OR "
			JoinProjects += "categories.category" + strconv.Itoa(key+1) + " ='1'" + " OR "
		}
	}
	JoinArticles += "1=0) "
	JoinDissertations += "1=0) "
	JoinBooks += "1=0) "
	JoinProjects += "1=0) "

	var rawSql = "SELECT SQL_CALC_FOUND_ROWS title,id,TableName,author,time" +
		" FROM (" + " SELECT title,articles.id,articles.author,articles.create_date AS time" + ",'articles' AS TableName" +
		" ,MATCH(title) AGAINST (? IN BOOLEAN MODE) AS relevance" +
		" FROM articles" + JoinArticles +
		" WHERE MATCH(title) AGAINST (? IN BOOLEAN MODE)" +
		" UNION ALL" +
		" SELECT title,books.id,books.author,books.time" + ",'books' AS TableName" +
		" ,MATCH(title) AGAINST (? IN BOOLEAN MODE) AS relevance" +
		" FROM books" + JoinBooks +
		" WHERE MATCH(title) AGAINST (? IN BOOLEAN MODE)" +
		" UNION ALL" +
		" SELECT title,dissertations.id,dissertations.author" + ",'dissertations' AS TableName" +
		" ,MATCH(title) AGAINST (? IN BOOLEAN MODE) AS relevance" +
		" FROM dissertations" + JoinDissertations +
		" WHERE MATCH(title) AGAINST (? IN BOOLEAN MODE)" +
		" UNION ALL" +
		" SELECT title,projects.id , projects.superintendent AS author" + ",'projects' AS TableName" +
		" ,MATCH(title) AGAINST (? IN BOOLEAN MODE) AS relevance" +
		" FROM projects" + JoinProjects +
		" WHERE MATCH(title) AGAINST (? IN BOOLEAN MODE)" +
		") AS results"

	rawSql += orderString(form.Name, form.Order)

	limit := " LIMIT " + strconv.Itoa((form.Page-1)*form.Limit) + " , " + strconv.Itoa(form.Limit)
	rawSql += limit

	model.DB.Raw(rawSql,
		form.Title, form.Title,
		form.Title, form.Title,
		form.Title, form.Title,
		form.Title, form.Title).Scan(&res)
	return
}

func (SearchService) SearchCommonDBProject(form model.SearchCommonDBRequest) (res []model.SearchCommonDBPreview, err error) {
	JoinProjects := " INNER JOIN categories ON categories.foreign_key =projects.id AND categories.tablee = 'projects' AND ( "
	for key, value := range form.Category {
		if value == 49 { // ascii 1 = 49
			JoinProjects += "categories.category" + strconv.Itoa(key+1) + " ='1'" + " OR "
		}
	}
	JoinProjects += "1=0) "

	var RawSql = "SELECT title,id,TableName,author,time" +
		" FROM (" + " SELECT title,projects.id , projects.superintendent AS author" + ",'projects' AS TableName" +
		" ,MATCH(title) AGAINST (? IN BOOLEAN MODE) AS relevance" +
		" FROM projects" + JoinProjects +
		" WHERE MATCH(title) AGAINST (? IN BOOLEAN MODE)" +
		") AS results"

	RawSql += orderString(form.Name, form.Order)

	limit := " LIMIT " + strconv.Itoa((form.Page-1)*form.Limit) + " , " + strconv.Itoa(form.Limit)
	RawSql += limit
	model.DB.Raw(RawSql, form.Title, form.Title).Scan(&res)

	return
}

func (SearchService) SearchCommonDBArticle(form model.SearchCommonDBRequest) (res []model.SearchCommonDBPreview, err error) {
	JoinArticles := " INNER JOIN categories ON categories.foreign_key = articles.id AND categories.tablee = 'articles' AND ( "
	for key, value := range form.Category {
		if value == 49 { // ascii 1 = 49
			JoinArticles += "categories.category" + strconv.Itoa(key+1) + " ='1'" + " OR "
		}
	}
	JoinArticles += "1=0) "

	var RawSql = "SELECT title,id,TableName,author,time" +
		" FROM (" + " SELECT title,articles.id , articles.author" + ",'articles' AS TableName" +
		" ,MATCH(title) AGAINST (? IN BOOLEAN MODE) AS relevance" +
		" FROM articles" + JoinArticles +
		" WHERE MATCH(title) AGAINST (? IN BOOLEAN MODE)" +
		") AS results"

	RawSql += orderString(form.Name, form.Order)

	limit := " LIMIT " + strconv.Itoa((form.Page-1)*form.Limit) + " , " + strconv.Itoa(form.Limit)
	RawSql += limit
	model.DB.Raw(RawSql, form.Title, form.Title).Scan(&res)
	return
}

func (SearchService) SearchCommonDBDissertation(form model.SearchCommonDBRequest) (res []model.SearchCommonDBPreview, err error) {
	JoinDissertations := " INNER JOIN categories ON categories.foreign_key = dissertations.id AND categories.tablee = 'dissertations' AND ( "
	for key, value := range form.Category {
		if value == 49 { // ascii 1 = 49
			JoinDissertations += "categories.category" + strconv.Itoa(key+1) + " ='1'" + " OR "
		}
	}
	JoinDissertations += "1=0) "
	var RawSql = "SELECT title,id,TableName,author,time" +
		" FROM (" + " SELECT title,dissertations.id , dissertations.author" + ",'dissertations' AS TableName" +
		" ,MATCH(title) AGAINST (? IN BOOLEAN MODE) AS relevance" +
		" FROM dissertations" + JoinDissertations +
		" WHERE MATCH(title) AGAINST (? IN BOOLEAN MODE)" +
		") AS results"

	RawSql += orderString(form.Name, form.Order)

	limit := " LIMIT " + strconv.Itoa((form.Page-1)*form.Limit) + " , " + strconv.Itoa(form.Limit)
	RawSql += limit
	model.DB.Raw(RawSql, form.Title, form.Title).Scan(&res)
	return
}

func (SearchService) SearchCommonDBBook(form model.SearchCommonDBRequest) (res []model.SearchCommonDBPreview, err error) {
	JoinBooks := " INNER JOIN categories ON categories.foreign_key = books.id AND categories.tablee = 'books' AND ( "
	for key, value := range form.Category {
		if value == 49 { // ascii 1 = 49
			JoinBooks += "categories.category" + strconv.Itoa(key+1) + " ='1'" + " OR "
		}
	}
	JoinBooks += "1=0) "
	var RawSql = "SELECT title,id,TableName,author,time" +
		" FROM (" + " SELECT title,books.id , books.author" + ",'dissertations' AS TableName" +
		" ,MATCH(title) AGAINST (? IN BOOLEAN MODE) AS relevance" +
		" FROM books" + JoinBooks +
		" WHERE MATCH(title) AGAINST (? IN BOOLEAN MODE)" +
		") AS results"

	RawSql += orderString(form.Name, form.Order)

	limit := " LIMIT " + strconv.Itoa((form.Page-1)*form.Limit) + " , " + strconv.Itoa(form.Limit)
	RawSql += limit
	model.DB.Raw(RawSql, form.Title, form.Title).Scan(&res)
	return
}

func (SearchService) SearchTeamwork() {

}

func (SearchService) SearchUniqueDB() {

}
