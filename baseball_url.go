package thesportsgo

const (
	// 基礎資料
	baseballCategoryPath = "/baseball/category/list" // 分類
	baseballCountryPath  = "/baseball/country/list"  // 國家/地區

	// 競賽與球隊資料
	baseballUniqueTournamentPath = "/baseball/unique_tournament/list" // 賽事（原 competition/additional/list）
	baseballTeamPath             = "/baseball/team/list"              // 球隊（原 team/additional/list）
	baseballPlayerPath           = "/baseball/player/list"            // 球員（原 player/with_stat/list）
	baseballVenuePath            = "/baseball/venue/list"             // 場館
	baseballSeasonPath           = "/baseball/season/list"            // 賽季

	// 比賽資料
	baseballMatchListPath       = "/baseball/match/list"        // 比賽列表
	baseballMatchDetailLivePath = "/baseball/match/detail_live" // 即時比賽詳情（原 match/detail）
	baseballMatchDiaryPath      = "/baseball/match/diary"       // 賽程與結果-日期查詢（新增）
	baseballMatchSeasonPath     = "/baseball/match/season"      // 賽程與結果-賽季查詢（新增）
	baseballMatchLiveHistoryPath = "/baseball/match/live/history" // 歷史比賽統計（新增）

	// 排名與統計
	baseballSeasonTablePath       = "/baseball/season/table/detail"       // 賽季排名（原 standings/list）
	baseballSeasonTeamStatsPath   = "/baseball/season/team/stats/detail"  // 賽季球隊統計（新增）
	baseballSeasonPlayerStatsPath = "/baseball/season/player/stats/detail" // 賽季球員統計（新增）
	baseballSeasonCoachStatsPath  = "/baseball/season/coach/stats/detail"  // 賽季教練統計（新增）

	// 榮譽與陣容
	baseballHonorPath       = "/baseball/honor/list"        // 榮譽列表（新增）
	baseballTeamHonorPath   = "/baseball/team/honor/list"   // 球隊榮譽（新增）
	baseballPlayerHonorPath = "/baseball/player/honor/list" // 球員榮譽（新增）
	baseballCoachHonorPath  = "/baseball/coach/honor/list"  // 教練榮譽（新增）
	baseballTeamSquadPath   = "/baseball/team/squad/list"   // 球隊陣容（新增）
	baseballTeamInjuryPath  = "/baseball/team/injury/list"  // 球隊傷病（新增）

	// 其他
	baseballBracketSeasonPath = "/baseball/bracket/season" // 對戰圖（新增）
	baseballDataUpdatePath    = "/baseball/data/update"    // 資料更新（新增）
	baseballDeletedPath       = "/baseball/deleted"        // 刪除資料（新增）

	// 賠率
	baseballOddsLivePath    = "/baseball/odds/live"    // 即時賠率（新增）
	baseballOddsHistoryPath = "/baseball/odds/history" // 單場比賽賠率（新增）
)
