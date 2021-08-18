package model

import (
	"context"
)

type AuthToken struct {
	projectID        string
	projectKey       string
	chainAccountAddr string
}

func NewAuthToken(projectID, projectKey, chainAccountAddr string) AuthToken {
	return AuthToken{
		projectID, projectKey, chainAccountAddr,
	}
}

func (a *AuthToken) GetRequestMetadata(context.Context, ...string) (
	map[string]string, error,
) {
	return map[string]string{"projectIdHeader": a.projectID, "projectKeyHeader": a.projectKey, "chainAccountAddressHeader": a.chainAccountAddr}, nil
}

func (a *AuthToken) RequireTransportSecurity() bool {
	return false
}