package thesportsgo

import (
	"context"
)

// ——— 基本資訊 ———

// EsportsRegionResponse 電競地區回應型別
type EsportsRegionResponse = Response[EsportsListWrapper[EsportsRegionResponseData]]

// EsportsRegion 取得電競地區
func (c *Client) EsportsRegion(ctx context.Context) (*EsportsRegionResponse, error) {
	return secretGet[EsportsRegionResponse](ctx, c, esportRegionPath)
}

// EsportsStageResponse 電競階段回應型別
type EsportsStageResponse = Response[EsportsListWrapper[EsportsStageResponseData]]

// EsportsStage 取得電競階段
func (c *Client) EsportsStage(ctx context.Context) (*EsportsStageResponse, error) {
	return secretGet[EsportsStageResponse](ctx, c, esportStagePath)
}

// EsportsMapResponse 電競地圖回應型別
type EsportsMapResponse = Response[EsportsListWrapper[EsportsMapResponseData]]

// EsportsMap 取得電競地圖
func (c *Client) EsportsMap(ctx context.Context) (*EsportsMapResponse, error) {
	return secretGet[EsportsMapResponse](ctx, c, esportMapPath)
}

// EsportsCountryResponse 電競國家/地區回應型別
type EsportsCountryResponse = Response[EsportsListWrapper[EsportsCountryResponseData]]

// EsportsCountry 取得電競國家/地區
func (c *Client) EsportsCountry(ctx context.Context) (*EsportsCountryResponse, error) {
	return secretGet[EsportsCountryResponse](ctx, c, esportCountryPath)
}

// ——— 賽事 ———

// EsportsTournamentQuery 電競賽事查詢參數
type EsportsTournamentQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// EsportsTournamentResponse 電競賽事回應型別
type EsportsTournamentResponse = Response[EsportsListWrapper[EsportsTournamentResponseData]]

// EsportsTournament 取得電競賽事
func (c *Client) EsportsTournament(ctx context.Context, query EsportsTournamentQuery) (*EsportsTournamentResponse, error) {
	return secretGet[EsportsTournamentResponse](ctx, c, esportTournamentPath, toQuery(query))
}

// ——— 戰隊 ———

// EsportsTeamQuery 電競戰隊查詢參數
type EsportsTeamQuery struct {
	Page         *int    `json:"page,omitempty"`
	Time         *int    `json:"time,omitempty"`
	UUID         *string `json:"uuid,omitempty"`
	TournamentID *string `json:"tournament_id,omitempty"`
}

// EsportsTeamResponse 電競戰隊回應型別
type EsportsTeamResponse = Response[EsportsListWrapper[EsportsTeamResponseData]]

// EsportsTeam 取得電競戰隊
func (c *Client) EsportsTeam(ctx context.Context, query EsportsTeamQuery) (*EsportsTeamResponse, error) {
	return secretGet[EsportsTeamResponse](ctx, c, esportTeamPath, toQuery(query))
}

// ——— 選手 ———

// EsportsPlayerQuery 電競選手查詢參數
type EsportsPlayerQuery struct {
	Page   *int    `json:"page,omitempty"`
	Time   *int    `json:"time,omitempty"`
	UUID   *string `json:"uuid,omitempty"`
	TeamID *string `json:"team_id,omitempty"`
}

// EsportsPlayerResponse 電競選手回應型別
type EsportsPlayerResponse = Response[EsportsListWrapper[EsportsPlayerResponseData]]

// EsportsPlayer 取得電競選手
func (c *Client) EsportsPlayer(ctx context.Context, query EsportsPlayerQuery) (*EsportsPlayerResponse, error) {
	return secretGet[EsportsPlayerResponse](ctx, c, esportPlayerPath, toQuery(query))
}

// ——— 比賽 ———

// EsportsMatchQuery 電競比賽查詢參數
type EsportsMatchQuery struct {
	Page         *int    `json:"page,omitempty"`
	Time         *int    `json:"time,omitempty"`
	UUID         *string `json:"uuid,omitempty"`
	TournamentID *string `json:"tournament_id,omitempty"`
}

// EsportsMatchResponse 電競比賽列表回應型別
type EsportsMatchResponse = Response[EsportsListWrapper[EsportsMatchResponseData]]

// EsportsMatch 取得電競比賽列表
func (c *Client) EsportsMatch(ctx context.Context, query EsportsMatchQuery) (*EsportsMatchResponse, error) {
	return secretGet[EsportsMatchResponse](ctx, c, esportMatchPath, toQuery(query))
}

// EsportsMatchDetailQuery 電競比賽詳情查詢參數
type EsportsMatchDetailQuery struct {
	MatchID *string `json:"match_id,omitempty"`
}

// EsportsMatchDetailResponse 電競比賽詳情回應型別（單一物件）
type EsportsMatchDetailResponse = Response[EsportsMatchDetailResponseData]

// EsportsMatchDetail 取得電競比賽詳情
func (c *Client) EsportsMatchDetail(ctx context.Context, query EsportsMatchDetailQuery) (*EsportsMatchDetailResponse, error) {
	return secretGet[EsportsMatchDetailResponse](ctx, c, esportMatchDetailPath, toQuery(query))
}

// EsportsMatchLineupQuery 電競比賽陣容查詢參數
type EsportsMatchLineupQuery struct {
	MatchID *string `json:"match_id,omitempty"`
}

// EsportsMatchLineupResponse 電競比賽陣容回應型別
type EsportsMatchLineupResponse = Response[EsportsMatchLineupResponseData]

// EsportsMatchLineup 取得電競比賽陣容
func (c *Client) EsportsMatchLineup(ctx context.Context, query EsportsMatchLineupQuery) (*EsportsMatchLineupResponse, error) {
	return secretGet[EsportsMatchLineupResponse](ctx, c, esportMatchLineupPath, toQuery(query))
}

// ——— 積分榜 ———

// EsportsStandingsQuery 電競積分榜查詢參數
type EsportsStandingsQuery struct {
	TournamentID *string `json:"tournament_id,omitempty"`
}

// EsportsStandingsResponse 電競積分榜回應型別
type EsportsStandingsResponse = Response[EsportsListWrapper[EsportsStandingsResponseData]]

// EsportsStandings 取得電競積分榜
func (c *Client) EsportsStandings(ctx context.Context, query EsportsStandingsQuery) (*EsportsStandingsResponse, error) {
	return secretGet[EsportsStandingsResponse](ctx, c, esportStandingsPath, toQuery(query))
}
