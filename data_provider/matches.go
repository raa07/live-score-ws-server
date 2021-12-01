package data_provider

import "sync"

type MatchesStorage struct {
	mx         sync.RWMutex
	serieDatas map[int]SerieData
}

type SerieData struct {
	SelectedMatch int `json:"selected_match"`
	DireWins      int `json:"dire_wins"`
	RadiantWins   int `json:"radiant_wins"`
	Serie         Serie `json:"serie"`
	Matches       map[int]Match `json:"matches"`
}

type Match struct {
	IdMatch        int         `json:"id_match"`
	Buildings      Buildings   `json:"buildings"`
	Result         Result      `json:"result"`
	PlayersDire    []TeamPlayer    `json:"playersDire"`
	PlayersRadiant []TeamPlayer    `json:"playersRadiant"`
	RadiantBans    []SingleBan `json:"radiantBans"`
	DireBans       []SingleBan `json:"direBans"`
}

type TeamPlayer struct {
	IdAcc          int     `json:"id_acc"`
	MainName       string  `json:"main_name"`
	HeroIdExternal int     `json:"hero_id_external,omitempty"`
	HeroImage      string  `json:"hero_image"`
	HeroName       *string `json:"hero_name"`
	HeroRealName   string  `json:"hero_real_name"`
	Items          []Items `json:"items"`
}

type Items struct {
	Name     string `json:"name"`
	RealName string `json:"real_name"`
	Image    string `json:"image"`
}

type Team struct {
	Id         int    `json:"id"`
	IdExternal int    `json:"id_external"`
	Name       string `json:"name"`
	Logo       string `json:"logo"`
}

type Buildings struct {
	Radiant         string `json:"radiant"`
	Dire            string `json:"dire"`
	DireBarracks    string `json:"dire_barracks"`
	RadiantBarracks string `json:"radiant_barracks"`
}

type Hero struct {
	IdExternal string `json:"id_external"`
	Image      string `json:"image"`
	RealName   string `json:"real_name"`
}

type Player struct {
	MainName   string `json:"main_name"`
	IdExternal int    `json:"id_external"`
}

type Serie struct {
	IdExternal    int  `json:"id_external"`
	IsDeactivated int  `json:"is_deactivated"`
	Type          int  `json:"type"`
	IdTeamWinner  int  `json:"id_team_winner"`
	TeamRadiant   Team `json:"teamRadiant"`
	TeamDire      Team `json:"teamDire"`
}

type SingleBan struct {
	HeroLogo string `json:"hero_logo"`
}

type Result struct {
	IdMatch      int    `json:"id_match"`
	Gold         int    `json:"gold"`
	KillsDire    int    `json:"kills_dire"`
	KillsRadiant int    `json:"kills_radiant"`
	Buildings    string `json:"buildings"`
	Barracks     string `json:"barracks"`
	Towers       string `json:"towers"`
	GameTime     string `json:"game_time"`
}

func NewMatches(serieDatas map[int]SerieData) (*MatchesStorage, error) {
	m := &MatchesStorage{
		serieDatas: serieDatas,
	}

	return m, nil
}

func (m *MatchesStorage) Load(idMatch int) (SerieData, bool) {
	m.mx.RLock()
	defer m.mx.RUnlock()
	val, ok := m.serieDatas[idMatch]

	return val, ok
}

func (m *MatchesStorage) Store(idMatch int, serieData SerieData) {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.serieDatas[idMatch] = serieData
}
