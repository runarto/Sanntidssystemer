package main


import (
    "fmt"
    "Exercise3/elevio"
)


func nullButtons() {
    elevio.SetStopLamp(Off)
    for f := 0; f < numFloors; f++ {
        for b := 0; b < numButtons; b++ {
			elevio.SetButtonLamp(elevio.ButtonType(b), f, Off)

        }
    }
}


/* func initElevator() {
    nullButtons()
    elevatorDoorState(Close)
    floor := elevio.GetFloor()


    for elevio.GetFloor() != 0 {
        if floor > 0 || floor == -1 {
            moveElevator(elevio.MD_Down)
            CurrentDirection = elevio.MD_Down
        }
        elevatorStill()
    }
    fmt.Println("Elevator is ready for use")
} */


func elevatorDoorState(state int) {
    if state == Open {
        elevio.SetDoorOpenLamp(On)
        IsDoorOpen = Open
    } else {
        elevio.SetDoorOpenLamp(Off)
        IsDoorOpen = Close
    }

}

func stopElevator() {
	elevio.SetMotorDirection(elevio.MD_Stop)
    nullButtons()
    if elevio.GetFloor() != -1 {
        elevatorDoorState(Open)
    }
    CurrentState = Stop
}

func elevatorStill() {
	elevio.SetMotorDirection(elevio.MD_Stop)
    CurrentState = Still
}


func floorLights() {
    currentFloor := elevio.GetFloor();
    if (currentFloor >= 0 && currentFloor <= 3) {
        elevio.SetFloorIndicator(currentFloor);
        LastDefinedFloor = currentFloor;
    }
    
}

func moveElevator(Direction elevio.MotorDirection) {
    if (Direction == elevio.MD_Down) {
        elevio.SetMotorDirection(Direction);
        CurrentDirection = int(Direction);
        CurrentState = Moving;
        fmt.Println("Now moving down\n");
    }
    if (Direction == elevio.MD_Up) {
        elevio.SetMotorDirection(Direction);
        CurrentDirection = int(Direction);
        CurrentState = Moving;
        fmt.Println("Now moving up\n");
    }
}