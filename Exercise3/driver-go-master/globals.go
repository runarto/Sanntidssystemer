package main

const (
    MaxOrders       = 12
    True            = true
    False           = false
    On              = true
    Off             = false
    Open            = 1
    Close           = 0
    NotDefined      = -1
    Up              = 1
    Down            = -1
    numFloors       = 4
    numButtons      = 3
    Yes             = true
    No              = false
    Cab            = 1
    Hall           = 0
)

type State int

const (
    Stop State = iota
    Moving
    Still
)

var (
    CurrentState      State
    CurrentDirection  int
    LastDefinedFloor  int
    IsDoorOpen        int
    OrderComplete     bool
    OrderArray        [MaxOrders][3]int
	//gOrderArray er bestående etasje orderen kommer fra/skal til, type knapp, og om ordren kom innvendig fra (true false)
)


