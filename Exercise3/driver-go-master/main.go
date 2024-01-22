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

    initElevator()
    initializeQueue()
    nullButtons()
    
    fmt.Println("inint")

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


            fmt.Println("New order.")
            if btn.Button == elevio.BT_Cab {
                addToQueueCab(btn.Floor)
                elevio.SetButtonLamp(btn.Button, btn.Floor, true)
            } else {
                addToQueueFromFloorPanel(btn.Floor, btn.Button)
                elevio.SetButtonLamp(btn.Button, btn.Floor, true)
        
            }
            
            printOrderArray()

        case floor := <-drv_floors:

            fmt.Println("Arrived at new floor")
            LastDefinedFloor = floor
            floorLights()

            switch (CurrentState) {
            case Moving:

                checkNewOrders()

                if orderCompleteCheck() != 0 {
                    elevatorStill()
                    elevatorDoorState(Open)
                    time.Sleep(3 * time.Second) // Delay in Go
                    elevatorDoorState(Close)
                }
            
            case Still:

                checkNewOrders()
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