package main

import (
    "fmt"
    "elevio"
	"requests"
	"globals"
)

func nullButtons() {
    elevio.SetStopLamp(globals.Off)
    for f := 0; f < elevio._numFloors; f++ {
        for b := 0; b < elevio._numButtons; b++ {
			elevio.SetButtonLamp(b, f, globals.Off)

        }
    }
}

func initElevator() {
    nullButtons()
    ElevatorDoorState(globals.Close)
    floor := elevio.GetFloor()


    for elevio.FloorSensor() != 0 {
        if floor > 0 || floor == -1 {
            MoveElevator(elevio.MD_Down)
            globals.CurrentDirection = elevio.MD_Down
        }
        ElevatorStill()
    }
    fmt.Println("Elevator is ready for use")
}


func elevatorDoorState(state int) {
    elevio.SetDoorOpenLamp(state)
    globals.IsDoorOpen = state
}

func stopElevator() {
	elevio.SetMotorDirection(elevio.MD_Stop)
    nullButtons()
    if elevio.GetFloor() != -1 {
        elevatorDoorState(globals.Open)
    }
    globals.CurrentState = globals.Stop
}

func elevatorStill() {
	elevio.SetMotorDirection(elevio.MD_Stop)
    globals.CurrentState = globals.Still
}


func floorLights() {
    currentFloor := elevio.GetFloor();
    if (currentFloor >= 0 && currentFloor <= 3) {
        elevio.SetFloorIndicator(currentFloor);
        globals.LastDefinedFloor = currentFloor;
    }
    
}

func moveElevator(MotorDirection Direction) {
    if (Direction == elevio.MD_Down)
    {
        elevio.SetMotorDirection(Direction);
        globals.CurrentDirection = Direction;
        globals.CurrentState = Moving;
        fmt.Println("Now moving down\n");
    }
    else if (Direction == elevio.MD_Up)
    {
        elevio.SetMotorDirection(Direction);
        globals.CurrentDirection = Direction;
        globals.CurrentState = Moving;
        fmt.Println("Now moving up\n");
    }
}