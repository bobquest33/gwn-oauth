package service

import (
	"errors"
	"fmt"

	"github.com/helderfarias/gwn-oauth/model"
	oauthmodel "github.com/helderfarias/oauthprovider-go/model"
)

type PGScopeStorage struct {
	user model.User
}

func (this *PGScopeStorage) Find(scope, clientId string) (*oauthmodel.Scope, error) {
	for _, role := range this.user.Roles {
		if role == scope {
			return &oauthmodel.Scope{Name: role}, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Role %s not found", scope))
}
