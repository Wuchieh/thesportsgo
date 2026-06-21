package thesportsgo

import "encoding/json"

//	===  基本清單端點結構  ===

// BasketballCategoryResponseData 籃球分類回應資料
type BasketballCategoryResponseData struct {
	// 分類 id
	ID string `json:"id"`
	// 分類名稱
	Name string `json:"name"`
	// 更新時間
	UpdatedAt int `json:"updated_at"`
}

// BasketballCountryResponseData 籃球國家/地區回應資料
type BasketballCountryResponseData struct {
	// 分類 id
	CategoryID string `json:"category_id"`
	// 國家/地區 id
	ID string `json:"id"`
	// 國家/地區標誌
	Logo string `json:"logo"`
	// 國家/地區名稱
	Name string `json:"name"`
	// 更新時間
	UpdatedAt int64 `json:"updated_at"`
}

// BasketballCompetitionResponseData 籃球賽事回應資料
type BasketballCompetitionResponseData struct {
	// 賽事 id
	ID string `json:"id"`
	// 分類 id
	CategoryID string `json:"category_id"`
	// 國家/地區 id
	CountryID string `json:"country_id"`
	// 賽事名稱
	Name string `json:"name"`
	// 賽事簡稱
	ShortName string `json:"short_name"`
	// 賽事標誌
	Logo string `json:"logo"`
	// 賽事類型，0-未知，1-聯賽，2-盃賽
	Type int `json:"type"`
	// 性別 1-男，2-女
	Gender int `json:"gender"`
	// 當前賽季 id
	CurSeasonID string `json:"cur_season_id"`
	// 合併後賽事 id
	UID string `json:"uid,omitempty"`
	// 更新時間
	UpdatedAt int64 `json:"updated_at"`
}

// BasketballTeamResponseData 籃球球隊回應資料
type BasketballTeamResponseData struct {
	// 球隊 id
	ID string `json:"id"`
	// 賽事 id
	CompetitionID string `json:"competition_id"`
	// 國家/地區 id
	CountryID string `json:"country_id"`
	// 分區 id，1-大西洋區，2-中央區，3-東南區，4-太平洋區，5-西北區，6-西南區，7-A組，8-B組，9-C組，10-D組，0-無
	ConferenceID int `json:"conference_id,omitempty"`
	// 球隊名稱
	Name string `json:"name"`
	// 球隊簡稱
	ShortName string `json:"short_name"`
	// 球隊標誌
	Logo string `json:"logo"`
	// 是否為國家隊，1-是，0-否
	National int `json:"national,omitempty"`
	// 成立時間
	FoundationTime int64 `json:"foundation_time,omitempty"`
	// 教練 id
	CoachID string `json:"coach_id,omitempty"`
	// 場館 id
	VenueID string `json:"venue_id,omitempty"`
	// 是否為佔位球隊，1-是，0-否
	Virtual int `json:"virtual,omitempty"`
	// 性別 1-男，2-女
	Gender int `json:"gender,omitempty"`
	// 合併後球隊 id
	UID string `json:"uid,omitempty"`
	// 更新時間
	UpdatedAt int64 `json:"updated_at"`
}

// BasketballPlayerResponseData 籃球球員回應資料
type BasketballPlayerResponseData struct {
	// 球員 id
	ID string `json:"id"`
	// 球員名稱
	Name string `json:"name"`
	// 球員簡稱
	ShortName string `json:"short_name"`
	// 球員照片
	Logo string `json:"logo"`
	// 所屬球隊 id
	TeamID string `json:"team_id"`
	// 年齡
	Age int `json:"age,omitempty"`
	// 生日（時間戳）
	Birthday int64 `json:"birthday,omitempty"`
	// 體重 (kg)
	Weight int `json:"weight,omitempty"`
	// 身高 (cm)
	Height int `json:"height,omitempty"`
	// 選秀順位
	Drafted string `json:"drafted,omitempty"`
	// 聯盟生涯年資
	LeagueCareerAge int `json:"league_career_age,omitempty"`
	// 畢業學校
	School string `json:"school,omitempty"`
	// 出生地
	City string `json:"city,omitempty"`
	// 國家/地區 id
	CountryID string `json:"country_id"`
	// 年薪
	Salary int `json:"salary,omitempty"`
	// 球衣號碼
	ShirtNumber int `json:"shirt_number,omitempty"`
	// 場上位置，C-中鋒，SF-小前鋒，PF-大前鋒，SG-得分後衛，PG-控球後衛，F-前鋒，G-後衛
	Position string `json:"position,omitempty"`
	// 合併後球員 id
	UID string `json:"uid,omitempty"`
	// 逝世時間（不存在則無此欄位）
	Deathday int64 `json:"deathday,omitempty"`
	// 退役時間（不存在則無此欄位）
	RetireTime int64 `json:"retire_time,omitempty"`
	// 更新時間
	UpdatedAt int64 `json:"updated_at"`
}

// BasketballCoachResponseData 籃球教練回應資料
type BasketballCoachResponseData struct {
	// 教練 id
	ID string `json:"id"`
	// 執教球隊 id
	TeamID string `json:"team_id"`
	// 教練名稱
	Name string `json:"name"`
	// 教練簡稱
	ShortName string `json:"short_name"`
	// 教練照片
	Logo string `json:"logo"`
	// 類型 1-總教練，2-助理教練，3-總經理
	Type int `json:"type,omitempty"`
	// 生日（時間戳）
	Birthday int64 `json:"birthday,omitempty"`
	// 國家/地區 id
	CountryID string `json:"country_id"`
	// 合併後教練 id
	UID string `json:"uid,omitempty"`
	// 逝世時間（不存在則無此欄位）
	Deathday int64 `json:"deathday,omitempty"`
	// 更新時間
	UpdatedAt int64 `json:"updated_at"`
}

// BasketballVenueResponseData 籃球場館回應資料
type BasketballVenueResponseData struct {
	// 場館 id
	ID string `json:"id"`
	// 場館名稱
	Name string `json:"name"`
	// 場館容量
	Capacity int `json:"capacity"`
	// 所在城市
	City string `json:"city"`
	// 合併後場館 id
	UID string `json:"uid,omitempty"`
	// 更新時間
	UpdatedAt int64 `json:"updated_at"`
}

// BasketballSeasonResponseData 籃球賽季回應資料
type BasketballSeasonResponseData struct {
	// 賽季 id
	ID string `json:"id"`
	// 賽事 id
	CompetitionID string `json:"competition_id"`
	// 賽季年度
	Year string `json:"year"`
	// 是否有球員統計，1-是，0-否
	HasPlayerStats int `json:"has_player_stats"`
	// 是否有球隊統計，1-是，0-否
	HasTeamStats int `json:"has_team_stats"`
	// 是否有積分榜，1-是，0-否
	HasTable int `json:"has_table"`
	// 是否為最新賽季，1-是，0-否
	IsCurrent int `json:"is_current"`
	// 賽季開始時間
	StartTime int64 `json:"start_time"`
	// 賽季結束時間
	EndTime int64 `json:"end_time"`
	// 更新時間
	UpdatedAt int64 `json:"updated_at"`
}

// BasketballStageResponseData 籃球階段回應資料
type BasketballStageResponseData struct {
	// 階段 id
	ID string `json:"id"`
	// 賽季 id
	SeasonID string `json:"season_id"`
	// 階段名稱
	Name string `json:"name"`
	// 比賽模式，1-積分制，2-淘汰制
	Mode int `json:"mode"`
	// 總組數
	GroupCount int `json:"group_count"`
	// 階段排序
	Order int `json:"order"`
	// 更新時間
	UpdatedAt int64 `json:"updated_at"`
}

//	===  比賽相關端點結構  ===

// BasketballMatchResponseData 籃球比賽回應資料（match/recent/list、match/diary、match/season/recent）
type BasketballMatchResponseData struct {
	// 比賽 id
	ID string `json:"id"`
	// 賽事 id
	CompetitionID string `json:"competition_id"`
	// 主隊 id
	HomeTeamID string `json:"home_team_id"`
	// 客隊 id
	AwayTeamID string `json:"away_team_id"`
	// 類型 id，1-例行賽，2-季後賽，3-季前賽，4-全明星，5-盃賽，6-附加賽，0-無
	Kind int `json:"kind"`
	// 比賽狀態，參考 Status Code -> Match Status
	StatusID int `json:"status_id"`
	// 比賽時間
	MatchTime int64 `json:"match_time"`
	// 是否中立場地，1-是，0-否
	Neutral int `json:"neutral"`
	// 主隊各節得分 [Q1, Q2, Q3, Q4, OT]
	HomeScores []int `json:"home_scores"`
	// 客隊各節得分 [Q1, Q2, Q3, Q4, OT]
	AwayScores []int `json:"away_scores"`
	// 比賽總節數（不含延長賽）
	PeriodCount int `json:"period_count"`
	// 延長賽得分（僅 1 次以上延長賽存在），[[主隊OT1,OT2,OTn], [客隊OT1,OT2,OTn]]
	OverTimeScores [][]int `json:"over_time_scores,omitempty"`
	// 動畫/直播覆蓋
	Coverage *MatchCoverage `json:"coverage"`
	// 賽季 id（無數據則不存在）
	SeasonID string `json:"season_id,omitempty"`
	// 場館 id（無數據則不存在）
	VenueID string `json:"venue_id,omitempty"`
	// 階段資訊（無數據則不存在）
	Round *MatchRound `json:"round,omitempty"`
	// 排名資訊（無數據則不存在）
	Position *MatchPosition `json:"position,omitempty"`
	// 比賽時間是否待定，1-是（存在則返回）
	TBD int `json:"tbd,omitempty"`
	// 結束時間（存在則返回）
	Ended int64 `json:"ended,omitempty"`
	// 比賽直接判負，1-是（存在則返回）
	Loss int `json:"loss,omitempty"`
	// 更新時間
	UpdatedAt int64 `json:"updated_at"`
}

// MatchCoverage 比賽直播覆蓋資訊
type MatchCoverage struct {
	// 是否有動畫，1-是，0-否
	Mlive int `json:"mlive"`
}

// MatchRound 比賽階段資訊
type MatchRound struct {
	// 階段 id
	StageID string `json:"stage_id"`
	// 組別（1-A，2-B 等，無數據則不存在）
	GroupNum int `json:"group_num,omitempty"`
	// 輪次（無數據則不存在）
	RoundNum int `json:"round_num,omitempty"`
}

// MatchPosition 比賽排名資訊
type MatchPosition struct {
	// 主隊排名
	Home string `json:"home"`
	// 客隊排名
	Away string `json:"away"`
}

// BasketballMatchDetailLiveResponseData 籃球即時數據回應資料（match/detail_live）
type BasketballMatchDetailLiveResponseData struct {
	// 比賽 id
	ID string `json:"id"`
	// 比分 [比賽id, 比賽狀態, 剩餘秒數, [主Q1~OT], [客Q1~OT]]
	Score json.RawMessage `json:"score"`
	// 計時器 [是否走秒, 是否倒數, 更新時間, 剩餘秒數]
	Timer json.RawMessage `json:"timer"`
	// 延長賽得分
	OverTimeScores json.RawMessage `json:"over_time_scores,omitempty"`
	// 比賽統計 [[類型, 主隊值, 客隊值], ...]
	Stats json.RawMessage `json:"stats,omitempty"`
	// 比賽文字直播 [[描述, ...], ...]
	Tlive json.RawMessage `json:"tlive,omitempty"`
}

// BasketballMatchLineupLiveResponseData 籃球即時陣容回應資料（match/lineup_live）
type BasketballMatchLineupLiveResponseData struct {
	// 比賽 id
	ID string `json:"id"`
	// 球員陣容 [主隊, 客隊, 主隊總計, 客隊總計]
	Players json.RawMessage `json:"players"`
}

// BasketballMatchTrendData 籃球比賽走勢資料
type BasketballMatchTrendData struct {
	// 比賽 id
	MatchID string `json:"match_id"`
	// 走勢
	Trend TrendDetail `json:"trend"`
}

// TrendDetail 走勢詳情
type TrendDetail struct {
	// 條數
	Count int `json:"count"`
	// 每條持續時間
	Per int `json:"per"`
	// 走勢變化數據 [Q1數據, Q2數據, Q3數據, Q4數據]
	Data json.RawMessage `json:"data"`
}

// BasketballMatchShootPointData 籃球投籃點資料
type BasketballMatchShootPointData struct {
	// 比賽 id
	MatchID string `json:"match_id"`
	// 投籃數據 [[隊伍, 投籃球員id, 助攻球員id, 是否命中, 得分, x, y, 節數, 節內剩餘時間], ...]
	Shoot json.RawMessage `json:"shoot"`
}

// BasketballMatchAnalysisResponseData 籃球比賽 H2H 分析回應資料
type BasketballMatchAnalysisResponseData struct {
	// 比賽資訊
	Info json.RawMessage `json:"info"`
	// 歷史對戰/近期戰績
	History BasketballMatchAnalysisHistory `json:"history"`
	// 未來賽程
	Future BasketballMatchAnalysisFuture `json:"future"`
}

// BasketballMatchAnalysisHistory 歷史對戰/近期戰績
type BasketballMatchAnalysisHistory struct {
	// 歷史對戰
	VS json.RawMessage `json:"vs"`
	// 主隊近期戰績
	Home json.RawMessage `json:"home"`
	// 客隊近期戰績
	Away json.RawMessage `json:"away"`
}

// BasketballMatchAnalysisFuture 未來賽程
type BasketballMatchAnalysisFuture struct {
	// 主隊未來賽程
	Home json.RawMessage `json:"home"`
	// 客隊未來賽程
	Away json.RawMessage `json:"away"`
}

//	===  排名相關端點結構  ===

// BasketballSeasonTableDetailNewResult 賽季排名（含升降級）結果
type BasketballSeasonTableDetailNewResult struct {
	// 升降級列表
	Promotions []SeasonTablePromotion `json:"promotions"`
	// 排名表列表
	Tables []BasketballSeasonTableResponseData `json:"tables"`
}

// SeasonTablePromotion 升降級資訊
type SeasonTablePromotion struct {
	// 升降級 id
	ID string `json:"id"`
	// 升降級名稱
	Name string `json:"name"`
	// 顏色值
	Color string `json:"color"`
}

// BasketballSeasonTableRowData 賽季排名行資料
type BasketballSeasonTableRowData struct {
	// 球隊 id
	TeamID string `json:"team_id"`
	// 升降級 id（含升降級端點存在）
	PromotionID string `json:"promotion_id,omitempty"`
	// 排名
	Position int `json:"position"`
	// 勝場
	Won int `json:"won"`
	// 敗場
	Lost int `json:"lost"`
	// 勝率
	WonRate float64 `json:"won_rate"`
	// 勝差（可能不存在）
	GameBack string `json:"game_back,omitempty"`
	// 場均得分（可能不存在）
	PointsAvg float64 `json:"points_avg,omitempty"`
	// 場均失分（可能不存在）
	PointsAgainstAvg float64 `json:"points_against_avg,omitempty"`
	// 場均淨勝分（可能不存在）
	DiffAvg float64 `json:"diff_avg,omitempty"`
	// 近期戰績，正數為連勝、負數為連敗（可能不存在）
	Streaks int `json:"streaks,omitempty"`
	// 主場戰績：主場勝-主場敗（可能不存在）
	Home string `json:"home,omitempty"`
	// 客場戰績：客場勝-客場敗（可能不存在）
	Away string `json:"away,omitempty"`
	// 分區勝負（可能不存在）
	Division string `json:"division,omitempty"`
	// 東西區勝負（可能不存在）
	Conference string `json:"conference,omitempty"`
	// 近十場勝負（可能不存在）
	Last10 string `json:"last_10,omitempty"`
	// 積分（盃賽存在）
	Points int `json:"points,omitempty"`
	// 總得分（盃賽存在）
	PointsFor int `json:"points_for,omitempty"`
	// 總失分（盃賽存在）
	PointsAgt int `json:"points_agt,omitempty"`
	// 更新時間
	UpdatedAt int `json:"updated_at"`
}

// BasketballSeasonTableResponseData 賽季排名表回應資料
type BasketballSeasonTableResponseData struct {
	// 排名表 id
	ID string `json:"id"`
	// 統計範圍，1-賽季，2-資格賽，3-小組賽，4-季前賽，5-例行賽，6-淘汰賽（季後賽）
	Scope int `json:"scope"`
	// 排名表名稱
	Name string `json:"name"`
	// 階段 id
	StageID string `json:"stage_id"`
	// 球隊積分列表
	Rows []BasketballSeasonTableRowData `json:"rows"`
}

//	===  數據更新端點結構  ===

// BasketballDataUpdateResult 數據更新結果（key 為數據類型 id 字串）
type BasketballDataUpdateResult map[string][]BasketballDataUpdateItem

// BasketballDataUpdateItem 數據更新項目
type BasketballDataUpdateItem struct {
	// 賽季 id（bracket、season standing 等欄位存在）
	SeasonID string `json:"season_id,omitempty"`
	// 排名發布時間（fiba men、fiba women 欄位存在）
	PubTime int64 `json:"pub_time,omitempty"`
	// 更新時間
	UpdateTime int64 `json:"update_time"`
}

//	===  刪除數據端點結構  ===

// BasketballDeletedResult 已刪除數據結果
type BasketballDeletedResult struct {
	// 已刪除比賽 id 列表
	Match []string `json:"match,omitempty"`
	// 已刪除賽事 id 列表
	Competition []string `json:"competition,omitempty"`
	// 已刪除球隊 id 列表
	Team []string `json:"team,omitempty"`
	// 已刪除球員 id 列表
	Player []string `json:"player,omitempty"`
	// 已刪除賽季 id 列表
	Season []string `json:"season,omitempty"`
	// 已刪除階段 id 列表
	Stage []string `json:"stage,omitempty"`
}

//	===  賠率相關端點結構  ===

// BasketballOddsLiveResult 即時賠率結果（key 為公司 id，值為賠率陣列的原始 JSON）
type BasketballOddsLiveResult map[string]json.RawMessage

//	===  歷史補償端點結構  ===

// BasketballCompensationResponseData 歷史賠率補償回應資料
type BasketballCompensationResponseData struct {
	// 比賽 id
	ID string `json:"id"`
	// 歷史對戰
	History CompensationHistory `json:"history"`
	// 近期戰績
	Recent CompensationRecent `json:"recent"`
	// 相同賠率盤口
	Similar CompensationSimilar `json:"similar"`
	// 更新時間
	UpdatedAt int64 `json:"updated_at"`
}

// CompensationHistory 歷史對戰
type CompensationHistory struct {
	Home CompensationWinLoss `json:"home"`
	Away CompensationWinLoss `json:"away"`
}

// CompensationRecent 近期戰績
type CompensationRecent struct {
	Home CompensationWinLoss `json:"home"`
	Away CompensationWinLoss `json:"away"`
}

// CompensationWinLoss 勝負統計
type CompensationWinLoss struct {
	WonCount  int     `json:"won_count"`
	LostCount int     `json:"lost_count"`
	Rate      float64 `json:"rate"`
}

// CompensationSimilar 相同賠率盤口
type CompensationSimilar struct {
	Teams        []CompensationTeam        `json:"teams"`
	Competitions []CompensationCompetition `json:"competitions"`
	Companies    []CompensationCompany     `json:"companies"`
	Europe       []CompensationEurope      `json:"europe"`
	Analysis     []CompensationAnalysis    `json:"analysis"`
}

// CompensationTeam 相同賠率涉及球隊
type CompensationTeam struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CompensationCompetition 相同賠率涉及賽事
type CompensationCompetition struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CompensationCompany 相同賠率涉及公司
type CompensationCompany struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// CompensationEurope 歐賠相同賠率盤口
type CompensationEurope struct {
	// 賠率公司 id
	ID int `json:"id"`
	// 總場次
	Total int `json:"total"`
	// 勝場
	Won int `json:"won"`
	// 敗場
	Loss int `json:"loss"`
	// 初始賠率 [主勝, 平局, 客勝]
	Odds []float64 `json:"odds"`
	// 比賽列表
	Matches []CompensationMatch `json:"matches"`
}

// CompensationMatch 賠率相同比賽
type CompensationMatch struct {
	ID            string    `json:"id"`
	CompetitionID string    `json:"competition_id"`
	HomeTeamID    string    `json:"home_team_id"`
	AwayTeamID    string    `json:"away_team_id"`
	MatchTime     int64     `json:"match_time"`
	HomeScore     int       `json:"home_score"`
	AwayScore     int       `json:"away_score"`
	BeginOdds     []float64 `json:"begin_odds"`
	ImmediateOdds []float64 `json:"immediate_odds"`
}

// CompensationAnalysis 賠率勝負概率分析
type CompensationAnalysis struct {
	// 賠率公司 id
	ID int `json:"id"`
	// 勝率
	WinRate float64 `json:"win_rate"`
	// 敗率
	LossRate float64 `json:"loss_rate"`
}

//	===  多語言端點結構  ===

// BasketballLanguageResponseData 籃球多語言回應資料
type BasketballLanguageResponseData struct {
	// 對應數據 id
	ID string `json:"id"`
	// 更新時間
	UpdatedAt int `json:"updated_at"`
}
