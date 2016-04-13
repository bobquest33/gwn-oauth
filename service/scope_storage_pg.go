package service

import (
	"errors"
	"fmt"

	"github.com/helderfarias/gwn-oauth/model"
	oauthmodel "github.com/helderfarias/oauthprovider-go/model"
)

type PGScopeStorage struct {
	roles []model.Role
}

func (c *PGScopeStorage) Find(scope, clientId string) (*oauthmodel.Scope, error) {
	for _, role := range c.roles {
		if role.Name == scope {
			return &oauthmodel.Scope{Name: role.Name}, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Role %s not found", scope))
}
