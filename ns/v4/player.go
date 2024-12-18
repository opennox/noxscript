package ns

import (
	"github.com/opennox/libs/player"
	"github.com/opennox/libs/types"
)

// PlayerJoinFunc is a callback type for OnPlayerJoin.
type PlayerJoinFunc func(p Player) bool

// PlayerLeaveFunc is a callback type for OnPlayerLeave.
type PlayerLeaveFunc func(p Player)

// PlayerDeathFunc is a callback type for OnPlayerDeath.
type PlayerDeathFunc func(p Player, killer Obj)

// OnPlayerJoin registers a callback function that will be called when player joins the server.
func OnPlayerJoin(fnc PlayerJoinFunc) {
	if impl == nil {
		return
	}
	impl.OnPlayerJoin(fnc)
}

// OnPlayerLeave registers a callback function that will be called when player leaves the server.
func OnPlayerLeave(fnc PlayerLeaveFunc) {
	if impl == nil {
		return
	}
	impl.OnPlayerLeave(fnc)
}

// OnPlayerDeath registers a callback function that will be called when player dies.
func OnPlayerDeath(fnc PlayerDeathFunc) {
	if impl == nil {
		return
	}
	impl.OnPlayerDeath(fnc)
}

type Player interface {
	// Name returns player's name.
	Name() string
	// Unit returns player's character unit.
	Unit() Obj
	// GetScore gets player's score.
	GetScore() int
	// ChangeScore adds or removes player's score(s).
	ChangeScore(score int)
	// Print displays a localized string on the screen of the player.
	//
	// If the string is not in the string database, it will instead print an error message with "MISSING:".
	Print(message StringID)
	// PrintStr displays a string on the screen of the player. It does not localize the string.
	PrintStr(message string)
	// Blind or unblind the player.
	Blind(blind bool)
	// HasTeam checks if a player belongs to a team.
	HasTeam(t Team) bool
	// Team returns current team of a player, if any.
	Team() Team
	// CursorPos returns aim cursor position for the player.
	CursorPos() types.Pointf
	// Store returns a player storage with a given type.
	Store(typ StorageType) Storage
}

// GetHost gets host's player object.
func GetHost() Obj {
	if impl == nil {
		return nil
	}
	return impl.GetHost()
}

// HostPlayer gets host player.
func HostPlayer() Player {
	if impl == nil {
		return nil
	}
	return impl.HostPlayer()
}

// Players returns a list of active players.
func Players() []Player {
	if impl == nil {
		return nil
	}
	return impl.Players()
}

// GetCharacterData gets information about the loaded character.
func GetCharacterData(field int) int {
	if impl == nil {
		return 0
	}
	return impl.GetCharacterData(field)
}

// StringID is an ID of a localized string in the string database.
type StringID = string

// Print displays a localized string on the screen of Other.
//
// If the string is not in the string database, it will instead print an error message with "MISSING:".
func Print(message StringID) {
	if impl == nil {
		return
	}
	impl.Print(message)
}

// PrintStr displays a string on the screen of Other. It does not localize the string.
func PrintStr(message string) {
	if impl == nil {
		return
	}
	impl.PrintStr(message)
}

// PrintToAll displays a localized string to everyone.
//
// If the string is not in the string database, it will instead print an error message with "MISSING:".
func PrintToAll(message StringID) {
	if impl == nil {
		return
	}
	impl.PrintToAll(message)
}

// PrintStrToAll displays a string to everyone. It does not localize the string.
func PrintStrToAll(message string) {
	if impl == nil {
		return
	}
	impl.PrintStrToAll(message)
}

// ClearMessages clears messages on player's screen.
func ClearMessages(player Obj) {
	if impl == nil {
		return
	}
	impl.ClearMessages(player)
}

// UnBlind the host.
func UnBlind() {
	if impl == nil {
		return
	}
	impl.UnBlind()
}

// Blind the host.
func Blind() {
	if impl == nil {
		return
	}
	impl.Blind()
}

// WideScreen enables or disables cinematic wide-screen effect.
func WideScreen(enable bool) {
	if impl == nil {
		return
	}
	impl.WideScreen(enable)
}

// IsTalking gets whether host is talking.
func IsTalking() bool {
	if impl == nil {
		return false
	}
	return impl.IsTalking()
}

// IsTrading returns whether the host is currently talking to shopkeeper.
func IsTrading() bool {
	if impl == nil {
		return false
	}
	return impl.IsTrading()
}

type HalberdLevel int

const (
	OblivionHalberd   = HalberdLevel(0)
	OblivionHeart     = HalberdLevel(1)
	OblivionWierdling = HalberdLevel(2)
	OblivionOrb       = HalberdLevel(3)
)

// SetHalberd upgrades host's oblivion staff.
func SetHalberd(upgrade HalberdLevel) {
	if impl == nil {
		return
	}
	impl.SetHalberd(upgrade)
}

// ImmediateBlind immediately blinds the host.
func ImmediateBlind() {
	if impl == nil {
		return
	}
	impl.ImmediateBlind()
}

// EndGame end of game for a specific class.
func EndGame(class player.Class) {
	if impl == nil {
		return
	}
	impl.EndGame(class)
}
