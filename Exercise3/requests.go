package main

import "globals.go"
import "fmt"
import "elevio"

func initializeQueue() {
    for i := range globals.OrderArray {
        globals.OrderArray[i][0] = globals.NotDefined
        globals.OrderArray[i][1] = globals.NotDefined
        globals.OrderArray[i][2] = globals.NotDefined
    }
}


func printOrderArray() {
    for i := 0; i < globals.MaxOrders; i++ {
        for j := 0; j < 3; j++ {
            fmt.Println(globals.OrderArray[i][j])
        }
    }
}

// fromFloor: etasje input kommer fra, button: type (opp, ned, cab)

func addToQueueFromFloorPanel(fromFloor int, button ButtonType) {
    for i := 0; i < MaxOrders; i++ {
        if globals.OrderArray[i][0] == fromFloor && globals.OrderArray[i][1] == int(button) 
		&& globals.OrderArray[i][2] == globals.NotDefined {
            return //Ordre kommer fra heisen
        }

        if globals.OrderArray[i][0] == globals.NotDefined && globals.OrderArray[i][1] == globals.NotDefined
		 && globals.OrderArray[i][2] == globals.NotDefined {
            gOrderArray[i][0] = fromFloor
            gOrderArray[i][1] = int(button)
            return
        }
    }
}

func addToQueueCab(toFloor int) {
    for i := 0; i < globals.MaxOrders; i++ {
        if globals.OrderArray[i][0] == toFloor && globals.OrderArray[i][2] == true {
            fmt.Println("Check1")
            return
        }

        if globals.OrderArray[i][0] == globals.NotDefined && globals.OrderArray[i][1] == globals.NotDefined
		 && globals.OrderArray[i][2] == globals.NotDefined {
            OrderArray[i][0] = toFloor
            fmt.Println("Check2")

            if globals.LastDefinedFloor < toFloor || (globals.LastDefinedFloor == toFloor && globals.CurrentDirection == globals.Down) {
                gOrderArray[i][1] = globals.Up
                gOrderArray[i][2] = true
                fmt.Println("Order added successfully")
            }

            if globals.LastDefinedFloor > toFloor || (globals.LastDefinedFloor == toFloor && globals.CurrentDirection == globals.Up) {
                gOrderArray[i][1] = globals.Down
                gOrderArray[i][2] = true
                fmt.Println("Order added successfully")
            }
        }
    }
}

func checkNewOrders() {
	for f := 0; f < elevio._numFloors; f++ {
		for b := 0; b < elevio._numButtons; b++ {
			bool btnPressed = elevio.GetButton(b,f)
			if btnPressed == true && f != elevio.GetFloor() {
				fmt.Println("Button was pressed");
				elevio.SetButtonLamp(b, f, globals.On)
				if b == elevio.BT_Cab {
					addToQueueCab(f)
					fmt.Println("Order added from inside cab")
				} else {
					addToQueueFromFloorPanel(f, b)
					fmt.Println("Order added from outside elevator")
				}
			}
		}
	}
}

func orderCompleteCheck() int {
	OrderComplete := 0
	int currentFloor = elevio.GetFloor()

	for i := 0; i < globals.MaxOrders; i++ {
		if currentFloor != -1 && currentFloor == globals.OrderArray[i][0] {
			if globals.OrderArray == globals.Up && globals.LastDefinedFloor < globals.OrderArray[i][0] {
				OrderComplete = removeOrdersAtFloor(globals.OrderArray[i][0])
			} 
			else if globals.OrderArray == globals.Down && globals.LastDefinedFloor > globals.OrderArray[i][0] {
				OrderComplete = removeOrdersAtFloor(g_OrderArray[i][0])
				break
			}
		}
	}
	return OrderComplete
}


func removeOrdersAtFloor(floor int) int {
	OrderComplete := 0
	int currentFloor = elevio.GetFloor()

	for i := 0; i < globals.MaxOrders; i++ {
		if currentFloor != -1 && floor == globals.OrderArray[i][0] {
			if globals.OrderArray[i][2] == true {
				elevio.SetButtonLamp(elevio.BT_Cab, floor, globals.Off)
				OrderComplete++
			}
			else [
				elevio.SetButtonLamp(globals.OrderArray[i][1], f, globals.Off)
				OrderComplete++
			]
		}

		for j := 0; j < 3; j++ {
			globals.OrderArray[i][j] == globals.NotDefined
			return OrderComplete
		}
	}
}

func amountOfOrders() int {
    orderAmount := 0
    for i := 0; i < globals.MaxOrders; i++ {
        if globals.OrderArray[i][0] == globals.NotDefined &&
           globals.OrderArray[i][1] == globals.NotDefined &&
           globals.OrderArray[i][2] == globals.NotDefined {
            continue
        } else {
            orderAmount++
        }
    }
    return orderAmount
}

func elevatorDirection() MotorDirection {
    if globals.IsDoorOpen == globals.Close && amountOfOrders() > 0 {
        if globals.CurrentDirection == globals.Up {
            for i := 0; i < MaxOrders; i++ {
                if globals.OrderArray[i][0] > globals.LastDefinedFloor && globals.OrderArray[i][0] != globals.NotDefined {
                    return elevio.MD_Up
                }
            }
            for i := 0; i < MaxOrders; i++ {
                if globals.OrderArray[i][0] != globals.NotDefined {
                    return elevio.MD_Down
                }
            }
        } else if globals.CurrentDirection == globals.Down {
            for i := 0; i < MaxOrders; i++ {
                if globals.OrderArray[i][0] < globals.LastDefinedFloor && globals.OrderArray[i][0] != globals.NotDefined {
                    return elevio.MD_Down
                }
            }
            for i := 0; i < MaxOrders; i++ {
                if globals.OrderArray[i][0] != globals.NotDefined {
                    return elevio.MD_Up
                }
            }
        }
    }
    return elevio.MD_Stop
}


func Obstruction() {
    for elevio.GetObstruction() == true {
        elevio.SetDoorOpenLamp(globals.Open)
    }

    elevio.SetDoorOpenLamp(globals.Close)
}
