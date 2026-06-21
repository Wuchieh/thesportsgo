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

// FootballCountry 取得足球國家/地區
func (c *Client) FootballCountry(ctx context.Context) (*FootballCountryResponse, error) {
	return secretGet[FootballCountryResponse](ctx, c, footballCountryPath)
}

type FootballCompetitionQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

type FootballCompetitionResponse = Response[[]FootballCompetitionResponseData]

// FootballCompetition 取得足球賽事
func (c *Client) FootballCompetition(ctx context.Context, query FootballCompetitionQuery) (*FootballCompetitionResponse, error) {
	return secretGet[FootballCompetitionResponse](ctx, c, footballCompetitionPath, toQuery(query))
}

type FootballTeamQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

type FootballTeamResponse = Response[[]FootballTeamResponseData]

// FootballTeam 取得足球球隊
func (c *Client) FootballTeam(ctx context.Context, query FootballTeamQuery) (*FootballTeamResponse, error) {
	return secretGet[FootballTeamResponse](ctx, c, footballTeamPath, toQuery(query))
}

type FootballPlayerQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

type FootballPlayerResponse = Response[[]FootballPlayerResponseData]

// FootballPlayer 取得足球球員
func (c *Client) FootballPlayer(ctx context.Context, query FootballPlayerQuery) (*FootballPlayerResponse, error) {
	return secretGet[FootballPlayerResponse](ctx, c, footballPlayerPath, toQuery(query))
}

type FootballCoachQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

type FootballCoachResponse = Response[[]FootballCoachResponseData]

// FootballCoach 取得足球教練
func (c *Client) FootballCoach(ctx context.Context, query FootballCoachQuery) (*FootballCoachResponse, error) {
	return secretGet[FootballCoachResponse](ctx, c, footballCoachPath, toQuery(query))
}

type FootballRefereeQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

type FootballRefereeResponse = Response[[]FootballRefereeResponseData]

// FootballReferee 取得足球裁判
func (c *Client) FootballReferee(ctx context.Context, query FootballRefereeQuery) (*FootballRefereeResponse, error) {
	return secretGet[FootballRefereeResponse](ctx, c, footballRefereePath, toQuery(query))
}

type FootballVenueQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

type FootballVenueResponse = Response[[]FootballVenueResponseData]

// FootballVenue 取得足球場館
func (c *Client) FootballVenue(ctx context.Context, query FootballVenueQuery) (*FootballVenueResponse, error) {
	return secretGet[FootballVenueResponse](ctx, c, footballVenuePath, toQuery(query))
}

type FootballSeasonQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

type FootballSeasonResponse = Response[[]FootballSeasonResponseData]

// FootballSeason 取得足球賽季
func (c *Client) FootballSeason(ctx context.Context, query FootballSeasonQuery) (*FootballSeasonResponse, error) {
	return secretGet[FootballSeasonResponse](ctx, c, footballSeasonPath, toQuery(query))
}

type FootballStageQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

type FootballStageResponse = Response[[]FootballStageResponseData]

// FootballStage 取得足球階段
func (c *Client) FootballStage(ctx context.Context, query FootballStageQuery) (*FootballStageResponse, error) {
	return secretGet[FootballStageResponse](ctx, c, footballStagePath, toQuery(query))
}
