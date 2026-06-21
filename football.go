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

// ====== 非 MATCH 端點 ======

// FootballSeasonUUIDQuery 賽季 UUID 查詢參數（必填）
type FootballSeasonUUIDQuery struct {
	UUID string `json:"uuid"`
}

// --- Season table ---

// FootballSeasonRecentTableResponse 最新賽季積分榜回應
type FootballSeasonRecentTableResponse = Response[FootballSeasonTableResponseData]

// FootballSeasonRecentTable 取得最新賽季積分榜
func (c *Client) FootballSeasonRecentTable(ctx context.Context, query FootballSeasonUUIDQuery) (*FootballSeasonRecentTableResponse, error) {
	return secretGet[FootballSeasonRecentTableResponse](ctx, c, footballSeasonRecentTablePath, toQuery(query))
}

// FootballSeasonTableResponse 所有賽季積分榜回應
type FootballSeasonTableResponse = Response[FootballSeasonTableResponseData]

// FootballSeasonTable 取得所有賽季積分榜
func (c *Client) FootballSeasonTable(ctx context.Context, query FootballSeasonUUIDQuery) (*FootballSeasonTableResponse, error) {
	return secretGet[FootballSeasonTableResponse](ctx, c, footballSeasonTablePath, toQuery(query))
}

// --- Table live ---

// FootballTableLiveResponse 即時積分榜回應
type FootballTableLiveResponse = Response[[]FootballTableLiveData]

// FootballTableLive 取得即時積分榜
func (c *Client) FootballTableLive(ctx context.Context) (*FootballTableLiveResponse, error) {
	return secretGet[FootballTableLiveResponse](ctx, c, footballTableLivePath)
}

// --- Season team stat ---

// FootballSeasonTeamStatResponse 賽季球隊統計回應
type FootballSeasonTeamStatResponse = Response[[]FootballSeasonTeamStatResponseData]

// FootballSeasonTeamStat 取得賽季球隊統計（所有賽季）
func (c *Client) FootballSeasonTeamStat(ctx context.Context, query FootballSeasonUUIDQuery) (*FootballSeasonTeamStatResponse, error) {
	return secretGet[FootballSeasonTeamStatResponse](ctx, c, footballSeasonTeamStatPath, toQuery(query))
}

// --- Season player stat ---

// FootballSeasonPlayerStatResponse 賽季球員統計回應
type FootballSeasonPlayerStatResponse = Response[[]FootballSeasonPlayerStatResponseData]

// FootballSeasonPlayerStat 取得賽季球員統計（所有賽季）
func (c *Client) FootballSeasonPlayerStat(ctx context.Context, query FootballSeasonUUIDQuery) (*FootballSeasonPlayerStatResponse, error) {
	return secretGet[FootballSeasonPlayerStatResponse](ctx, c, footballSeasonPlayerStatPath, toQuery(query))
}

// --- Season shooter stat ---

// FootballSeasonShooterStatResponse 賽季射手榜回應
type FootballSeasonShooterStatResponse = Response[[]FootballSeasonShooterStatResponseData]

// FootballSeasonShooterStat 取得賽季射手榜（所有賽季）
func (c *Client) FootballSeasonShooterStat(ctx context.Context, query FootballSeasonUUIDQuery) (*FootballSeasonShooterStatResponse, error) {
	return secretGet[FootballSeasonShooterStatResponse](ctx, c, footballSeasonShooterStatPath, toQuery(query))
}

// --- Season recent team stat ---

// FootballSeasonRecentTeamStat 取得賽季球隊統計（最新賽季）
func (c *Client) FootballSeasonRecentTeamStat(ctx context.Context, query FootballSeasonUUIDQuery) (*FootballSeasonTeamStatResponse, error) {
	return secretGet[FootballSeasonTeamStatResponse](ctx, c, footballSeasonRecentTeamStatPath, toQuery(query))
}

// --- Season recent player stat ---

// FootballSeasonRecentPlayerStat 取得賽季球員統計（最新賽季）
func (c *Client) FootballSeasonRecentPlayerStat(ctx context.Context, query FootballSeasonUUIDQuery) (*FootballSeasonPlayerStatResponse, error) {
	return secretGet[FootballSeasonPlayerStatResponse](ctx, c, footballSeasonRecentPlayerStatPath, toQuery(query))
}

// --- Season recent shooter stat ---

// FootballSeasonRecentShooterStat 取得賽季射手榜（最新賽季）
func (c *Client) FootballSeasonRecentShooterStat(ctx context.Context, query FootballSeasonUUIDQuery) (*FootballSeasonShooterStatResponse, error) {
	return secretGet[FootballSeasonShooterStatResponse](ctx, c, footballSeasonRecentShooterStatPath, toQuery(query))
}

// --- Bracket ---

// FootballBracketSeasonResponse 盃賽回應
type FootballBracketSeasonResponse = Response[FootballBracketSeasonResponseData]

// FootballBracketSeason 取得盃賽
func (c *Client) FootballBracketSeason(ctx context.Context, query FootballSeasonUUIDQuery) (*FootballBracketSeasonResponse, error) {
	return secretGet[FootballBracketSeasonResponse](ctx, c, footballBracketSeasonPath, toQuery(query))
}

// --- Ranking ---

// FootballRankingFifaQuery FIFA 排名查詢參數
type FootballRankingFifaQuery struct {
	PubTime *int64 `json:"pub_time,omitempty"`
}

// FootballRankingFifaMenResponse FIFA 男子排名回應
type FootballRankingFifaMenResponse = Response[FootballRankingFifaResponseData]

// FootballRankingFifaMen 取得 FIFA 男子排名
func (c *Client) FootballRankingFifaMen(ctx context.Context, query FootballRankingFifaQuery) (*FootballRankingFifaMenResponse, error) {
	return secretGet[FootballRankingFifaMenResponse](ctx, c, footballRankingFifaMenPath, toQuery(query))
}

// FootballRankingFifaWomenResponse FIFA 女子排名回應
type FootballRankingFifaWomenResponse = Response[FootballRankingFifaResponseData]

// FootballRankingFifaWomen 取得 FIFA 女子排名
func (c *Client) FootballRankingFifaWomen(ctx context.Context, query FootballRankingFifaQuery) (*FootballRankingFifaWomenResponse, error) {
	return secretGet[FootballRankingFifaWomenResponse](ctx, c, footballRankingFifaWomenPath, toQuery(query))
}

// FootballRankingClubResponse 俱樂部排名回應
type FootballRankingClubResponse = Response[[]FootballRankingClubItem]

// FootballRankingClub 取得俱樂部排名
func (c *Client) FootballRankingClub(ctx context.Context) (*FootballRankingClubResponse, error) {
	return secretGet[FootballRankingClubResponse](ctx, c, footballRankingClubPath)
}

// FootballRankingFifaLiveResponse FIFA 即時排名回應
type FootballRankingFifaLiveResponse = Response[[]FootballRankingFifaLiveItem]

// FootballRankingFifaLive 取得 FIFA 即時排名
func (c *Client) FootballRankingFifaLive(ctx context.Context) (*FootballRankingFifaLiveResponse, error) {
	return secretGet[FootballRankingFifaLiveResponse](ctx, c, footballRankingFifaLivePath)
}

// --- Deleted ---

// FootballDeletedResponse 已刪除資料回應
type FootballDeletedResponse = Response[FootballDeletedResponseData]

// FootballDeleted 取得已刪除資料
func (c *Client) FootballDeleted(ctx context.Context) (*FootballDeletedResponse, error) {
	return secretGet[FootballDeletedResponse](ctx, c, footballDeletedPath)
}

// --- Data update ---

// FootballDataUpdateResponse 資料更新回應
type FootballDataUpdateResponse = Response[FootballDataUpdateResponseData]

// FootballDataUpdate 取得資料更新
func (c *Client) FootballDataUpdate(ctx context.Context) (*FootballDataUpdateResponse, error) {
	return secretGet[FootballDataUpdateResponse](ctx, c, footballDataUpdatePath)
}

// --- Language ---

// FootballLanguageResponse 語言回應
type FootballLanguageResponse = Response[[]FootballLanguageResponseData]

// FootballLanguage 取得語言列表
func (c *Client) FootballLanguage(ctx context.Context) (*FootballLanguageResponse, error) {
	return secretGet[FootballLanguageResponse](ctx, c, footballLanguagePath)
}

// --- Honor ---

// FootballHonorResponse 榮譽回應
type FootballHonorResponse = Response[[]FootballHonorResponseData]

// FootballHonor 取得榮譽列表
func (c *Client) FootballHonor(ctx context.Context) (*FootballHonorResponse, error) {
	return secretGet[FootballHonorResponse](ctx, c, footballHonorPath)
}

// --- Compensation ---

// FootballCompensationResponse 歷史賠償回應
type FootballCompensationResponse = Response[[]FootballCompensationResponseData]

// FootballCompensation 取得歷史賠償
func (c *Client) FootballCompensation(ctx context.Context) (*FootballCompensationResponse, error) {
	return secretGet[FootballCompensationResponse](ctx, c, footballCompensationPath)
}

// --- Odds ---

// FootballOddsLiveResponse 即時賠率回應
type FootballOddsLiveResponse = Response[FootballOddsLiveResponseData]

// FootballOddsLive 取得即時賠率
func (c *Client) FootballOddsLive(ctx context.Context) (*FootballOddsLiveResponse, error) {
	return secretGet[FootballOddsLiveResponse](ctx, c, footballOddsLivePath)
}

// FootballOddsHistoryQuery 單場賠率查詢參數
type FootballOddsHistoryQuery struct {
	UUID string `json:"uuid"`
}

// FootballOddsHistoryResponse 單場賠率回應
type FootballOddsHistoryResponse = Response[FootballOddsHistoryResponseData]

// FootballOddsHistory 取得單場賠率
func (c *Client) FootballOddsHistory(ctx context.Context, query FootballOddsHistoryQuery) (*FootballOddsHistoryResponse, error) {
	return secretGet[FootballOddsHistoryResponse](ctx, c, footballOddsHistoryPath, toQuery(query))
}

// FootballOddsUpdateResponse 賠率更新回應
type FootballOddsUpdateResponse = Response[FootballOddsUpdateResponseData]

// FootballOddsUpdate 取得賠率更新
func (c *Client) FootballOddsUpdate(ctx context.Context) (*FootballOddsUpdateResponse, error) {
	return secretGet[FootballOddsUpdateResponse](ctx, c, footballOddsUpdatePath)
}

// ====== TEAM / PLAYER / COACH 相關端點 ======

// FootballPageTimeUUIDQuery 通用查詢參數（page / time / uuid 可選）
type FootballPageTimeUUIDQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// --- Team squad ---

// FootballTeamSquadResponse 球隊陣容回應
type FootballTeamSquadResponse = Response[[]FootballTeamSquadResponseData]

// FootballTeamSquad 取得球隊陣容
func (c *Client) FootballTeamSquad(ctx context.Context, query FootballPageTimeUUIDQuery) (*FootballTeamSquadResponse, error) {
	return secretGet[FootballTeamSquadResponse](ctx, c, footballTeamSquadPath, toQuery(query))
}

// --- Team squad history ---

// FootballTeamSquadHistoryQuery 球隊歷史陣容查詢參數
type FootballTeamSquadHistoryQuery struct {
	UUID string `json:"uuid"`
}

// FootballTeamSquadHistoryResponse 球隊歷史陣容回應
type FootballTeamSquadHistoryResponse = Response[[]FootballTeamSquadHistoryResponseData]

// FootballTeamSquadHistory 取得球隊歷史陣容
func (c *Client) FootballTeamSquadHistory(ctx context.Context, query FootballTeamSquadHistoryQuery) (*FootballTeamSquadHistoryResponse, error) {
	return secretGet[FootballTeamSquadHistoryResponse](ctx, c, footballTeamSquadHistoryPath, toQuery(query))
}

// --- Team injury ---

// FootballTeamInjuryResponse 球隊傷病回應
type FootballTeamInjuryResponse = Response[[]FootballTeamInjuryResponseData]

// FootballTeamInjury 取得球隊傷病
func (c *Client) FootballTeamInjury(ctx context.Context, query FootballPageTimeUUIDQuery) (*FootballTeamInjuryResponse, error) {
	return secretGet[FootballTeamInjuryResponse](ctx, c, footballTeamInjuryPath, toQuery(query))
}

// --- Team honor ---

// FootballTeamHonorResponse 球隊榮譽回應
type FootballTeamHonorResponse = Response[[]FootballTeamHonorResponseData]

// FootballTeamHonor 取得球隊榮譽
func (c *Client) FootballTeamHonor(ctx context.Context, query FootballPageTimeUUIDQuery) (*FootballTeamHonorResponse, error) {
	return secretGet[FootballTeamHonorResponse](ctx, c, footballTeamHonorPath, toQuery(query))
}

// --- Team ability ---

// FootballTeamAbilityResponse 球隊能力回應
type FootballTeamAbilityResponse = Response[[]FootballTeamAbilityResponseData]

// FootballTeamAbility 取得球隊能力
func (c *Client) FootballTeamAbility(ctx context.Context, query FootballPageTimeUUIDQuery) (*FootballTeamAbilityResponse, error) {
	return secretGet[FootballTeamAbilityResponse](ctx, c, footballTeamAbilityPath, toQuery(query))
}

// --- Team record ---

// FootballTeamRecordResponse 球隊記錄回應
type FootballTeamRecordResponse = Response[[]FootballTeamRecordResponseData]

// FootballTeamRecord 取得球隊記錄
func (c *Client) FootballTeamRecord(ctx context.Context, query FootballPageTimeUUIDQuery) (*FootballTeamRecordResponse, error) {
	return secretGet[FootballTeamRecordResponse](ctx, c, footballTeamRecordPath, toQuery(query))
}

// --- Player ability ---

// FootballPlayerAbilityResponse 球員能力回應
type FootballPlayerAbilityResponse = Response[[]FootballPlayerAbilityResponseData]

// FootballPlayerAbility 取得球員能力
func (c *Client) FootballPlayerAbility(ctx context.Context, query FootballPageTimeUUIDQuery) (*FootballPlayerAbilityResponse, error) {
	return secretGet[FootballPlayerAbilityResponse](ctx, c, footballPlayerAbilityPath, toQuery(query))
}

// --- Player honor ---

// FootballPlayerHonorResponse 球員榮譽回應
type FootballPlayerHonorResponse = Response[[]FootballPlayerHonorResponseData]

// FootballPlayerHonor 取得球員榮譽
func (c *Client) FootballPlayerHonor(ctx context.Context, query FootballPageTimeUUIDQuery) (*FootballPlayerHonorResponse, error) {
	return secretGet[FootballPlayerHonorResponse](ctx, c, footballPlayerHonorPath, toQuery(query))
}

// --- Player market ---

// FootballPlayerMarketResponse 球員市場價值回應
type FootballPlayerMarketResponse = Response[[]FootballPlayerMarketResponseData]

// FootballPlayerMarket 取得球員市場價值
func (c *Client) FootballPlayerMarket(ctx context.Context, query FootballPageTimeUUIDQuery) (*FootballPlayerMarketResponse, error) {
	return secretGet[FootballPlayerMarketResponse](ctx, c, footballPlayerMarketPath, toQuery(query))
}

// --- Player transfer ---

// FootballPlayerTransferResponse 球員轉會回應
type FootballPlayerTransferResponse = Response[[]FootballPlayerTransferResponseData]

// FootballPlayerTransfer 取得球員轉會
func (c *Client) FootballPlayerTransfer(ctx context.Context, query FootballPageTimeUUIDQuery) (*FootballPlayerTransferResponse, error) {
	return secretGet[FootballPlayerTransferResponse](ctx, c, footballPlayerTransferPath, toQuery(query))
}

// --- Coach history ---

// FootballCoachHistoryResponse 教練執教經歷回應
type FootballCoachHistoryResponse = Response[[]FootballCoachHistoryResponseData]

// FootballCoachHistory 取得教練執教經歷
func (c *Client) FootballCoachHistory(ctx context.Context, query FootballPageTimeUUIDQuery) (*FootballCoachHistoryResponse, error) {
	return secretGet[FootballCoachHistoryResponse](ctx, c, footballCoachHistoryPath, toQuery(query))
}

// --- Coach honor ---

// FootballCoachHonorResponse 教練榮譽回應
type FootballCoachHonorResponse = Response[[]FootballCoachHonorResponseData]

// FootballCoachHonor 取得教練榮譽
func (c *Client) FootballCoachHonor(ctx context.Context, query FootballPageTimeUUIDQuery) (*FootballCoachHonorResponse, error) {
	return secretGet[FootballCoachHonorResponse](ctx, c, footballCoachHonorPath, toQuery(query))
}

// --- Competition best lineup ---

// FootballCompetitionBestLineupResponse 最佳陣容回應
type FootballCompetitionBestLineupResponse = Response[[]FootballCompetitionBestLineupResponseData]

// FootballCompetitionBestLineup 取得最佳陣容
func (c *Client) FootballCompetitionBestLineup(ctx context.Context, query FootballPageTimeUUIDQuery) (*FootballCompetitionBestLineupResponse, error) {
	return secretGet[FootballCompetitionBestLineupResponse](ctx, c, footballCompetitionBestLineupPath, toQuery(query))
}

// --- Competition best lineup detail ---

// FootballCompetitionBestLineupDetailResponse 最佳陣容詳情回應
type FootballCompetitionBestLineupDetailResponse = Response[[]FootballCompetitionBestLineupDetailResponseData]

// FootballCompetitionBestLineupDetail 取得最佳陣容詳情
func (c *Client) FootballCompetitionBestLineupDetail(ctx context.Context, query FootballSeasonUUIDQuery) (*FootballCompetitionBestLineupDetailResponse, error) {
	return secretGet[FootballCompetitionBestLineupDetailResponse](ctx, c, footballCompetitionBestLineupDetailPath, toQuery(query))
}

// --- Competition record ---

// FootballCompetitionRecordResponse 賽事紀錄回應
type FootballCompetitionRecordResponse = Response[[]FootballCompetitionRecordResponseData]

// FootballCompetitionRecord 取得賽事紀錄
func (c *Client) FootballCompetitionRecord(ctx context.Context, query FootballPageTimeUUIDQuery) (*FootballCompetitionRecordResponse, error) {
	return secretGet[FootballCompetitionRecordResponse](ctx, c, footballCompetitionRecordPath, toQuery(query))
}
