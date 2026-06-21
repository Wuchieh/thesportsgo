package thesportsgo_test

import (
	"context"
	"testing"

	"github.com/wuchieh/thesportsgo"
)

// TestEsportsBasicInfo 測試電競基本資訊端點
func TestEsportsBasicInfo(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	debug(client.EsportsCountry(ctx))(t, "EsportsCountry")
	debug(client.EsportsTournament(ctx, thesportsgo.EsportsTournamentQuery{}))(t, "EsportsTournament")
	debug(client.EsportsTeam(ctx, thesportsgo.EsportsTeamQuery{}))(t, "EsportsTeam")
	debug(client.EsportsPlayer(ctx, thesportsgo.EsportsPlayerQuery{}))(t, "EsportsPlayer")
	debug(client.EsportsMatch(ctx, thesportsgo.EsportsMatchQuery{}))(t, "EsportsMatch")
	debug(client.EsportsMatchDetail(ctx, thesportsgo.EsportsMatchDetailQuery{}))(t, "EsportsMatchDetail")
	debug(client.EsportsMatchLineup(ctx, thesportsgo.EsportsMatchLineupQuery{}))(t, "EsportsMatchLineup")
	debug(client.EsportsStandings(ctx, thesportsgo.EsportsStandingsQuery{}))(t, "EsportsStandings")
}
