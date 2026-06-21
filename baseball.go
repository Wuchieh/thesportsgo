package thesportsgo

import (
	"context"
	"encoding/json"
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

// BaseballUniqueTournamentQuery 棒球賽事查詢參數
type BaseballUniqueTournamentQuery struct {
	UUID *string `json:"uuid,omitempty"`
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
}

// BaseballUniqueTournamentResponse 棒球賽事回應
type BaseballUniqueTournamentResponse = Response[[]BaseballUniqueTournamentResponseData]

// BaseballUniqueTournament 取得棒球賽事
func (c *Client) BaseballUniqueTournament(ctx context.Context, query BaseballUniqueTournamentQuery) (*BaseballUniqueTournamentResponse, error) {
	return secretGet[BaseballUniqueTournamentResponse](ctx, c, baseballUniqueTournamentPath, toQuery(query))
}

// BaseballTeamQuery 棒球球隊查詢參數
type BaseballTeamQuery struct {
	UUID *string `json:"uuid,omitempty"`
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
}

// BaseballTeamResponse 棒球球隊回應
type BaseballTeamResponse = Response[[]BaseballTeamResponseData]

// BaseballTeam 取得棒球球隊
func (c *Client) BaseballTeam(ctx context.Context, query BaseballTeamQuery) (*BaseballTeamResponse, error) {
	return secretGet[BaseballTeamResponse](ctx, c, baseballTeamPath, toQuery(query))
}

// BaseballPlayerQuery 棒球球員查詢參數
type BaseballPlayerQuery struct {
	UUID *string `json:"uuid,omitempty"`
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
}

// BaseballPlayerResponse 棒球球員回應
type BaseballPlayerResponse = Response[[]BaseballPlayerResponseData]

// BaseballPlayer 取得棒球球員
func (c *Client) BaseballPlayer(ctx context.Context, query BaseballPlayerQuery) (*BaseballPlayerResponse, error) {
	return secretGet[BaseballPlayerResponse](ctx, c, baseballPlayerPath, toQuery(query))
}

// BaseballVenueQuery 棒球場館查詢參數
type BaseballVenueQuery struct {
	UUID *string `json:"uuid,omitempty"`
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
}

// BaseballVenueResponse 棒球場館回應
type BaseballVenueResponse = Response[[]BaseballVenueResponseData]

// BaseballVenue 取得棒球場館
func (c *Client) BaseballVenue(ctx context.Context, query BaseballVenueQuery) (*BaseballVenueResponse, error) {
	return secretGet[BaseballVenueResponse](ctx, c, baseballVenuePath, toQuery(query))
}

// BaseballSeasonQuery 棒球賽季查詢參數
type BaseballSeasonQuery struct {
	UUID *string `json:"uuid,omitempty"`
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
}

// BaseballSeasonResponse 棒球賽季回應
type BaseballSeasonResponse = Response[[]BaseballSeasonResponseData]

// BaseballSeason 取得棒球賽季
func (c *Client) BaseballSeason(ctx context.Context, query BaseballSeasonQuery) (*BaseballSeasonResponse, error) {
	return secretGet[BaseballSeasonResponse](ctx, c, baseballSeasonPath, toQuery(query))
}

// BaseballMatchListQuery 棒球比賽列表查詢參數
type BaseballMatchListQuery struct {
	UUID *string `json:"uuid,omitempty"`
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
}

// BaseballMatchListResponse 棒球比賽列表回應
type BaseballMatchListResponse = Response[[]BaseballMatchResponseData]

// BaseballMatchList 取得棒球比賽列表
func (c *Client) BaseballMatchList(ctx context.Context, query BaseballMatchListQuery) (*BaseballMatchListResponse, error) {
	return secretGet[BaseballMatchListResponse](ctx, c, baseballMatchListPath, toQuery(query))
}

// BaseballMatchDetailLiveResponse 棒球即時比賽詳情回應
type BaseballMatchDetailLiveResponse = Response[[]BaseballMatchDetailLiveResultData]

// BaseballMatchDetailLive 取得棒球即時比賽詳情
func (c *Client) BaseballMatchDetailLive(ctx context.Context) (*BaseballMatchDetailLiveResponse, error) {
	return secretGet[BaseballMatchDetailLiveResponse](ctx, c, baseballMatchDetailLivePath)
}

// BaseballMatchDiaryQuery 棒球比賽日程查詢參數
type BaseballMatchDiaryQuery struct {
	TSP int `json:"tsp"`
}

// BaseballMatchDiaryResponse 棒球比賽日程回應
type BaseballMatchDiaryResponse = Response[[]BaseballMatchResponseData]

// BaseballMatchDiary 取得棒球比賽日程（日期查詢）
func (c *Client) BaseballMatchDiary(ctx context.Context, query BaseballMatchDiaryQuery) (*BaseballMatchDiaryResponse, error) {
	return secretGet[BaseballMatchDiaryResponse](ctx, c, baseballMatchDiaryPath, toQuery(query))
}

// BaseballMatchSeasonQuery 棒球比賽賽季查詢參數
type BaseballMatchSeasonQuery struct {
	UUID string `json:"uuid"`
}

// BaseballMatchSeasonResponse 棒球比賽賽季回應
type BaseballMatchSeasonResponse = Response[[]BaseballMatchResponseData]

// BaseballMatchSeason 取得棒球比賽賽季（賽季查詢）
func (c *Client) BaseballMatchSeason(ctx context.Context, query BaseballMatchSeasonQuery) (*BaseballMatchSeasonResponse, error) {
	return secretGet[BaseballMatchSeasonResponse](ctx, c, baseballMatchSeasonPath, toQuery(query))
}

// BaseballMatchLiveHistoryQuery 棒球歷史比賽統計查詢參數
type BaseballMatchLiveHistoryQuery struct {
	UUID string `json:"uuid"`
}

// BaseballMatchLiveHistoryResponse 棒球歷史比賽統計回應
type BaseballMatchLiveHistoryResponse = Response[BaseballMatchLiveHistoryResultData]

// BaseballMatchLiveHistory 取得棒球歷史比賽統計
func (c *Client) BaseballMatchLiveHistory(ctx context.Context, query BaseballMatchLiveHistoryQuery) (*BaseballMatchLiveHistoryResponse, error) {
	return secretGet[BaseballMatchLiveHistoryResponse](ctx, c, baseballMatchLiveHistoryPath, toQuery(query))
}

// BaseballSeasonTableQuery 棒球賽季排名查詢參數
type BaseballSeasonTableQuery struct {
	UUID string `json:"uuid"`
}

// BaseballSeasonTableResponse 棒球賽季排名回應
type BaseballSeasonTableResponse = Response[BaseballSeasonTableResultData]

// BaseballSeasonTable 取得棒球賽季排名
func (c *Client) BaseballSeasonTable(ctx context.Context, query BaseballSeasonTableQuery) (*BaseballSeasonTableResponse, error) {
	return secretGet[BaseballSeasonTableResponse](ctx, c, baseballSeasonTablePath, toQuery(query))
}

// BaseballSeasonTeamStatsQuery 棒球賽季球隊統計查詢參數
type BaseballSeasonTeamStatsQuery struct {
	UUID string `json:"uuid"`
}

// BaseballSeasonTeamStatsResponse 棒球賽季球隊統計回應
type BaseballSeasonTeamStatsResponse = Response[[]BaseballSeasonTeamStatsResultData]

// BaseballSeasonTeamStats 取得棒球賽季球隊統計
func (c *Client) BaseballSeasonTeamStats(ctx context.Context, query BaseballSeasonTeamStatsQuery) (*BaseballSeasonTeamStatsResponse, error) {
	return secretGet[BaseballSeasonTeamStatsResponse](ctx, c, baseballSeasonTeamStatsPath, toQuery(query))
}

// BaseballSeasonPlayerStatsQuery 棒球賽季球員統計查詢參數
type BaseballSeasonPlayerStatsQuery struct {
	UUID string `json:"uuid"`
}

// BaseballSeasonPlayerStatsResponse 棒球賽季球員統計回應
type BaseballSeasonPlayerStatsResponse = Response[[]BaseballSeasonPlayerStatsResultData]

// BaseballSeasonPlayerStats 取得棒球賽季球員統計
func (c *Client) BaseballSeasonPlayerStats(ctx context.Context, query BaseballSeasonPlayerStatsQuery) (*BaseballSeasonPlayerStatsResponse, error) {
	return secretGet[BaseballSeasonPlayerStatsResponse](ctx, c, baseballSeasonPlayerStatsPath, toQuery(query))
}

// BaseballSeasonCoachStatsQuery 棒球賽季教練統計查詢參數
type BaseballSeasonCoachStatsQuery struct {
	UUID string `json:"uuid"`
}

// BaseballSeasonCoachStatsResponse 棒球賽季教練統計回應
type BaseballSeasonCoachStatsResponse = Response[[]BaseballSeasonCoachStatsResultData]

// BaseballSeasonCoachStats 取得棒球賽季教練統計
func (c *Client) BaseballSeasonCoachStats(ctx context.Context, query BaseballSeasonCoachStatsQuery) (*BaseballSeasonCoachStatsResponse, error) {
	return secretGet[BaseballSeasonCoachStatsResponse](ctx, c, baseballSeasonCoachStatsPath, toQuery(query))
}

// BaseballTeamSquadQuery 棒球球隊陣容查詢參數
type BaseballTeamSquadQuery struct {
	UUID *string `json:"uuid,omitempty"`
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
}

// BaseballTeamSquadResponse 棒球球隊陣容回應
type BaseballTeamSquadResponse = Response[[]BaseballTeamSquadResponseData]

// BaseballTeamSquad 取得棒球球隊陣容
func (c *Client) BaseballTeamSquad(ctx context.Context, query BaseballTeamSquadQuery) (*BaseballTeamSquadResponse, error) {
	return secretGet[BaseballTeamSquadResponse](ctx, c, baseballTeamSquadPath, toQuery(query))
}

// BaseballTeamInjuryQuery 棒球球隊傷病查詢參數
type BaseballTeamInjuryQuery struct {
	UUID *string `json:"uuid,omitempty"`
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
}

// BaseballTeamInjuryResponse 棒球球隊傷病回應
type BaseballTeamInjuryResponse = Response[[]BaseballTeamInjuryResponseData]

// BaseballTeamInjury 取得棒球球隊傷病
func (c *Client) BaseballTeamInjury(ctx context.Context, query BaseballTeamInjuryQuery) (*BaseballTeamInjuryResponse, error) {
	return secretGet[BaseballTeamInjuryResponse](ctx, c, baseballTeamInjuryPath, toQuery(query))
}

// BaseballTeamHonorQuery 棒球球隊榮譽查詢參數
type BaseballTeamHonorQuery struct {
	UUID *string `json:"uuid,omitempty"`
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
}

// BaseballTeamHonorResponse 棒球球隊榮譽回應
type BaseballTeamHonorResponse = Response[[]BaseballTeamHonorResponseData]

// BaseballTeamHonor 取得棒球球隊榮譽
func (c *Client) BaseballTeamHonor(ctx context.Context, query BaseballTeamHonorQuery) (*BaseballTeamHonorResponse, error) {
	return secretGet[BaseballTeamHonorResponse](ctx, c, baseballTeamHonorPath, toQuery(query))
}

// BaseballPlayerHonorQuery 棒球球員榮譽查詢參數
type BaseballPlayerHonorQuery struct {
	UUID *string `json:"uuid,omitempty"`
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
}

// BaseballPlayerHonorResponse 棒球球員榮譽回應
type BaseballPlayerHonorResponse = Response[[]BaseballPlayerHonorResponseData]

// BaseballPlayerHonor 取得棒球球員榮譽
func (c *Client) BaseballPlayerHonor(ctx context.Context, query BaseballPlayerHonorQuery) (*BaseballPlayerHonorResponse, error) {
	return secretGet[BaseballPlayerHonorResponse](ctx, c, baseballPlayerHonorPath, toQuery(query))
}

// BaseballCoachHonorQuery 棒球教練榮譽查詢參數
type BaseballCoachHonorQuery struct {
	UUID *string `json:"uuid,omitempty"`
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
}

// BaseballCoachHonorResponse 棒球教練榮譽回應
type BaseballCoachHonorResponse = Response[[]BaseballCoachHonorResponseData]

// BaseballCoachHonor 取得棒球教練榮譽
func (c *Client) BaseballCoachHonor(ctx context.Context, query BaseballCoachHonorQuery) (*BaseballCoachHonorResponse, error) {
	return secretGet[BaseballCoachHonorResponse](ctx, c, baseballCoachHonorPath, toQuery(query))
}

// BaseballHonorQuery 棒球榮譽列表查詢參數
type BaseballHonorQuery struct {
	UUID *string `json:"uuid,omitempty"`
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
}

// BaseballHonorResponse 棒球榮譽列表回應
type BaseballHonorResponse = Response[[]BaseballHonorResponseData]

// BaseballHonor 取得棒球榮譽列表
func (c *Client) BaseballHonor(ctx context.Context, query BaseballHonorQuery) (*BaseballHonorResponse, error) {
	return secretGet[BaseballHonorResponse](ctx, c, baseballHonorPath, toQuery(query))
}

// BaseballBracketSeasonQuery 棒球對戰圖查詢參數
type BaseballBracketSeasonQuery struct {
	UUID string `json:"uuid"`
}

// BaseballBracketSeasonResponse 棒球對戰圖回應
type BaseballBracketSeasonResponse = Response[BaseballBracketSeasonResultData]

// BaseballBracketSeason 取得棒球對戰圖
func (c *Client) BaseballBracketSeason(ctx context.Context, query BaseballBracketSeasonQuery) (*BaseballBracketSeasonResponse, error) {
	return secretGet[BaseballBracketSeasonResponse](ctx, c, baseballBracketSeasonPath, toQuery(query))
}

// BaseballDataUpdateResponse 棒球資料更新回應
type BaseballDataUpdateResponse = Response[json.RawMessage]

// BaseballDataUpdate 取得棒球資料更新
func (c *Client) BaseballDataUpdate(ctx context.Context) (*BaseballDataUpdateResponse, error) {
	return secretGet[BaseballDataUpdateResponse](ctx, c, baseballDataUpdatePath)
}

// BaseballDeletedResponse 棒球刪除資料回應
type BaseballDeletedResponse = Response[json.RawMessage]

// BaseballDeleted 取得棒球刪除資料
func (c *Client) BaseballDeleted(ctx context.Context) (*BaseballDeletedResponse, error) {
	return secretGet[BaseballDeletedResponse](ctx, c, baseballDeletedPath)
}

// BaseballOddsLiveResponse 棒球即時賠率回應
type BaseballOddsLiveResponse = Response[json.RawMessage]

// BaseballOddsLive 取得棒球即時賠率
func (c *Client) BaseballOddsLive(ctx context.Context) (*BaseballOddsLiveResponse, error) {
	return secretGet[BaseballOddsLiveResponse](ctx, c, baseballOddsLivePath)
}

// BaseballOddsHistoryQuery 棒球歷史賠率查詢參數
type BaseballOddsHistoryQuery struct {
	UUID string `json:"uuid"`
}

// BaseballOddsHistoryResponse 棒球歷史賠率回應
type BaseballOddsHistoryResponse = Response[json.RawMessage]

// BaseballOddsHistory 取得棒球歷史賠率
func (c *Client) BaseballOddsHistory(ctx context.Context, query BaseballOddsHistoryQuery) (*BaseballOddsHistoryResponse, error) {
	return secretGet[BaseballOddsHistoryResponse](ctx, c, baseballOddsHistoryPath, toQuery(query))
}
