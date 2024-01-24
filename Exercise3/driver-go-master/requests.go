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
           (OrderArray[i][1] == int(button)) {
            return
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

        if (OrderArray[i][0] == toFloor) &&
           (OrderArray[i][2] == Cab) {
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

func orderCompleteCheck(currentFloor int) int {
	OrderComplete := 0
    fmt.Println("The current floor is:", currentFloor)


    switch (CurrentState) {
    

    case Moving:
        switch (CurrentDirection) {

        case Up:
            fmt.Println("The current state is", CurrentState, ",and the direction is", CurrentDirection)
            for i := 0; i < MaxOrders; i++ {
                if (OrderArray[i][1] == Up) && (currentFloor == OrderArray[i][0]) {
                    OrderComplete = removeOrder(OrderArray[i][1], i, currentFloor)
                }  else if (OrderArray[i][1] == Down) && (currentFloor == OrderArray[i][0]) {
                    OrderComplete = removeOrder(OrderArray[i][1], i, currentFloor)
                }  else if (OrderArray[i][2] == Cab) {
                    OrderComplete = removeOrder(OrderArray[i][2], i, currentFloor)
                }
            }

        case Down: 
            for i := 0; i < MaxOrders; i++ {
                if (OrderArray[i][1] == Down) && (currentFloor == OrderArray[i][0]) {
                    OrderComplete = removeOrder(OrderArray[i][1], i, currentFloor)
                }  else if (OrderArray[i][1] == Up) && (currentFloor == OrderArray[i][0]) {
                    OrderComplete = removeOrder(OrderArray[i][1], i, currentFloor)
                } else if (OrderArray[i][2] == Cab) {
                    OrderComplete = removeOrder(OrderArray[i][2], i, currentFloor)
                }
            }
        }

    case Still:

        for i := 0; i < MaxOrders; i++ {
            if (currentFloor == OrderArray[i][0]) {
                OrderComplete = removeOrder(OrderArray[i][1], i, currentFloor)
                OrderComplete = OrderComplete + removeOrder(OrderArray[i][2], i, currentFloor)
            }
        }


    }
	return OrderComplete
}


func removeOrder(ButtonType int, entry int, currentFloor int) int {
	OrderComplete := 0

    elevio.SetButtonLamp(elevio.ButtonType(ButtonType), currentFloor, Off)

    for j := 0; j < 3; j++ {
        OrderArray[entry][j] = NotDefined
    }
	
    OrderComplete++
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
