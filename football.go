package thesportsgo

import (
	"context"
	"encoding/json"
)

// ====== 既有基礎端點 ======

// FootballCategoryResponse 足球分類回應
type FootballCategoryResponse = Response[[]FootballCategoryResponseData]

// FootballCategory 取得足球分類
func (c *Client) FootballCategory(ctx context.Context) (*FootballCategoryResponse, error) {
	return secretGet[FootballCategoryResponse](ctx, c, footballCategoryPath)
}

// FootballCountryResponse 足球國家/地區回應
type FootballCountryResponse = Response[[]FootballCountryResponseData]

// FootballCountry 取得足球國家/地區
func (c *Client) FootballCountry(ctx context.Context) (*FootballCountryResponse, error) {
	return secretGet[FootballCountryResponse](ctx, c, footballCountryPath)
}

// FootballCompetitionQuery 足球賽事查詢參數
type FootballCompetitionQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// FootballCompetitionResponse 足球賽事回應
type FootballCompetitionResponse = Response[[]FootballCompetitionResponseData]

// FootballCompetition 取得足球賽事
func (c *Client) FootballCompetition(ctx context.Context, query FootballCompetitionQuery) (*FootballCompetitionResponse, error) {
	return secretGet[FootballCompetitionResponse](ctx, c, footballCompetitionPath, toQuery(query))
}

// FootballTeamQuery 足球球隊查詢參數
type FootballTeamQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// FootballTeamResponse 足球球隊回應
type FootballTeamResponse = Response[[]FootballTeamResponseData]

// FootballTeam 取得足球球隊
func (c *Client) FootballTeam(ctx context.Context, query FootballTeamQuery) (*FootballTeamResponse, error) {
	return secretGet[FootballTeamResponse](ctx, c, footballTeamPath, toQuery(query))
}

// FootballPlayerQuery 足球球員查詢參數
type FootballPlayerQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// FootballPlayerResponse 足球球員回應
type FootballPlayerResponse = Response[[]FootballPlayerResponseData]

// FootballPlayer 取得足球球員
func (c *Client) FootballPlayer(ctx context.Context, query FootballPlayerQuery) (*FootballPlayerResponse, error) {
	return secretGet[FootballPlayerResponse](ctx, c, footballPlayerPath, toQuery(query))
}

// FootballCoachQuery 足球教練查詢參數
type FootballCoachQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// FootballCoachResponse 足球教練回應
type FootballCoachResponse = Response[[]FootballCoachResponseData]

// FootballCoach 取得足球教練
func (c *Client) FootballCoach(ctx context.Context, query FootballCoachQuery) (*FootballCoachResponse, error) {
	return secretGet[FootballCoachResponse](ctx, c, footballCoachPath, toQuery(query))
}

// FootballRefereeQuery 足球裁判查詢參數
type FootballRefereeQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// FootballRefereeResponse 足球裁判回應
type FootballRefereeResponse = Response[[]FootballRefereeResponseData]

// FootballReferee 取得足球裁判
func (c *Client) FootballReferee(ctx context.Context, query FootballRefereeQuery) (*FootballRefereeResponse, error) {
	return secretGet[FootballRefereeResponse](ctx, c, footballRefereePath, toQuery(query))
}

// FootballVenueQuery 足球場館查詢參數
type FootballVenueQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// FootballVenueResponse 足球場館回應
type FootballVenueResponse = Response[[]FootballVenueResponseData]

// FootballVenue 取得足球場館
func (c *Client) FootballVenue(ctx context.Context, query FootballVenueQuery) (*FootballVenueResponse, error) {
	return secretGet[FootballVenueResponse](ctx, c, footballVenuePath, toQuery(query))
}

// FootballSeasonQuery 足球賽季查詢參數
type FootballSeasonQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// FootballSeasonResponse 足球賽季回應
type FootballSeasonResponse = Response[[]FootballSeasonResponseData]

// FootballSeason 取得足球賽季
func (c *Client) FootballSeason(ctx context.Context, query FootballSeasonQuery) (*FootballSeasonResponse, error) {
	return secretGet[FootballSeasonResponse](ctx, c, footballSeasonPath, toQuery(query))
}

// FootballStageQuery 足球階段查詢參數
type FootballStageQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// FootballStageResponse 足球階段回應
type FootballStageResponse = Response[[]FootballStageResponseData]

// FootballStage 取得足球階段
func (c *Client) FootballStage(ctx context.Context, query FootballStageQuery) (*FootballStageResponse, error) {
	return secretGet[FootballStageResponse](ctx, c, footballStagePath, toQuery(query))
}

// ====== MATCH 相關端點 ======

// FootballMatchQuery 比賽查詢參數（page/time/uuid 可選）
type FootballMatchQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// FootballMatchResponse 比賽回應
type FootballMatchResponse = Response[[]FootballMatchResponseData]

// FootballMatchList 取得比賽列表
func (c *Client) FootballMatchList(ctx context.Context, query FootballMatchQuery) (*FootballMatchResponse, error) {
	return secretGet[FootballMatchResponse](ctx, c, footballMatchListPath, toQuery(query))
}

// FootballMatchRecentList 取得近期比賽列表
func (c *Client) FootballMatchRecentList(ctx context.Context, query FootballMatchQuery) (*FootballMatchResponse, error) {
	return secretGet[FootballMatchResponse](ctx, c, footballMatchRecentListPath, toQuery(query))
}

// FootballMatchDetailLiveResponse 即時資料回應
type FootballMatchDetailLiveResponse = Response[[]FootballMatchDetailLiveResponseData]

// FootballMatchDetailLive 取得即時資料
func (c *Client) FootballMatchDetailLive(ctx context.Context) (*FootballMatchDetailLiveResponse, error) {
	return secretGet[FootballMatchDetailLiveResponse](ctx, c, footballMatchDetailLivePath)
}

// FootballMatchDiaryQuery 比賽日曆查詢參數
type FootballMatchDiaryQuery struct {
	TSP int64 `json:"tsp"`
}

// FootballMatchDiary 取得比賽日曆（按日期查詢賽程）
func (c *Client) FootballMatchDiary(ctx context.Context, query FootballMatchDiaryQuery) (*FootballMatchResponse, error) {
	return secretGet[FootballMatchResponse](ctx, c, footballMatchDiaryPath, toQuery(query))
}

// FootballMatchUUIDQuery 比賽 UUID 查詢參數（必填 uuid）
type FootballMatchUUIDQuery struct {
	UUID string `json:"uuid"`
}

// FootballMatchSeason 取得賽季比賽（所有賽季）
func (c *Client) FootballMatchSeason(ctx context.Context, query FootballMatchUUIDQuery) (*FootballMatchResponse, error) {
	return secretGet[FootballMatchResponse](ctx, c, footballMatchSeasonPath, toQuery(query))
}

// FootballMatchSeasonRecent 取得賽季比賽（最新賽季）
func (c *Client) FootballMatchSeasonRecent(ctx context.Context, query FootballMatchUUIDQuery) (*FootballMatchResponse, error) {
	return secretGet[FootballMatchResponse](ctx, c, footballMatchSeasonRecentPath, toQuery(query))
}

// FootballMatchLineupDetailResponse 單場陣容回應
type FootballMatchLineupDetailResponse = Response[FootballMatchLineupResponseData]

// FootballMatchLineupDetail 取得單場陣容
func (c *Client) FootballMatchLineupDetail(ctx context.Context, query FootballMatchUUIDQuery) (*FootballMatchLineupDetailResponse, error) {
	return secretGet[FootballMatchLineupDetailResponse](ctx, c, footballMatchLineupDetailPath, toQuery(query))
}

// FootballMatchTrendResponse 比賽走勢回應（detail 端點）
type FootballMatchTrendResponse = Response[FootballMatchTrendData]

// FootballMatchTrendDetail 取得比賽走勢詳情
func (c *Client) FootballMatchTrendDetail(ctx context.Context, query FootballMatchUUIDQuery) (*FootballMatchTrendResponse, error) {
	return secretGet[FootballMatchTrendResponse](ctx, c, footballMatchTrendDetailPath, toQuery(query))
}

// FootballMatchTrendLiveResponse 即時比賽走勢回應
type FootballMatchTrendLiveResponse = Response[[]FootballMatchTrendDetailResponseData]

// FootballMatchTrendLive 取得即時比賽走勢
func (c *Client) FootballMatchTrendLive(ctx context.Context) (*FootballMatchTrendLiveResponse, error) {
	return secretGet[FootballMatchTrendLiveResponse](ctx, c, footballMatchTrendLivePath)
}

// FootballMatchPlayerStatsListResponse 比賽球員統計列表回應
type FootballMatchPlayerStatsListResponse = Response[[]FootballMatchPlayerStatsResponseData]

// FootballMatchPlayerStatsList 取得比賽球員統計
func (c *Client) FootballMatchPlayerStatsList(ctx context.Context) (*FootballMatchPlayerStatsListResponse, error) {
	return secretGet[FootballMatchPlayerStatsListResponse](ctx, c, footballMatchPlayerStatsListPath)
}

// FootballMatchPlayerStatsDetailResponse 比賽球員統計詳情回應
type FootballMatchPlayerStatsDetailResponse = Response[[]FootballMatchPlayerStat]

// FootballMatchPlayerStatsDetail 取得比賽球員統計詳情
func (c *Client) FootballMatchPlayerStatsDetail(ctx context.Context, query FootballMatchUUIDQuery) (*FootballMatchPlayerStatsDetailResponse, error) {
	return secretGet[FootballMatchPlayerStatsDetailResponse](ctx, c, footballMatchPlayerStatsDetailPath, toQuery(query))
}

// FootballMatchTeamStatsListResponse 比賽球隊統計列表回應
type FootballMatchTeamStatsListResponse = Response[[]FootballMatchTeamStatsResponseData]

// FootballMatchTeamStatsList 取得比賽球隊統計
func (c *Client) FootballMatchTeamStatsList(ctx context.Context) (*FootballMatchTeamStatsListResponse, error) {
	return secretGet[FootballMatchTeamStatsListResponse](ctx, c, footballMatchTeamStatsListPath)
}

// FootballMatchTeamStatsDetailResponse 比賽球隊統計詳情回應
type FootballMatchTeamStatsDetailResponse = Response[[]FootballMatchTeamStat]

// FootballMatchTeamStatsDetail 取得比賽球隊統計詳情
func (c *Client) FootballMatchTeamStatsDetail(ctx context.Context, query FootballMatchUUIDQuery) (*FootballMatchTeamStatsDetailResponse, error) {
	return secretGet[FootballMatchTeamStatsDetailResponse](ctx, c, footballMatchTeamStatsDetailPath, toQuery(query))
}

// FootballMatchHalfTeamStatsListResponse 比賽半場球隊統計列表回應
type FootballMatchHalfTeamStatsListResponse = Response[json.RawMessage]

// FootballMatchHalfTeamStatsList 取得比賽半場球隊統計
func (c *Client) FootballMatchHalfTeamStatsList(ctx context.Context) (*FootballMatchHalfTeamStatsListResponse, error) {
	return secretGet[FootballMatchHalfTeamStatsListResponse](ctx, c, footballMatchHalfTeamStatsListPath)
}

// FootballMatchHalfTeamStatsDetailResponse 比賽半場球隊統計詳情回應
type FootballMatchHalfTeamStatsDetailResponse = Response[json.RawMessage]

// FootballMatchHalfTeamStatsDetail 取得比賽半場球隊統計詳情
func (c *Client) FootballMatchHalfTeamStatsDetail(ctx context.Context, query FootballMatchUUIDQuery) (*FootballMatchHalfTeamStatsDetailResponse, error) {
	return secretGet[FootballMatchHalfTeamStatsDetailResponse](ctx, c, footballMatchHalfTeamStatsDetailPath, toQuery(query))
}

// FootballMatchAnalysisResponse H2H 分析回應
type FootballMatchAnalysisResponse = Response[FootballMatchAnalysisResponseData]

// FootballMatchAnalysis 取得 H2H 分析
func (c *Client) FootballMatchAnalysis(ctx context.Context, query FootballMatchUUIDQuery) (*FootballMatchAnalysisResponse, error) {
	return secretGet[FootballMatchAnalysisResponse](ctx, c, footballMatchAnalysisPath, toQuery(query))
}

// FootballMatchLiveHistoryResponse 比賽歷史統計回應
type FootballMatchLiveHistoryResponse = Response[FootballMatchLiveHistoryResponseData]

// FootballMatchLiveHistory 取得比賽歷史統計
func (c *Client) FootballMatchLiveHistory(ctx context.Context, query FootballMatchUUIDQuery) (*FootballMatchLiveHistoryResponse, error) {
	return secretGet[FootballMatchLiveHistoryResponse](ctx, c, footballMatchLiveHistoryPath, toQuery(query))
}

// FootballMatchGoalLineDetailResponse 比賽進球線回應
type FootballMatchGoalLineDetailResponse = Response[[]FootballMatchGoalLineResponseData]

// FootballMatchGoalLineDetail 取得比賽進球線
func (c *Client) FootballMatchGoalLineDetail(ctx context.Context, query FootballMatchUUIDQuery) (*FootballMatchGoalLineDetailResponse, error) {
	return secretGet[FootballMatchGoalLineDetailResponse](ctx, c, footballMatchGoalLineDetailPath, toQuery(query))
}

// FootballMatchHighlightDetailResponse 比賽精華 GIF 回應
type FootballMatchHighlightDetailResponse = Response[[]FootballMatchHighlightResponseData]

// FootballMatchHighlightDetail 取得比賽精華 GIF
func (c *Client) FootballMatchHighlightDetail(ctx context.Context, query FootballMatchUUIDQuery) (*FootballMatchHighlightDetailResponse, error) {
	return secretGet[FootballMatchHighlightDetailResponse](ctx, c, footballMatchHighlightDetailPath, toQuery(query))
}

// FootballMatchTVListQuery 比賽電視頻道查詢參數
type FootballMatchTVListQuery = FootballMatchQuery

// FootballMatchTVListResponse 比賽電視頻道回應
type FootballMatchTVListResponse = Response[[]FootballMatchTVResponseData]

// FootballMatchTVList 取得比賽電視頻道
func (c *Client) FootballMatchTVList(ctx context.Context, query FootballMatchTVListQuery) (*FootballMatchTVListResponse, error) {
	return secretGet[FootballMatchTVListResponse](ctx, c, footballMatchTVListPath, toQuery(query))
}
