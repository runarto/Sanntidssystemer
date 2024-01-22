package main


import (
    "fmt"
    "Exercise3/elevio"
)


func initializeQueue() {
    for i := range OrderArray {
        OrderArray[i][0] = NotDefined
        OrderArray[i][1] = NotDefined
        OrderArray[i][2] = NotDefined
    }
}


func printOrderArray() {
    for i := 0; i < MaxOrders; i++ {
        for j := 0; j < 3; j++ {
            fmt.Println(OrderArray[i][j])
        }
    }
}

// fromFloor: etasje input kommer fra, button: type (opp, ned, cab)

func addToQueueFromFloorPanel(fromFloor int, button elevio.ButtonType) {
    for i := 0; i < MaxOrders; i++ {
        if (OrderArray[i][0] == fromFloor) && 
           (OrderArray[i][1] == int(button)) && 
           (OrderArray[i][2] == NotDefined) {
            return //Ordre kommer fra heisen
        }

        if (OrderArray[i][0] == NotDefined) && 
           (OrderArray[i][1] == NotDefined) && 
           (OrderArray[i][2] == NotDefined) {
            OrderArray[i][0] = fromFloor
            OrderArray[i][1] = int(button)
            return
        }
    }
}

func addToQueueCab(toFloor int) {
    for i := 0; i < MaxOrders; i++ {
        if OrderArray[i][0] == toFloor && OrderArray[i][2] == Cab {
            fmt.Println("Check1")
            return
        }

        if (OrderArray[i][0] == NotDefined) && 
           (OrderArray[i][1] == NotDefined) && 
           (OrderArray[i][2] == NotDefined) {
            OrderArray[i][0] = toFloor
            fmt.Println("Check2")

            if (LastDefinedFloor < toFloor) || 
               (LastDefinedFloor == toFloor && 
                CurrentDirection == Down) {
                OrderArray[i][1] = Up
                OrderArray[i][2] = Cab
                fmt.Println("Order added successfully")
            }

            if (LastDefinedFloor > toFloor) || 
               (LastDefinedFloor == toFloor && 
                CurrentDirection == Up) {
                OrderArray[i][1] = Down
                OrderArray[i][2] = Cab
                fmt.Println("Order added successfully")
            }
        }
    }
}

/* func checkNewOrders() {
	for f := 0; f < numFloors; f++ {
		for b := 0; b < numButtons; b++ {
			btnPressed := elevio.GetButton(elevio.ButtonType(b),f)
			if btnPressed && f != elevio.GetFloor() {
				fmt.Println("Button was pressed");
				elevio.SetButtonLamp(elevio.ButtonType(b), f, On)
				if b == elevio.BT_Cab {
					addToQueueCab(f)
					fmt.Println("Order added from inside cab")
				} else {
					addToQueueFromFloorPanel(f, elevio.ButtonType(b))
					fmt.Println("Order added from outside elevator")
				}
			}
		}
	}
} */

func orderCompleteCheck() int {
	OrderComplete := 0
	currentFloor := elevio.GetFloor()

	for i := 0; i < MaxOrders; i++ {
		if currentFloor != -1 && currentFloor == OrderArray[i][0] {
			if OrderArray[i][1] == Up && LastDefinedFloor < OrderArray[i][0] {
				OrderComplete = removeOrdersAtFloor(OrderArray[i][0])
			} 
			if OrderArray[i][1] == Down && LastDefinedFloor > OrderArray[i][0] {
				OrderComplete = removeOrdersAtFloor(OrderArray[i][0])
				break
			}
		}
	}
	return OrderComplete
}


func removeOrdersAtFloor(floor int) int {
	OrderComplete := 0
	currentFloor := elevio.GetFloor()

	for i := 0; i < MaxOrders; i++ {
		if currentFloor != -1 && floor == OrderArray[i][0] {
			if OrderArray[i][2] == Cab {
				elevio.SetButtonLamp(elevio.BT_Cab, floor, Off)
				OrderComplete++
			} else {
				elevio.SetButtonLamp(elevio.ButtonType(OrderArray[i][1]), floor, Off)
				OrderComplete++
            }
            
            for j := 0; j < 3; j++ {
                OrderArray[i][j] = NotDefined
            }
		}
	}

	return OrderComplete
}


func amountOfOrders() int {
    orderAmount := 0
    for i := 0; i < MaxOrders; i++ {
        if OrderArray[i][0] == NotDefined &&
           OrderArray[i][1] == NotDefined &&
           OrderArray[i][2] == NotDefined {
            continue
        } else {
            orderAmount++
        }
    }
    return orderAmount
}

func elevatorDirection() elevio.MotorDirection {
    if IsDoorOpen == Close && amountOfOrders() > 0 {
        if CurrentDirection == Up {
            for i := 0; i < MaxOrders; i++ {
                if OrderArray[i][0] > LastDefinedFloor && OrderArray[i][0] != NotDefined {
                    return elevio.MD_Up
                }
            }
            for i := 0; i < MaxOrders; i++ {
                if OrderArray[i][0] != NotDefined {
                    return elevio.MD_Down
                }
            }
        } else if CurrentDirection == Down {
            for i := 0; i < MaxOrders; i++ {
                if OrderArray[i][0] < LastDefinedFloor && OrderArray[i][0] != NotDefined {
                    return elevio.MD_Down
                }
            }
            for i := 0; i < MaxOrders; i++ {
                if OrderArray[i][0] != NotDefined {
                    return elevio.MD_Up
                }
            }
        }
    }
    return elevio.MD_Stop
}


func Obstruction() {
    for elevio.GetObstruction() {
        elevio.SetDoorOpenLamp(On)
    }

    elevio.SetDoorOpenLamp(Off)
}
