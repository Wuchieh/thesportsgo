package thesportsgo_test

import (
	"context"
	"testing"

	"github.com/wuchieh/thesportsgo"
)

func TestBasketballBasicInfo(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	debug(client.BasketballCategory(ctx))(t, "BasketballCategory")
	debug(client.BasketballCountry(ctx))(t, "BasketballCountry")
	debug(client.BasketballCompetition(ctx, thesportsgo.BasketballCompetitionQuery{}))(t, "BasketballCompetition")
	debug(client.BasketballTeam(ctx, thesportsgo.BasketballTeamQuery{}))(t, "BasketballTeam")
	debug(client.BasketballPlayer(ctx, thesportsgo.BasketballPlayerQuery{}))(t, "BasketballPlayer")
	debug(client.BasketballVenue(ctx, thesportsgo.BasketballVenueQuery{}))(t, "BasketballVenue")
	debug(client.BasketballSeason(ctx, thesportsgo.BasketballSeasonQuery{}))(t, "BasketballSeason")
	debug(client.BasketballStage(ctx, thesportsgo.BasketballStageQuery{}))(t, "BasketballStage")
}

func TestBasketballMatch(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	debug(client.BasketballMatch(ctx, thesportsgo.BasketballMatchQuery{}))(t, "BasketballMatch")
	debug(client.BasketballStandings(ctx, thesportsgo.BasketballStandingsQuery{}))(t, "BasketballStandings")
}

func TestBasketballMatchDetail(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	debug(client.BasketballMatchDetail(ctx, thesportsgo.BasketballMatchDetailQuery{}))(t, "BasketballMatchDetail")
	debug(client.BasketballMatchLineup(ctx, thesportsgo.BasketballMatchLineupQuery{}))(t, "BasketballMatchLineup")
	debug(client.BasketballMatchStatistics(ctx, thesportsgo.BasketballMatchStatisticsQuery{}))(t, "BasketballMatchStatistics")
}
