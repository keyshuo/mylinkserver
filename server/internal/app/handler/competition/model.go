package netcompete

type Player struct {
	//player name
	username string
	//player prepare status
	prepare bool
	//player connection status
	connection bool
	//player game time
	//gametime
}

type Competetion struct {
	//checkerboard
	checkerboard [][]int
	//the first player
	player1 Player
	//the second player
	player2 Player
}
