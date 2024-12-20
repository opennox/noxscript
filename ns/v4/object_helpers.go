package ns

import (
	"github.com/opennox/noxscript/ns/v4/abil"
	"github.com/opennox/noxscript/ns/v4/enchant"
	"github.com/opennox/noxscript/ns/v4/spell"
)

// CreateGold creates gold item with a given amount of gold.
func CreateGold(pos Positioner, amount int) Obj {
	return NewGold(pos, amount)
}

// NewGold creates gold item with a given amount of gold.
func NewGold(pos Positioner, amount int) Obj {
	g := CreateObject("Gold", pos)
	g.SetGold(amount)
	return g
}

// CreateGoldChest creates gold chest item with a given amount of gold.
func CreateGoldChest(pos Positioner, amount int) Obj {
	return NewGoldChest(pos, amount)
}

// NewGoldChest creates gold chest item with a given amount of gold.
func NewGoldChest(pos Positioner, amount int) Obj {
	g := CreateObject("QuestGoldChest", pos)
	g.SetGold(amount)
	return g
}

// CreateGoldPile creates gold pile item with a given amount of gold.
func CreateGoldPile(pos Positioner, amount int) Obj {
	return NewGoldPile(pos, amount)
}

// NewGoldPile creates gold pile item with a given amount of gold.
func NewGoldPile(pos Positioner, amount int) Obj {
	g := CreateObject("QuestGoldPile", pos)
	g.SetGold(amount)
	return g
}

// NewHealthPotion creates a health potion that restores a given amount.
func NewHealthPotion(pos Positioner, amount int) Obj {
	if impl == nil {
		return nil
	}
	return impl.NewHealthPotion(pos, amount)
}

// NewManaPotion creates a mana potion that restores a given amount.
func NewManaPotion(pos Positioner, amount int) Obj {
	if impl == nil {
		return nil
	}
	return impl.NewManaPotion(pos, amount)
}

// NewSpellBook creates a spell book.
func NewSpellBook(pos Positioner, spell spell.Spell) Obj {
	if impl == nil {
		return nil
	}
	return impl.NewSpellBook(pos, spell)
}

// NewAbilityBook creates an ability book.
func NewAbilityBook(pos Positioner, abil abil.Ability) Obj {
	if impl == nil {
		return nil
	}
	return impl.NewAbilityBook(pos, abil)
}

// NewFieldGuide creates a field guide (beast scroll) for a given creature.
func NewFieldGuide(pos Positioner, creature ObjTypeName) Obj {
	if impl == nil {
		return nil
	}
	return impl.NewFieldGuide(pos, creature)
}

// NewEnchantUseItem creates an item that applies an enchant on use.
func NewEnchantUseItem(typ ObjTypeName, pos Positioner, enc enchant.Enchant, dur Duration) Obj {
	if impl == nil {
		return nil
	}
	return impl.NewEnchantUseItem(typ, pos, enc, dur)
}

// NewSpellUseItem creates an item that casts a spell on use.
func NewSpellUseItem(typ ObjTypeName, pos Positioner, spell spell.Spell, lvl int) Obj {
	if impl == nil {
		return nil
	}
	return impl.NewSpellUseItem(typ, pos, spell, lvl)
}
