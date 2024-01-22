package main

import (
    "fmt"
    "Exercise3/elevio"
    "time"
)



func main() {
    numFloors := 4

    // Initialize the elevator
    elevio.Init("localhost:15657", numFloors)
    //var d elevio.MotorDirection = elevio.MD_Up

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

            if btn.Button == elevio.BT_Cab {
                addToQueueCab(btn.Floor)
                elevio.SetButtonLamp(btn.Button, btn.Floor, true)
            } else {
                addToQueueFromFloorPanel(btn.Floor, btn.Button)
                elevio.SetButtonLamp(btn.Button, btn.Floor, true)
        
            }

        case floor := <-drv_floors:

            if floor == -1 || floor == 0 {
                fmt.Println("lol")
            }

            switch (CurrentState) {
            case Moving:
                if orderCompleteCheck() != 0 {
                    elevatorStill()
                    elevatorDoorState(Open)
                    time.Sleep(3 * time.Second) // Delay in Go
                    elevatorDoorState(Close)
                }
            
            case Still:
                moveElevator(elevatorDirection())
            }

        case obstr := <-drv_obstr:
            if obstr {
                Obstruction()
            }
        case stop := <-drv_stop:
            if stop {
                stopElevator()
            }
        }
    }
}

// Define additional functions like handleButtonEvent, handleFloorEvent, etc.
