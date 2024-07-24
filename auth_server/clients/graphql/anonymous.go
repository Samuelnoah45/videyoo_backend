package graphqlClient

import (
 "net/http"
 "os"
 "github.com/hasura/go-graphql-client"
)


func AnonymousClient() *graphql.Client { 
 // Set up the HTTP client with the request headers
 headers := http.Header{}
 // An HTTP transport that adds headers to requests
 httpClient := &http.Client{Transport: &headersTransport{headers, http.DefaultTransport}}
 // Set up the GraphQL client
 newClient :=  graphql.NewClient( os.Getenv("HASURA_GRAPHQL_URL"), httpClient)
 return newClient
}

