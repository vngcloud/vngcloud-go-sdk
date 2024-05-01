package secgroup_rule

import (
	"fmt"
	lStr "strings"

	vconError "github.com/vngcloud/vngcloud-go-sdk/error"
)

const (
	errSecurityGroupRuleExists = "securitygroupruleexists"
)

// *********************************************** ErrorResult for pool ************************************************

type ErrorResolver struct {
	Message string `json:"message"`
}

func (s *ErrorResolver) ToError() error {
	msg := lStr.TrimSpace(lStr.ToLower(s.Message))

	if lStr.Contains(msg, errSecurityGroupRuleExists) {
		return NewErrSecurityGroupRuleExists()
	}

	return fmt.Errorf(s.Message)
}

// ************************************************* ErrPoolNotFound ***************************************************

func NewErrSecurityGroupRuleExists() vconError.IErrorBuilder {
	return new(ErrSecurityGroupRuleExists)
}

func IsErrSecurityGroupRuleExists(pErr error) bool {
	_, ok := pErr.(*ErrSecurityGroupRuleExists)
	return ok
}

type ErrSecurityGroupRuleExists struct {
	vconError.BaseError
}

func (s *ErrSecurityGroupRuleExists) Error() string {
	s.DefaultError = "Security group rule already exists"
	return s.ChoseErrString()
}
