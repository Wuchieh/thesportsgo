package thesportsgo

import (
	"encoding/json"
	"strconv"
	"strings"
)

// ——— CS:GO 基本資料 ———

// CsgoCountryResponseData CS:GO 國家/地區回應資料
type CsgoCountryResponseData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Logo      string `json:"logo"`
	UpdatedAt int64  `json:"updated_at"`
}

// CsgoTournamentResponseData CS:GO 賽事回應資料
type CsgoTournamentResponseData struct {
	ID        string   `json:"id"`
	Type      int      `json:"type"`
	Name      string   `json:"name"`
	Abbr      string   `json:"abbr"`
	Logo      string   `json:"logo"`
	Cover     string   `json:"cover"`
	Address   string   `json:"address"`
	StatusID  int      `json:"status_id"`
	StartTime int64    `json:"start_time"`
	EndTime   int64    `json:"end_time"`
	PrizePool string   `json:"prize_pool"`
	MapPool   []string `json:"map_pool"`
	UpdatedAt int64    `json:"updated_at"`
}

// CsgoTeamResponseData CS:GO 戰隊回應資料
type CsgoTeamResponseData struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Abbr       string `json:"abbr"`
	Logo       string `json:"logo"`
	CountryID  string `json:"country_id"`
	CreateTime int64  `json:"create_time"`
	UpdatedAt  int64  `json:"updated_at"`
}

// CsgoPlayerResponseData CS:GO 選手回應資料
type CsgoPlayerResponseData struct {
	ID        string `json:"id"`
	CountryID string `json:"country_id"`
	TeamID    string `json:"team_id"`
	Name      string `json:"name"`
	RealName  string `json:"real_name"`
	Logo      string `json:"logo"`
	Birthday  int64  `json:"birthday"`
	UpdatedAt int64  `json:"updated_at"`
}

// CsgoStageResponseData CS:GO 階段回應資料
type CsgoStageResponseData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	UpdatedAt int64  `json:"updated_at"`
}

// ——— CS:GO 比賽 ———

// CsgoCoverageData CS:GO 比賽動畫覆蓋資料
type CsgoCoverageData struct {
	Mlive int `json:"mlive"`
}

// CsgoBpData CS:GO 地圖 bp 資料
type CsgoBpData struct {
	Type   int    `json:"type"`
	MapID  string `json:"map_id"`
	TeamID string `json:"team_id"`
}

// CsgoMatchResponseData CS:GO 比賽回應資料
type CsgoMatchResponseData struct {
	ID           string           `json:"id"`
	TournamentID string           `json:"tournament_id"`
	StageID      string           `json:"stage_id"`
	HomeTeamID   string           `json:"home_team_id"`
	AwayTeamID   string           `json:"away_team_id"`
	Box          int              `json:"box"`
	Round        int              `json:"round"`
	HomeScore    int              `json:"home_score"`
	AwayScore    int              `json:"away_score"`
	StatusID     int              `json:"status_id"`
	MatchTime    int64            `json:"match_time"`
	Coverage     CsgoCoverageData `json:"coverage"`
	BpData       []CsgoBpData     `json:"bp_data"`
	Description  string           `json:"description"`
	UpdatedAt    int64            `json:"updated_at"`
}

// ——— CS:GO 賽程日誌（含 results_extra）———

// CsgoMatchDiaryExtraTournament CS:GO 賽程日誌附帶賽事資料
type CsgoMatchDiaryExtraTournament struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

// CsgoMatchDiaryExtraTeam CS:GO 賽程日誌附帶戰隊資料
type CsgoMatchDiaryExtraTeam struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

// CsgoMatchDiaryExtraStage CS:GO 賽程日誌附帶階段資料
type CsgoMatchDiaryExtraStage struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CsgoMatchDiaryResultsExtra CS:GO 賽程日誌附帶資料
type CsgoMatchDiaryResultsExtra struct {
	Tournament []CsgoMatchDiaryExtraTournament `json:"tournament"`
	Team       []CsgoMatchDiaryExtraTeam       `json:"team"`
	Stage      []CsgoMatchDiaryExtraStage      `json:"stage"`
}

// CsgoMatchDiaryResponse 賽程日誌完整回應（含 results_extra）
type CsgoMatchDiaryResponse struct {
	Code         int                          `json:"code"`
	Results      []CsgoMatchResponseData      `json:"results"`
	ResultsExtra CsgoMatchDiaryResultsExtra   `json:"results_extra"`
	Err          *string                      `json:"err,omitempty"`
	Msg          *string                      `json:"msg,omitempty"`
}

// ——— CS:GO 單場比賽 ———

// CsgoSingleMatchResponseData CS:GO 單場比賽回應資料
type CsgoSingleMatchResponseData struct {
	ID          string          `json:"id"`
	MatchID     string          `json:"match_id"`
	BoxNum      int             `json:"box_num"`
	MapID       string          `json:"map_id"`
	StatusID    int             `json:"status_id"`
	MatchTime   int64           `json:"match_time"`
	HomeTeamID  string          `json:"home_team_id"`
	HomeScore   int             `json:"home_score"`
	AwayTeamID  string          `json:"away_team_id"`
	AwayScore   int             `json:"away_score"`
	Score       json.RawMessage `json:"score"`
	Result      json.RawMessage `json:"result"`
	OverResult  json.RawMessage `json:"over_result"`
	Rating      json.RawMessage `json:"rating"`
	UpdatedAt   int64           `json:"updated_at"`
}

// ——— CS:GO 單場選手統計 ———

// CsgoSingleMatchPlayerStatResponseData CS:GO 單場選手統計回應資料
type CsgoSingleMatchPlayerStatResponseData struct {
	ID             string          `json:"id"`
	MatchSingleID  string          `json:"match_single_id"`
	TeamID         string          `json:"team_id"`
	PlayerID       string          `json:"player_id"`
	Kill           int             `json:"kill"`
	Death          int             `json:"death"`
	Hs             int             `json:"hs"`
	Assists        int             `json:"assists"`
	FAssists       int             `json:"f_assists"`
	Kast           string          `json:"kast"`
	Adr            float64         `json:"adr"`
	KfDiff         string          `json:"kf_diff"`
	AwpKill        int             `json:"awp_kill"`
	FirstKill      int             `json:"first_kill"`
	Impact         float64         `json:"impact"`
	Mks            int             `json:"mks"`
	Onevsx         int             `json:"onevsx"`
	Swing          string          `json:"swing"`
	Rating         float64         `json:"rating"`
	WeaponsKill    json.RawMessage `json:"weapons_kill"`
	UpdatedAt      int64           `json:"updated_at"`
}

// ——— CS:GO 即時數據 ———

// CsgoDetailLivePlayer CS:GO 即時數據選手資料
type CsgoDetailLivePlayer struct {
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Stats    json.RawMessage `json:"stats"`
	WeaponID string          `json:"weapon_id"`
}

// CsgoDetailLiveSide CS:GO 即時數據單邊資料
type CsgoDetailLiveSide struct {
	ID       string                  `json:"id"`
	Side     int                     `json:"side"`
	Score    int                     `json:"score"`
	Stats    json.RawMessage         `json:"stats"`
	WinIcon  json.RawMessage         `json:"win_icon"`
	Players  []CsgoDetailLivePlayer  `json:"players"`
	FhScore  int                     `json:"fh_score"`
	ShScore  int                     `json:"sh_score"`
	MScore   int                     `json:"m_score"`
	CScore   int                     `json:"c_score"`
}

// CsgoDetailLiveEvent CS:GO 即時事件資料
type CsgoDetailLiveEvent struct {
	Index        int    `json:"index"`
	Typeid       int    `json:"typeid"`
	Killername   string `json:"killername"`
	Killerside   int    `json:"killerside"`
	Victimname   string `json:"victimname"`
	Victimside   int    `json:"victimside"`
	Assistername string `json:"assistername"`
	Assisterside int    `json:"assisterside"`
	WeaponID     string `json:"weapon_id"`
	Headshot     int    `json:"headshot"`
	Wintype      int    `json:"wintype"`
	Ctplayers    int    `json:"ctplayers"`
	Tplayers     int    `json:"tplayers"`
	Winner       string `json:"winner"`
	WinnerSide   int    `json:"winner_side"`
	TScore       int    `json:"t_score"`
	CtScore      int    `json:"ct_score"`
	Playername   string `json:"playername"`
	Playerside   int    `json:"playerside"`
	Flashername  string `json:"flashername"`
	Flasherside  int    `json:"flasherside"`
}

// CsgoDetailLiveResponseData CS:GO 即時數據回應資料
type CsgoDetailLiveResponseData struct {
	MatchID         string                  `json:"match_id"`
	BoxNum          int                     `json:"box_num"`
	StatusID        int                     `json:"status_id"`
	SingleStatusID  int                     `json:"single_status_id"`
	MapID           string                  `json:"map_id"`
	RoundNum        int                     `json:"round_num"`
	Timer           json.RawMessage         `json:"timer"`
	Home            CsgoDetailLiveSide      `json:"home"`
	Away            CsgoDetailLiveSide      `json:"away"`
	Event           []CsgoDetailLiveEvent   `json:"event"`
}

// ——— CS:GO 刪除 ———

// CsgoDeleteResponseData CS:GO 刪除回應資料
type CsgoDeleteResponseData struct {
	Type int      `json:"type"`
	IDs  []string `json:"ids"`
}

// ——— CS:GO 戰隊統計 ———

// CsgoTeamStatResponseData CS:GO 戰隊統計回應資料
type CsgoTeamStatResponseData struct {
	ID            string  `json:"id"`
	TeamID        string  `json:"team_id"`
	MatchCount    int     `json:"match_count"`
	Win           int     `json:"win"`
	Lose          int     `json:"lose"`
	FiveKillsRate float64 `json:"five_kills_rate"`
	TenKillsRate  float64 `json:"ten_kills_rate"`
	R1Rate        float64 `json:"r1_rate"`
	R16Rate       float64 `json:"r16_rate"`
	RoundMore26   float64 `json:"round_more_26"`
	UpdatedAt     int64   `json:"updated_at"`
}

// ——— CS:GO 選手統計 ———

// CsgoPlayerStatResponseData CS:GO 選手統計回應資料
type CsgoPlayerStatResponseData struct {
	ID         string  `json:"id"`
	PlayerID   string  `json:"player_id"`
	MatchCount int     `json:"match_count"`
	K          float64 `json:"k"`
	D          float64 `json:"d"`
	Adr        float64 `json:"adr"`
	Kast       float64 `json:"kast"`
	Rating     float64 `json:"rating"`
	UpdatedAt  int64   `json:"updated_at"`
}

// ——— CS:GO 賽事統計 ———

// CsgoTournamentStatTeam CS:GO 賽事統計戰隊資料
type CsgoTournamentStatTeam struct {
	ID    string   `json:"id"`
	Stats []string `json:"stats"`
}

// CsgoTournamentStatPlayer CS:GO 賽事統計選手資料
type CsgoTournamentStatPlayer struct {
	ID    string   `json:"id"`
	Stats []string `json:"stats"`
}

// CsgoTournamentStatMap CS:GO 賽事統計地圖資料
type CsgoTournamentStatMap struct {
	ID    string   `json:"id"`
	Stats []string `json:"stats"`
}

// CsgoTournamentStatResponseData CS:GO 賽事統計回應資料
type CsgoTournamentStatResponseData struct {
	Teams   []CsgoTournamentStatTeam   `json:"teams"`
	Players []CsgoTournamentStatPlayer `json:"players"`
	Maps    []CsgoTournamentStatMap    `json:"maps"`
}

// ——— CS:GO 賽事戰隊 ———

// CsgoTournamentTeamResponseData CS:GO 賽事戰隊回應資料
type CsgoTournamentTeamResponseData struct {
	ID           string `json:"id"`
	TournamentID string `json:"tournament_id"`
	TeamID       string `json:"team_id"`
	UpdatedAt    int64  `json:"updated_at"`
}

// ——— CS:GO 賽事回合 ———

// CsgoTournamentRoundResponseData CS:GO 賽事回合回應資料
type CsgoTournamentRoundResponseData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	UpdatedAt int64  `json:"updated_at"`
}

// ——— CS:GO 賽事對戰表 ———

// CsgoTournamentBracketResponseData CS:GO 賽事對戰表回應資料
type CsgoTournamentBracketResponseData struct {
	ID             string `json:"id"`
	TournamentID   string `json:"tournament_id"`
	StageID        string `json:"stage_id"`
	PartStageID    string `json:"part_stage_id"`
	RoundID        string `json:"round_id"`
	MatchID        string `json:"match_id"`
	MatchTime      int64  `json:"match_time"`
	HomeTeamID     string `json:"home_team_id"`
	AwayTeamID     string `json:"away_team_id"`
	HomeScore      int    `json:"home_score"`
	AwayScore      int    `json:"away_score"`
	Round          int    `json:"round"`
	Sequence       int    `json:"sequence"`
	SequenceType   int    `json:"sequence_type"`
	IsPromotion    int    `json:"is_promotion"`
	Num            int    `json:"num"`
	UpdatedAt      int64  `json:"updated_at"`
}

// ——— CS:GO 賽事對戰表連線 ———

// CsgoTournamentBracketLineResponseData CS:GO 賽事對戰表連線回應資料
type CsgoTournamentBracketLineResponseData struct {
	ID             string          `json:"id"`
	TournamentID   string          `json:"tournament_id"`
	StageID        string          `json:"stage_id"`
	PartStageID    string          `json:"part_stage_id"`
	Lines          json.RawMessage `json:"lines"`
	UpdatedAt      int64           `json:"updated_at"`
}

// ——— CS:GO 賽事積分表 ———

// CsgoTournamentTableGroup CS:GO 賽事積分表分組資料
type CsgoTournamentTableGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CsgoTournamentTableResponseData CS:GO 賽事積分表回應資料
type CsgoTournamentTableResponseData struct {
	ID           string                    `json:"id"`
	TournamentID string                    `json:"tournament_id"`
	StageID      string                    `json:"stage_id"`
	TeamID       string                    `json:"team_id"`
	Group        CsgoTournamentTableGroup  `json:"group"`
	Win          int                       `json:"win"`
	Lose         int                       `json:"lose"`
	Equal        int                       `json:"equal"`
	Score        int                       `json:"score"`
	RoundSub     string                    `json:"round_sub"`
	Position     int                       `json:"position"`
	UpdatedAt    int64                     `json:"updated_at"`
}

// ——— CS:GO 賠率 ———

// CsgoOddsLiveResponseData CS:GO 即時賠率回應資料
type CsgoOddsLiveResponseData struct {
	PlatformID int             `json:"platform_id"`
	Odds       json.RawMessage `json:"odds"`
}

// CsgoOddsHistoryBox CS:GO 賠率歷史局數資料
type CsgoOddsHistoryBox struct {
	BoxNum int             `json:"box_num"`
	Odds   json.RawMessage `json:"odds"`
}

// CsgoOddsHistoryPlay CS:GO 賠率歷史玩法資料
type CsgoOddsHistoryPlay struct {
	PlayID int                    `json:"play_id"`
	Boxes  []CsgoOddsHistoryBox   `json:"boxes"`
}

// CsgoOddsHistoryResponseData CS:GO 賠率歷史回應資料
type CsgoOddsHistoryResponseData struct {
	PlatformID int                       `json:"platform_id"`
	Plays      []CsgoOddsHistoryPlay     `json:"plays"`
}

// ——— Dota2 基本資料 ———

// Dota2HeroResponseData Dota2 英雄回應資料
type Dota2HeroResponseData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Logo      string `json:"logo"`
	Icon      string `json:"icon"`
	VertLogo  string `json:"vert_logo"`
	UpdatedAt int64  `json:"updated_at"`
}

// Dota2EquipmentResponseData Dota2 裝備回應資料
type Dota2EquipmentResponseData struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Logo       string `json:"logo"`
	Cost       int    `json:"cost"`
	SecretShop int    `json:"secret_shop"`
	SideShop   int    `json:"side_shop"`
	Recipe     int    `json:"recipe"`
	UpdatedAt  int64  `json:"updated_at"`
}

// ——— LoL 基本資料 ———

// LolSpellResponseData LoL 召喚師技能回應資料
type LolSpellResponseData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Logo      string `json:"logo"`
	UpdatedAt int64  `json:"updated_at"`
}

// LolRuneResponseData LoL 符文回應資料
type LolRuneResponseData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Logo      string `json:"logo"`
	Type      int    `json:"type"`
	UpdatedAt int64  `json:"updated_at"`
}

// LolMatchEventResponseData LoL 比賽事件回應資料
type LolMatchEventResponseData struct {
	ID          string `json:"id"`
	MatchID     string `json:"match_id"`
	BoxNum      int    `json:"box_num"`
	Type        int    `json:"type"`
	Side        int    `json:"side"`
	Time        int    `json:"time"`
	Killer      string `json:"killer"`
	Victim      string `json:"victim"`
	MonsterType int    `json:"monster_type"`
	UpdatedAt   int64  `json:"updated_at"`
}

// ——— CsgoMatchDiaryResponse 錯誤實作 ———

// Error 實作 error 介面
func (r CsgoMatchDiaryResponse) Error() string {
	var msg []string
	if r.Code != 0 {
		msg = append(msg, "code: "+strconv.Itoa(r.Code))
	}
	if r.Err != nil && *r.Err != "" {
		msg = append(msg, "err: "+*r.Err)
	}
	if r.Msg != nil && *r.Msg != "" {
		msg = append(msg, "msg: "+*r.Msg)
	}
	msg = append(msg, "raw: diary response")
	return strings.Join(msg, ", ")
}
