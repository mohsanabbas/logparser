package parser

// GameInterface defines the behaviors for Game struct
type GameInterface interface {
	GetID() int
	GetTotalKills() int
	GetPlayers() map[string]PlayerInterface
	GetKillsByMeans() map[string]int
}

// Game holds game-related data
type Game struct {
	ID           int
	TotalKills   int
	Players      map[string]*Player
	KillsByMeans map[string]int
}

// GetID returns the ID of the game
func (g *Game) GetID() int {
	return g.ID
}

// GetTotalKills returns the total number of kills in the game
func (g *Game) GetTotalKills() int {
	return g.TotalKills
}

// GetPlayers returns players interface type
func (g *Game) GetPlayers() map[string]PlayerInterface {
	players := make(map[string]PlayerInterface)
	for name, player := range g.Players {
		players[name] = player
	}
	return players
}

// GetKillsByMeans returns a copy of the kill counts by means
func (g *Game) GetKillsByMeans() map[string]int {
	killsByMeans := make(map[string]int, len(g.KillsByMeans))
	for k, v := range g.KillsByMeans {
		killsByMeans[k] = v
	}
	return killsByMeans
}
