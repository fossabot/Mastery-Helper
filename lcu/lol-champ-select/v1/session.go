package v1

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
)

// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

// GetSession returns champion select session data.
func GetSession(token string, port uint64) *Session {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://127.0.0.1:%d/lol-champ-select/v1/session", port), nil)
	if err != nil {
		return nil
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("riot:%s", token)))))

	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if resp.StatusCode != 200 {
		return nil
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	session, err := UnmarshalSession(bytes)
	if err != nil {
		return nil
	}

	return &session
}

func UnmarshalSession(data []byte) (Session, error) {
	var r Session
	err := jsoniter.Unmarshal(data, &r)
	return r, err
}

type Session struct {
	Actions              [][]Action           `json:"actions"`
	AllowBattleBoost     bool                 `json:"allowBattleBoost"`
	AllowDuplicatePicks  bool                 `json:"allowDuplicatePicks"`
	AllowLockedEvents    bool                 `json:"allowLockedEvents"`
	AllowRerolling       bool                 `json:"allowRerolling"`
	AllowSkinSelection   bool                 `json:"allowSkinSelection"`
	Bans                 Bans                 `json:"bans"`
	BenchChampionIDS     []int                `json:"benchChampionIds"`
	BenchEnabled         bool                 `json:"benchEnabled"`
	BoostableSkinCount   int64                `json:"boostableSkinCount"`
	ChatDetails          ChatDetails          `json:"chatDetails"`
	Counter              int64                `json:"counter"`
	EntitledFeatureState EntitledFeatureState `json:"entitledFeatureState"`
	GameID               int64                `json:"gameId"`
	HasSimultaneousBans  bool                 `json:"hasSimultaneousBans"`
	HasSimultaneousPicks bool                 `json:"hasSimultaneousPicks"`
	IsCustomGame         bool                 `json:"isCustomGame"`
	IsSpectating         bool                 `json:"isSpectating"`
	LocalPlayerCellID    int64                `json:"localPlayerCellId"`
	LockedEventIndex     int64                `json:"lockedEventIndex"`
	MyTeam               []MyTeam             `json:"myTeam"`
	RecoveryCounter      int64                `json:"recoveryCounter"`
	RerollsRemaining     int64                `json:"rerollsRemaining"`
	SkipChampionSelect   bool                 `json:"skipChampionSelect"`
	TheirTeam            []interface{}        `json:"theirTeam"`
	Timer                Timer                `json:"timer"`
	Trades               []interface{}        `json:"trades"`
}

type Action struct {
	ActorCellID  int64  `json:"actorCellId"`
	ChampionID   int64  `json:"championId"`
	Completed    bool   `json:"completed"`
	ID           int64  `json:"id"`
	IsAllyAction bool   `json:"isAllyAction"`
	IsInProgress bool   `json:"isInProgress"`
	PickTurn     int64  `json:"pickTurn"`
	Type         string `json:"type"`
}

type Bans struct {
	MyTeamBans    []interface{} `json:"myTeamBans"`
	NumBans       int64         `json:"numBans"`
	TheirTeamBans []interface{} `json:"theirTeamBans"`
}

type ChatDetails struct {
	ChatRoomName     string `json:"chatRoomName"`
	ChatRoomPassword string `json:"chatRoomPassword"`
}

type EntitledFeatureState struct {
	AdditionalRerolls int64         `json:"additionalRerolls"`
	UnlockedSkinIDS   []interface{} `json:"unlockedSkinIds"`
}

type MyTeam struct {
	AssignedPosition    string `json:"assignedPosition"`
	CellID              int64  `json:"cellId"`
	ChampionID          int64  `json:"championId"`
	ChampionPickIntent  int64  `json:"championPickIntent"`
	EntitledFeatureType string `json:"entitledFeatureType"`
	SelectedSkinID      int64  `json:"selectedSkinId"`
	Spell1ID            int64  `json:"spell1Id"`
	Spell2ID            int64  `json:"spell2Id"`
	SummonerID          int64  `json:"summonerId"`
	Team                int64  `json:"team"`
	WardSkinID          int64  `json:"wardSkinId"`
}

type Timer struct {
	AdjustedTimeLeftInPhase int64  `json:"adjustedTimeLeftInPhase"`
	InternalNowInEpochMS    int64  `json:"internalNowInEpochMs"`
	IsInfinite              bool   `json:"isInfinite"`
	Phase                   string `json:"phase"`
	TotalTimeInPhase        int64  `json:"totalTimeInPhase"`
}