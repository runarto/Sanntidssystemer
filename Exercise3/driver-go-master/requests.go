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
            if (OrderArray[i][0] != NotDefined) {
                fmt.Println(OrderArray[i][j])
                }
        }
    }
}

// fromFloor: etasje input kommer fra, button: type (opp, ned, cab)

func addToQueueFromFloorPanel(fromFloor int, button int) {
    for i := 0; i < MaxOrders; i++ {
        if (OrderArray[i][0] == fromFloor) &&
           (OrderArray[i][1] == button) {
            return
           }

        if (OrderArray[i][0] == NotDefined) && 
           (OrderArray[i][1] == NotDefined) && 
           (OrderArray[i][2] == NotDefined) {
            OrderArray[i][0] = fromFloor
            OrderArray[i][1] = button
            return
        }
    }
}

func addToQueueCab(toFloor int) {
    for i := 0; i < MaxOrders; i++ {

        if (OrderArray[i][0] == toFloor) &&
           (OrderArray[i][2] == True) {
         return
        }


        if (OrderArray[i][0] == NotDefined) && 
           (OrderArray[i][1] == NotDefined) && 
           (OrderArray[i][2] == NotDefined) {
            OrderArray[i][0] = toFloor
            fmt.Println("Check2")

            if (LastDefinedFloor < toFloor) || 
               ((LastDefinedFloor == toFloor) && 
                (CurrentDirection == ElevDown)) {
                OrderArray[i][1] = Up
                OrderArray[i][2] = True
                fmt.Println("Order added successfully")
            }

            if (LastDefinedFloor > toFloor) || 
               (LastDefinedFloor == toFloor && 
                CurrentDirection == ElevUp) {
                OrderArray[i][1] = Down
                OrderArray[i][2] = True
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


func checkOrderCompletion() int {
    completedOrders := 0

    for i := 0; i < MaxOrders; i++ {
        floor := OrderArray[i][0]
        direction := OrderArray[i][1]
        fromCab := OrderArray[i][2]

        // Check if the current floor matches the order floor
        if currentFloor == floor {
            // Process the order if it's from the cab, or if the elevator's direction matches the order's direction
            if fromCab == 1 || (direction == CurrentDirection && CurrentDirection != 0) {
                elevio.SetButtonLamp(elevio.ButtonType(2), floor, Off)
                removeOrder(i)
                completedOrders++
            } else if CurrentState == Still {
                // If the elevator is stationary, process any order on the current floor
                elevio.SetButtonLamp(elevio.ButtonType(2), floor, Off)
                removeOrder(i)
                completedOrders++
            }
        } else if (currentFloor-1 == floor && direction == Up && CurrentDirection != Down) || 
                  (currentFloor+1 == floor && direction == Down && CurrentDirection != Up) {
            // Special handling for orders from adjacent floors where direction matches the intended travel
            // This check avoids picking up passengers who want to go in the opposite direction
            elevio.SetButtonLamp(elevio.ButtonType(direction), floor, Off)
            removeOrder(i)
            completedOrders++
        }
    }

    return completedOrders
}



func removeOrder(index int) {
    // Set all elements of the order to an 'empty' or 'inactive' state
    OrderArray[index][0] = NotDefined // Assuming -1 indicates an inactive order
    OrderArray[index][1] = NotDefined
    OrderArray[index][2] = NotDefined
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
        if CurrentDirection == ElevUp {
            for i := 0; i < MaxOrders; i++ {
                if OrderArray[i][0] > LastDefinedFloor && OrderArray[i][0] != NotDefined {
                    fmt.Println("Elevator go up")
                    return elevio.MD_Up
                } else if OrderArray[i][0] < LastDefinedFloor && OrderArray[i][0] != NotDefined {
                    fmt.Println("Elevator go down")
                    return elevio.MD_Down
                }
            }
        } else if CurrentDirection == ElevDown {
            for i := 0; i < MaxOrders; i++ {
                if OrderArray[i][0] < LastDefinedFloor && OrderArray[i][0] != NotDefined {
                    fmt.Println("Elevator go down")
                    return elevio.MD_Down
                } else if OrderArray[i][0] > LastDefinedFloor && OrderArray[i][0] != NotDefined {
                    fmt.Println("Elevator go up")
                    return elevio.MD_Up
                }
            }
        }
    }
    fmt.Println("oops, elevator has no orders")
    return elevio.MD_Stop
}



func Obstruction() {
    for elevio.GetObstruction() {
        elevio.SetDoorOpenLamp(On)
    }

    elevio.SetDoorOpenLamp(Off)
}
