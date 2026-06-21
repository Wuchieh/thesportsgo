package thesportsgo

import "encoding/json"

// EsportsListWrapper 電競列表端點通用包裝結構
type EsportsListWrapper[T any] struct {
	Total int  `json:"total"`
	Data  *[]T `json:"data"`
}

// ——— 基本資訊 ———

// EsportsCountryResponseData 電競國家/地區回應資料
type EsportsCountryResponseData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Logo        string `json:"logo"`
	CountryName string `json:"country_name"`
}

// EsportsRegionResponseData 電競地區回應資料
type EsportsRegionResponseData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// EsportsStageResponseData 電競階段回應資料
type EsportsStageResponseData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type int    `json:"type"`
}

// EsportsMapResponseData 電競地圖回應資料
type EsportsMapResponseData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ——— 賽事 ———

// EsportsTournamentResponseData 電競賽事回應資料
type EsportsTournamentResponseData struct {
	ID             string `json:"id"`
	CategoryID     string `json:"category_id"`
	CountryID      string `json:"country_id"`
	Name           string `json:"name"`
	ShortName      string `json:"short_name"`
	Logo           string `json:"logo"`
	Type           int    `json:"type"`
	Gender         int    `json:"gender"`
	Status         int    `json:"status"`
	PrimaryColor   string `json:"primary_color"`
	SecondaryColor string `json:"secondary_color"`
	UpdatedAt      int64  `json:"updated_at"`
}

// ——— 戰隊 ———

// EsportsTeamResponseData 電競戰隊回應資料
type EsportsTeamResponseData struct {
	ID        string `json:"id"`
	CountryID string `json:"country_id"`
	Name      string `json:"name"`
	Logo      string `json:"logo"`
	RegionID  string `json:"region_id"`
	UpdatedAt int64  `json:"updated_at"`
}

// ——— 選手 ———

// EsportsPlayerResponseData 電競選手回應資料
type EsportsPlayerResponseData struct {
	ID         string `json:"id"`
	TeamID     string `json:"team_id"`
	Name       string `json:"name"`
	CommonName string `json:"common_name"`
	Logo       string `json:"logo"`
	CountryID  string `json:"country_id"`
	Position   string `json:"position"`
	UpdatedAt  int64  `json:"updated_at"`
}

// ——— 比賽 ———

// EsportsMatchResponseData 電競比賽回應資料
type EsportsMatchResponseData struct {
	ID           string `json:"id"`
	TournamentID string `json:"tournament_id"`
	HomeTeamID   string `json:"home_team_id"`
	AwayTeamID   string `json:"away_team_id"`
	Status       int    `json:"status"`
	Scheduled    int64  `json:"scheduled"`
	MapID        string `json:"map_id"`
	UpdatedAt    int64  `json:"updated_at"`
}

// EsportsMatchDetailResponseData 電競比賽詳情回應資料（單一物件，非列表）
type EsportsMatchDetailResponseData struct {
	ID           string          `json:"id"`
	TournamentID string          `json:"tournament_id"`
	HomeTeamID   string          `json:"home_team_id"`
	AwayTeamID   string          `json:"away_team_id"`
	HomeScore    int             `json:"home_score"`
	AwayScore    int             `json:"away_score"`
	Status       int             `json:"status"`
	Scheduled    int64           `json:"scheduled"`
	Stages       json.RawMessage `json:"stages"`
	Events       json.RawMessage `json:"events"`
	Stats        json.RawMessage `json:"stats"`
	UpdatedAt    int64           `json:"updated_at"`
}

// EsportsMatchLineupPlayer 電競比賽陣容選手資料
type EsportsMatchLineupPlayer struct {
	PlayerID string `json:"player_id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Order    int    `json:"order"`
}

// EsportsMatchLineupSide 電競比賽陣容單邊資料
type EsportsMatchLineupSide struct {
	Starting    []EsportsMatchLineupPlayer `json:"starting"`
	Substitutes []EsportsMatchLineupPlayer `json:"substitutes"`
}

// EsportsMatchLineupResponseData 電競比賽陣容回應資料
type EsportsMatchLineupResponseData struct {
	Home EsportsMatchLineupSide `json:"home"`
	Away EsportsMatchLineupSide `json:"away"`
}

// ——— 積分榜 ———

// EsportsStandingsResponseData 電競積分榜回應資料
type EsportsStandingsResponseData struct {
	ID             string `json:"id"`
	TournamentID   string `json:"tournament_id"`
	TeamID         string `json:"team_id"`
	TeamName       string `json:"team_name"`
	Position       int    `json:"position"`
	Played         int    `json:"played"`
	Won            int    `json:"won"`
	Drawn          int    `json:"drawn"`
	Lost           int    `json:"lost"`
	GoalsFor       int    `json:"goals_for"`
	GoalsAgainst   int    `json:"goals_against"`
	GoalDifference int    `json:"goal_difference"`
	Points         int    `json:"points"`
	UpdatedAt      int64  `json:"updated_at"`
}
