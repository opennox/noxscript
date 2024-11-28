// Code generated by 'yaegi extract github.com/noxworld-dev/noxscript/ns/v4/class'. DO NOT EDIT.

package eval

import (
	"go/constant"
	"go/token"
	"reflect"

	"github.com/opennox/noxscript/ns/v4/class"
)

func init() {
	Symbols["github.com/noxworld-dev/noxscript/ns/v4/class/class"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ARMOR":            reflect.ValueOf(constant.MakeFromLiteral("\"ARMOR\"", token.STRING, 0)),
		"CLIENT_PERSIST":   reflect.ValueOf(constant.MakeFromLiteral("\"CLIENT_PERSIST\"", token.STRING, 0)),
		"CLIENT_PREDICT":   reflect.ValueOf(constant.MakeFromLiteral("\"CLIENT_PREDICT\"", token.STRING, 0)),
		"COMPLEX":          reflect.ValueOf(constant.MakeFromLiteral("\"COMPLEX\"", token.STRING, 0)),
		"DANGEROUS":        reflect.ValueOf(constant.MakeFromLiteral("\"DANGEROUS\"", token.STRING, 0)),
		"DOOR":             reflect.ValueOf(constant.MakeFromLiteral("\"DOOR\"", token.STRING, 0)),
		"ELEVATOR":         reflect.ValueOf(constant.MakeFromLiteral("\"ELEVATOR\"", token.STRING, 0)),
		"ELEVATOR_SHAFT":   reflect.ValueOf(constant.MakeFromLiteral("\"ELEVATOR_SHAFT\"", token.STRING, 0)),
		"EXIT":             reflect.ValueOf(constant.MakeFromLiteral("\"EXIT\"", token.STRING, 0)),
		"FIRE":             reflect.ValueOf(constant.MakeFromLiteral("\"FIRE\"", token.STRING, 0)),
		"FLAG":             reflect.ValueOf(constant.MakeFromLiteral("\"FLAG\"", token.STRING, 0)),
		"FOOD":             reflect.ValueOf(constant.MakeFromLiteral("\"FOOD\"", token.STRING, 0)),
		"HOLE":             reflect.ValueOf(constant.MakeFromLiteral("\"HOLE\"", token.STRING, 0)),
		"IMMOBILE":         reflect.ValueOf(constant.MakeFromLiteral("\"IMMOBILE\"", token.STRING, 0)),
		"INFO_BOOK":        reflect.ValueOf(constant.MakeFromLiteral("\"INFO_BOOK\"", token.STRING, 0)),
		"KEY":              reflect.ValueOf(constant.MakeFromLiteral("\"KEY\"", token.STRING, 0)),
		"LIGHT":            reflect.ValueOf(constant.MakeFromLiteral("\"LIGHT\"", token.STRING, 0)),
		"MISSILE":          reflect.ValueOf(constant.MakeFromLiteral("\"MISSILE\"", token.STRING, 0)),
		"MONSTER":          reflect.ValueOf(constant.MakeFromLiteral("\"MONSTER\"", token.STRING, 0)),
		"MONSTERGENERATOR": reflect.ValueOf(constant.MakeFromLiteral("\"MONSTERGENERATOR\"", token.STRING, 0)),
		"NOT_STACKABLE":    reflect.ValueOf(constant.MakeFromLiteral("\"NOT_STACKABLE\"", token.STRING, 0)),
		"OBSTACLE":         reflect.ValueOf(constant.MakeFromLiteral("\"OBSTACLE\"", token.STRING, 0)),
		"PICKUP":           reflect.ValueOf(constant.MakeFromLiteral("\"PICKUP\"", token.STRING, 0)),
		"PLAYER":           reflect.ValueOf(constant.MakeFromLiteral("\"PLAYER\"", token.STRING, 0)),
		"READABLE":         reflect.ValueOf(constant.MakeFromLiteral("\"READABLE\"", token.STRING, 0)),
		"SIMPLE":           reflect.ValueOf(constant.MakeFromLiteral("\"SIMPLE\"", token.STRING, 0)),
		"TRANSPORTER":      reflect.ValueOf(constant.MakeFromLiteral("\"TRANSPORTER\"", token.STRING, 0)),
		"TREASURE":         reflect.ValueOf(constant.MakeFromLiteral("\"TREASURE\"", token.STRING, 0)),
		"TRIGGER":          reflect.ValueOf(constant.MakeFromLiteral("\"TRIGGER\"", token.STRING, 0)),
		"VISIBLE_ENABLE":   reflect.ValueOf(constant.MakeFromLiteral("\"VISIBLE_ENABLE\"", token.STRING, 0)),
		"WAND":             reflect.ValueOf(constant.MakeFromLiteral("\"WAND\"", token.STRING, 0)),
		"WEAPON":           reflect.ValueOf(constant.MakeFromLiteral("\"WEAPON\"", token.STRING, 0)),

		// type definitions
		"Class": reflect.ValueOf((*class.Class)(nil)),
	}
}
