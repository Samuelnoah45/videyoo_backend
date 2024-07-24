package graphqlClient

import (
	"net/http"
	"server/config"

	"github.com/hasura/go-graphql-client"
)

func AuthClient(xHasuraRole string, tokenString string) *graphql.Client {
	// Set up the HTTP client with the request headers
	headers := http.Header{}
	headers.Add("X-Hasura-Role", xHasuraRole)
	headers.Add("Authorization", tokenString)
	//  headers.Add("Authorization", "")
	// An HTTP transport that adds headers to requests
	httpClient := &http.Client{Transport: &headersTransport{headers, http.DefaultTransport}}
	// Set up the GraphQL client
	newClient := graphql.NewClient(config.HASURA_GRAPHQL_URL, httpClient)
	return newClient
}
