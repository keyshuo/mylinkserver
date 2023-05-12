package netcompete

type Player struct {
	//player name
	name string
	//player prepare status
	prepare bool
	//player connection status
	connection bool
}

type Competetion struct {
	//checkerboard
	checkerboard [][]int
	//the first player
	player1 Player
	//the second player
	player2 Player
}
