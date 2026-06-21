package thesportsgo

import "encoding/json"

// FootballCategoryResponseData 足球分類
type FootballCategoryResponseData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	UpdatedAt int    `json:"updated_at"`
}

// FootballCountryResponseData 足球國家/地區
type FootballCountryResponseData struct {
	CategoryID string `json:"category_id"`
	ID         string `json:"id"`
	Logo       string `json:"logo"`
	Name       string `json:"name"`
	UpdatedAt  int64  `json:"updated_at"`
}

// FootballCompetitionHost 賽事主辦方
type FootballCompetitionHost struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

// FootballCompetitionResponseData 足球賽事
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

// FootballTeamResponseData 足球球隊
type FootballTeamResponseData struct {
	ID                  string `json:"id"`
	CompetitionID       string `json:"competition_id"`
	CountryID           string `json:"country_id"`
	Name                string `json:"name"`
	ShortName           string `json:"short_name"`
	Logo                string `json:"logo"`
	National            int    `json:"national"`
	CountryLogo         string `json:"country_logo"`
	FoundationTime      int64  `json:"foundation_time"`
	Website             string `json:"website"`
	CoachID             string `json:"coach_id"`
	VenueID             string `json:"venue_id"`
	MarketValue         int    `json:"market_value"`
	MarketValueCurrency string `json:"market_value_currency"`
	TotalPlayers        int    `json:"total_players"`
	ForeignPlayers      int    `json:"foreign_players"`
	NationalPlayers     int    `json:"national_players"`
	UID                 string `json:"uid"`
	Virtual             int    `json:"virtual"`
	Gender              int    `json:"gender"`
	UpdatedAt           int64  `json:"updated_at"`
}

// FootballPlayerResponseData 足球球員
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
	Deathday            int64           `json:"deathday"`
	RetireTime          int64           `json:"retire_time"`
	UpdatedAt           int64           `json:"updated_at"`
}

// FootballCoachResponseData 足球教練
type FootballCoachResponseData struct {
	ID                 string `json:"id"`
	TeamID             string `json:"team_id"`
	Name               string `json:"name"`
	ShortName          string `json:"short_name"`
	Logo               string `json:"logo"`
	Type               int    `json:"type"`
	Birthday           int64  `json:"birthday"`
	Age                int    `json:"age"`
	PreferredFormation string `json:"preferred_formation"`
	CountryID          string `json:"country_id"`
	Nationality        string `json:"nationality"`
	Joined             int64  `json:"joined"`
	ContractUntil      int64  `json:"contract_until"`
	UID                string `json:"uid"`
	Deathday           int64  `json:"deathday"`
	UpdatedAt          int64  `json:"updated_at"`
}

// FootballRefereeResponseData 足球裁判
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

// FootballVenueResponseData 足球場館
type FootballVenueResponseData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Capacity  int    `json:"capacity"`
	CountryID string `json:"country_id"`
	City      string `json:"city"`
	Country   string `json:"country"`
	RealName  string `json:"real_name"`
	UID       string `json:"uid"`
	UpdatedAt int64  `json:"updated_at"`
}

// FootballSeasonResponseData 足球賽季
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

// FootballStageResponseData 足球階段
type FootballStageResponseData struct {
	ID         string `json:"id"`
	SeasonID   string `json:"season_id"`
	Name       string `json:"name"`
	Mode       int    `json:"mode"`
	GroupCount int    `json:"group_count"`
	RoundCount int    `json:"round_count"`
	Order      int    `json:"order"`
	UpdatedAt  int64  `json:"updated_at"`
}

// --- Match-related response types ---

// FootballMatchCoverage 比賽覆蓋資料
type FootballMatchCoverage struct {
	Mlive  int `json:"mlive"`
	Lineup int `json:"lineup"`
	Gif    int `json:"gif"`
}

// FootballMatchRound 比賽階段
type FootballMatchRound struct {
	StageID  string `json:"stage_id"`
	GroupNum int    `json:"group_num"`
	RoundNum int    `json:"round_num"`
}

// FootballMatchEnvironment 比賽環境
type FootballMatchEnvironment struct {
	Weather     int    `json:"weather"`
	Pressure    string `json:"pressure"`
	Temperature string `json:"temperature"`
	Wind        string `json:"wind"`
	Humidity    string `json:"humidity"`
}

// FootballMatchResponseData 比賽資料（共用於 match/list, match/recent/list, match/diary, match/season, match/season/recent）
type FootballMatchResponseData struct {
	ID            string                    `json:"id"`
	SeasonID      string                    `json:"season_id"`
	CompetitionID string                    `json:"competition_id"`
	HomeTeamID    string                    `json:"home_team_id"`
	AwayTeamID    string                    `json:"away_team_id"`
	StatusID      int                       `json:"status_id"`
	MatchTime     int64                     `json:"match_time"`
	VenueID       string                    `json:"venue_id"`
	RefereeID     string                    `json:"referee_id"`
	Neutral       int                       `json:"neutral"`
	Note          string                    `json:"note"`
	HomeScores    json.RawMessage           `json:"home_scores"`
	AwayScores    json.RawMessage           `json:"away_scores"`
	HomePosition  string                    `json:"home_position"`
	AwayPosition  string                    `json:"away_position"`
	Coverage      FootballMatchCoverage     `json:"coverage"`
	Round         FootballMatchRound        `json:"round"`
	RelatedID     string                    `json:"related_id"`
	AggScore      json.RawMessage           `json:"agg_score"`
	Environment   *FootballMatchEnvironment `json:"environment"`
	TBD           *int                      `json:"tbd"`
	HasOT         *int                      `json:"has_ot"`
	Ended         *int64                    `json:"ended"`
	TeamReverse   *int                      `json:"team_reverse"`
	Loss          *int                      `json:"loss"`
	UpdatedAt     int64                     `json:"updated_at"`
}

// --- Match Detail Live ---

// FootballMatchDetailLiveStat 即時統計
type FootballMatchDetailLiveStat struct {
	Type int `json:"type"`
	Home int `json:"home"`
	Away int `json:"away"`
}

// FootballMatchDetailLiveIncident 即時事件
type FootballMatchDetailLiveIncident struct {
	Type          int     `json:"type"`
	Position      int     `json:"position"`
	Time          int     `json:"time"`
	Second        *int    `json:"second"`
	AddTime       *int    `json:"add_time"`
	PlayerID      *string `json:"player_id"`
	PlayerName    *string `json:"player_name"`
	Assist1ID     *string `json:"assist1_id"`
	Assist1Name   *string `json:"assist1_name"`
	Assist2ID     *string `json:"assist2_id"`
	Assist2Name   *string `json:"assist2_name"`
	HomeScore     *int    `json:"home_score"`
	AwayScore     *int    `json:"away_score"`
	InPlayerID    *string `json:"in_player_id"`
	InPlayerName  *string `json:"in_player_name"`
	OutPlayerID   *string `json:"out_player_id"`
	OutPlayerName *string `json:"out_player_name"`
	VarReason     *int    `json:"var_reason"`
	VarResult     *int    `json:"var_result"`
	ReasonType    *int    `json:"reason_type"`
}

// FootballMatchDetailLiveTlive 即時文字直播
type FootballMatchDetailLiveTlive struct {
	Time     string `json:"time"`
	Data     string `json:"data"`
	Position int    `json:"position"`
}

// FootballMatchDetailLiveResponseData 即時資料
type FootballMatchDetailLiveResponseData struct {
	ID        string                            `json:"id"`
	Score     json.RawMessage                   `json:"score"`
	Stats     []FootballMatchDetailLiveStat     `json:"stats"`
	Incidents []FootballMatchDetailLiveIncident `json:"incidents"`
	Tlive     []FootballMatchDetailLiveTlive    `json:"tlive"`
}

// --- Match Trend ---

// FootballMatchTrendData 比賽走勢資料
type FootballMatchTrendData struct {
	Count int             `json:"count"`
	Per   int             `json:"per"`
	Data  json.RawMessage `json:"data"`
}

// FootballMatchTrendDetailResponseData 比賽走勢詳情
type FootballMatchTrendDetailResponseData struct {
	MatchID string                 `json:"match_id"`
	Trend   FootballMatchTrendData `json:"trend"`
}

// --- Match Lineup ---

// FootballMatchLineupPlayer 球員
type FootballMatchLineupPlayer struct {
	ID          string                        `json:"id"`
	First       int                           `json:"first"`
	Captain     int                           `json:"captain"`
	Name        string                        `json:"name"`
	Logo        string                        `json:"logo"`
	ShirtNumber int                           `json:"shirt_number"`
	Position    string                        `json:"position"`
	X           *int                          `json:"x"`
	Y           *int                          `json:"y"`
	Rating      string                        `json:"rating"`
	Incidents   []FootballMatchLineupIncident `json:"incidents"`
}

// FootballMatchLineupIncident 陣容事件
type FootballMatchLineupIncident struct {
	Type      int                           `json:"type"`
	Time      string                        `json:"time"`
	Minute    int                           `json:"minute"`
	Addtime   *int                          `json:"addtime"`
	Belong    int                           `json:"belong"`
	HomeScore int                           `json:"home_score"`
	AwayScore int                           `json:"away_score"`
	Player    *FootballMatchLineupPlayerRef `json:"player"`
	Assist1   *FootballMatchLineupPlayerRef `json:"assist1"`
	Assist2   *FootballMatchLineupPlayerRef `json:"assist2"`
	InPlayer  *FootballMatchLineupPlayerRef `json:"in_player"`
	OutPlayer *FootballMatchLineupPlayerRef `json:"out_player"`
}

// FootballMatchLineupPlayerRef 陣容球員引用
type FootballMatchLineupPlayerRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// FootballMatchLineupCoach 教練
type FootballMatchLineupCoach struct {
	Home string `json:"home"`
	Away string `json:"away"`
}

// FootballMatchLineupInjury 傷病
type FootballMatchLineupInjury struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Position      string `json:"position"`
	Logo          string `json:"logo"`
	Type          int    `json:"type"`
	Reason        string `json:"reason"`
	StartTime     int64  `json:"start_time"`
	EndTime       int64  `json:"end_time"`
	MissedMatches int    `json:"missed_matches"`
}

// FootballMatchLineupResponseData 單場陣容
type FootballMatchLineupResponseData struct {
	Confirmed     int                            `json:"confirmed"`
	HomeFormation string                         `json:"home_formation"`
	AwayFormation string                         `json:"away_formation"`
	CoachID       FootballMatchLineupCoach       `json:"coach_id"`
	Lineup        FootballMatchLineupLineup      `json:"lineup"`
	Injury        *FootballMatchLineupInjuryData `json:"injury"`
}

// FootballMatchLineupLineup 陣容
type FootballMatchLineupLineup struct {
	Home []FootballMatchLineupPlayer `json:"home"`
	Away []FootballMatchLineupPlayer `json:"away"`
}

// FootballMatchLineupInjuryData 傷病資料
type FootballMatchLineupInjuryData struct {
	Home []FootballMatchLineupInjury `json:"home"`
	Away []FootballMatchLineupInjury `json:"away"`
}

// --- Match Player Stats ---

// FootballMatchPlayerStat 球員統計
type FootballMatchPlayerStat struct {
	PlayerID          string  `json:"player_id"`
	TeamID            string  `json:"team_id"`
	First             int     `json:"first"`
	Goals             int     `json:"goals"`
	Penalty           int     `json:"penalty"`
	Assists           int     `json:"assists"`
	MinutesPlayed     int     `json:"minutes_played"`
	RedCards          int     `json:"red_cards"`
	YellowCards       int     `json:"yellow_cards"`
	Shots             int     `json:"shots"`
	ShotsOnTarget     int     `json:"shots_on_target"`
	Dribble           int     `json:"dribble"`
	DribbleSucc       int     `json:"dribble_succ"`
	Clearances        int     `json:"clearances"`
	BlockedShots      int     `json:"blocked_shots"`
	Interceptions     int     `json:"interceptions"`
	Tackles           int     `json:"tackles"`
	Passes            int     `json:"passes"`
	PassesAccuracy    int     `json:"passes_accuracy"`
	KeyPasses         int     `json:"key_passes"`
	Crosses           int     `json:"crosses"`
	CrossesAccuracy   int     `json:"crosses_accuracy"`
	LongBalls         int     `json:"long_balls"`
	LongBallsAccuracy int     `json:"long_balls_accuracy"`
	Duels             int     `json:"duels"`
	DuelsWon          int     `json:"duels_won"`
	Dispossessed      int     `json:"dispossessed"`
	Fouls             int     `json:"fouls"`
	WasFouled         int     `json:"was_fouled"`
	Offsides          int     `json:"offsides"`
	Yellow2RedCards   int     `json:"yellow2red_cards"`
	Saves             int     `json:"saves"`
	Punches           int     `json:"punches"`
	RunsOut           int     `json:"runs_out"`
	RunsOutSucc       int     `json:"runs_out_succ"`
	GoodHighClaim     int     `json:"good_high_claim"`
	Rating            float64 `json:"rating"`
	Freekicks         int     `json:"freekicks"`
	FreekickGoals     int     `json:"freekick_goals"`
	HitWoodwork       int     `json:"hit_woodwork"`
	Fastbreaks        int     `json:"fastbreaks"`
	FastbreakShots    int     `json:"fastbreak_shots"`
	FastbreakGoals    int     `json:"fastbreak_goals"`
	PossLosts         int     `json:"poss_losts"`
	AerialWon         int     `json:"aerial_won"`
	AerialLost        int     `json:"aerial_lost"`
	GroundWon         int     `json:"ground_won"`
	GroundLost        int     `json:"ground_lost"`
	DuelWon           int     `json:"duel_won"`
	DuelLost          int     `json:"duel_lost"`
	BigChanceMissed   int     `json:"big_chance_missed"`
	BigChanceCreated  int     `json:"big_chance_created"`
}

// FootballMatchPlayerStatsResponseData 比賽球員統計（list）
type FootballMatchPlayerStatsResponseData struct {
	ID          string                    `json:"id"`
	PlayerStats []FootballMatchPlayerStat `json:"player_stats"`
}

// --- Match Team Stats ---

// FootballMatchTeamStat 球隊統計
type FootballMatchTeamStat struct {
	TeamID            string `json:"team_id"`
	Goals             int    `json:"goals"`
	Penalty           int    `json:"penalty"`
	Assists           int    `json:"assists"`
	RedCards          int    `json:"red_cards"`
	YellowCards       int    `json:"yellow_cards"`
	Shots             int    `json:"shots"`
	ShotsOnTarget     int    `json:"shots_on_target"`
	Dribble           int    `json:"dribble"`
	DribbleSucc       int    `json:"dribble_succ"`
	Clearances        int    `json:"clearances"`
	BlockedShots      int    `json:"blocked_shots"`
	Interceptions     int    `json:"interceptions"`
	Tackles           int    `json:"tackles"`
	Passes            int    `json:"passes"`
	PassesAccuracy    int    `json:"passes_accuracy"`
	KeyPasses         int    `json:"key_passes"`
	Crosses           int    `json:"crosses"`
	CrossesAccuracy   int    `json:"crosses_accuracy"`
	LongBalls         int    `json:"long_balls"`
	LongBallsAccuracy int    `json:"long_balls_accuracy"`
	Duels             int    `json:"duels"`
	DuelsWon          int    `json:"duels_won"`
	Fouls             int    `json:"fouls"`
	WasFouled         int    `json:"was_fouled"`
	GoalsAgainst      int    `json:"goals_against"`
	Offsides          int    `json:"offsides"`
	Yellow2RedCards   int    `json:"yellow2red_cards"`
	CornerKicks       int    `json:"corner_kicks"`
	BallPossession    int    `json:"ball_possession"`
	Freekicks         int    `json:"freekicks"`
	FreekickGoals     int    `json:"freekick_goals"`
	BigChanceMissed   int    `json:"big_chance_missed"`
	BigChanceCreated  int    `json:"big_chance_created"`
	HitWoodwork       int    `json:"hit_woodwork"`
	Fastbreaks        int    `json:"fastbreaks"`
	FastbreakShots    int    `json:"fastbreak_shots"`
	FastbreakGoals    int    `json:"fastbreak_goals"`
	PossLosts         int    `json:"poss_losts"`
	AerialWon         int    `json:"aerial_won"`
	AerialLost        int    `json:"aerial_lost"`
	GroundWon         int    `json:"ground_won"`
	GroundLost        int    `json:"ground_lost"`
	DuelWon           int    `json:"duel_won"`
	DuelLost          int    `json:"duel_lost"`
}

// FootballMatchTeamStatsResponseData 比賽球隊統計（list）
type FootballMatchTeamStatsResponseData struct {
	ID    string                  `json:"id"`
	Stats []FootballMatchTeamStat `json:"stats"`
}

// --- Match Analysis (H2H) ---

// FootballMatchAnalysisResponseData H2H 分析
// info/history/future 為複雜巢狀陣列結構，以 json.RawMessage 保留原始資料
type FootballMatchAnalysisResponseData struct {
	Info             json.RawMessage `json:"info"`
	History          json.RawMessage `json:"history"`
	Future           json.RawMessage `json:"future"`
	GoalDistribution json.RawMessage `json:"goal_distribution"`
}

// --- Match Live History ---

// FootballMatchLiveHistoryResponseData 比賽歷史統計
type FootballMatchLiveHistoryResponseData struct {
	ID        string                            `json:"id"`
	Score     json.RawMessage                   `json:"score"`
	Stats     []FootballMatchDetailLiveStat     `json:"stats"`
	Incidents []FootballMatchDetailLiveIncident `json:"incidents"`
	Tlive     []FootballMatchDetailLiveTlive    `json:"tlive"`
}

// --- Match Goal Line ---

// FootballMatchGoalLinePass 傳球線
type FootballMatchGoalLinePass struct {
	Belong      int    `json:"belong"`
	PlayerID    string `json:"player_id"`
	ShirtNumber string `json:"shirt_number"`
	X           string `json:"x"`
	Y           string `json:"y"`
	Shooter     int    `json:"shooter"`
	Assist      int    `json:"assist"`
}

// FootballMatchGoalLineGoalkeeper 守門員
type FootballMatchGoalLineGoalkeeper struct {
	Belong      int    `json:"belong"`
	PlayerID    string `json:"player_id"`
	ShirtNumber string `json:"shirt_number"`
}

// FootballMatchGoalLineResponseData 進球線
type FootballMatchGoalLineResponseData struct {
	Number     int                              `json:"number"`
	Time       int                              `json:"time"`
	GoalX      string                           `json:"goal_x"`
	GoalY      string                           `json:"goal_y"`
	OwnGoal    int                              `json:"own_goal"`
	Pass       []FootballMatchGoalLinePass      `json:"pass"`
	Goalkeeper *FootballMatchGoalLineGoalkeeper `json:"goalkeeper"`
}

// --- Match Highlight ---

// FootballMatchHighlightResponseData 比賽精華 GIF
type FootballMatchHighlightResponseData struct {
	Type     int    `json:"type"`
	Time     int    `json:"time"`
	Position int    `json:"position"`
	Home     int    `json:"home"`
	Away     int    `json:"away"`
	Gif      string `json:"gif"`
	Cover    string `json:"cover"`
}

// --- Match TV ---

// FootballMatchTVChannel 電視頻道
type FootballMatchTVChannel struct {
	CountryID string   `json:"country_id"`
	Names     []string `json:"names"`
}

// FootballMatchTVResponseData 比賽電視頻道
type FootballMatchTVResponseData struct {
	ID        string                   `json:"id"`
	TV        []FootballMatchTVChannel `json:"tv"`
	UpdatedAt int64                    `json:"updated_at"`
}

// --- Season table ---

// FootballTablePromotion 升降級資訊
type FootballTablePromotion struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

// FootballTableRow 積分榜行
type FootballTableRow struct {
	TeamID           string `json:"team_id"`
	PromotionID      string `json:"promotion_id"`
	Points           int    `json:"points"`
	Position         int    `json:"position"`
	DeductPoints     int    `json:"deduct_points"`
	Note             string `json:"note"`
	Total            int    `json:"total"`
	Won              int    `json:"won"`
	Draw             int    `json:"draw"`
	Loss             int    `json:"loss"`
	Goals            int    `json:"goals"`
	GoalsAgainst     int    `json:"goals_against"`
	GoalDiff         int    `json:"goal_diff"`
	HomePoints       int    `json:"home_points"`
	HomePosition     int    `json:"home_position"`
	HomeTotal        int    `json:"home_total"`
	HomeWon          int    `json:"home_won"`
	HomeDraw         int    `json:"home_draw"`
	HomeLoss         int    `json:"home_loss"`
	HomeGoals        int    `json:"home_goals"`
	HomeGoalsAgainst int    `json:"home_goals_against"`
	HomeGoalDiff     int    `json:"home_goal_diff"`
	AwayPoints       int    `json:"away_points"`
	AwayPosition     int    `json:"away_position"`
	AwayTotal        int    `json:"away_total"`
	AwayWon          int    `json:"away_won"`
	AwayDraw         int    `json:"away_draw"`
	AwayLoss         int    `json:"away_loss"`
	AwayGoals        int    `json:"away_goals"`
	AwayGoalsAgainst int    `json:"away_goals_against"`
	AwayGoalDiff     int    `json:"away_goal_diff"`
	UpdatedAt        int64  `json:"updated_at"`
}

// FootballTableData 積分表
type FootballTableData struct {
	ID         string             `json:"id"`
	Conference string             `json:"conference"`
	Group      int                `json:"group"`
	StageID    string             `json:"stage_id"`
	Rows       []FootballTableRow `json:"rows"`
}

// FootballSeasonTableResponseData 賽季積分榜
type FootballSeasonTableResponseData struct {
	Promotions []FootballTablePromotion `json:"promotions"`
	Tables     []FootballTableData      `json:"tables"`
}

// --- Table live ---

// FootballTableLiveData 即時積分榜資料
type FootballTableLiveData struct {
	SeasonID   string                   `json:"season_id"`
	Promotions []FootballTablePromotion `json:"promotions"`
	Tables     []FootballTableData      `json:"tables"`
	UpdatedAt  int64                    `json:"updated_at"`
}

// --- Season team stat ---

// FootballTeamStatTeam 球隊統計隊伍資訊
type FootballTeamStatTeam struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

// FootballSeasonTeamStatResponseData 賽季球隊統計
type FootballSeasonTeamStatResponseData struct {
	Team              FootballTeamStatTeam `json:"team"`
	Matches           int                  `json:"matches"`
	Goals             int                  `json:"goals"`
	Penalty           int                  `json:"penalty"`
	Assists           int                  `json:"assists"`
	RedCards          int                  `json:"red_cards"`
	YellowCards       int                  `json:"yellow_cards"`
	Shots             int                  `json:"shots"`
	ShotsOnTarget     int                  `json:"shots_on_target"`
	Dribble           int                  `json:"dribble"`
	DribbleSucc       int                  `json:"dribble_succ"`
	Clearances        int                  `json:"clearances"`
	BlockedShots      int                  `json:"blocked_shots"`
	Tackles           int                  `json:"tackles"`
	Passes            int                  `json:"passes"`
	PassesAccuracy    int                  `json:"passes_accuracy"`
	KeyPasses         int                  `json:"key_passes"`
	Crosses           int                  `json:"crosses"`
	CrossesAccuracy   int                  `json:"crosses_accuracy"`
	LongBalls         int                  `json:"long_balls"`
	LongBallsAccuracy int                  `json:"long_balls_accuracy"`
	Duels             int                  `json:"duels"`
	DuelsWon          int                  `json:"duels_won"`
	Fouls             int                  `json:"fouls"`
	WasFouled         int                  `json:"was_fouled"`
	GoalsAgainst      int                  `json:"goals_against"`
	Interceptions     int                  `json:"interceptions"`
	Offsides          int                  `json:"offsides"`
	Yellow2RedCards   int                  `json:"yellow2red_cards"`
	CornerKicks       int                  `json:"corner_kicks"`
	BallPossession    int                  `json:"ball_possession"`
	Freekicks         int                  `json:"freekicks"`
	FreekickGoals     int                  `json:"freekick_goals"`
	HitWoodwork       int                  `json:"hit_woodwork"`
	Fastbreaks        int                  `json:"fastbreaks"`
	FastbreakShots    int                  `json:"fastbreak_shots"`
	FastbreakGoals    int                  `json:"fastbreak_goals"`
	PossLosts         int                  `json:"poss_losts"`
	Saves             int                  `json:"saves"`
	PenaltyConceded   int                  `json:"penalty_conceded"`
	BigChanceCreated  int                  `json:"big_chance_created"`
	BigChanceMissed   int                  `json:"big_chance_missed"`
	AerialWon         int                  `json:"aerial_won"`
	AerialLost        int                  `json:"aerial_lost"`
	GroundWon         int                  `json:"ground_won"`
	GroundLost        int                  `json:"ground_lost"`
	UpdatedAt         int64                `json:"updated_at"`
}

// --- Season player stat ---

// FootballSeasonPlayerStatPlayer 賽季球員統計球員資訊
type FootballSeasonPlayerStatPlayer struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Logo        string `json:"logo"`
	Position    string `json:"position"`
	CountryID   string `json:"country_id"`
	Nationality string `json:"nationality"`
}

// FootballSeasonPlayerStatTeam 賽季球員統計隊伍資訊
type FootballSeasonPlayerStatTeam struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

// FootballSeasonPlayerStatResponseData 賽季球員統計
type FootballSeasonPlayerStatResponseData struct {
	Player            FootballSeasonPlayerStatPlayer `json:"player"`
	Team              FootballSeasonPlayerStatTeam   `json:"team"`
	Matches           int                            `json:"matches"`
	Court             int                            `json:"court"`
	First             int                            `json:"first"`
	Goals             int                            `json:"goals"`
	Penalty           int                            `json:"penalty"`
	Assists           int                            `json:"assists"`
	MinutesPlayed     int                            `json:"minutes_played"`
	RedCards          int                            `json:"red_cards"`
	YellowCards       int                            `json:"yellow_cards"`
	Shots             int                            `json:"shots"`
	ShotsOnTarget     int                            `json:"shots_on_target"`
	Dribble           int                            `json:"dribble"`
	DribbleSucc       int                            `json:"dribble_succ"`
	Clearances        int                            `json:"clearances"`
	BlockedShots      int                            `json:"blocked_shots"`
	Interceptions     int                            `json:"interceptions"`
	Tackles           int                            `json:"tackles"`
	Passes            int                            `json:"passes"`
	PassesAccuracy    int                            `json:"passes_accuracy"`
	KeyPasses         int                            `json:"key_passes"`
	Crosses           int                            `json:"crosses"`
	CrossesAccuracy   int                            `json:"crosses_accuracy"`
	LongBalls         int                            `json:"long_balls"`
	LongBallsAccuracy int                            `json:"long_balls_accuracy"`
	Duels             int                            `json:"duels"`
	DuelsWon          int                            `json:"duels_won"`
	Dispossessed      int                            `json:"dispossessed"`
	Fouls             int                            `json:"fouls"`
	WasFouled         int                            `json:"was_fouled"`
	GoalsAgainst      int                            `json:"goals_against"`
	Offsides          int                            `json:"offsides"`
	Yellow2RedCards   int                            `json:"yellow2red_cards"`
	CornerKicks       int                            `json:"corner_kicks"`
	BallPossession    int                            `json:"ball_possession"`
	Freekicks         int                            `json:"freekicks"`
	FreekickGoals     int                            `json:"freekick_goals"`
	HitWoodwork       int                            `json:"hit_woodwork"`
	Fastbreaks        int                            `json:"fastbreaks"`
	FastbreakShots    int                            `json:"fastbreak_shots"`
	FastbreakGoals    int                            `json:"fastbreak_goals"`
	PossLosts         int                            `json:"poss_losts"`
	Saves             int                            `json:"saves"`
	PenaltyConceded   int                            `json:"penalty_conceded"`
	BigChanceCreated  int                            `json:"big_chance_created"`
	BigChanceMissed   int                            `json:"big_chance_missed"`
	AerialWon         int                            `json:"aerial_won"`
	AerialLost        int                            `json:"aerial_lost"`
	GroundWon         int                            `json:"ground_won"`
	GroundLost        int                            `json:"ground_lost"`
	UpdatedAt         int64                          `json:"updated_at"`
}

// --- Season shooter stat ---

// FootballSeasonShooterStatPlayer 射手榜球員資訊
type FootballSeasonShooterStatPlayer struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Logo        string `json:"logo"`
	Position    string `json:"position"`
	CountryID   string `json:"country_id"`
	Nationality string `json:"nationality"`
}

// FootballSeasonShooterStatResponseData 賽季射手榜
type FootballSeasonShooterStatResponseData struct {
	Position      int                             `json:"position"`
	Player        FootballSeasonShooterStatPlayer `json:"player"`
	Team          FootballSeasonPlayerStatTeam    `json:"team"`
	Goals         int                             `json:"goals"`
	Penalty       int                             `json:"penalty"`
	Assists       int                             `json:"assists"`
	MinutesPlayed int                             `json:"minutes_played"`
	UpdatedAt     int64                           `json:"updated_at"`
}

// --- Bracket ---

// FootballBracketSeasonBracket 盃賽結構
type FootballBracketSeasonBracket struct {
	ID            string `json:"id"`
	CompetitionID string `json:"competition_id"`
	Name          string `json:"name"`
	Abbr          string `json:"abbr"`
	StartTime     int64  `json:"start_time"`
	EndTime       int64  `json:"end_time"`
}

// FootballBracketSeasonGroup 盃賽分組
type FootballBracketSeasonGroup struct {
	ID        string `json:"id"`
	BracketID string `json:"bracket_id"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	TypeID    int    `json:"type_id"`
	Number    int    `json:"number"`
}

// FootballBracketSeasonRound 盃賽輪次
type FootballBracketSeasonRound struct {
	ID        string `json:"id"`
	BracketID string `json:"bracket_id"`
	GroupID   string `json:"group_id"`
	Name      string `json:"name"`
	Abbr      string `json:"abbr"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	Number    int    `json:"number"`
}

// FootballBracketSeasonMatchUp 盃賽對陣
type FootballBracketSeasonMatchUp struct {
	ID           string          `json:"id"`
	RoundID      string          `json:"round_id"`
	Number       int             `json:"number"`
	TypeID       int             `json:"type_id"`
	StateID      int             `json:"state_id"`
	HomeTeamID   string          `json:"home_team_id"`
	AwayTeamID   string          `json:"away_team_id"`
	WinnerTeamID string          `json:"winner_team_id"`
	HomeScore    int             `json:"home_score"`
	AwayScore    int             `json:"away_score"`
	ParentIDs    json.RawMessage `json:"parent_ids"`
	ChildrenIDs  json.RawMessage `json:"children_ids"`
	MatchIDs     json.RawMessage `json:"match_ids"`
	Note         string          `json:"note"`
}

// FootballBracketSeasonResponseData 盃賽資料
type FootballBracketSeasonResponseData struct {
	Brackets []FootballBracketSeasonBracket `json:"brackets"`
	Groups   []FootballBracketSeasonGroup   `json:"groups"`
	Rounds   []FootballBracketSeasonRound   `json:"rounds"`
	MatchUps []FootballBracketSeasonMatchUp `json:"match_ups"`
}

// --- Ranking ---

// FootballRankingTeam 排名隊伍資訊
type FootballRankingTeam struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Logo        string `json:"logo"`
	CountryLogo string `json:"country_logo"`
}

// FootballRankingFifaItem FIFA 排名項目
type FootballRankingFifaItem struct {
	Team            FootballRankingTeam `json:"team"`
	RegionID        int                 `json:"region_id"`
	Ranking         int                 `json:"ranking"`
	Points          int                 `json:"points"`
	PreviousPoints  int                 `json:"previous_points"`
	PositionChanged int                 `json:"position_changed"`
}

// FootballRankingFifaResponseData FIFA 排名
type FootballRankingFifaResponseData struct {
	PubTimes []int64                   `json:"pub_times"`
	PubTime  int64                     `json:"pub_time"`
	Items    []FootballRankingFifaItem `json:"items"`
}

// --- Club ranking ---

// FootballRankingClubItem 俱樂部排名項目
type FootballRankingClubItem struct {
	Team    FootballRankingTeam `json:"team"`
	Ranking int                 `json:"ranking"`
	Points  float64             `json:"points"`
}

// FootballRankingClubResponseData 俱樂部排名
type FootballRankingClubResponseData struct {
	Items []FootballRankingClubItem `json:"items"`
}

// --- FIFA live ranking ---

// FootballRankingFifaLiveItem FIFA 即時排名項目
type FootballRankingFifaLiveItem struct {
	Team     FootballRankingTeam `json:"team"`
	RegionID int                 `json:"region_id"`
	Ranking  int                 `json:"ranking"`
	Points   float64             `json:"points"`
}

// FootballRankingFifaLiveResponseData FIFA 即時排名
type FootballRankingFifaLiveResponseData struct {
	Items []FootballRankingFifaLiveItem `json:"items"`
}

// --- Deleted ---

// FootballDeletedResponseData 已刪除資料
type FootballDeletedResponseData struct {
	Match       json.RawMessage `json:"match"`
	Team        json.RawMessage `json:"team"`
	Player      json.RawMessage `json:"player"`
	Competition json.RawMessage `json:"competition"`
	Season      json.RawMessage `json:"season"`
	Stage       json.RawMessage `json:"stage"`
}

// --- Data update ---

// FootballDataUpdateItem 資料更新項目
type FootballDataUpdateItem struct {
	MatchID    string `json:"match_id"`
	SeasonID   string `json:"season_id"`
	PubTime    int64  `json:"pub_time"`
	TeamID     string `json:"team_id"`
	UpdateTime int64  `json:"update_time"`
}

// FootballDataUpdateResponseData 資料更新，key 為類型 ID
type FootballDataUpdateResponseData map[string][]FootballDataUpdateItem

// --- Language ---

// FootballLanguageResponseData 語言
type FootballLanguageResponseData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// --- Honor ---

// FootballHonorResponseData 榮譽
type FootballHonorResponseData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// --- Compensation ---

// FootballCompensationResponseData 歷史賠償
type FootballCompensationResponseData struct {
	ID        string  `json:"id"`
	StartTime int64   `json:"start_time"`
	EndTime   int64   `json:"end_time"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
}

// --- Odds ---

// FootballOddsLiveResponseData 即時賠率，key 為賠率公司 ID
type FootballOddsLiveResponseData map[string]json.RawMessage

// FootballOddsHistoryCompany 單場賠率公司資料
type FootballOddsHistoryCompany struct {
	Asia json.RawMessage `json:"asia"`
	EU   json.RawMessage `json:"eu"`
	BS   json.RawMessage `json:"bs"`
	CR   json.RawMessage `json:"cr"`
}

// FootballOddsHistoryResponseData 單場賠率，key 為賠率公司 ID
type FootballOddsHistoryResponseData map[string]FootballOddsHistoryCompany

// FootballOddsUpdateResponseData 賠率更新，key 為賠率公司 ID
type FootballOddsUpdateResponseData map[string]json.RawMessage

// --- Team squad ---

// FootballTeamSquadTeam 陣容隊伍資訊
type FootballTeamSquadTeam struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// FootballTeamSquadPlayer 陣容球員資訊
type FootballTeamSquadPlayer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// FootballTeamSquadItem 陣容球員
type FootballTeamSquadItem struct {
	Player      FootballTeamSquadPlayer `json:"player"`
	Position    string                  `json:"position"`
	ShirtNumber int                     `json:"shirt_number"`
}

// FootballTeamSquadResponseData 球隊陣容
type FootballTeamSquadResponseData struct {
	ID        string                  `json:"id"`
	Team      FootballTeamSquadTeam   `json:"team"`
	Squad     []FootballTeamSquadItem `json:"squad"`
	UpdatedAt int64                   `json:"updated_at"`
}

// --- Team squad history ---

// FootballTeamSquadHistoryMatch 歷史陣容比賽資訊
type FootballTeamSquadHistoryMatch struct {
	ID        string `json:"id"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

// FootballTeamSquadHistoryItem 歷史陣容球員
type FootballTeamSquadHistoryItem struct {
	Player      FootballTeamSquadPlayer       `json:"player"`
	Position    string                        `json:"position"`
	ShirtNumber int                           `json:"shirt_number"`
	Match       FootballTeamSquadHistoryMatch `json:"match"`
}

// FootballTeamSquadHistoryResponseData 球隊歷史陣容
type FootballTeamSquadHistoryResponseData struct {
	ID        string                         `json:"id"`
	Team      FootballTeamSquadTeam          `json:"team"`
	Squad     []FootballTeamSquadHistoryItem `json:"squad"`
	UpdatedAt int64                          `json:"updated_at"`
}

// --- Team injury ---

// FootballTeamInjuryPlayer 傷病球員資訊
type FootballTeamInjuryPlayer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// FootballTeamInjuryResponseData 球隊傷病
type FootballTeamInjuryResponseData struct {
	ID        string                   `json:"id"`
	Team      FootballTeamSquadTeam    `json:"team"`
	Player    FootballTeamInjuryPlayer `json:"player"`
	Reason    string                   `json:"reason"`
	StartTime int64                    `json:"start_time"`
	EndTime   int64                    `json:"end_time"`
	UpdatedAt int64                    `json:"updated_at"`
}

// --- Team honor ---

// FootballTeamHonorResponseData 球隊榮譽
type FootballTeamHonorResponseData struct {
	ID        string `json:"id"`
	TeamID    string `json:"team_id"`
	HonorID   string `json:"honor_id"`
	Name      string `json:"name"`
	Count     int    `json:"count"`
	SeasonID  string `json:"season_id"`
	UpdatedAt int64  `json:"updated_at"`
}

// --- Team ability ---

// FootballTeamAbilityResponseData 球隊能力
type FootballTeamAbilityResponseData struct {
	ID        string          `json:"id"`
	TeamID    string          `json:"team_id"`
	Ability   json.RawMessage `json:"ability"`
	UpdatedAt int64           `json:"updated_at"`
}

// --- Team record ---

// FootballTeamRecordItem 球隊對戰記錄
type FootballTeamRecordItem struct {
	TeamID       string `json:"team_id"`
	TeamName     string `json:"team_name"`
	Won          int    `json:"won"`
	Draw         int    `json:"draw"`
	Loss         int    `json:"loss"`
	Goals        int    `json:"goals"`
	GoalsAgainst int    `json:"goals_against"`
}

// FootballTeamRecordResponseData 球隊記錄
type FootballTeamRecordResponseData struct {
	ID        string                   `json:"id"`
	TeamID    string                   `json:"team_id"`
	Record    []FootballTeamRecordItem `json:"record"`
	UpdatedAt int64                    `json:"updated_at"`
}

// --- Player ability ---

// FootballPlayerAbilityResponseData 球員能力
type FootballPlayerAbilityResponseData struct {
	ID        string          `json:"id"`
	PlayerID  string          `json:"player_id"`
	Ability   json.RawMessage `json:"ability"`
	UpdatedAt int64           `json:"updated_at"`
}

// --- Player honor ---

// FootballPlayerHonorResponseData 球員榮譽
type FootballPlayerHonorResponseData struct {
	ID        string `json:"id"`
	PlayerID  string `json:"player_id"`
	HonorID   string `json:"honor_id"`
	Name      string `json:"name"`
	Count     int    `json:"count"`
	SeasonID  string `json:"season_id"`
	UpdatedAt int64  `json:"updated_at"`
}

// --- Player market ---

// FootballPlayerMarketHistoryItem 球員市場價值歷史
type FootballPlayerMarketHistoryItem struct {
	MarketValue int   `json:"market_value"`
	ChangeTime  int64 `json:"change_time"`
}

// FootballPlayerMarketResponseData 球員市場價值
type FootballPlayerMarketResponseData struct {
	ID        string                            `json:"id"`
	History   []FootballPlayerMarketHistoryItem `json:"history"`
	UpdatedAt int64                             `json:"updated_at"`
}

// --- Player transfer ---

// FootballPlayerTransferItem 球員轉會項目
type FootballPlayerTransferItem struct {
	FromTeamID   string `json:"from_team_id"`
	FromTeamName string `json:"from_team_name"`
	ToTeamID     string `json:"to_team_id"`
	ToTeamName   string `json:"to_team_name"`
	TransferType int    `json:"transfer_type"`
	TransferTime int64  `json:"transfer_time"`
	TransferFee  int    `json:"transfer_fee"`
	TransferDesc string `json:"transfer_desc"`
}

// FootballPlayerTransferResponseData 球員轉會
type FootballPlayerTransferResponseData struct {
	ID        string                       `json:"id"`
	Player    FootballTeamSquadPlayer      `json:"player"`
	Transfer  []FootballPlayerTransferItem `json:"transfer"`
	UpdatedAt int64                        `json:"updated_at"`
}

// --- Coach history ---

// FootballCoachHistoryResponseData 教練執教經歷
type FootballCoachHistoryResponseData struct {
	ID        string `json:"id"`
	CoachID   string `json:"coach_id"`
	TeamID    string `json:"team_id"`
	TeamName  string `json:"team_name"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	UpdatedAt int64  `json:"updated_at"`
}

// --- Coach honor ---

// FootballCoachHonorResponseData 教練榮譽
type FootballCoachHonorResponseData struct {
	ID        string `json:"id"`
	CoachID   string `json:"coach_id"`
	HonorID   string `json:"honor_id"`
	Name      string `json:"name"`
	Count     int    `json:"count"`
	SeasonID  string `json:"season_id"`
	UpdatedAt int64  `json:"updated_at"`
}

// --- Competition best lineup ---

// FootballCompetitionBestLineupMember 最佳陣容成員
type FootballCompetitionBestLineupMember struct {
	Player      FootballTeamSquadPlayer `json:"player"`
	Position    string                  `json:"position"`
	ShirtNumber int                     `json:"shirt_number"`
}

// FootballCompetitionBestLineupResponseData 最佳陣容
type FootballCompetitionBestLineupResponseData struct {
	ID        string                                `json:"id"`
	SeasonID  string                                `json:"season_id"`
	RoundNum  int                                   `json:"round_num"`
	Lineup    []FootballCompetitionBestLineupMember `json:"lineup"`
	UpdatedAt int64                                 `json:"updated_at"`
}

// --- Competition best lineup detail ---

// FootballCompetitionBestLineupPlayer 最佳陣容詳情球員
type FootballCompetitionBestLineupPlayer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// FootballCompetitionBestLineupDetailResponseData 最佳陣容詳情
type FootballCompetitionBestLineupDetailResponseData struct {
	ID        string                              `json:"id"`
	LineupID  string                              `json:"lineup_id"`
	Player    FootballCompetitionBestLineupPlayer `json:"player"`
	Position  string                              `json:"position"`
	Points    int                                 `json:"points"`
	UpdatedAt int64                               `json:"updated_at"`
}

// --- Competition record ---

// FootballCompetitionRecordResponseData 賽事紀錄
type FootballCompetitionRecordResponseData struct {
	ID            string `json:"id"`
	CompetitionID string `json:"competition_id"`
	TeamID        string `json:"team_id"`
	TeamName      string `json:"team_name"`
	Won           int    `json:"won"`
	Draw          int    `json:"draw"`
	Loss          int    `json:"loss"`
	Goals         int    `json:"goals"`
	GoalsAgainst  int    `json:"goals_against"`
	UpdatedAt     int64  `json:"updated_at"`
}
