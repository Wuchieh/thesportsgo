package thesportsgo_test

import (
	"context"
	"testing"

	"github.com/wuchieh/thesportsgo"
)

func TestFootballBasicIngo(t *testing.T) {
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
