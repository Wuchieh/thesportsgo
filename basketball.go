package thesportsgo

import (
	"context"
)

//	===  基本清單端點（無查詢參數）  ===

// BasketballCategoryResponse 籃球分類回應型別
type BasketballCategoryResponse = Response[[]BasketballCategoryResponseData]

// BasketballCategory 取得籃球分類
func (c *Client) BasketballCategory(ctx context.Context) (*BasketballCategoryResponse, error) {
	return secretGet[BasketballCategoryResponse](ctx, c, basketballCategoryPath)
}

// BasketballCountryResponse 籃球國家/地區回應型別
type BasketballCountryResponse = Response[[]BasketballCountryResponseData]

// BasketballCountry 取得籃球國家/地區
func (c *Client) BasketballCountry(ctx context.Context) (*BasketballCountryResponse, error) {
	return secretGet[BasketballCountryResponse](ctx, c, basketballCountryPath)
}

//	===  基本清單端點（有查詢參數）  ===

// BasketballCompetitionQuery 籃球賽事查詢參數
type BasketballCompetitionQuery struct {
	Page      *int    `json:"page,omitempty"`
	Time      *int    `json:"time,omitempty"`
	UUID      *string `json:"uuid,omitempty"`
	CountryID *string `json:"country_id,omitempty"`
}

// BasketballCompetitionResponse 籃球賽事回應型別
type BasketballCompetitionResponse = Response[[]BasketballCompetitionResponseData]

// BasketballCompetition 取得籃球賽事
func (c *Client) BasketballCompetition(ctx context.Context, query BasketballCompetitionQuery) (*BasketballCompetitionResponse, error) {
	return secretGet[BasketballCompetitionResponse](ctx, c, basketballCompetitionPath, toQuery(query))
}

// BasketballTeamQuery 籃球球隊查詢參數
type BasketballTeamQuery struct {
	Page          *int    `json:"page,omitempty"`
	Time          *int    `json:"time,omitempty"`
	UUID          *string `json:"uuid,omitempty"`
	CompetitionID *string `json:"competition_id,omitempty"`
}

// BasketballTeamResponse 籃球球隊回應型別
type BasketballTeamResponse = Response[[]BasketballTeamResponseData]

// BasketballTeam 取得籃球球隊
func (c *Client) BasketballTeam(ctx context.Context, query BasketballTeamQuery) (*BasketballTeamResponse, error) {
	return secretGet[BasketballTeamResponse](ctx, c, basketballTeamPath, toQuery(query))
}

// BasketballPlayerQuery 籃球球員查詢參數
type BasketballPlayerQuery struct {
	Page   *int    `json:"page,omitempty"`
	Time   *int    `json:"time,omitempty"`
	UUID   *string `json:"uuid,omitempty"`
	TeamID *string `json:"team_id,omitempty"`
}

// BasketballPlayerResponse 籃球球員回應型別
type BasketballPlayerResponse = Response[[]BasketballPlayerResponseData]

// BasketballPlayer 取得籃球球員
func (c *Client) BasketballPlayer(ctx context.Context, query BasketballPlayerQuery) (*BasketballPlayerResponse, error) {
	return secretGet[BasketballPlayerResponse](ctx, c, basketballPlayerPath, toQuery(query))
}

// BasketballVenueQuery 籃球場館查詢參數
type BasketballVenueQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// BasketballVenueResponse 籃球場館回應型別
type BasketballVenueResponse = Response[[]BasketballVenueResponseData]

// BasketballVenue 取得籃球場館
func (c *Client) BasketballVenue(ctx context.Context, query BasketballVenueQuery) (*BasketballVenueResponse, error) {
	return secretGet[BasketballVenueResponse](ctx, c, basketballVenuePath, toQuery(query))
}

// BasketballSeasonQuery 籃球賽季查詢參數
type BasketballSeasonQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// BasketballSeasonResponse 籃球賽季回應型別
type BasketballSeasonResponse = Response[[]BasketballSeasonResponseData]

// BasketballSeason 取得籃球賽季
func (c *Client) BasketballSeason(ctx context.Context, query BasketballSeasonQuery) (*BasketballSeasonResponse, error) {
	return secretGet[BasketballSeasonResponse](ctx, c, basketballSeasonPath, toQuery(query))
}

// BasketballStageQuery 籃球階段查詢參數
type BasketballStageQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// BasketballStageResponse 籃球階段回應型別
type BasketballStageResponse = Response[[]BasketballStageResponseData]

// BasketballStage 取得籃球階段
func (c *Client) BasketballStage(ctx context.Context, query BasketballStageQuery) (*BasketballStageResponse, error) {
	return secretGet[BasketballStageResponse](ctx, c, basketballStagePath, toQuery(query))
}

//	===  比賽相關端點  ===

// BasketballMatchQuery 籃球比賽查詢參數
type BasketballMatchQuery struct {
	CompetitionID *string `json:"competition_id,omitempty"`
	Date          *int    `json:"date,omitempty"`
	SeasonID      *string `json:"season_id,omitempty"`
}

// BasketballMatchResponse 籃球比賽回應型別
type BasketballMatchResponse = Response[[]BasketballMatchResponseData]

// BasketballMatch 取得籃球比賽
func (c *Client) BasketballMatch(ctx context.Context, query BasketballMatchQuery) (*BasketballMatchResponse, error) {
	return secretGet[BasketballMatchResponse](ctx, c, basketballMatchPath, toQuery(query))
}

// BasketballMatchDetailQuery 籃球比賽詳情查詢參數
type BasketballMatchDetailQuery struct {
	MatchID *string `json:"match_id,omitempty"`
}

// BasketballMatchDetailResponse 籃球比賽詳情回應型別
type BasketballMatchDetailResponse = Response[BasketballMatchDetailResponseData]

// BasketballMatchDetail 取得籃球比賽詳情
func (c *Client) BasketballMatchDetail(ctx context.Context, query BasketballMatchDetailQuery) (*BasketballMatchDetailResponse, error) {
	return secretGet[BasketballMatchDetailResponse](ctx, c, basketballMatchDetailPath, toQuery(query))
}

// BasketballMatchLineupQuery 籃球比賽陣容查詢參數
type BasketballMatchLineupQuery struct {
	MatchID *string `json:"match_id,omitempty"`
}

// BasketballMatchLineupResponse 籃球比賽陣容回應型別
type BasketballMatchLineupResponse = Response[BasketballMatchLineupResponseData]

// BasketballMatchLineup 取得籃球比賽陣容
func (c *Client) BasketballMatchLineup(ctx context.Context, query BasketballMatchLineupQuery) (*BasketballMatchLineupResponse, error) {
	return secretGet[BasketballMatchLineupResponse](ctx, c, basketballMatchLineupPath, toQuery(query))
}

// BasketballMatchStatisticsQuery 籃球比賽統計查詢參數
type BasketballMatchStatisticsQuery struct {
	MatchID *string `json:"match_id,omitempty"`
}

// BasketballMatchStatisticsResponse 籃球比賽統計回應型別
type BasketballMatchStatisticsResponse = Response[BasketballMatchStatisticsResponseData]

// BasketballMatchStatistics 取得籃球比賽統計
func (c *Client) BasketballMatchStatistics(ctx context.Context, query BasketballMatchStatisticsQuery) (*BasketballMatchStatisticsResponse, error) {
	return secretGet[BasketballMatchStatisticsResponse](ctx, c, basketballMatchStatisticsPath, toQuery(query))
}

//	===  排名相關端點  ===

// BasketballStandingsQuery 籃球排名查詢參數
type BasketballStandingsQuery struct {
	CompetitionID *string `json:"competition_id,omitempty"`
}

// BasketballStandingsResponse 籃球排名回應型別
type BasketballStandingsResponse = Response[[]BasketballStandingsResponseData]

// BasketballStandings 取得籃球排名
func (c *Client) BasketballStandings(ctx context.Context, query BasketballStandingsQuery) (*BasketballStandingsResponse, error) {
	return secretGet[BasketballStandingsResponse](ctx, c, basketballStandingsPath, toQuery(query))
}
