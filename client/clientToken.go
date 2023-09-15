package main

import "context"

type clientTokenAuth struct{}

func (c clientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		// "authorization": "Bearer " + token,
		"appId":  "myApp",
		"appKey": "myAppKey",
	}, nil
}

func (c clientTokenAuth) RequireTransportSecurity() bool {
	return false
}
