package thesportsgo_test

import (
	"context"
	"testing"

	"github.com/wuchieh/thesportsgo"
)

// TestCsgoBasicData 測試 CS:GO 基本資料端點
func TestCsgoBasicData(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	q := thesportsgo.CsgoListQuery{}

	debug(client.CsgoCountry(ctx, q))(t, "CsgoCountry")
	debug(client.CsgoTournament(ctx, q))(t, "CsgoTournament")
	debug(client.CsgoTeam(ctx, q))(t, "CsgoTeam")
	debug(client.CsgoPlayer(ctx, q))(t, "CsgoPlayer")
	debug(client.CsgoStage(ctx, q))(t, "CsgoStage")
}

// TestCsgoMatch 測試 CS:GO 比賽相關端點
func TestCsgoMatch(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	q := thesportsgo.CsgoListQuery{}

	debug(client.CsgoMatch(ctx, q))(t, "CsgoMatch")
	debug(client.CsgoMatchSingle(ctx, q))(t, "CsgoMatchSingle")
	debug(client.CsgoMatchSinglePlayerStat(ctx, q))(t, "CsgoMatchSinglePlayerStat")
}

// TestCsgoRealtime 測試 CS:GO 即時端點
func TestCsgoRealtime(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	debug(client.CsgoMatchDetailLive(ctx))(t, "CsgoMatchDetailLive")
	debug(client.CsgoDelete(ctx))(t, "CsgoDelete")
	debug(client.CsgoOddsLive(ctx))(t, "CsgoOddsLive")
}

// TestCsgoSchedule 測試 CS:GO 賽程相關端點
func TestCsgoSchedule(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	// 賽程日誌會返回自訂結構（含 results_extra）
	diaryResp, err := client.CsgoMatchDiary(ctx, thesportsgo.CsgoMatchDiaryQuery{})
	if err != nil {
		t.Log("CsgoMatchDiary error (可能缺少 tsp 參數):", err)
	} else {
		debug(diaryResp, nil)(t, "CsgoMatchDiary")
	}
}

// TestCsgoStats 測試 CS:GO 統計端點
func TestCsgoStats(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	q := thesportsgo.CsgoListQuery{}

	debug(client.CsgoTeamStat(ctx, q))(t, "CsgoTeamStat")
	debug(client.CsgoPlayerStat(ctx, q))(t, "CsgoPlayerStat")
}

// TestCsgoTournamentExtras 測試 CS:GO 賽事附加端點
func TestCsgoTournamentExtras(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	q := thesportsgo.CsgoListQuery{}

	debug(client.CsgoTournamentTeam(ctx, q))(t, "CsgoTournamentTeam")
	debug(client.CsgoTournamentRound(ctx, q))(t, "CsgoTournamentRound")
	debug(client.CsgoTournamentBracket(ctx, q))(t, "CsgoTournamentBracket")
	debug(client.CsgoTournamentBracketLine(ctx, q))(t, "CsgoTournamentBracketLine")
	debug(client.CsgoTournamentTable(ctx, q))(t, "CsgoTournamentTable")
}

// TestDota2BasicData 測試 Dota2 基本資料端點
func TestDota2BasicData(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	q := thesportsgo.CsgoListQuery{}

	debug(client.Dota2Hero(ctx, q))(t, "Dota2Hero")
	debug(client.Dota2Equipment(ctx, q))(t, "Dota2Equipment")
}

// TestLolBasicData 測試 LoL 基本資料端點
func TestLolBasicData(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	q := thesportsgo.CsgoListQuery{}

	debug(client.LolSpell(ctx, q))(t, "LolSpell")
	debug(client.LolRune(ctx, q))(t, "LolRune")
	debug(client.LolMatchEvent(ctx, q))(t, "LolMatchEvent")
}
