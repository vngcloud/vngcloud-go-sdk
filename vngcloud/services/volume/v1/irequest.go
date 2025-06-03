package v1

type IGetVolumeTypeByIdRequest interface {
	GetVolumeTypeId() string
}

type IGetListVolumeTypeRequest interface {
	GetVolumeTypeZoneId() string
}

type IGetVolumeTypeZonesRequest interface {
	ToQuery() (string, error)
	GetDefaultQuery() string
}
