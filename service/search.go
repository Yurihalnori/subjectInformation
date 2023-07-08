package service

import "subjectInformation/model"

type SearchService struct {
}

func (SearchService) SearchInCommonDB(form model.SearchCommonDBRequest) {
	var rawSql = "SELECT title" +
		" FROM (" +
		" SELECT title" + " ,MATCH(title) AGAINST ('" + form.Title + "' IN BOOLEAN MODE) AS relevance" +
		" FROM articles" +
		" WHERE MATCH(title) AGAINST ('" + form.Title + "' IN BOOLEAN MODE)" +
		" UNION ALL" +
		" SELECT title" + " ,MATCH(title) AGAINST ('" + form.Title + "' IN BOOLEAN MODE) AS relevance" +
		" FROM books" +
		" WHERE MATCH(title) AGAINST ('" + form.Title + "' IN BOOLEAN MODE)" +
		" UNION ALL" +
		" SELECT title" + " ,MATCH(title) AGAINST ('" + form.Title + "' IN BOOLEAN MODE) AS relevance" +
		" FROM dissertations" +
		" WHERE MATCH(title) AGAINST ('" + form.Title + "' IN BOOLEAN MODE)" +
		" UNION ALL" +
		" SELECT title" + " ,MATCH(title) AGAINST ('" + form.Title + "' IN BOOLEAN MODE) AS relevance" +
		" FROM projects" +
		" WHERE MATCH(title) AGAINST ('" + form.Title + "' IN BOOLEAN MODE)" +
		") AS results" +
		" ORDER BY relevance DESC"
	var res []model.SearchCommonDBPreview
	model.DB.Raw(rawSql).Scan(&res)
}
