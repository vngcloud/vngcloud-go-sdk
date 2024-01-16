package policy

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

// ******************************************************* MISC ********************************************************

type ResponseData struct {
	UUID             string   `json:"uuid"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	RedirectPoolID   string   `json:"redirectPoolId"`
	RedirectPoolName string   `json:"redirectPoolName"`
	Action           string   `json:"action"`
	RedirectURL      string   `json:"redirectUrl"`
	RedirectHTTPCode string   `json:"redirectHttpCode"`
	KeepQueryString  string   `json:"keepQueryString"`
	Position         int      `json:"position"`
	L7Rules          []L7Rule `json:"l7Rules"`
	DisplayStatus    string   `json:"displayStatus"`
	CreatedAt        string   `json:"createdAt"`
	UpdatedAt        string   `json:"updatedAt"`
	ProgressStatus   string   `json:"progressStatus"`
}

type L7Rule struct {
	UUID               string `json:"uuid"`
	CompareType        string `json:"compareType"`
	RuleValue          string `json:"ruleValue"`
	RuleType           string `json:"ruleType"`
	CreatedAt          string `json:"createdAt"`
	UpdatedAt          string `json:"updatedAt"`
	ProvisioningStatus string `json:"provisioningStatus"`
	OperatingStatus    string `json:"operatingStatus"`
}

// *********************************************** Response of List API ************************************************

type ListResponse struct {
	ListData  []ResponseData `json:"data"`
	Page      int            `json:"page"`
	PageSize  int            `json:"pageSize"`
	TotalPage int            `json:"totalPage"`
	TotalItem int            `json:"totalItem"`
}

func (s *ListResponse) ToListPolicyObjects() []*objects.Policy {
	if s == nil || s.ListData == nil || len(s.ListData) < 1 {
		return nil
	}

	var result []*objects.Policy
	for _, itemLb := range s.ListData {
		l7rule := make([]*objects.L7Rule, 0)
		for _, itemL7 := range itemLb.L7Rules {
			l7rule = append(l7rule, &objects.L7Rule{
				UUID:               itemL7.UUID,
				CompareType:        itemL7.CompareType,
				RuleValue:          itemL7.RuleValue,
				RuleType:           itemL7.RuleType,
				ProvisioningStatus: itemL7.ProvisioningStatus,
				OperatingStatus:    itemL7.OperatingStatus,
			})
		}
		result = append(result, &objects.Policy{
			UUID:             itemLb.UUID,
			Name:             itemLb.Name,
			Description:      itemLb.Description,
			RedirectPoolID:   itemLb.RedirectPoolID,
			RedirectPoolName: itemLb.RedirectPoolName,
			Action:           itemLb.Action,
			RedirectURL:      itemLb.RedirectURL,
			RedirectHTTPCode: itemLb.RedirectHTTPCode,
			KeepQueryString:  itemLb.KeepQueryString,
			Position:         itemLb.Position,
			L7Rules:          l7rule,
			DisplayStatus:    itemLb.DisplayStatus,
			CreatedAt:        itemLb.CreatedAt,
			UpdatedAt:        itemLb.UpdatedAt,
			ProgressStatus:   itemLb.ProgressStatus,
		})
	}

	return result
}
