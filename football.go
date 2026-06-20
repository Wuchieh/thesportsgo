package thesportsgo

import (
	"context"
)

type FootballCategoryResponse = Response[[]FootballCategoryResponseData]

// FootballCategory 取得足球分類
func (c *Client) FootballCategory(ctx context.Context) (*FootballCategoryResponse, error) {
	return secretGet[FootballCategoryResponse](ctx, c, footballCategoryPath)
}

type FootballCountryResponse = Response[[]FootballCountryResponseData]

// FootballCountry 取得足球國家
func (c *Client) FootballCountry(ctx context.Context) (*FootballCountryResponse, error) {
	return secretGet[FootballCountryResponse](ctx, c, footballCountryPath)
}

type FootballCompetitionQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// FootballCompetition 取得足球比賽
func (c *Client) FootballCompetition(ctx context.Context, query FootballCompetitionQuery) (*FootballCountryResponse, error) {
	return secretGet[FootballCountryResponse](ctx, c, footballCountryPath, toQuery(query))
}
