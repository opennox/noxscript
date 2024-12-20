package ns

// CreateGold creates gold item with a given amount of gold.
func CreateGold(pos Positioner, amount int) Obj {
	g := CreateObject("Gold", pos)
	g.SetGold(amount)
	return g
}

// CreateGoldChest creates gold chest item with a given amount of gold.
func CreateGoldChest(pos Positioner, amount int) Obj {
	g := CreateObject("QuestGoldChest", pos)
	g.SetGold(amount)
	return g
}

// CreateGoldPile creates gold pile item with a given amount of gold.
func CreateGoldPile(pos Positioner, amount int) Obj {
	g := CreateObject("QuestGoldPile", pos)
	g.SetGold(amount)
	return g
}
