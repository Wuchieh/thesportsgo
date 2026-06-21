package thesportsgo_test

import (
	"context"
	"testing"

	"github.com/wuchieh/thesportsgo"
)

func TestFootballBasicInfo(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	debug(client.FootballCategory(ctx))(t, "FootballCategory")
	debug(client.FootballCountry(ctx))(t, "FootballCountry")
	debug(client.FootballCompetition(ctx, thesportsgo.FootballCompetitionQuery{}))(t, "FootballCompetition")
	debug(client.FootballTeam(ctx, thesportsgo.FootballTeamQuery{}))(t, "FootballTeam")
	debug(client.FootballPlayer(ctx, thesportsgo.FootballPlayerQuery{}))(t, "FootballPlayer")
	debug(client.FootballCoach(ctx, thesportsgo.FootballCoachQuery{}))(t, "FootballCoach")
	debug(client.FootballReferee(ctx, thesportsgo.FootballRefereeQuery{}))(t, "FootballReferee")
	debug(client.FootballVenue(ctx, thesportsgo.FootballVenueQuery{}))(t, "FootballVenue")
	debug(client.FootballSeason(ctx, thesportsgo.FootballSeasonQuery{}))(t, "FootballSeason")
	debug(client.FootballStage(ctx, thesportsgo.FootballStageQuery{}))(t, "FootballStage")
}

func TestFootballMatch(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	debug(client.FootballMatchList(ctx, thesportsgo.FootballMatchQuery{}))(t, "FootballMatchList")
	debug(client.FootballMatchRecentList(ctx, thesportsgo.FootballMatchQuery{}))(t, "FootballMatchRecentList")
	debug(client.FootballMatchDetailLive(ctx))(t, "FootballMatchDetailLive")
	debug(client.FootballMatchTrendLive(ctx))(t, "FootballMatchTrendLive")
	debug(client.FootballMatchPlayerStatsList(ctx))(t, "FootballMatchPlayerStatsList")
	debug(client.FootballMatchTeamStatsList(ctx))(t, "FootballMatchTeamStatsList")
	debug(client.FootballMatchHalfTeamStatsList(ctx))(t, "FootballMatchHalfTeamStatsList")
	debug(client.FootballMatchTVList(ctx, thesportsgo.FootballMatchQuery{}))(t, "FootballMatchTVList")
}

func TestFootballSeason(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	q := thesportsgo.FootballSeasonUUIDQuery{UUID: "test"}
	debug(client.FootballSeasonRecentTable(ctx, q))(t, "FootballSeasonRecentTable")
	debug(client.FootballSeasonTable(ctx, q))(t, "FootballSeasonTable")
	debug(client.FootballSeasonTeamStat(ctx, q))(t, "FootballSeasonTeamStat")
	debug(client.FootballSeasonPlayerStat(ctx, q))(t, "FootballSeasonPlayerStat")
	debug(client.FootballSeasonShooterStat(ctx, q))(t, "FootballSeasonShooterStat")
	debug(client.FootballSeasonRecentTeamStat(ctx, q))(t, "FootballSeasonRecentTeamStat")
	debug(client.FootballSeasonRecentPlayerStat(ctx, q))(t, "FootballSeasonRecentPlayerStat")
	debug(client.FootballSeasonRecentShooterStat(ctx, q))(t, "FootballSeasonRecentShooterStat")
}

func TestFootballOther(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	debug(client.FootballTableLive(ctx))(t, "FootballTableLive")
	debug(client.FootballBracketSeason(ctx, thesportsgo.FootballSeasonUUIDQuery{UUID: "test"}))(t, "FootballBracketSeason")
	debug(client.FootballRankingFifaMen(ctx, thesportsgo.FootballRankingFifaQuery{}))(t, "FootballRankingFifaMen")
	debug(client.FootballRankingFifaWomen(ctx, thesportsgo.FootballRankingFifaQuery{}))(t, "FootballRankingFifaWomen")
	debug(client.FootballRankingClub(ctx))(t, "FootballRankingClub")
	debug(client.FootballRankingFifaLive(ctx))(t, "FootballRankingFifaLive")
	debug(client.FootballDeleted(ctx))(t, "FootballDeleted")
	debug(client.FootballDataUpdate(ctx))(t, "FootballDataUpdate")
	debug(client.FootballLanguage(ctx))(t, "FootballLanguage")
	debug(client.FootballHonor(ctx))(t, "FootballHonor")
	debug(client.FootballCompensation(ctx))(t, "FootballCompensation")
	debug(client.FootballOddsLive(ctx))(t, "FootballOddsLive")
	debug(client.FootballOddsHistory(ctx, thesportsgo.FootballOddsHistoryQuery{UUID: "test"}))(t, "FootballOddsHistory")
	debug(client.FootballOddsUpdate(ctx))(t, "FootballOddsUpdate")
}

func TestFootballTeamPlayerCoach(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	pq := thesportsgo.FootballPageTimeUUIDQuery{}

	debug(client.FootballTeamSquad(ctx, pq))(t, "FootballTeamSquad")
	debug(client.FootballTeamSquadHistory(ctx, thesportsgo.FootballTeamSquadHistoryQuery{UUID: "test"}))(t, "FootballTeamSquadHistory")
	debug(client.FootballTeamInjury(ctx, pq))(t, "FootballTeamInjury")
	debug(client.FootballTeamHonor(ctx, pq))(t, "FootballTeamHonor")
	debug(client.FootballTeamAbility(ctx, pq))(t, "FootballTeamAbility")
	debug(client.FootballTeamRecord(ctx, pq))(t, "FootballTeamRecord")

	debug(client.FootballPlayerAbility(ctx, pq))(t, "FootballPlayerAbility")
	debug(client.FootballPlayerHonor(ctx, pq))(t, "FootballPlayerHonor")
	debug(client.FootballPlayerMarket(ctx, pq))(t, "FootballPlayerMarket")
	debug(client.FootballPlayerTransfer(ctx, pq))(t, "FootballPlayerTransfer")

	debug(client.FootballCoachHistory(ctx, pq))(t, "FootballCoachHistory")
	debug(client.FootballCoachHonor(ctx, pq))(t, "FootballCoachHonor")

	debug(client.FootballCompetitionBestLineup(ctx, pq))(t, "FootballCompetitionBestLineup")
	debug(client.FootballCompetitionBestLineupDetail(ctx, thesportsgo.FootballSeasonUUIDQuery{UUID: "test"}))(t, "FootballCompetitionBestLineupDetail")
	debug(client.FootballCompetitionRecord(ctx, pq))(t, "FootballCompetitionRecord")
}
