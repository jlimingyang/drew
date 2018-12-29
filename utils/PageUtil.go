package utils

type Page struct {
	Page       int
	PageSize   int
	TotalPage  int
	TotalCount int
	FirstPage  bool
	LastPage   bool
	List       interface{}
}

func PageUtil(total int, page int, pageSize int, list interface{}) Page {
	tp := total / pageSize
	if total%pageSize > 0 {
		tp = total/pageSize + 1
	}
	return Page{Page: page, PageSize: pageSize, TotalPage: tp, TotalCount: total, FirstPage: page == 1, LastPage: page == tp, List: list}
}
