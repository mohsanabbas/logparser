package parser

// PlayerInterface defines the behaviors for Player struct
type PlayerInterface interface {
	GetName() string
	GetKills() int
}

// Player holds player-related data
type Player struct {
	Name  string
	Kills int
}

// GetName returns the name of the player
func (p *Player) GetName() string {
	return p.Name
}

// GetKills returns the number of kills the player has made
func (p *Player) GetKills() int {
	return p.Kills
}
