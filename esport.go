package thesportsgo

import (
	"context"
	"encoding/json"
	"io"
)

// ——— 通用查詢參數 ———

// CsgoListQuery CS:GO 列表通用查詢參數（含 page / time / uuid）
type CsgoListQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// ——— CS:GO 國家/地區 ———

// CsgoCountryResponse CS:GO 國家/地區回應型別
type CsgoCountryResponse = Response[[]CsgoCountryResponseData]

// CsgoCountry 取得 CS:GO 國家/地區列表
func (c *Client) CsgoCountry(ctx context.Context, query CsgoListQuery) (*CsgoCountryResponse, error) {
	return secretGet[CsgoCountryResponse](ctx, c, csgoCountryPath, toQuery(query))
}

// ——— CS:GO 賽事 ———

// CsgoTournamentResponse CS:GO 賽事回應型別
type CsgoTournamentResponse = Response[[]CsgoTournamentResponseData]

// CsgoTournament 取得 CS:GO 賽事列表
func (c *Client) CsgoTournament(ctx context.Context, query CsgoListQuery) (*CsgoTournamentResponse, error) {
	return secretGet[CsgoTournamentResponse](ctx, c, csgoTournamentPath, toQuery(query))
}

// ——— CS:GO 戰隊 ———

// CsgoTeamResponse CS:GO 戰隊回應型別
type CsgoTeamResponse = Response[[]CsgoTeamResponseData]

// CsgoTeam 取得 CS:GO 戰隊列表
func (c *Client) CsgoTeam(ctx context.Context, query CsgoListQuery) (*CsgoTeamResponse, error) {
	return secretGet[CsgoTeamResponse](ctx, c, csgoTeamPath, toQuery(query))
}

// ——— CS:GO 選手 ———

// CsgoPlayerResponse CS:GO 選手回應型別
type CsgoPlayerResponse = Response[[]CsgoPlayerResponseData]

// CsgoPlayer 取得 CS:GO 選手列表
func (c *Client) CsgoPlayer(ctx context.Context, query CsgoListQuery) (*CsgoPlayerResponse, error) {
	return secretGet[CsgoPlayerResponse](ctx, c, csgoPlayerPath, toQuery(query))
}

// ——— CS:GO 階段 ———

// CsgoStageResponse CS:GO 階段回應型別
type CsgoStageResponse = Response[[]CsgoStageResponseData]

// CsgoStage 取得 CS:GO 階段列表
func (c *Client) CsgoStage(ctx context.Context, query CsgoListQuery) (*CsgoStageResponse, error) {
	return secretGet[CsgoStageResponse](ctx, c, csgoStagePath, toQuery(query))
}

// ——— CS:GO 比賽 ———

// CsgoMatchResponse CS:GO 比賽回應型別
type CsgoMatchResponse = Response[[]CsgoMatchResponseData]

// CsgoMatch 取得 CS:GO 比賽列表
func (c *Client) CsgoMatch(ctx context.Context, query CsgoListQuery) (*CsgoMatchResponse, error) {
	return secretGet[CsgoMatchResponse](ctx, c, csgoMatchPath, toQuery(query))
}

// ——— CS:GO 賽程日誌（依日期查詢）———

// CsgoMatchDiaryQuery CS:GO 賽程日誌查詢參數
type CsgoMatchDiaryQuery struct {
	Tsp int `json:"tsp"`
}

// CsgoMatchDiary 取得 CS:GO 賽程與結果（依日期查詢）
// 注意：此端點回應含 results_extra，需自訂反序列化
func (c *Client) CsgoMatchDiary(ctx context.Context, query CsgoMatchDiaryQuery) (*CsgoMatchDiaryResponse, error) {
	resp, err := c.SecretGet(ctx, csgoMatchDiaryPath, toQuery(query))
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result CsgoMatchDiaryResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	if result.Err != nil && *result.Err != "" {
		return nil, result
	}
	return &result, nil
}

// ——— CS:GO 賽程日誌（依賽事查詢）———

// CsgoMatchTournamentQuery CS:GO 賽程依賽事查詢參數
type CsgoMatchTournamentQuery struct {
	UUID string `json:"uuid"`
}

// CsgoMatchTournament 取得 CS:GO 賽程與結果（依賽事查詢）
func (c *Client) CsgoMatchTournament(ctx context.Context, query CsgoMatchTournamentQuery) (*CsgoMatchResponse, error) {
	return secretGet[CsgoMatchResponse](ctx, c, csgoMatchTournamentPath, toQuery(query))
}

// ——— CS:GO 單場比賽 ———

// CsgoMatchSingleResponse CS:GO 單場比賽回應型別
type CsgoMatchSingleResponse = Response[[]CsgoSingleMatchResponseData]

// CsgoMatchSingle 取得 CS:GO 單場比賽列表
func (c *Client) CsgoMatchSingle(ctx context.Context, query CsgoListQuery) (*CsgoMatchSingleResponse, error) {
	return secretGet[CsgoMatchSingleResponse](ctx, c, csgoMatchSinglePath, toQuery(query))
}

// ——— CS:GO 單場選手統計 ———

// CsgoMatchSinglePlayerStatResponse CS:GO 單場選手統計回應型別
type CsgoMatchSinglePlayerStatResponse = Response[[]CsgoSingleMatchPlayerStatResponseData]

// CsgoMatchSinglePlayerStat 取得 CS:GO 單場選手統計
func (c *Client) CsgoMatchSinglePlayerStat(ctx context.Context, query CsgoListQuery) (*CsgoMatchSinglePlayerStatResponse, error) {
	return secretGet[CsgoMatchSinglePlayerStatResponse](ctx, c, csgoMatchSinglePlayerStatPath, toQuery(query))
}

// ——— CS:GO 即時數據 ———

// CsgoMatchDetailLiveResponse CS:GO 即時數據回應型別
type CsgoMatchDetailLiveResponse = Response[[]CsgoDetailLiveResponseData]

// CsgoMatchDetailLive 取得 CS:GO 即時數據
func (c *Client) CsgoMatchDetailLive(ctx context.Context) (*CsgoMatchDetailLiveResponse, error) {
	return secretGet[CsgoMatchDetailLiveResponse](ctx, c, csgoMatchDetailLivePath)
}

// ——— CS:GO 刪除記錄 ———

// CsgoDeleteResponse CS:GO 刪除記錄回應型別
type CsgoDeleteResponse = Response[[]CsgoDeleteResponseData]

// CsgoDelete 取得 CS:GO 刪除記錄
func (c *Client) CsgoDelete(ctx context.Context) (*CsgoDeleteResponse, error) {
	return secretGet[CsgoDeleteResponse](ctx, c, csgoDeletePath)
}

// ——— CS:GO 戰隊統計 ———

// CsgoTeamStatResponse CS:GO 戰隊統計回應型別
type CsgoTeamStatResponse = Response[[]CsgoTeamStatResponseData]

// CsgoTeamStat 取得 CS:GO 戰隊統計
func (c *Client) CsgoTeamStat(ctx context.Context, query CsgoListQuery) (*CsgoTeamStatResponse, error) {
	return secretGet[CsgoTeamStatResponse](ctx, c, csgoTeamStatPath, toQuery(query))
}

// ——— CS:GO 選手統計 ———

// CsgoPlayerStatResponse CS:GO 選手統計回應型別
type CsgoPlayerStatResponse = Response[[]CsgoPlayerStatResponseData]

// CsgoPlayerStat 取得 CS:GO 選手統計
func (c *Client) CsgoPlayerStat(ctx context.Context, query CsgoListQuery) (*CsgoPlayerStatResponse, error) {
	return secretGet[CsgoPlayerStatResponse](ctx, c, csgoPlayerStatPath, toQuery(query))
}

// ——— CS:GO 賽事統計 ———

// CsgoTournamentStatQuery CS:GO 賽事統計查詢參數
type CsgoTournamentStatQuery struct {
	UUID string `json:"uuid"`
}

// CsgoTournamentStatResponse CS:GO 賽事統計回應型別
type CsgoTournamentStatResponse = Response[CsgoTournamentStatResponseData]

// CsgoTournamentStat 取得 CS:GO 賽事統計
func (c *Client) CsgoTournamentStat(ctx context.Context, query CsgoTournamentStatQuery) (*CsgoTournamentStatResponse, error) {
	return secretGet[CsgoTournamentStatResponse](ctx, c, csgoTournamentStatPath, toQuery(query))
}

// ——— CS:GO 賽事戰隊 ———

// CsgoTournamentTeamResponse CS:GO 賽事戰隊回應型別
type CsgoTournamentTeamResponse = Response[[]CsgoTournamentTeamResponseData]

// CsgoTournamentTeam 取得 CS:GO 賽事戰隊列表
func (c *Client) CsgoTournamentTeam(ctx context.Context, query CsgoListQuery) (*CsgoTournamentTeamResponse, error) {
	return secretGet[CsgoTournamentTeamResponse](ctx, c, csgoTournamentTeamPath, toQuery(query))
}

// ——— CS:GO 賽事回合 ———

// CsgoTournamentRoundResponse CS:GO 賽事回合回應型別
type CsgoTournamentRoundResponse = Response[[]CsgoTournamentRoundResponseData]

// CsgoTournamentRound 取得 CS:GO 賽事回合列表
func (c *Client) CsgoTournamentRound(ctx context.Context, query CsgoListQuery) (*CsgoTournamentRoundResponse, error) {
	return secretGet[CsgoTournamentRoundResponse](ctx, c, csgoTournamentRoundPath, toQuery(query))
}

// ——— CS:GO 賽事對戰表 ———

// CsgoTournamentBracketResponse CS:GO 賽事對戰表回應型別
type CsgoTournamentBracketResponse = Response[[]CsgoTournamentBracketResponseData]

// CsgoTournamentBracket 取得 CS:GO 賽事對戰表
func (c *Client) CsgoTournamentBracket(ctx context.Context, query CsgoListQuery) (*CsgoTournamentBracketResponse, error) {
	return secretGet[CsgoTournamentBracketResponse](ctx, c, csgoTournamentBracketPath, toQuery(query))
}

// ——— CS:GO 賽事對戰表連線 ———

// CsgoTournamentBracketLineResponse CS:GO 賽事對戰表連線回應型別
type CsgoTournamentBracketLineResponse = Response[[]CsgoTournamentBracketLineResponseData]

// CsgoTournamentBracketLine 取得 CS:GO 賽事對戰表連線
func (c *Client) CsgoTournamentBracketLine(ctx context.Context, query CsgoListQuery) (*CsgoTournamentBracketLineResponse, error) {
	return secretGet[CsgoTournamentBracketLineResponse](ctx, c, csgoTournamentBracketLinePath, toQuery(query))
}

// ——— CS:GO 賽事積分表 ———

// CsgoTournamentTableResponse CS:GO 賽事積分表回應型別
type CsgoTournamentTableResponse = Response[[]CsgoTournamentTableResponseData]

// CsgoTournamentTable 取得 CS:GO 賽事積分表
func (c *Client) CsgoTournamentTable(ctx context.Context, query CsgoListQuery) (*CsgoTournamentTableResponse, error) {
	return secretGet[CsgoTournamentTableResponse](ctx, c, csgoTournamentTablePath, toQuery(query))
}

// ——— CS:GO 賠率 ———

// CsgoOddsLiveResponse CS:GO 即時賠率回應型別
type CsgoOddsLiveResponse = Response[[]CsgoOddsLiveResponseData]

// CsgoOddsLive 取得 CS:GO 即時賠率
func (c *Client) CsgoOddsLive(ctx context.Context) (*CsgoOddsLiveResponse, error) {
	return secretGet[CsgoOddsLiveResponse](ctx, c, csgoOddsLivePath)
}

// CsgoOddsHistoryQuery CS:GO 賠率歷史查詢參數
type CsgoOddsHistoryQuery struct {
	UUID string `json:"uuid"`
}

// CsgoOddsHistoryResponse CS:GO 賠率歷史回應型別
type CsgoOddsHistoryResponse = Response[[]CsgoOddsHistoryResponseData]

// CsgoOddsHistory 取得 CS:GO 單場賠率歷史
func (c *Client) CsgoOddsHistory(ctx context.Context, query CsgoOddsHistoryQuery) (*CsgoOddsHistoryResponse, error) {
	return secretGet[CsgoOddsHistoryResponse](ctx, c, csgoOddsHistoryPath, toQuery(query))
}

// ——— Dota2 英雄 ———

// Dota2HeroResponse Dota2 英雄回應型別
type Dota2HeroResponse = Response[[]Dota2HeroResponseData]

// Dota2Hero 取得 Dota2 英雄列表
func (c *Client) Dota2Hero(ctx context.Context, query CsgoListQuery) (*Dota2HeroResponse, error) {
	return secretGet[Dota2HeroResponse](ctx, c, dota2HeroPath, toQuery(query))
}

// ——— Dota2 裝備 ———

// Dota2EquipmentResponse Dota2 裝備回應型別
type Dota2EquipmentResponse = Response[[]Dota2EquipmentResponseData]

// Dota2Equipment 取得 Dota2 裝備列表
func (c *Client) Dota2Equipment(ctx context.Context, query CsgoListQuery) (*Dota2EquipmentResponse, error) {
	return secretGet[Dota2EquipmentResponse](ctx, c, dota2EquipmentPath, toQuery(query))
}

// ——— LoL 召喚師技能 ———

// LolSpellResponse LoL 召喚師技能回應型別
type LolSpellResponse = Response[[]LolSpellResponseData]

// LolSpell 取得 LoL 召喚師技能列表
func (c *Client) LolSpell(ctx context.Context, query CsgoListQuery) (*LolSpellResponse, error) {
	return secretGet[LolSpellResponse](ctx, c, lolSpellPath, toQuery(query))
}

// ——— LoL 符文 ———

// LolRuneResponse LoL 符文回應型別
type LolRuneResponse = Response[[]LolRuneResponseData]

// LolRune 取得 LoL 符文列表
func (c *Client) LolRune(ctx context.Context, query CsgoListQuery) (*LolRuneResponse, error) {
	return secretGet[LolRuneResponse](ctx, c, lolRunePath, toQuery(query))
}

// ——— LoL 比賽事件 ———

// LolMatchEventResponse LoL 比賽事件回應型別
type LolMatchEventResponse = Response[[]LolMatchEventResponseData]

// LolMatchEvent 取得 LoL 比賽事件列表
func (c *Client) LolMatchEvent(ctx context.Context, query CsgoListQuery) (*LolMatchEventResponse, error) {
	return secretGet[LolMatchEventResponse](ctx, c, lolMatchEventPath, toQuery(query))
}
