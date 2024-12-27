package service

import (
	"context"
	"log"

	"github.com/machinebox/graphql"
)

type Country struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Currency string `json:"currency"`
}

func FetchCountries() ([]Country, error) {
	graphqlClient := graphql.NewClient("https://countries.trevorblades.com/")
	graphqlRequest := graphql.NewRequest(`
        {
            countries {
                code
                name
                currency
            }
        }
    `)

	var response struct {
		Countries []Country `json:"countries"`
	}

	if err := graphqlClient.Run(context.Background(), graphqlRequest, &response); err != nil {
		log.Printf("Error executing GraphQL Request: %v", err)
		return nil, err
	}

	return response.Countries, nil
}
