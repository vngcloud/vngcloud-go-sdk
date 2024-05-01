package pagination

type IPage interface {
	NextPageURL(pPageOpts IPageOpts) (string, error)
	GetBody() interface{}
	IsEmpty() (bool, error)
}

type IPageOpts interface {
	ToListQuery() (string, error)
	ToListQueryWithParams(*map[string]interface{}) (string, error)
}
