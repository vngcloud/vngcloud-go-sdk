package snapshot

type IListOptsBuilder interface {
	ToListQuery() (string, error)
	ToListQueryWithParams(*map[string]interface{}) (string, error)
	GetVolumeID() string
	GetStatus() string
	GetName() string
}
