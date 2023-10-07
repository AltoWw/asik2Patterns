package main

import third_party "asik2/thrird_party"

func main() {
	basicSoldier := third_party.BasicSoldier{}
	third_party.DisplayStats(basicSoldier)

	soldierWithSword := SoldierWithSword{soldier: basicSoldier}
	third_party.DisplayStats(soldierWithSword)

	soldierWithShield := SoldierWithShield{soldier: basicSoldier}
	third_party.DisplayStats(soldierWithShield)

	soldierWithShieldWithSword := soldierWithShield{
		soldier: soldierWithSword{
			soldier: basicSoldier,
		},
	}
	third_party.DisplayStats(soldierWithShieldWithSword)
}

type SoldierWithSword struct {
	soldier third_party.InterfaceSoldier
}
