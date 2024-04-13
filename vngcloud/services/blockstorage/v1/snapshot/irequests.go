package snapshot

type IListVolumeOptsBuilder interface {
	GetProjectID() string
	ToListQuery() (string, error)
	GetDefaultQuery() string
}
