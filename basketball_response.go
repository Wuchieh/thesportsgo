package thesportsgo

import "encoding/json"

//	===  基本清單端點結構  ===

// BasketballCategoryResponseData 籃球分類回應資料
type BasketballCategoryResponseData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	UpdatedAt int    `json:"updated_at"`
}

// BasketballCountryResponseData 籃球國家/地區回應資料
type BasketballCountryResponseData struct {
	CategoryID string `json:"category_id"`
	ID         string `json:"id"`
	Logo       string `json:"logo"`
	Name       string `json:"name"`
	UpdatedAt  int64  `json:"updated_at"`
}

// BasketballCompetitionResponseData 籃球賽事回應資料
type BasketballCompetitionResponseData struct {
	ID         string `json:"id"`
	CategoryID string `json:"category_id"`
	CountryID  string `json:"country_id"`
	Name       string `json:"name"`
	ShortName  string `json:"short_name"`
	Logo       string `json:"logo"`
	Type       int    `json:"type"`
	// 當前賽季
	CurSeasonID string `json:"cur_season_id"`
	// 當前階段
	CurStageID string `json:"cur_stage_id"`
	// 當前輪次
	CurRound   int `json:"cur_round"`
	RoundCount int `json:"round_count"`
	// 賽事主辦國／城市
	Host      *FootballCompetitionHost `json:"host"`
	Gender    int                      `json:"gender"`
	UID       string                   `json:"uid"`
	UpdatedAt int64                    `json:"updated_at"`
}

// BasketballTeamResponseData 籃球球隊回應資料
type BasketballTeamResponseData struct {
	ID            string `json:"id"`
	CompetitionID string `json:"competition_id"`
	CountryID     string `json:"country_id"`
	Name          string `json:"name"`
	ShortName     string `json:"short_name"`
	Logo          string `json:"logo"`
	// 主場場館 ID
	VenueID string `json:"venue_id"`
	// 教練 ID
	CoachID   string `json:"coach_id"`
	UID       string `json:"uid"`
	UpdatedAt int64  `json:"updated_at"`
}

// BasketballPlayerResponseData 籃球球員回應資料
type BasketballPlayerResponseData struct {
	ID        string `json:"id"`
	TeamID    string `json:"team_id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Logo      string `json:"logo"`
	// 年齡
	Age int `json:"age"`
	// 生日
	Birthday int64 `json:"birthday"`
	// 體重 (kg)
	Weight int `json:"weight"`
	// 身高 (cm)
	Height    int    `json:"height"`
	CountryID string `json:"country_id"`
	// 國籍
	Nationality string `json:"nationality"`
	// 場上位置 (PG/SG/SF/PF/C)
	Position string `json:"position"`
	// 球衣號碼
	JerseyNumber int `json:"jersey_number"`
	// 能力值
	Ability json.RawMessage `json:"ability"`
	UID     string          `json:"uid"`
	// 更新時間
	UpdatedAt int64 `json:"updated_at"`
}

// BasketballVenueResponseData 籃球場館回應資料
type BasketballVenueResponseData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CountryID string `json:"country_id"`
	// 實際名稱（當地語言）
	RealName string `json:"real_name"`
	// 容量
	Capacity  int    `json:"capacity"`
	UID       string `json:"uid"`
	UpdatedAt int64  `json:"updated_at"`
}

// BasketballSeasonResponseData 籃球賽季回應資料
type BasketballSeasonResponseData struct {
	ID            string `json:"id"`
	CompetitionID string `json:"competition_id"`
	Year          string `json:"year"`
	// 是否有球員統計
	HasPlayerStats int `json:"has_player_stats"`
	// 是否有球隊統計
	HasTeamStats int `json:"has_team_stats"`
	// 是否有積分榜
	HasTable  int `json:"has_table"`
	IsCurrent int `json:"is_current"`
	// 賽季開始時間
	StartTime int64 `json:"start_time"`
	// 賽季結束時間
	EndTime   int64 `json:"end_time"`
	UpdatedAt int64 `json:"updated_at"`
}

// BasketballStageResponseData 籃球階段回應資料
type BasketballStageResponseData struct {
	ID        string `json:"id"`
	SeasonID  string `json:"season_id"`
	Name      string `json:"name"`
	UpdatedAt int64  `json:"updated_at"`
}

//	===  比賽相關端點結構  ===

// BasketballMatchResponseData 籃球比賽回應資料
type BasketballMatchResponseData struct {
	ID     string `json:"id"`
	HomeID string `json:"home_id"`
	AwayID string `json:"away_id"`
	// 比賽狀態
	Status int `json:"status"`
	// 比賽日期
	Date int64 `json:"date"`
	// 主場分數
	HomeScore int `json:"home_score"`
	// 客場分數
	AwayScore int `json:"away_score"`
	// 第一節比分（主場）
	HomeQ1 int `json:"home_q1"`
	// 第一節比分（客場）
	AwayQ1 int `json:"away_q1"`
	// 第二節比分（主場）
	HomeQ2 int `json:"home_q2"`
	// 第二節比分（客場）
	AwayQ2 int `json:"away_q2"`
	// 第三節比分（主場）
	HomeQ3 int `json:"home_q3"`
	// 第三節比分（客場）
	AwayQ3 int `json:"away_q3"`
	// 第四節比分（主場）
	HomeQ4 int `json:"home_q4"`
	// 第四節比分（客場）
	AwayQ4 int `json:"away_q4"`
	// 延長賽比分（主場），無延長賽為 null
	HomeOT *int `json:"home_ot,omitempty"`
	// 延長賽比分（客場），無延長賽為 null
	AwayOT   *int      `json:"away_ot,omitempty"`
	UID      string    `json:"uid"`
	HomeTeam *TeamInfo `json:"home_team"`
	AwayTeam *TeamInfo `json:"away_team"`
	// 更新時間
	UpdatedAt int64 `json:"updated_at"`
}

// TeamInfo 球隊基本資訊（內嵌於比賽回應）
type TeamInfo struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Logo      string `json:"logo"`
}

// BasketballMatchDetailResponseData 籃球比賽詳情回應資料
type BasketballMatchDetailResponseData struct {
	ID        string `json:"id"`
	HomeID    string `json:"home_id"`
	AwayID    string `json:"away_id"`
	Date      int64  `json:"date"`
	HomeScore int    `json:"home_score"`
	AwayScore int    `json:"away_score"`
	// 各節比分
	HomeQ1 int  `json:"home_q1"`
	AwayQ1 int  `json:"away_q1"`
	HomeQ2 int  `json:"home_q2"`
	AwayQ2 int  `json:"away_q2"`
	HomeQ3 int  `json:"home_q3"`
	AwayQ3 int  `json:"away_q3"`
	HomeQ4 int  `json:"home_q4"`
	AwayQ4 int  `json:"away_q4"`
	HomeOT *int `json:"home_ot,omitempty"`
	AwayOT *int `json:"away_ot,omitempty"`
	Status int  `json:"status"`
	// 比賽場次（例：G1）
	Round         int    `json:"round"`
	SeasonID      string `json:"season_id"`
	CompetitionID string `json:"competition_id"`
	VenueID       string `json:"venue_id"`
	RefereeID     string `json:"referee_id"`
	UID           string `json:"uid"`
	// 比賽流程文字直播
	PlayByPlay json.RawMessage `json:"play_by_play"`
	// 走勢資料
	Trends    json.RawMessage `json:"trends"`
	HomeTeam  *TeamInfo       `json:"home_team"`
	AwayTeam  *TeamInfo       `json:"away_team"`
	UpdatedAt int64           `json:"updated_at"`
}

// BasketballMatchLineupResponseData 籃球比賽陣容回應資料
type BasketballMatchLineupResponseData struct {
	// 主場先發陣容
	HomeLineup []LineupPlayer `json:"home_lineup"`
	// 客場先發陣容
	AwayLineup []LineupPlayer `json:"away_lineup"`
	// 主場替補陣容
	HomeBench []LineupPlayer `json:"home_bench"`
	// 客場替補陣容
	AwayBench []LineupPlayer `json:"away_bench"`
}

// LineupPlayer 上場球員資訊
type LineupPlayer struct {
	PlayerID string `json:"player_id"`
	Name     string `json:"name"`
	// 球衣號碼
	JerseyNumber int `json:"jersey_number"`
	// 場上位置
	Position string `json:"position"`
	// 是否先發
	IsStarter int `json:"is_starter"`
}

// BasketballMatchStatisticsResponseData 籃球比賽統計回應資料
type BasketballMatchStatisticsResponseData struct {
	MatchID string        `json:"match_id"`
	Home    *TeamBoxScore `json:"home"`
	Away    *TeamBoxScore `json:"away"`
}

// TeamBoxScore 球隊 box score（比賽統計）
type TeamBoxScore struct {
	// 各節得分 (Q1, Q2, Q3, Q4, OT)
	Periods []PeriodScore `json:"periods"`
	// 球員個人統計
	Players []PlayerBoxScore `json:"players"`
	// 球隊總計
	Totals TeamTotals `json:"totals"`
}

// PeriodScore 單節得分
type PeriodScore struct {
	Period int `json:"period"`
	Score  int `json:"score"`
}

// PlayerBoxScore 球員個人 box score
type PlayerBoxScore struct {
	PlayerID string `json:"player_id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	// 上場時間（分鐘）
	Minutes int `json:"minutes"`
	// 得分
	Points int `json:"points"`
	// 二分球命中／出手
	TwoPointMade    int `json:"two_point_made"`
	TwoPointAttempt int `json:"two_point_attempt"`
	TwoPointPercent int `json:"two_point_percent"`
	// 三分球命中／出手
	ThreePointMade    int `json:"three_point_made"`
	ThreePointAttempt int `json:"three_point_attempt"`
	ThreePointPercent int `json:"three_point_percent"`
	// 罰球命中／出手
	FreeThrowMade    int `json:"free_throw_made"`
	FreeThrowAttempt int `json:"free_throw_attempt"`
	FreeThrowPercent int `json:"free_throw_percent"`
	// 籃板
	OffensiveRebound int `json:"offensive_rebound"`
	DefensiveRebound int `json:"defensive_rebound"`
	TotalRebound     int `json:"total_rebound"`
	// 助攻
	Assists int `json:"assists"`
	// 抄截
	Steals int `json:"steals"`
	// 阻攻
	Blocks int `json:"blocks"`
	// 失誤
	Turnovers int `json:"turnovers"`
	// 個人犯規
	PersonalFouls int `json:"personal_fouls"`
	// 正負值
	PlusMinus int `json:"plus_minus"`
}

// TeamTotals 球隊總計統計
type TeamTotals struct {
	Points            int `json:"points"`
	TwoPointMade      int `json:"two_point_made"`
	TwoPointAttempt   int `json:"two_point_attempt"`
	TwoPointPercent   int `json:"two_point_percent"`
	ThreePointMade    int `json:"three_point_made"`
	ThreePointAttempt int `json:"three_point_attempt"`
	ThreePointPercent int `json:"three_point_percent"`
	FreeThrowMade     int `json:"free_throw_made"`
	FreeThrowAttempt  int `json:"free_throw_attempt"`
	FreeThrowPercent  int `json:"free_throw_percent"`
	OffensiveRebound  int `json:"offensive_rebound"`
	DefensiveRebound  int `json:"defensive_rebound"`
	TotalRebound      int `json:"total_rebound"`
	Assists           int `json:"assists"`
	Steals            int `json:"steals"`
	Blocks            int `json:"blocks"`
	Turnovers         int `json:"turnovers"`
	PersonalFouls     int `json:"personal_fouls"`
}

//	===  排名相關端點結構  ===

// BasketballStandingsResponseData 籃球排名回應資料
type BasketballStandingsResponseData struct {
	ID       string `json:"id"`
	SeasonID string `json:"season_id"`
	TeamID   string `json:"team_id"`
	TeamName string `json:"team_name"`
	TeamLogo string `json:"team_logo"`
	// 排名
	Position int `json:"position"`
	// 已賽場次
	Played int `json:"played"`
	// 勝場
	Won int `json:"won"`
	// 敗場
	Lost int `json:"lost"`
	// 得分
	PointsFor int `json:"points_for"`
	// 失分
	PointsAgainst int `json:"points_against"`
	// 主場勝場
	HomeWon int `json:"home_won"`
	// 主場敗場
	HomeLost int `json:"home_lost"`
	// 客場勝場
	AwayWon int `json:"away_won"`
	// 客場敗場
	AwayLost int `json:"away_lost"`
	// 勝率
	WinPercent string `json:"win_percent"`
	// 勝差
	GamesBehind string `json:"games_behind"`
	// 近期十場戰績
	LastTen string `json:"last_ten"`
	// 連勝/連敗
	Streak    string `json:"streak"`
	UID       string `json:"uid"`
	UpdatedAt int64  `json:"updated_at"`
}
