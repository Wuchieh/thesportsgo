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
	debug(client.BaseballCompetition(ctx, thesportsgo.BaseballCompetitionQuery{}))(t, "BaseballCompetition")
	debug(client.BaseballTeam(ctx, thesportsgo.BaseballTeamQuery{}))(t, "BaseballTeam")
	debug(client.BaseballPlayer(ctx, thesportsgo.BaseballPlayerQuery{}))(t, "BaseballPlayer")
	debug(client.BaseballCoach(ctx, thesportsgo.BaseballCoachQuery{}))(t, "BaseballCoach")
	debug(client.BaseballReferee(ctx, thesportsgo.BaseballRefereeQuery{}))(t, "BaseballReferee")
	debug(client.BaseballVenue(ctx, thesportsgo.BaseballVenueQuery{}))(t, "BaseballVenue")
	debug(client.BaseballSeason(ctx, thesportsgo.BaseballSeasonQuery{}))(t, "BaseballSeason")
	debug(client.BaseballStage(ctx, thesportsgo.BaseballStageQuery{}))(t, "BaseballStage")
	debug(client.BaseballMatchList(ctx, thesportsgo.BaseballMatchListQuery{}))(t, "BaseballMatchList")
	debug(client.BaseballMatchDetail(ctx, thesportsgo.BaseballMatchDetailQuery{}))(t, "BaseballMatchDetail")
	debug(client.BaseballStandings(ctx, thesportsgo.BaseballStandingsQuery{}))(t, "BaseballStandings")
}
