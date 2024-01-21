package main

import (
    "fmt"
    "elevio"
    "globals"
    "requests"
    "FSM"
)

handleButtonEvent(btn elevio.ButtonEvent) {
    if btn.Button == elevio.BT_Cab {
        requests.addToQueueCab(btn.floor)
        elevio.SetButtonLamp(btn.Button, btn.floor, globals.On)
    } else {
        requests.addToQueueFromFloorPanel(btn.floor, btn.Button)
        elevio.SetButtonLamp(btn.Button, btn.floor, globals.On)

    }
}

func handleFloorEvent(floor int, MotorDirection *elevio.MotorDirection) {
    fmt.Printf("Reached floor: %d\n", floor)
    elevio.SetFloorIndicator(floor)
    globals.LastDefinedFloor = floor

    // Example logic: Determine if the elevator should stop at this floor
    ''
    if requests.OrderComplete() > 0 {
        FSM.elevatorStill()
        FSM.elevatorDoorState(globals.Open)
        // wait x seconds
        FSM.elevatorDoorState(globals.Close)
    }

}

func handleObstructionEvent(obstr bool, MotorDirection *elevio.MotorDirection) {
    if obstr {
        fmt.Println("Obstruction detected")
        *d = elevio.MD_Stop
        elevatorDoorState(globals.Open)
    } else {
        fmt.Println("Obstruction cleared")
        // Resume operation if necessary
    }
}


func handleStopEvent(stop bool) {
    if stop {
        fmt.Println("Stop button pressed")
        stopElevator()
    } else {
        fmt.Println("Stop button released")
        // Resume normal operation if needed
    }
}




func main() {
    numFloors := 4

    // Initialize the elevator
    elevio.Init("localhost:15657", numFloors)
    var d elevio.MotorDirection = elevio.MD_Up

    // Create channels for handling events
    drv_buttons := make(chan elevio.ButtonEvent)
    drv_floors := make(chan int)
    drv_obstr := make(chan bool)
    drv_stop := make(chan bool)

    // Start polling functions in separate goroutines
    go elevio.PollButtons(drv_buttons)
    go elevio.PollFloorSensor(drv_floors)
    go elevio.PollObstructionSwitch(drv_obstr)
    go elevio.PollStopButton(drv_stop)

    // Main event loop
    for {
        select {
        case btn := <-drv_buttons:
            handleButtonEvent(btn)
        case floor := <-drv_floors:
            handleFloorEvent(floor, &MotorDirection)
        case obstr := <-drv_obstr:
            handleObstructionEvent(obstr, &MotorDirection)
        case stop := <-drv_stop:
            handleStopEvent(stop)
        }
    }
}

// Define additional functions like handleButtonEvent, handleFloorEvent, etc.
