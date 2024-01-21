package main

import (
    "fmt"
    "elevio"
    "globals"
    "requests"
    "FSM"
    "time"
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


func handleObstructionEvent() {
    requests.Obstruction()
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

            if btn == elevio.BT_Cab {
                requests.addToQueueCab(btn.floor)
                elevio.SetButtonLamp(btn.Button, btn.floor, globals.On)
            } else {
                requests.addToQueueFromFloorPanel(btn.floor, btn.Button)
                elevio.SetButtonLamp(btn.Button, btn.floor, globals.On)
        
            }

        case floor := <-drv_floors:

            switch (globals.CurrentState) {
            case globals.Moving:
                if requests.orderCompleteCheck() != 0 {
                    FSM.elevatorStill()
                    FSM.elevatorDoorState(globals.Open)
                    time.Sleep(3 * time.Second) // Delay in Go
                    FSM.elevatorDoorState(globals.Close)
                }
                break; 
            
            case.globals.Still:
                moveElevator(requests.elevatorDirection())
                break
            }

        case obstr := <-drv_obstr:
            handleObstructionEvent()
        case stop := <-drv_stop:
            FSM.stopElevator()
        }
    }
}

// Define additional functions like handleButtonEvent, handleFloorEvent, etc.
