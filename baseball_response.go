package thesportsgo

import "encoding/json"

// BaseballCategoryResponseData 棒球分類資料
type BaseballCategoryResponseData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	UpdatedAt int    `json:"updated_at"`
}

// BaseballCountryResponseData 棒球國家/地區資料
type BaseballCountryResponseData struct {
	CategoryID string `json:"category_id"`
	ID         string `json:"id"`
	Logo       string `json:"logo"`
	Name       string `json:"name"`
	UpdatedAt  int64  `json:"updated_at"`
}

// BaseballCompetitionHost 棒球賽事主辦資訊
type BaseballCompetitionHost struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

// BaseballCompetitionResponseData 棒球賽事資料
type BaseballCompetitionResponseData struct {
	ID             string                   `json:"id"`
	CategoryID     string                   `json:"category_id"`
	CountryID      string                   `json:"country_id"`
	Name           string                   `json:"name"`
	ShortName      string                   `json:"short_name"`
	Logo           string                   `json:"logo"`
	Type           int                      `json:"type"`
	CurSeasonID    string                   `json:"cur_season_id"`
	CurStageID     string                   `json:"cur_stage_id"`
	CurRound       int                      `json:"cur_round"`
	RoundCount     int                      `json:"round_count"`
	TitleHolder    json.RawMessage          `json:"title_holder"`
	MostTitles     json.RawMessage          `json:"most_titles"`
	Newcomers      json.RawMessage          `json:"newcomers"`
	Divisions      json.RawMessage          `json:"divisions"`
	Host           *BaseballCompetitionHost `json:"host"`
	Gender         int                      `json:"gender"`
	PrimaryColor   string                   `json:"primary_color"`
	SecondaryColor string                   `json:"secondary_color"`
	UID            string                   `json:"uid"`
	UpdatedAt      int64                    `json:"updated_at"`
}

// BaseballTeamResponseData 棒球球隊資料
type BaseballTeamResponseData struct {
	ID            string `json:"id"`
	CompetitionID string `json:"competition_id"`
	CountryID     string `json:"country_id"`
	Name          string `json:"name"`
	ShortName     string `json:"short_name"`
	CoachID       string `json:"coach_id"`
	VenueID       string `json:"venue_id"`
	UID           string `json:"uid"`
	UpdatedAt     int64  `json:"updated_at"`
}

// BaseballPlayerResponseData 棒球球員資料
type BaseballPlayerResponseData struct {
	ID              string          `json:"id"`
	TeamID          string          `json:"team_id"`
	Name            string          `json:"name"`
	ShortName       string          `json:"short_name"`
	Logo            string          `json:"logo"`
	NationalLogo    string          `json:"national_logo"`
	Age             int             `json:"age"`
	Birthday        int64           `json:"birthday"`
	Weight          int             `json:"weight"`
	Height          int             `json:"height"`
	CountryID       string          `json:"country_id"`
	Nationality     string          `json:"nationality"`
	CoachID         string          `json:"coach_id"`
	JerseyNumber    int             `json:"jersey_number"`
	Position        string          `json:"position"`
	Positions       json.RawMessage `json:"positions"`
	Bats            int             `json:"bats"`
	Throws          int             `json:"throws"`
	Ability         json.RawMessage `json:"ability"`
	Characteristics json.RawMessage `json:"characteristics"`
	UID             string          `json:"uid"`
	UpdatedAt       int64           `json:"updated_at"`
}

// BaseballCoachResponseData 棒球教練資料
type BaseballCoachResponseData struct {
	ID        string `json:"id"`
	TeamID    string `json:"team_id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	CountryID string `json:"country_id"`
	UID       string `json:"uid"`
	UpdatedAt int64  `json:"updated_at"`
}

// BaseballRefereeResponseData 棒球裁判資料
type BaseballRefereeResponseData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Logo      string `json:"logo"`
	Birthday  int64  `json:"birthday"`
	CountryID string `json:"country_id"`
	UID       string `json:"uid"`
	UpdatedAt int64  `json:"updated_at"`
}

// BaseballVenueResponseData 棒球場館資料
type BaseballVenueResponseData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CountryID string `json:"country_id"`
	RealName  string `json:"real_name"`
	UID       string `json:"uid"`
	UpdatedAt int64  `json:"updated_at"`
}

// BaseballSeasonResponseData 棒球賽季資料
type BaseballSeasonResponseData struct {
	ID             string `json:"id"`
	CompetitionID  string `json:"competition_id"`
	Year           string `json:"year"`
	HasPlayerStats int    `json:"has_player_stats"`
	HasTeamStats   int    `json:"has_team_stats"`
	HasTable       int    `json:"has_table"`
	IsCurrent      int    `json:"is_current"`
	StartTime      int64  `json:"start_time"`
	EndTime        int64  `json:"end_time"`
	UpdatedAt      int64  `json:"updated_at"`
}

// BaseballStageResponseData 棒球階段資料
type BaseballStageResponseData struct {
	ID        string `json:"id"`
	SeasonID  string `json:"season_id"`
	Name      string `json:"name"`
	UpdatedAt int64  `json:"updated_at"`
}

// BaseballMatchResponseData 棒球比賽列表資料
type BaseballMatchResponseData struct {
	ID            string          `json:"id"`
	CompetitionID string          `json:"competition_id"`
	SeasonID      string          `json:"season_id"`
	StageID       string          `json:"stage_id"`
	Round         int             `json:"round"`
	HomeTeamID    string          `json:"home_team_id"`
	AwayTeamID    string          `json:"away_team_id"`
	HomeScore     json.RawMessage `json:"home_score"`
	AwayScore     json.RawMessage `json:"away_score"`
	MatchTime     int64           `json:"match_time"`
	StatusID      int             `json:"status_id"`
	VenueID       string          `json:"venue_id"`
	RefereeID     string          `json:"referee_id"`
	Neutral       int             `json:"neutral"`
	Note          string          `json:"note"`
	HomeScores    json.RawMessage `json:"home_scores"`
	AwayScores    json.RawMessage `json:"away_scores"`
	UpdatedAt     int64           `json:"updated_at"`
}

// BaseballMatchDetailResponseData 棒球比賽詳情資料
type BaseballMatchDetailResponseData struct {
	BaseballMatchResponseData
	Lineup   json.RawMessage `json:"lineup"`
	Stats    json.RawMessage `json:"stats"`
	Timeline json.RawMessage `json:"timeline"`
}

// BaseballStandingsResponseData 棒球排名資料
type BaseballStandingsResponseData struct {
	CompetitionID string `json:"competition_id"`
	SeasonID      string `json:"season_id"`
	StageID       string `json:"stage_id"`
	TeamID        string `json:"team_id"`
	Rank          int    `json:"rank"`
	Played        int    `json:"played"`
	Wins          int    `json:"wins"`
	Losses        int    `json:"losses"`
	WinPct        string `json:"win_pct"`
	HomeRecord    string `json:"home_record"`
	AwayRecord    string `json:"away_record"`
	Last10Record  string `json:"last10_record"`
	Streak        string `json:"streak"`
	RunsFor       int    `json:"runs_for"`
	RunsAgainst   int    `json:"runs_against"`
	RunDiff       int    `json:"run_diff"`
	GroupID       string `json:"group_id"`
	PromotionID   string `json:"promotion_id"`
	PromotionName string `json:"promotion_name"`
	Points        int    `json:"points"`
	UpdatedAt     int64  `json:"updated_at"`
}
