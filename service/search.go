package service

import (
	"strconv"
	"subjectInformation/model"
)

type SearchService struct {
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

	var rawSql = "SELECT title,id,TableName,author,time" +
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
		") AS results" +
		" ORDER BY relevance DESC"

	model.DB.Raw(rawSql,
		form.Title, form.Title,
		form.Title, form.Title,
		form.Title, form.Title,
		form.Title, form.Title).Scan(&res)
	return
}

func (SearchService) SearchCommonDBProject(form model.SearchCommonDBRequest) (res []model.SearchCommonDBPreview) {
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
		") AS results" +
		" ORDER BY relevance DESC"

	model.DB.Raw(RawSql).Scan(&res)
	return
}

func (SearchService) SearchCommonDBArticle() {

}

func (SearchService) SearchCommonDBDissertation() {

}

func (SearchService) SearchCommonDBBook() {

}

func (SearchService) SearchTeamwork() {

}

func (SearchService) SearchUniqueDB() {

}
