package thesportsgo

import (
	"context"
)

// BaseballCategoryResponse 棒球分類回應
type BaseballCategoryResponse = Response[[]BaseballCategoryResponseData]

// BaseballCategory 取得棒球分類
func (c *Client) BaseballCategory(ctx context.Context) (*BaseballCategoryResponse, error) {
	return secretGet[BaseballCategoryResponse](ctx, c, baseballCategoryPath)
}

// BaseballCountryResponse 棒球國家/地區回應
type BaseballCountryResponse = Response[[]BaseballCountryResponseData]

// BaseballCountry 取得棒球國家/地區
func (c *Client) BaseballCountry(ctx context.Context) (*BaseballCountryResponse, error) {
	return secretGet[BaseballCountryResponse](ctx, c, baseballCountryPath)
}

// BaseballCompetitionQuery 棒球賽事查詢參數
type BaseballCompetitionQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// BaseballCompetitionResponse 棒球賽事回應
type BaseballCompetitionResponse = Response[[]BaseballCompetitionResponseData]

// BaseballCompetition 取得棒球賽事
func (c *Client) BaseballCompetition(ctx context.Context, query BaseballCompetitionQuery) (*BaseballCompetitionResponse, error) {
	return secretGet[BaseballCompetitionResponse](ctx, c, baseballCompetitionPath, toQuery(query))
}

// BaseballTeamQuery 棒球球隊查詢參數
type BaseballTeamQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// BaseballTeamResponse 棒球球隊回應
type BaseballTeamResponse = Response[[]BaseballTeamResponseData]

// BaseballTeam 取得棒球球隊
func (c *Client) BaseballTeam(ctx context.Context, query BaseballTeamQuery) (*BaseballTeamResponse, error) {
	return secretGet[BaseballTeamResponse](ctx, c, baseballTeamPath, toQuery(query))
}

// BaseballPlayerQuery 棒球球員查詢參數
type BaseballPlayerQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// BaseballPlayerResponse 棒球球員回應
type BaseballPlayerResponse = Response[[]BaseballPlayerResponseData]

// BaseballPlayer 取得棒球球員
func (c *Client) BaseballPlayer(ctx context.Context, query BaseballPlayerQuery) (*BaseballPlayerResponse, error) {
	return secretGet[BaseballPlayerResponse](ctx, c, baseballPlayerPath, toQuery(query))
}

// BaseballCoachQuery 棒球教練查詢參數
type BaseballCoachQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// BaseballCoachResponse 棒球教練回應
type BaseballCoachResponse = Response[[]BaseballCoachResponseData]

// BaseballCoach 取得棒球教練
func (c *Client) BaseballCoach(ctx context.Context, query BaseballCoachQuery) (*BaseballCoachResponse, error) {
	return secretGet[BaseballCoachResponse](ctx, c, baseballCoachPath, toQuery(query))
}

// BaseballRefereeQuery 棒球裁判查詢參數
type BaseballRefereeQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// BaseballRefereeResponse 棒球裁判回應
type BaseballRefereeResponse = Response[[]BaseballRefereeResponseData]

// BaseballReferee 取得棒球裁判
func (c *Client) BaseballReferee(ctx context.Context, query BaseballRefereeQuery) (*BaseballRefereeResponse, error) {
	return secretGet[BaseballRefereeResponse](ctx, c, baseballRefereePath, toQuery(query))
}

// BaseballVenueQuery 棒球場館查詢參數
type BaseballVenueQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// BaseballVenueResponse 棒球場館回應
type BaseballVenueResponse = Response[[]BaseballVenueResponseData]

// BaseballVenue 取得棒球場館
func (c *Client) BaseballVenue(ctx context.Context, query BaseballVenueQuery) (*BaseballVenueResponse, error) {
	return secretGet[BaseballVenueResponse](ctx, c, baseballVenuePath, toQuery(query))
}

// BaseballSeasonQuery 棒球賽季查詢參數
type BaseballSeasonQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// BaseballSeasonResponse 棒球賽季回應
type BaseballSeasonResponse = Response[[]BaseballSeasonResponseData]

// BaseballSeason 取得棒球賽季
func (c *Client) BaseballSeason(ctx context.Context, query BaseballSeasonQuery) (*BaseballSeasonResponse, error) {
	return secretGet[BaseballSeasonResponse](ctx, c, baseballSeasonPath, toQuery(query))
}

// BaseballStageQuery 棒球階段查詢參數
type BaseballStageQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// BaseballStageResponse 棒球階段回應
type BaseballStageResponse = Response[[]BaseballStageResponseData]

// BaseballStage 取得棒球階段
func (c *Client) BaseballStage(ctx context.Context, query BaseballStageQuery) (*BaseballStageResponse, error) {
	return secretGet[BaseballStageResponse](ctx, c, baseballStagePath, toQuery(query))
}

// BaseballMatchListQuery 棒球比賽列表查詢參數
type BaseballMatchListQuery struct {
	Page     *int    `json:"page,omitempty"`
	Time     *int    `json:"time,omitempty"`
	UUID     *string `json:"uuid,omitempty"`
	SeasonID *string `json:"season_id,omitempty"`
	StageID  *string `json:"stage_id,omitempty"`
}

// BaseballMatchListResponse 棒球比賽列表回應
type BaseballMatchListResponse = Response[[]BaseballMatchResponseData]

// BaseballMatchList 取得棒球比賽列表
func (c *Client) BaseballMatchList(ctx context.Context, query BaseballMatchListQuery) (*BaseballMatchListResponse, error) {
	return secretGet[BaseballMatchListResponse](ctx, c, baseballMatchListPath, toQuery(query))
}

// BaseballMatchDetailQuery 棒球比賽詳情查詢參數
type BaseballMatchDetailQuery struct {
	UUID *string `json:"uuid,omitempty"`
}

// BaseballMatchDetailResponse 棒球比賽詳情回應
type BaseballMatchDetailResponse = Response[[]BaseballMatchDetailResponseData]

// BaseballMatchDetail 取得棒球比賽詳情
func (c *Client) BaseballMatchDetail(ctx context.Context, query BaseballMatchDetailQuery) (*BaseballMatchDetailResponse, error) {
	return secretGet[BaseballMatchDetailResponse](ctx, c, baseballMatchDetailPath, toQuery(query))
}

// BaseballStandingsQuery 棒球排名查詢參數
type BaseballStandingsQuery struct {
	UUID          *string `json:"uuid,omitempty"`
	CompetitionID *string `json:"competition_id,omitempty"`
	SeasonID      *string `json:"season_id,omitempty"`
	StageID       *string `json:"stage_id,omitempty"`
}

// BaseballStandingsResponse 棒球排名回應
type BaseballStandingsResponse = Response[[]BaseballStandingsResponseData]

// BaseballStandings 取得棒球排名
func (c *Client) BaseballStandings(ctx context.Context, query BaseballStandingsQuery) (*BaseballStandingsResponse, error) {
	return secretGet[BaseballStandingsResponse](ctx, c, baseballStandingsPath, toQuery(query))
}
