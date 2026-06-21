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
	debug(client.BasketballCoach(ctx, thesportsgo.BasketballCoachQuery{}))(t, "BasketballCoach")
	debug(client.BasketballVenue(ctx, thesportsgo.BasketballVenueQuery{}))(t, "BasketballVenue")
	debug(client.BasketballSeason(ctx, thesportsgo.BasketballSeasonQuery{}))(t, "BasketballSeason")
	debug(client.BasketballStage(ctx, thesportsgo.BasketballStageQuery{}))(t, "BasketballStage")
}

func TestBasketballMatch(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	debug(client.BasketballMatchRecent(ctx, thesportsgo.BasketballMatchRecentQuery{}))(t, "BasketballMatchRecent")
	debug(client.BasketballSeasonRecentTableDetail(ctx, thesportsgo.BasketballSeasonRecentTableDetailQuery{}))(t, "BasketballSeasonRecentTableDetail")
}

func TestBasketballMatchDetail(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	debug(client.BasketballMatchDetailLive(ctx))(t, "BasketballMatchDetailLive")
	debug(client.BasketballMatchLineupLive(ctx))(t, "BasketballMatchLineupLive")
	debug(client.BasketballMatchLiveHistory(ctx, thesportsgo.BasketballMatchLiveHistoryQuery{}))(t, "BasketballMatchLiveHistory")
}
