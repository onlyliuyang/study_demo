package main

const (
	mutexLocked = 1 << iota
	mutexWoken
	mutesStarving
	mutexWaiterShift = iota

	starvationThresholdNs = 1e6
)

func main() {

}
