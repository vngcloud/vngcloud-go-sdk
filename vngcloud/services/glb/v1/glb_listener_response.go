package v1

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type GlobalListenerResponse struct {
	CreatedAt            string  `json:"createdAt"`
	UpdatedAt            string  `json:"updatedAt"`
	DeletedAt            *string `json:"deletedAt"`
	ID                   string  `json:"id"`
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	Protocol             string  `json:"protocol"`
	Port                 int     `json:"port"`
	GlobalLoadBalancerID string  `json:"globalLoadBalancerId"`
	GlobalPoolID         string  `json:"globalPoolId"`
	TimeoutClient        int     `json:"timeoutClient"`
	TimeoutMember        int     `json:"timeoutMember"`
	TimeoutConnection    int     `json:"timeoutConnection"`
	AllowedCidrs         string  `json:"allowedCidrs"`
	Headers              *string `json:"headers"`
	Status               string  `json:"status"`
}

func (s *GlobalListenerResponse) ToEntityGlobalListener() *lsentity.GlobalListener {
	return &lsentity.GlobalListener{
		CreatedAt:            s.CreatedAt,
		UpdatedAt:            s.UpdatedAt,
		DeletedAt:            s.DeletedAt,
		ID:                   s.ID,
		Name:                 s.Name,
		Description:          s.Description,
		Protocol:             s.Protocol,
		Port:                 s.Port,
		GlobalLoadBalancerID: s.GlobalLoadBalancerID,
		GlobalPoolID:         s.GlobalPoolID,
		TimeoutClient:        s.TimeoutClient,
		TimeoutMember:        s.TimeoutMember,
		TimeoutConnection:    s.TimeoutConnection,
		AllowedCidrs:         s.AllowedCidrs,
		Headers:              s.Headers,
		Status:               s.Status,
	}
}

type ListGlobalListenersResponse []GlobalListenerResponse

func (s ListGlobalListenersResponse) ToEntityListGlobalListeners() *lsentity.ListGlobalListeners {
	listeners := &lsentity.ListGlobalListeners{}
	for _, itemListener := range s {
		listeners.Items = append(listeners.Items, itemListener.ToEntityGlobalListener())
	}
	return listeners
}

// --------------------------------------------------

type CreateGlobalListenerResponse GlobalListenerResponse

func (s *CreateGlobalListenerResponse) ToEntityGlobalListener() *lsentity.GlobalListener {
	return &lsentity.GlobalListener{
		CreatedAt:            s.CreatedAt,
		UpdatedAt:            s.UpdatedAt,
		DeletedAt:            s.DeletedAt,
		ID:                   s.ID,
		Name:                 s.Name,
		Description:          s.Description,
		Protocol:             s.Protocol,
		Port:                 s.Port,
		GlobalLoadBalancerID: s.GlobalLoadBalancerID,
		GlobalPoolID:         s.GlobalPoolID,
		TimeoutClient:        s.TimeoutClient,
		TimeoutMember:        s.TimeoutMember,
		TimeoutConnection:    s.TimeoutConnection,
		AllowedCidrs:         s.AllowedCidrs,
		Headers:              s.Headers,
		Status:               s.Status,
	}
}

// --------------------------------------------------

type UpdateGlobalListenerResponse GlobalListenerResponse

func (s *UpdateGlobalListenerResponse) ToEntityGlobalListener() *lsentity.GlobalListener {
	return &lsentity.GlobalListener{
		CreatedAt:            s.CreatedAt,
		UpdatedAt:            s.UpdatedAt,
		DeletedAt:            s.DeletedAt,
		ID:                   s.ID,
		Name:                 s.Name,
		Description:          s.Description,
		Protocol:             s.Protocol,
		Port:                 s.Port,
		GlobalLoadBalancerID: s.GlobalLoadBalancerID,
		GlobalPoolID:         s.GlobalPoolID,
		TimeoutClient:        s.TimeoutClient,
		TimeoutMember:        s.TimeoutMember,
		TimeoutConnection:    s.TimeoutConnection,
		AllowedCidrs:         s.AllowedCidrs,
		Headers:              s.Headers,
		Status:               s.Status,
	}
}

// --------------------------------------------------

type GetGlobalListenerResponse GlobalListenerResponse

func (s *GetGlobalListenerResponse) ToEntityGlobalListener() *lsentity.GlobalListener {
	return &lsentity.GlobalListener{
		CreatedAt:            s.CreatedAt,
		UpdatedAt:            s.UpdatedAt,
		DeletedAt:            s.DeletedAt,
		ID:                   s.ID,
		Name:                 s.Name,
		Description:          s.Description,
		Protocol:             s.Protocol,
		Port:                 s.Port,
		GlobalLoadBalancerID: s.GlobalLoadBalancerID,
		GlobalPoolID:         s.GlobalPoolID,
		TimeoutClient:        s.TimeoutClient,
		TimeoutMember:        s.TimeoutMember,
		TimeoutConnection:    s.TimeoutConnection,
		AllowedCidrs:         s.AllowedCidrs,
		Headers:              s.Headers,
		Status:               s.Status,
	}
}
