
1: Sharing a variable

If the functions simultaneously read the same value, they will both write back a different value from what was saved 
initally. Hence, only one of the values can actually be written to the memory location it is stored in. Because the functions do not take turn induvidually to increment and decrement the values, you end up with a result that is not zero in every (or almost no) case.

A mutex is a mutual exclusion lock that allows only one thread to access a resource at any given time. A semaphore is a signaling mechanism that can allow more than one thread to access a resource concurrently, depending on how the semaphore is initialized. However, a semaphore can be initialized with a value of one to behave like a mutex, known as a binary semaphore.

In our case, a mutex is generally more appropriate since we want to ensure that only one thread can modify the shared variable at a time.

runtime.GOMAXPROCS is a function in Go that sets the maximum number of operating system threads that can execute simultaneously. When setting runtime.GOMAXPROCS(1), the Go scheduler is configured to use only one OS thread for executing all user-level Go code.
This means that only one can execute at a time because there's only one thread available.