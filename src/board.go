package main

//Amount per player
const jadeAmount = 6
const roseAmount = 6
const chrysantherAmount = 6
const rhododendronAmount = 6
const jasmineAmount = 6
const lilyAmount = 6
const lotusAmount = 3
const dragonAmount = 3
const wheelAmount = 3
const boatAmount = 3
const rockAmount = 3
const knotweedAmount = 3

var board [256]stone
var firstPlayerStones [54]stone
var secondPlayerStones [54]stone

func initStones() {
	for jad := 0; jad < jadeAmount; jad++ {
		firstPlayerStones[jad] = stone{0, "White Jade", true}
		secondPlayerStones[jad] = stone{0, "White Jade", false}
	}

	for ros := 0; ros < roseAmount; ros++ {
		firstPlayerStones[ros] = stone{1, "Rose", true}
		secondPlayerStones[ros] = stone{1, "Rose", false}
	}

	for chr := 0; chr < chrysantherAmount; chr++ {
		firstPlayerStones[chr] = stone{2, "Chrysanther", true}
		secondPlayerStones[chr] = stone{2, "Chrysanther", false}
	}

	for rho := 0; rho < rhododendronAmount; rho++ {
		firstPlayerStones[rho] = stone{3, "Rhododendron", true}
		secondPlayerStones[rho] = stone{3, "Rhododendron", false}
	}

	for jas := 0; jas < jasmineAmount; jas++ {
		firstPlayerStones[jas] = stone{4, "Jasmin", true}
		secondPlayerStones[jas] = stone{4, "Jasmin", false}
	}

	for lil := 0; lil < lilyAmount; lil++ {
		firstPlayerStones[lil] = stone{5, "White Lily", true}
		secondPlayerStones[lil] = stone{5, "White Lily", false}
	}

	for lot := 0; lot < lotusAmount; lot++ {
		firstPlayerStones[lot] = stone{6, "White Lotus", true}
		secondPlayerStones[lot] = stone{6, "White Lotus", false}
	}

	for dra := 0; dra < dragonAmount; dra++ {
		firstPlayerStones[dra] = stone{7, "White Dragon", true}
		secondPlayerStones[dra] = stone{7, "White Dragon", false}
	}

	for whe := 0; whe < wheelAmount; whe++ {
		firstPlayerStones[whe] = stone{8, "Wheel", true}
		secondPlayerStones[whe] = stone{8, "Wheel", false}
	}

	for boa := 0; boa < boatAmount; boa++ {
		firstPlayerStones[boa] = stone{9, "Boat", true}
		secondPlayerStones[boa] = stone{9, "Boat", false}
	}

	for roc := 0; roc < rockAmount; roc++ {
		firstPlayerStones[roc] = stone{10, "Rock", true}
		secondPlayerStones[roc] = stone{10, "Rock", false}
	}

	for kno := 0; kno < knotweedAmount; kno++ {
		firstPlayerStones[kno] = stone{11, "Knotweed", true}
		secondPlayerStones[kno] = stone{11, "Knotweed", false}
	}
}
