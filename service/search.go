package service

import (
	"strconv"
	"subjectInformation/model"
)

type SearchService struct {
}

func orderString(name string, ord string) (sql string) {
	order := " ORDER BY "
	if name == "" || ord == "" || name == "relevance" {
		order += "relevance DESC"
	} else {
		switch name {
		case "time":
			switch ord {
			case "increase":
				{
					order += "time asc"
				}
			case "decrease":
				{
					order += "time desc"
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

	var rawSql = "SELECT " +
		"r.TableName, " + "r.title, " + "r.id, " + "r.author, " + "r.time " +
		"FROM ( " +
		"SELECT " +
		"'articles' AS TableName, " + "articles.title, " +
		"articles.id, " + "articles.author, " + "articles.create_date AS time " +
		" ,MATCH(articles.title) AGAINST (? IN BOOLEAN MODE) AS relevance " +
		"FROM articles " + JoinArticles +
		"WHERE MATCH(articles.title) AGAINST (? IN BOOLEAN MODE) " +
		"UNION ALL " +
		"SELECT " +
		"'books' AS TableName, " + "books.title, " + "books.id, " +
		"books.author, " + "books.time " +
		" ,MATCH(books.title) AGAINST (? IN BOOLEAN MODE) AS relevance " +
		"FROM books " + JoinBooks +
		"WHERE MATCH(books.title) AGAINST (? IN BOOLEAN MODE) " +
		"UNION ALL " + "SELECT " +
		"'dissertations' AS TableName, " + "dissertations.title, " + "dissertations.id, " +
		"dissertations.author, " + "dissertations.date AS time " +
		" ,MATCH(dissertations.title) AGAINST (? IN BOOLEAN MODE) AS relevance " +
		"FROM dissertations " + JoinDissertations +
		"WHERE MATCH(dissertations.title) AGAINST (? IN BOOLEAN MODE) " +
		"UNION ALL " +
		"SELECT " +
		"'projects' AS TableName, " + "projects.title, " + "projects.id, " +
		"superintendent AS author, " + "projects.create_date AS time " +
		" ,MATCH(projects.title) AGAINST (? IN BOOLEAN MODE) AS relevance " +
		"FROM projects " + JoinProjects +
		"WHERE MATCH(projects.title) AGAINST (? IN BOOLEAN MODE) " +
		") AS r"

	rawSql += orderString(form.Name, form.Order)

	limit := " LIMIT " + strconv.Itoa((form.Page-1)*form.Limit) + " , " + strconv.Itoa(form.Limit)
	rawSql += limit

	err = model.DB.Raw(rawSql, form.Title, form.Title, form.Title, form.Title,
		form.Title, form.Title, form.Title, form.Title).Scan(&res).Error
	return
}

func (SearchService) CountModuleInCommonDB(form model.SearchCommonDBRequest) (book int, dissertation int, article int, project int, err error) {
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

	var rawSql = "SELECT " +
		"'articles' AS TableName, COUNT(*) AS COUNT " +
		"FROM articles " + JoinArticles +
		"WHERE MATCH(articles.title) AGAINST (? IN BOOLEAN MODE) " +
		"UNION ALL " + "SELECT " +
		"'books' AS TableName, COUNT(*) AS COUNT " +
		"FROM books " + JoinBooks +
		"WHERE MATCH(books.title) AGAINST (? IN BOOLEAN MODE) " +
		"UNION ALL " + "SELECT " +
		"'dissertations' AS TableName, COUNT(*) AS COUNT " +
		"FROM dissertations " + JoinDissertations +
		"WHERE MATCH(dissertations.title) AGAINST (? IN BOOLEAN MODE) " +
		"UNION ALL " + "SELECT " +
		"'projects' AS TableName,  COUNT(*) AS COUNT " +
		"FROM projects " + JoinProjects +
		"WHERE MATCH(projects.title) AGAINST (? IN BOOLEAN MODE);"
	type CommonDBModuleCounts struct {
		TableName string
		COUNT     int
	}
	var res []CommonDBModuleCounts
	err = model.DB.Raw(rawSql, form.Title, form.Title, form.Title, form.Title).Scan(&res).Error
	for _, v := range res {
		switch v.TableName {
		case "articles":
			article = v.COUNT
		case "books":
			book = v.COUNT
		case "dissertations":
			dissertation = v.COUNT
		case "projects":
			project = v.COUNT
		}
	}
	return
}

func (SearchService) SearchCommonDBProject(form model.SearchCommonDBRequest) (res []model.SearchCommonDBPreview, total int, err error) {
	JoinProjects := " INNER JOIN categories ON categories.foreign_key =projects.id AND categories.tablee = 'projects' AND ( "
	for key, value := range form.Category {
		if value == 49 { // ascii 1 = 49
			JoinProjects += "categories.category" + strconv.Itoa(key+1) + " ='1'" + " OR "
		}
	}
	JoinProjects += "1=0) "

	var RawSql = "SELECT title,id,TableName,author,time" +
		" FROM (" + " SELECT title,projects.id , projects.superintendent AS author" +
		",'projects' AS TableName" + ",projects.create_date AS time " +
		" ,MATCH(title) AGAINST (? IN BOOLEAN MODE) AS relevance" +
		" FROM projects" + JoinProjects +
		" WHERE MATCH(title) AGAINST (? IN BOOLEAN MODE)" +
		") AS results"

	var countSql = "SELECT COUNT(*) as theCount" +
		" FROM projects" + JoinProjects +
		" WHERE MATCH(title) AGAINST (? IN BOOLEAN MODE)"

	var theCount = 0
	err = model.DB.Raw(countSql, form.Title).Scan(&theCount).Error
	total = theCount
	if err != nil {
		return nil, 0, err
	}
	RawSql += orderString(form.Name, form.Order)
	limit := " LIMIT " + strconv.Itoa((form.Page-1)*form.Limit) + " , " + strconv.Itoa(form.Limit)
	RawSql += limit
	err = model.DB.Raw(RawSql, form.Title, form.Title).Scan(&res).Error
	return
}

func (SearchService) SearchCommonDBArticle(form model.SearchCommonDBRequest) (res []model.SearchCommonDBPreview, total int, err error) {
	JoinArticles := " INNER JOIN categories ON categories.foreign_key = articles.id AND categories.tablee = 'articles' AND ( "
	for key, value := range form.Category {
		if value == 49 { // ascii 1 = 49
			JoinArticles += "categories.category" + strconv.Itoa(key+1) + " ='1'" + " OR "
		}
	}
	JoinArticles += "1=0) "

	var RawSql = "SELECT title,id,TableName,author,time" +
		" FROM (" + " SELECT title,articles.id , articles.author" +
		",'articles' AS TableName" + ",articles.create_date AS time " +
		" ,MATCH(title) AGAINST (? IN BOOLEAN MODE) AS relevance" +
		" FROM articles" + JoinArticles +
		" WHERE MATCH(title) AGAINST (? IN BOOLEAN MODE)" +
		") AS results"

	var countSql = "SELECT COUNT(*) as theCount" +
		" FROM articles" + JoinArticles +
		" WHERE MATCH(title) AGAINST (? IN BOOLEAN MODE)"
	var theCount = 0
	err = model.DB.Raw(countSql, form.Title).Scan(&theCount).Error
	total = theCount
	if err != nil {
		return nil, 0, err
	}
	RawSql += orderString(form.Name, form.Order)

	limit := " LIMIT " + strconv.Itoa((form.Page-1)*form.Limit) + " , " + strconv.Itoa(form.Limit)
	RawSql += limit
	err = model.DB.Raw(RawSql, form.Title, form.Title).Scan(&res).Error
	return
}

func (SearchService) SearchCommonDBDissertation(form model.SearchCommonDBRequest) (res []model.SearchCommonDBPreview, total int, err error) {
	JoinDissertations := " INNER JOIN categories ON categories.foreign_key = dissertations.id AND categories.tablee = 'dissertations' AND ( "
	for key, value := range form.Category {
		if value == 49 { // ascii 1 = 49
			JoinDissertations += "categories.category" + strconv.Itoa(key+1) + " ='1'" + " OR "
		}
	}
	JoinDissertations += "1=0) "
	var RawSql = "SELECT title,id,TableName,author,time" +
		" FROM (" + " SELECT title,dissertations.id , dissertations.author" +
		",'dissertations' AS TableName" + ",dissertations.date AS time " +
		" ,MATCH(title) AGAINST (? IN BOOLEAN MODE) AS relevance" +
		" FROM dissertations" + JoinDissertations +
		" WHERE MATCH(title) AGAINST (? IN BOOLEAN MODE)" +
		") AS results"
	var countSql = "SELECT COUNT(*) as theCount" +
		" FROM dissertations" + JoinDissertations +
		" WHERE MATCH(title) AGAINST (? IN BOOLEAN MODE)"

	var theCount = 0
	err = model.DB.Raw(countSql, form.Title).Scan(&theCount).Error
	total = theCount
	if err != nil {
		return nil, 0, err
	}
	RawSql += orderString(form.Name, form.Order)

	limit := " LIMIT " + strconv.Itoa((form.Page-1)*form.Limit) + " , " + strconv.Itoa(form.Limit)
	RawSql += limit
	err = model.DB.Raw(RawSql, form.Title, form.Title).Scan(&res).Error
	return
}

func (SearchService) SearchCommonDBBook(form model.SearchCommonDBRequest) (res []model.SearchCommonDBPreview, total int, err error) {
	JoinBooks := " INNER JOIN categories ON categories.foreign_key = books.id AND categories.tablee = 'books' AND ( "
	for key, value := range form.Category {
		if value == 49 { // ascii 1 = 49
			JoinBooks += "categories.category" + strconv.Itoa(key+1) + " ='1'" + " OR "
		}
	}
	JoinBooks += "1=0) "
	var RawSql = "SELECT title,id,TableName,author,time" +
		" FROM (" + " SELECT title,books.id , books.author" +
		",'books' AS TableName" + ",time" +
		" ,MATCH(title) AGAINST (? IN BOOLEAN MODE) AS relevance" +
		" FROM books" + JoinBooks +
		" WHERE MATCH(title) AGAINST (? IN BOOLEAN MODE)" +
		") AS results"
	var countSql = "SELECT COUNT(*) as theCount" +
		" FROM books" + JoinBooks +
		" WHERE MATCH(title) AGAINST (? IN BOOLEAN MODE)"

	var theCount = 0
	err = model.DB.Raw(countSql, form.Title).Scan(&theCount).Error
	total = theCount
	if err != nil {
		return nil, 0, err
	}
	RawSql += orderString(form.Name, form.Order)
	limit := " LIMIT " + strconv.Itoa((form.Page-1)*form.Limit) + " , " + strconv.Itoa(form.Limit)
	RawSql += limit
	err = model.DB.Raw(RawSql, form.Title, form.Title).Scan(&res).Error
	return
}

func (SearchService) SearchTeamwork() {

}

func (SearchService) SearchUniqueDB() {

}
