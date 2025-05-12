package oauth

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"
)

func ValidateGoogleToken(googleToken *oauth2.Token, userCtx *context.Context, googleAppClientID string) (*idtoken.Payload, error) {
	googleIDTokenString, ok := googleToken.Extra("id_token").(string)
	if !ok {
		return nil, fmt.Errorf("google ID token is missing from response")
	}

	validatedGoogleIDToken, err := idtoken.Validate(*userCtx, googleIDTokenString, googleAppClientID)
	if err != nil {
		return nil, err
	}

	return validatedGoogleIDToken, nil
}
