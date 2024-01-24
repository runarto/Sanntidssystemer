package main

import (
    "fmt"
    "Exercise3/elevio"
    "time"
    "sync"
)



func main() {
    numFloors := 4



    // Initialize the elevator
    elevio.Init("localhost:15657", numFloors)

    var wg sync.WaitGroup

    // Add 2 to the WaitGroup for the two functions you want to wait for
    wg.Add(2)

    go func() {
        defer wg.Done() // Signal that initElevator() is done
        initElevator()
    }()

    // Start a goroutine for initializeQueue()
    go func() {
        defer wg.Done() // Signal that initializeQueue() is done
        initializeQueue()
    }()

    // Wait for both functions to complete
    wg.Wait()

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
                fmt.Println("Order added from cab")
            } else {
                addToQueueFromFloorPanel(btn.Floor, btn.Button)
                fmt.Println("Order added from ")
            }
            elevio.SetButtonLamp(btn.Button, btn.Floor, true)



            switch (CurrentState) {
            case Moving:

                if orderCompleteCheck(LastDefinedFloor) != 0 {
                    elevatorStill()
                    elevatorDoorState(Open)
                    time.Sleep(3 * time.Second) // Delay in Go
                    elevatorDoorState(Close)
                }
            
            case Still:

                moveElevator(elevatorDirection())
            }

            
            printOrderArray()

        case floor := <-drv_floors:

            fmt.Println("Arrived at new floor")
            floorLights(floor)

            switch (CurrentState) {
            case Moving:
                fmt.Println("Elevator arrived at floor while moving.")

                if orderCompleteCheck(floor) != 0 {
                    elevatorStill()
                    elevatorDoorState(Open)
                    time.Sleep(3 * time.Second) // Delay in Go
                    elevatorDoorState(Close)
                }
            
            case Still:

                fmt.Println("Elevator arrived at floor while still.")
                moveElevator(elevatorDirection())
            }

        case obstr := <-drv_obstr:
            if obstr {
                Obstruction()
            }

            switch (CurrentState) {
            case Moving:

                if orderCompleteCheck(LastDefinedFloor) != 0 {
                    elevatorStill()
                    elevatorDoorState(Open)
                    time.Sleep(3 * time.Second) // Delay in Go
                    elevatorDoorState(Close)
                }
            
            case Still:

                moveElevator(elevatorDirection())
            }


        case stop := <-drv_stop:
            if stop {
                stopElevator()
            }
        }
    }
}
