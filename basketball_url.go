package thesportsgo

//	===  基本資訊端點（無 query 參數）  ===

const (
	// 分類列表
	basketballCategoryPath = "/basketball/category/list"
	// 國家/地區列表
	basketballCountryPath = "/basketball/country/list"
)

//	===  基本資訊端點（有 query 參數：page/time/uuid）  ===

const (
	// 賽事列表
	basketballCompetitionPath = "/basketball/competition/list"
	// 球隊列表
	basketballTeamPath = "/basketball/team/list"
	// 球員列表
	basketballPlayerPath = "/basketball/player/list"
	// 教練列表
	basketballCoachPath = "/basketball/coach/list"
	// 場館列表
	basketballVenuePath = "/basketball/venue/list"
	// 賽季列表
	basketballSeasonPath = "/basketball/season/list"
	// 階段列表
	basketballStagePath = "/basketball/stage/list"
)

//	===  比賽相關端點  ===

const (
	// 近期比賽列表（30 天內）
	basketballMatchRecentPath = "/basketball/match/recent/list"
	// 賽程與結果 - 日期查詢
	basketballMatchDiaryPath = "/basketball/match/diary"
	// 賽程與結果 - 賽季查詢（最新賽季）
	basketballMatchSeasonRecentPath = "/basketball/match/season/recent"
	// 即時數據
	basketballMatchDetailLivePath = "/basketball/match/detail_live"
	// 即時陣容與球員數據
	basketballMatchLineupLivePath = "/basketball/match/lineup_live"
	// 即時比賽走勢
	basketballMatchTrendLivePath = "/basketball/match/trend/live"
	// 即時投籃點
	basketballMatchShootPointLivePath = "/basketball/match/shoot/point/live"
	// 比賽走勢詳情
	basketballMatchTrendDetailPath = "/basketball/match/trend/detail"
	// 比賽投籃點
	basketballMatchShootPointPath = "/basketball/match/shoot/point"
	// H2H 比賽分析
	basketballMatchAnalysisPath = "/basketball/match/analysis"
	// 歷史比賽統計數據
	basketballMatchLiveHistoryPath = "/basketball/match/live/history"
)

//	===  排名相關端點  ===

const (
	// 賽季排名（最新賽季）
	basketballSeasonRecentTableDetailPath = "/basketball/season/recent/table/detail"
	// 賽季排名含升降級（最新賽季）
	basketballSeasonRecentTableDetailNewPath = "/basketball/season/recent/table/detail/new"
)

//	===  其他端點  ===

const (
	// 數據更新（120 秒內變更）
	basketballDataUpdatePath = "/basketball/data/update"
	// 歷史賠率補償
	basketballCompensationListPath = "/basketball/compensation/list"
	// 已刪除數據
	basketballDeletedPath = "/basketball/deleted"
	// 即時賠率
	basketballOddsLivePath = "/basketball/odds/live"
	// 單場賠率歷史
	basketballOddsHistoryPath = "/basketball/odds/history"
	// 賠率更新
	basketballOddsUpdatePath = "/basketball/odds/update"
	// 多語言數據
	basketballLanguageListPath = "/basketball/language/list"
)
