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

//	===  基本清單端點（有查詢參數：page/time/uuid）  ===

// BasketballCompetitionQuery 籃球賽事查詢參數
type BasketballCompetitionQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// BasketballCompetitionResponse 籃球賽事回應型別
type BasketballCompetitionResponse = Response[[]BasketballCompetitionResponseData]

// BasketballCompetition 取得籃球賽事
func (c *Client) BasketballCompetition(ctx context.Context, query BasketballCompetitionQuery) (*BasketballCompetitionResponse, error) {
	return secretGet[BasketballCompetitionResponse](ctx, c, basketballCompetitionPath, toQuery(query))
}

// BasketballTeamQuery 籃球球隊查詢參數
type BasketballTeamQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// BasketballTeamResponse 籃球球隊回應型別
type BasketballTeamResponse = Response[[]BasketballTeamResponseData]

// BasketballTeam 取得籃球球隊
func (c *Client) BasketballTeam(ctx context.Context, query BasketballTeamQuery) (*BasketballTeamResponse, error) {
	return secretGet[BasketballTeamResponse](ctx, c, basketballTeamPath, toQuery(query))
}

// BasketballPlayerQuery 籃球球員查詢參數
type BasketballPlayerQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// BasketballPlayerResponse 籃球球員回應型別
type BasketballPlayerResponse = Response[[]BasketballPlayerResponseData]

// BasketballPlayer 取得籃球球員
func (c *Client) BasketballPlayer(ctx context.Context, query BasketballPlayerQuery) (*BasketballPlayerResponse, error) {
	return secretGet[BasketballPlayerResponse](ctx, c, basketballPlayerPath, toQuery(query))
}

// BasketballCoachQuery 籃球教練查詢參數
type BasketballCoachQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// BasketballCoachResponse 籃球教練回應型別
type BasketballCoachResponse = Response[[]BasketballCoachResponseData]

// BasketballCoach 取得籃球教練
func (c *Client) BasketballCoach(ctx context.Context, query BasketballCoachQuery) (*BasketballCoachResponse, error) {
	return secretGet[BasketballCoachResponse](ctx, c, basketballCoachPath, toQuery(query))
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

// BasketballMatchRecentQuery 籃球近期比賽查詢參數
type BasketballMatchRecentQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// BasketballMatchRecentResponse 籃球近期比賽回應型別
type BasketballMatchRecentResponse = Response[[]BasketballMatchResponseData]

// BasketballMatchRecent 取得籃球近期比賽（30 天內）
func (c *Client) BasketballMatchRecent(ctx context.Context, query BasketballMatchRecentQuery) (*BasketballMatchRecentResponse, error) {
	return secretGet[BasketballMatchRecentResponse](ctx, c, basketballMatchRecentPath, toQuery(query))
}

// BasketballMatchDiaryQuery 籃球賽程日期查詢參數
type BasketballMatchDiaryQuery struct {
	// 查詢日期時間戳
	Tsp int `json:"tsp"`
}

// BasketballMatchDiaryResponse 籃球賽程日期查詢回應型別
type BasketballMatchDiaryResponse = Response[[]BasketballMatchResponseData]

// BasketballMatchDiary 取得籃球賽程與結果（日期查詢）
func (c *Client) BasketballMatchDiary(ctx context.Context, query BasketballMatchDiaryQuery) (*BasketballMatchDiaryResponse, error) {
	return secretGet[BasketballMatchDiaryResponse](ctx, c, basketballMatchDiaryPath, toQuery(query))
}

// BasketballMatchSeasonRecentQuery 籃球賽季賽程查詢參數（最新賽季）
type BasketballMatchSeasonRecentQuery struct {
	// 賽季 id
	UUID string `json:"uuid"`
}

// BasketballMatchSeasonRecentResponse 籃球賽季賽程回應型別（最新賽季）
type BasketballMatchSeasonRecentResponse = Response[[]BasketballMatchResponseData]

// BasketballMatchSeasonRecent 取得籃球賽程與結果（賽季查詢-最新賽季）
func (c *Client) BasketballMatchSeasonRecent(ctx context.Context, query BasketballMatchSeasonRecentQuery) (*BasketballMatchSeasonRecentResponse, error) {
	return secretGet[BasketballMatchSeasonRecentResponse](ctx, c, basketballMatchSeasonRecentPath, toQuery(query))
}

// BasketballMatchDetailLiveResponse 籃球即時數據回應型別
type BasketballMatchDetailLiveResponse = Response[[]BasketballMatchDetailLiveResponseData]

// BasketballMatchDetailLive 取得籃球即時數據（120 分鐘內比賽）
func (c *Client) BasketballMatchDetailLive(ctx context.Context) (*BasketballMatchDetailLiveResponse, error) {
	return secretGet[BasketballMatchDetailLiveResponse](ctx, c, basketballMatchDetailLivePath)
}

// BasketballMatchLineupLiveResponse 籃球即時陣容回應型別
type BasketballMatchLineupLiveResponse = Response[[]BasketballMatchLineupLiveResponseData]

// BasketballMatchLineupLive 取得籃球即時陣容與球員數據（120 分鐘內比賽）
func (c *Client) BasketballMatchLineupLive(ctx context.Context) (*BasketballMatchLineupLiveResponse, error) {
	return secretGet[BasketballMatchLineupLiveResponse](ctx, c, basketballMatchLineupLivePath)
}

// BasketballMatchTrendLiveResponse 籃球即時走勢回應型別
type BasketballMatchTrendLiveResponse = Response[[]BasketballMatchTrendData]

// BasketballMatchTrendLive 取得籃球即時比賽走勢
func (c *Client) BasketballMatchTrendLive(ctx context.Context) (*BasketballMatchTrendLiveResponse, error) {
	return secretGet[BasketballMatchTrendLiveResponse](ctx, c, basketballMatchTrendLivePath)
}

// BasketballMatchShootPointLiveResponse 籃球即時投籃點回應型別
type BasketballMatchShootPointLiveResponse = Response[[]BasketballMatchShootPointData]

// BasketballMatchShootPointLive 取得籃球即時比賽投籃點
func (c *Client) BasketballMatchShootPointLive(ctx context.Context) (*BasketballMatchShootPointLiveResponse, error) {
	return secretGet[BasketballMatchShootPointLiveResponse](ctx, c, basketballMatchShootPointLivePath)
}

// BasketballMatchTrendDetailQuery 籃球比賽走勢詳情查詢參數
type BasketballMatchTrendDetailQuery struct {
	// 比賽 id
	UUID string `json:"uuid"`
}

// BasketballMatchTrendDetailResponse 籃球比賽走勢詳情回應型別
type BasketballMatchTrendDetailResponse = Response[BasketballMatchTrendData]

// BasketballMatchTrendDetail 取得籃球比賽走勢詳情
func (c *Client) BasketballMatchTrendDetail(ctx context.Context, query BasketballMatchTrendDetailQuery) (*BasketballMatchTrendDetailResponse, error) {
	return secretGet[BasketballMatchTrendDetailResponse](ctx, c, basketballMatchTrendDetailPath, toQuery(query))
}

// BasketballMatchShootPointQuery 籃球比賽投籃點查詢參數
type BasketballMatchShootPointQuery struct {
	// 比賽 id
	UUID string `json:"uuid"`
}

// BasketballMatchShootPointResponse 籃球比賽投籃點回應型別
type BasketballMatchShootPointResponse = Response[[]BasketballMatchShootPointData]

// BasketballMatchShootPoint 取得籃球比賽投籃點
func (c *Client) BasketballMatchShootPoint(ctx context.Context, query BasketballMatchShootPointQuery) (*BasketballMatchShootPointResponse, error) {
	return secretGet[BasketballMatchShootPointResponse](ctx, c, basketballMatchShootPointPath, toQuery(query))
}

// BasketballMatchAnalysisQuery 籃球比賽 H2H 分析查詢參數
type BasketballMatchAnalysisQuery struct {
	// 比賽 id
	UUID string `json:"uuid"`
}

// BasketballMatchAnalysisResponse 籃球比賽 H2H 分析回應型別
type BasketballMatchAnalysisResponse = Response[BasketballMatchAnalysisResponseData]

// BasketballMatchAnalysis 取得籃球比賽 H2H 分析
func (c *Client) BasketballMatchAnalysis(ctx context.Context, query BasketballMatchAnalysisQuery) (*BasketballMatchAnalysisResponse, error) {
	return secretGet[BasketballMatchAnalysisResponse](ctx, c, basketballMatchAnalysisPath, toQuery(query))
}

// BasketballMatchLiveHistoryQuery 籃球歷史比賽統計查詢參數
type BasketballMatchLiveHistoryQuery struct {
	// 比賽 id
	UUID string `json:"uuid"`
}

// BasketballMatchLiveHistoryResponse 籃球歷史比賽統計回應型別
type BasketballMatchLiveHistoryResponse = Response[BasketballMatchDetailLiveResponseData]

// BasketballMatchLiveHistory 取得籃球歷史比賽統計數據
func (c *Client) BasketballMatchLiveHistory(ctx context.Context, query BasketballMatchLiveHistoryQuery) (*BasketballMatchLiveHistoryResponse, error) {
	return secretGet[BasketballMatchLiveHistoryResponse](ctx, c, basketballMatchLiveHistoryPath, toQuery(query))
}

//	===  排名相關端點  ===

// BasketballSeasonRecentTableDetailQuery 籃球賽季排名查詢參數（最新賽季）
type BasketballSeasonRecentTableDetailQuery struct {
	// 賽季 id
	UUID string `json:"uuid"`
}

// BasketballSeasonRecentTableDetailResponse 籃球賽季排名回應型別（最新賽季）
type BasketballSeasonRecentTableDetailResponse = Response[[]BasketballSeasonTableResponseData]

// BasketballSeasonRecentTableDetail 取得籃球賽季排名（最新賽季）
func (c *Client) BasketballSeasonRecentTableDetail(ctx context.Context, query BasketballSeasonRecentTableDetailQuery) (*BasketballSeasonRecentTableDetailResponse, error) {
	return secretGet[BasketballSeasonRecentTableDetailResponse](ctx, c, basketballSeasonRecentTableDetailPath, toQuery(query))
}

// BasketballSeasonRecentTableDetailNewQuery 籃球賽季排名（含升降級）查詢參數
type BasketballSeasonRecentTableDetailNewQuery struct {
	// 賽季 id
	UUID string `json:"uuid"`
}

// BasketballSeasonRecentTableDetailNewResponse 籃球賽季排名（含升降級）回應型別
type BasketballSeasonRecentTableDetailNewResponse = Response[BasketballSeasonTableDetailNewResult]

// BasketballSeasonRecentTableDetailNew 取得籃球賽季排名含升降級（最新賽季）
func (c *Client) BasketballSeasonRecentTableDetailNew(ctx context.Context, query BasketballSeasonRecentTableDetailNewQuery) (*BasketballSeasonRecentTableDetailNewResponse, error) {
	return secretGet[BasketballSeasonRecentTableDetailNewResponse](ctx, c, basketballSeasonRecentTableDetailNewPath, toQuery(query))
}

//	===  其他端點  ===

// BasketballDataUpdateResponse 籃球數據更新回應型別
type BasketballDataUpdateResponse = Response[BasketballDataUpdateResult]

// BasketballDataUpdate 取得籃球數據更新（120 秒內變更）
func (c *Client) BasketballDataUpdate(ctx context.Context) (*BasketballDataUpdateResponse, error) {
	return secretGet[BasketballDataUpdateResponse](ctx, c, basketballDataUpdatePath)
}

// BasketballCompensationQuery 籃球歷史賠率補償查詢參數
type BasketballCompensationQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// BasketballCompensationResponse 籃球歷史賠率補償回應型別
type BasketballCompensationResponse = Response[[]BasketballCompensationResponseData]

// BasketballCompensation 取得籃球歷史賠率補償
func (c *Client) BasketballCompensation(ctx context.Context, query BasketballCompensationQuery) (*BasketballCompensationResponse, error) {
	return secretGet[BasketballCompensationResponse](ctx, c, basketballCompensationListPath, toQuery(query))
}

// BasketballDeletedResponse 籃球已刪除數據回應型別
type BasketballDeletedResponse = Response[BasketballDeletedResult]

// BasketballDeleted 取得籃球已刪除數據（24 小時內）
func (c *Client) BasketballDeleted(ctx context.Context) (*BasketballDeletedResponse, error) {
	return secretGet[BasketballDeletedResponse](ctx, c, basketballDeletedPath)
}

// BasketballOddsLiveResponse 籃球即時賠率回應型別
type BasketballOddsLiveResponse = Response[BasketballOddsLiveResult]

// BasketballOddsLive 取得籃球即時賠率
func (c *Client) BasketballOddsLive(ctx context.Context) (*BasketballOddsLiveResponse, error) {
	return secretGet[BasketballOddsLiveResponse](ctx, c, basketballOddsLivePath)
}

// BasketballOddsHistoryQuery 籃球單場賠率歷史查詢參數
type BasketballOddsHistoryQuery struct {
	// 比賽 id
	UUID string `json:"uuid"`
}

// BasketballOddsHistoryResponse 籃球單場賠率歷史回應型別
type BasketballOddsHistoryResponse = Response[BasketballOddsLiveResult]

// BasketballOddsHistory 取得籃球單場賠率歷史
func (c *Client) BasketballOddsHistory(ctx context.Context, query BasketballOddsHistoryQuery) (*BasketballOddsHistoryResponse, error) {
	return secretGet[BasketballOddsHistoryResponse](ctx, c, basketballOddsHistoryPath, toQuery(query))
}

// BasketballOddsUpdateResponseData 籃球賠率更新回應資料
type BasketballOddsUpdateResponseData struct {
	// 比賽 id
	ID string `json:"id"`
	// 賠率公司 id
	CompanyID int `json:"company_id"`
	// 更新修復時間
	UpdateTime int64 `json:"update_time"`
}

// BasketballOddsUpdateResponse 籃球賠率更新回應型別
type BasketballOddsUpdateResponse = Response[[]BasketballOddsUpdateResponseData]

// BasketballOddsUpdate 取得籃球賠率更新
func (c *Client) BasketballOddsUpdate(ctx context.Context) (*BasketballOddsUpdateResponse, error) {
	return secretGet[BasketballOddsUpdateResponse](ctx, c, basketballOddsUpdatePath)
}

// BasketballLanguageQuery 籃球多語言查詢參數
type BasketballLanguageQuery struct {
	Page *int    `json:"page,omitempty"`
	Time *int    `json:"time,omitempty"`
	UUID *string `json:"uuid,omitempty"`
	// 列表類型，1-分類，2-國家/地區，3-賽事，4-球隊（必填）
	Type int `json:"type"`
}

// BasketballLanguageResponse 籃球多語言回應型別
type BasketballLanguageResponse = Response[[]BasketballLanguageResponseData]

// BasketballLanguage 取得籃球多語言數據
func (c *Client) BasketballLanguage(ctx context.Context, query BasketballLanguageQuery) (*BasketballLanguageResponse, error) {
	return secretGet[BasketballLanguageResponse](ctx, c, basketballLanguageListPath, toQuery(query))
}
