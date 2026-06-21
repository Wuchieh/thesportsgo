package thesportsgo

import "encoding/json"

// BaseballCategoryResponseData 棒球分類資料
type BaseballCategoryResponseData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CountryID string `json:"country_id"`
	UpdatedAt int64  `json:"updated_at"`
}

// BaseballCountryResponseData 棒球國家/地區資料
type BaseballCountryResponseData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Logo      string `json:"logo"`
	UpdatedAt int64  `json:"updated_at"`
}

// BaseballUniqueTournamentResponseData 棒球賽事資料（unique_tournament）
type BaseballUniqueTournamentResponseData struct {
	ID           string `json:"id"`
	CategoryID   string `json:"category_id"`
	CountryID    string `json:"country_id"`
	Name         string `json:"name"`
	ShortName    string `json:"short_name"`
	Logo         string `json:"logo"`
	CurSeasonID  string `json:"cur_season_id"`
	UID          string `json:"uid"`
	UpdatedAt    int64  `json:"updated_at"`
}

// BaseballTeamResponseData 棒球球隊資料
type BaseballTeamResponseData struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	ShortName string   `json:"short_name"`
	Abbr      string   `json:"abbr"`
	Logo      string   `json:"logo"`
	Gender    int      `json:"gender"`
	National  int      `json:"national"`
	Virtual   int      `json:"virtual"`
	CountryID string   `json:"country_id"`
	UID       string   `json:"uid"`
	CoachIDs  []string `json:"coach_ids"`
	UpdatedAt int64    `json:"updated_at"`
}

// BaseballPlayerDrafted 棒球球員選秀資訊
type BaseballPlayerDrafted struct {
	Year   int    `json:"year"`
	Round  int    `json:"round"`
	Pick   int    `json:"pick"`
	TeamID string `json:"team_id"`
}

// BaseballPlayerPreferredHand 棒球球員慣用手
type BaseballPlayerPreferredHand struct {
	Batter  int `json:"batter"`
	Pitcher int `json:"pitcher"`
}

// BaseballPlayerSchool 棒球球員學校資訊
type BaseballPlayerSchool struct {
	CountryID string `json:"country_id"`
	Name      string `json:"name"`
}

// BaseballPlayerResponseData 棒球球員資料
type BaseballPlayerResponseData struct {
	ID               string                       `json:"id"`
	CountryID        string                       `json:"country_id"`
	Name             string                       `json:"name"`
	ShortName        string                       `json:"short_name"`
	Logo             string                       `json:"logo"`
	Birthday         int64                        `json:"birthday"`
	UID              string                       `json:"uid"`
	Type             int                          `json:"type"`
	Height           int                          `json:"height"`
	Weight           int                          `json:"weight"`
	Position         string                       `json:"position"`
	Drafted          *BaseballPlayerDrafted       `json:"drafted"`
	PreferredHand    *BaseballPlayerPreferredHand `json:"preferred_hand"`
	School           *BaseballPlayerSchool        `json:"school"`
	Deathday         int64                        `json:"deathday"`
	LeagueCareerAge  int                          `json:"league_career_age"`
	UpdatedAt        int64                        `json:"updated_at"`
}

// BaseballVenueResponseData 棒球場館資料
type BaseballVenueResponseData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Capacity  int    `json:"capacity"`
	CountryID string `json:"country_id"`
	UpdatedAt int64  `json:"updated_at"`
}

// BaseballSeasonResponseData 棒球賽季資料
type BaseballSeasonResponseData struct {
	ID                 string `json:"id"`
	UniqueTournamentID string `json:"unique_tournament_id"`
	Year               string `json:"year"`
	UpdatedAt          int64  `json:"updated_at"`
}

// BaseballCoverage 棒球比賽動畫覆蓋資訊
type BaseballCoverage struct {
	Mlive int `json:"mlive"`
}

// BaseballWeather 棒球比賽天氣資訊
type BaseballWeather struct {
	Desc      string `json:"desc"`
	Temp      int    `json:"temp"`
	Wind      int    `json:"wind"`
	Humidity  int    `json:"humidity"`
	Pressure  int    `json:"pressure"`
}

// BaseballMatchResponseData 棒球比賽列表資料（match/list, match/season 共用）
type BaseballMatchResponseData struct {
	ID                 string            `json:"id"`
	UniqueTournamentID string            `json:"unique_tournament_id"`
	SeasonID           string            `json:"season_id"`
	TournamentID       string            `json:"tournament_id"`
	HomeTeamID         string            `json:"home_team_id"`
	AwayTeamID         string            `json:"away_team_id"`
	StatusID           int               `json:"status_id"`
	MatchTime          int64             `json:"match_time"`
	VenueID            string            `json:"venue_id"`
	Neutral            int               `json:"neutral"`
	Coverage           *BaseballCoverage `json:"coverage"`
	Scores             json.RawMessage   `json:"scores"`
	Description        string            `json:"description"`
	TBD                int               `json:"tbd"`
	Weather            *BaseballWeather  `json:"weather"`
	UpdatedAt          int64             `json:"updated_at"`
}

// BaseballMatchDetailLiveResultData 棒球即時比賽詳情資料
type BaseballMatchDetailLiveResultData struct {
	ID      string          `json:"id"`
	Score   json.RawMessage `json:"score"`
	Stats   json.RawMessage `json:"stats"`
	Players json.RawMessage `json:"players"`
	Extra   json.RawMessage `json:"extra"`
}

// BaseballMatchLiveHistoryResultData 棒球歷史比賽統計資料
type BaseballMatchLiveHistoryResultData struct {
	ID      string          `json:"id"`
	Score   json.RawMessage `json:"score"`
	Stats   json.RawMessage `json:"stats"`
	Players json.RawMessage `json:"players"`
}

// BaseballMatchDiaryResultExtraTeam 棒球比賽日程額外球隊資料
type BaseballMatchDiaryResultExtraTeam struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

// BaseballMatchDiaryResultExtraTournament 棒球比賽日程額外賽事資料
type BaseballMatchDiaryResultExtraTournament struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

// BaseballMatchDiaryResultsExtra 棒球比賽日程額外關聯資料
type BaseballMatchDiaryResultsExtra struct {
	UniqueTournament []BaseballMatchDiaryResultExtraTournament `json:"unique_tournament"`
	Team             []BaseballMatchDiaryResultExtraTeam       `json:"team"`
}

// BaseballSeasonTableRow 棒球賽季排名隊伍列
type BaseballSeasonTableRow struct {
	TeamID       string  `json:"team_id"`
	Position     int     `json:"position"`
	Total        int     `json:"total"`
	Win          int     `json:"win"`
	Draw         int     `json:"draw"`
	Loss         int     `json:"loss"`
	Goals        int     `json:"goals"`
	GoalsAgainst int     `json:"goals_against"`
	WinRate      float64 `json:"win_rate"`
}

// BaseballSeasonTable 棒球賽季排名表
type BaseballSeasonTable struct {
	ID      string                   `json:"id"`
	Name    string                   `json:"name"`
	StageID string                   `json:"stage_id"`
	Rows    []BaseballSeasonTableRow `json:"rows"`
}

// BaseballSeasonTableResultData 棒球賽季排名回應資料
type BaseballSeasonTableResultData struct {
	Tables []BaseballSeasonTable `json:"tables"`
}

// BaseballSeasonTeamStatsResultData 棒球賽季球隊統計資料
type BaseballSeasonTeamStatsResultData struct {
	TeamID       string          `json:"team_id"`
	TournamentID string          `json:"tournament_id"`
	Stats        json.RawMessage `json:"stats"`
}

// BaseballSeasonPlayerStatsResultData 棒球賽季球員統計資料
type BaseballSeasonPlayerStatsResultData struct {
	PlayerID     string          `json:"player_id"`
	TeamID       string          `json:"team_id"`
	TournamentID string          `json:"tournament_id"`
	Stats        json.RawMessage `json:"stats"`
}

// BaseballSeasonCoachStatsResultData 棒球賽季教練統計資料
type BaseballSeasonCoachStatsResultData struct {
	CoachID          string `json:"coach_id"`
	TeamID           string `json:"team_id"`
	Type             int    `json:"type"`
	RegularMatches   int    `json:"regular_matches"`
	RegularWin       int    `json:"regular_win"`
	RegularDeuce     int    `json:"regular_deuce"`
	RegularLose      int    `json:"regular_lose"`
	PlayoffsMatches  int    `json:"playoffs_matches"`
	PlayoffsWin      int    `json:"playoffs_win"`
	PlayoffsLose     int    `json:"playoffs_lose"`
}

// BaseballTeamSquadPlayer 棒球球隊陣容球員
type BaseballTeamSquadPlayer struct {
	PlayerID    string `json:"player_id"`
	Position    string `json:"position"`
	ShirtNumber int    `json:"shirt_number"`
}

// BaseballTeamSquadResponseData 棒球球隊陣容資料
type BaseballTeamSquadResponseData struct {
	ID        string                    `json:"id"`
	Squad     []BaseballTeamSquadPlayer `json:"squad"`
	UpdatedAt int64                     `json:"updated_at"`
}

// BaseballTeamInjuryPlayer 棒球球隊傷病球員
type BaseballTeamInjuryPlayer struct {
	PlayerID  string `json:"player_id"`
	Type      int    `json:"type"`
	Status    int    `json:"status"`
	Reason    string `json:"reason"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

// BaseballTeamInjuryResponseData 棒球球隊傷病資料
type BaseballTeamInjuryResponseData struct {
	ID        string                      `json:"id"`
	Injury    []BaseballTeamInjuryPlayer  `json:"injury"`
	UpdatedAt int64                       `json:"updated_at"`
}

// BaseballHonorItem 棒球榮譽項目
type BaseballHonorItem struct {
	HonorID             string `json:"honor_id"`
	Season              string `json:"season"`
	SeasonID            string `json:"season_id"`
	UniqueTournamentID  string `json:"unique_tournament_id"`
}

// BaseballPlayerHonorItem 棒球球員榮譽項目（含 team_id）
type BaseballPlayerHonorItem struct {
	BaseballHonorItem
	TeamID string `json:"team_id"`
}

// BaseballCoachHonorItem 棒球教練榮譽項目（含 team_id）
type BaseballCoachHonorItem struct {
	BaseballHonorItem
	TeamID string `json:"team_id"`
}

// BaseballTeamHonorResponseData 棒球球隊榮譽資料
type BaseballTeamHonorResponseData struct {
	ID        string              `json:"id"`
	Honor     []BaseballHonorItem `json:"honor"`
	UpdatedAt int64               `json:"updated_at"`
}

// BaseballPlayerHonorResponseData 棒球球員榮譽資料
type BaseballPlayerHonorResponseData struct {
	ID        string                  `json:"id"`
	Honor     []BaseballPlayerHonorItem `json:"honor"`
	UpdatedAt int64                   `json:"updated_at"`
}

// BaseballCoachHonorResponseData 棒球教練榮譽資料
type BaseballCoachHonorResponseData struct {
	ID        string                  `json:"id"`
	Honor     []BaseballCoachHonorItem `json:"honor"`
	UpdatedAt int64                   `json:"updated_at"`
}

// BaseballHonorResponseData 棒球榮譽列表資料
type BaseballHonorResponseData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Logo      string `json:"logo"`
	UpdatedAt int64  `json:"updated_at"`
}

// BaseballBracket 棒球對戰圖階段
type BaseballBracket struct {
	ID                 string `json:"id"`
	UniqueTournamentID string `json:"unique_tournament_id"`
	Name               string `json:"name"`
	Abbr               string `json:"abbr"`
	StartTime          int64  `json:"start_time"`
	EndTime            int64  `json:"end_time"`
}

// BaseballBracketGroup 棒球對戰圖分組
type BaseballBracketGroup struct {
	ID        string `json:"id"`
	BracketID string `json:"bracket_id"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	TypeID    int    `json:"type_id"`
	Number    int    `json:"number"`
}

// BaseballBracketRound 棒球對戰圖回合
type BaseballBracketRound struct {
	ID        string `json:"id"`
	BracketID string `json:"bracket_id"`
	GroupID   string `json:"group_id"`
	Name      string `json:"name"`
	Abbr      string `json:"abbr"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	Number    int    `json:"number"`
}

// BaseballBracketMatchUp 棒球對戰圖對戰組合
type BaseballBracketMatchUp struct {
	ID            string   `json:"id"`
	RoundID       string   `json:"round_id"`
	Number        int      `json:"number"`
	TypeID        int      `json:"type_id"`
	StateID       int      `json:"state_id"`
	HomeTeamID    string   `json:"home_team_id"`
	AwayTeamID    string   `json:"away_team_id"`
	WinnerTeamID  string   `json:"winner_team_id"`
	HomeScore     int      `json:"home_score"`
	AwayScore     int      `json:"away_score"`
	ParentIDs     []string `json:"parent_ids"`
	ChildrenIDs   []string `json:"children_ids"`
	MatchIDs      []string `json:"match_ids"`
	Note          string   `json:"note"`
}

// BaseballBracketSeasonResultData 棒球對戰圖回應資料
type BaseballBracketSeasonResultData struct {
	Brackets []BaseballBracket         `json:"brackets"`
	Groups   []BaseballBracketGroup    `json:"groups"`
	Rounds   []BaseballBracketRound    `json:"rounds"`
	MatchUps []BaseballBracketMatchUp  `json:"match_ups"`
}

// BaseballDataUpdateItem 棒球資料更新項目
type BaseballDataUpdateItem struct {
	SeasonID   string `json:"season_id"`
	UpdateTime int64  `json:"update_time"`
}

// BaseballOddsLiveEntry 棒球即時賠率單筆資料
// 格式：[match_id, odds_type, [change_time, home/over, draw/handicap, away/under, closed]]
type BaseballOddsLiveEntry json.RawMessage

// BaseballOddsHistoryEntry 棒球歷史賠率單筆資料
// 格式：[change_time, home, draw/handicap, away, closed] (asia/bs)
// 格式：[change_time, home, draw, away, closed] (eu)
type BaseballOddsHistoryEntry json.RawMessage

// BaseballOddsHistoryCompany 棒球歷史賠率公司資料
type BaseballOddsHistoryCompany struct {
	Asia []BaseballOddsHistoryEntry `json:"asia"`
	Eu   []BaseballOddsHistoryEntry `json:"eu"`
	Bs   []BaseballOddsHistoryEntry `json:"bs"`
}
