package thesportsgo_test

import (
	"context"
	"testing"

	"github.com/wuchieh/thesportsgo"
)

func TestBaseballBasicInfo(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	debug(client.BaseballCategory(ctx))(t, "BaseballCategory")
	debug(client.BaseballCountry(ctx))(t, "BaseballCountry")
	debug(client.BaseballVenue(ctx, thesportsgo.BaseballVenueQuery{}))(t, "BaseballVenue")
	debug(client.BaseballSeason(ctx, thesportsgo.BaseballSeasonQuery{}))(t, "BaseballSeason")
}

func TestBaseballCompetition(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	debug(client.BaseballUniqueTournament(ctx, thesportsgo.BaseballUniqueTournamentQuery{}))(t, "BaseballUniqueTournament")
	debug(client.BaseballTeam(ctx, thesportsgo.BaseballTeamQuery{}))(t, "BaseballTeam")
	debug(client.BaseballPlayer(ctx, thesportsgo.BaseballPlayerQuery{}))(t, "BaseballPlayer")
}

func TestBaseballMatch(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	debug(client.BaseballMatchList(ctx, thesportsgo.BaseballMatchListQuery{}))(t, "BaseballMatchList")
	debug(client.BaseballMatchDetailLive(ctx))(t, "BaseballMatchDetailLive")
	debug(client.BaseballMatchDiary(ctx, thesportsgo.BaseballMatchDiaryQuery{}))(t, "BaseballMatchDiary")
	debug(client.BaseballMatchSeason(ctx, thesportsgo.BaseballMatchSeasonQuery{}))(t, "BaseballMatchSeason")
	debug(client.BaseballMatchLiveHistory(ctx, thesportsgo.BaseballMatchLiveHistoryQuery{}))(t, "BaseballMatchLiveHistory")
}

func TestBaseballSeasonStats(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	debug(client.BaseballSeasonTable(ctx, thesportsgo.BaseballSeasonTableQuery{}))(t, "BaseballSeasonTable")
	debug(client.BaseballSeasonTeamStats(ctx, thesportsgo.BaseballSeasonTeamStatsQuery{}))(t, "BaseballSeasonTeamStats")
	debug(client.BaseballSeasonPlayerStats(ctx, thesportsgo.BaseballSeasonPlayerStatsQuery{}))(t, "BaseballSeasonPlayerStats")
	debug(client.BaseballSeasonCoachStats(ctx, thesportsgo.BaseballSeasonCoachStatsQuery{}))(t, "BaseballSeasonCoachStats")
}

func TestBaseballSquadAndInjury(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	debug(client.BaseballTeamSquad(ctx, thesportsgo.BaseballTeamSquadQuery{}))(t, "BaseballTeamSquad")
	debug(client.BaseballTeamInjury(ctx, thesportsgo.BaseballTeamInjuryQuery{}))(t, "BaseballTeamInjury")
}

func TestBaseballHonor(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	debug(client.BaseballHonor(ctx, thesportsgo.BaseballHonorQuery{}))(t, "BaseballHonor")
	debug(client.BaseballTeamHonor(ctx, thesportsgo.BaseballTeamHonorQuery{}))(t, "BaseballTeamHonor")
	debug(client.BaseballPlayerHonor(ctx, thesportsgo.BaseballPlayerHonorQuery{}))(t, "BaseballPlayerHonor")
	debug(client.BaseballCoachHonor(ctx, thesportsgo.BaseballCoachHonorQuery{}))(t, "BaseballCoachHonor")
}

func TestBaseballOther(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	debug(client.BaseballBracketSeason(ctx, thesportsgo.BaseballBracketSeasonQuery{}))(t, "BaseballBracketSeason")
	debug(client.BaseballDataUpdate(ctx))(t, "BaseballDataUpdate")
	debug(client.BaseballDeleted(ctx))(t, "BaseballDeleted")
}

func TestBaseballOdds(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	debug(client.BaseballOddsLive(ctx))(t, "BaseballOddsLive")
	debug(client.BaseballOddsHistory(ctx, thesportsgo.BaseballOddsHistoryQuery{}))(t, "BaseballOddsHistory")
}
