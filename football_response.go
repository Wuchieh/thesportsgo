package thesportsgo

import "encoding/json"

type FootballCategoryResponseData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	UpdatedAt int    `json:"updated_at"`
}

type FootballCountryResponseData struct {
	CategoryID string `json:"category_id"`
	ID         string `json:"id"`
	Logo       string `json:"logo"`
	Name       string `json:"name"`
	UpdatedAt  int64  `json:"updated_at"`
}

type FootballCompetitionHost struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

type FootballCompetitionResponseData struct {
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
	Host           *FootballCompetitionHost `json:"host"`
	Gender         int                      `json:"gender"`
	PrimaryColor   string                   `json:"primary_color"`
	SecondaryColor string                   `json:"secondary_color"`
	UID            string                   `json:"uid"`
	UpdatedAt      int64                    `json:"updated_at"`
}

type FootballTeamResponseData struct {
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

type FootballPlayerResponseData struct {
	ID                  string          `json:"id"`
	TeamID              string          `json:"team_id"`
	Name                string          `json:"name"`
	ShortName           string          `json:"short_name"`
	Logo                string          `json:"logo"`
	NationalLogo        string          `json:"national_logo"`
	Age                 int             `json:"age"`
	Birthday            int64           `json:"birthday"`
	Weight              int             `json:"weight"`
	Height              int             `json:"height"`
	CountryID           string          `json:"country_id"`
	Nationality         string          `json:"nationality"`
	CoachID             string          `json:"coach_id"`
	MarketValue         int             `json:"market_value"`
	MarketValueCurrency string          `json:"market_value_currency"`
	ContractUntil       int64           `json:"contract_until"`
	PreferredFoot       int             `json:"preferred_foot"`
	Ability             json.RawMessage `json:"ability"`
	Characteristics     json.RawMessage `json:"characteristics"`
	Position            string          `json:"position"`
	Positions           json.RawMessage `json:"positions"`
	UID                 string          `json:"uid"`
	UpdatedAt           int64           `json:"updated_at"`
}

type FootballCoachResponseData struct {
	ID        string `json:"id"`
	TeamID    string `json:"team_id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	CountryID string `json:"country_id"`
	UID       string `json:"uid"`
	UpdatedAt int64  `json:"updated_at"`
}

type FootballRefereeResponseData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Logo      string `json:"logo"`
	Birthday  int64  `json:"birthday"`
	CountryID string `json:"country_id"`
	UID       string `json:"uid"`
	UpdatedAt int64  `json:"updated_at"`
}

type FootballVenueResponseData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CountryID string `json:"country_id"`
	RealName  string `json:"real_name"`
	UID       string `json:"uid"`
	UpdatedAt int64  `json:"updated_at"`
}

type FootballSeasonResponseData struct {
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

type FootballStageResponseData struct {
	ID        string `json:"id"`
	SeasonID  string `json:"season_id"`
	Name      string `json:"name"`
	UpdatedAt int64  `json:"updated_at"`
}
