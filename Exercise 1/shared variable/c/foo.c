// Compile with `gcc foo.c -Wall -std=gnu99 -lpthread`, or use the makefile
// The executable will be named `foo` if you use the makefile, or `a.out` if you use gcc directly

#include <pthread.h>
#include <stdio.h>

int i = 0;

// Note the return type: void*
void* incrementingThreadFunction(){
    // TODO: increment i 1_000_000 times
    int *p = &i;
    for (int j = 0; j < 1000000; j++)
    {
        (*p)++; 
    }
    return NULL;
}

void* decrementingThreadFunction(){
    // TODO: decrement i 1_000_000 times

    int *p = &i; 
    for (int j = 0; j < 1000000; j++)
    {
        (*p)--; 
    }

    return NULL;
}


int main(){
    pthread_t increment;
    pthread_t decrement;
    

    // TODO: 
    // start the two functions as their own threads using `pthread_create`
    // Hint: search the web! Maybe try "pthread_create example"?
    
    // TODO:
    // wait for the two threads to be done before printing the final result
    // Hint: Use `pthread_join`   

    int incrementResult = pthread_create(&increment, NULL, incrementingThreadFunction, NULL);
    int decrementResult = pthread_create(&decrement, NULL, incrementingThreadFunction, NULL);

    if (incrementResult == 0) {
        // Successfully created the thread, now wait for it to finish
        pthread_join(increment, NULL);
    } else {
        // An error occurred
        fprintf(stderr, "Error: pthread_create returned %d\n", incrementResult);
    }

    if (decrementResult == 0) {
        // Successfully created the thread, now wait for it to finish
        pthread_join(decrement, NULL);
    } else {
        // An error occurred
        fprintf(stderr, "Error: pthread_create returned %d\n", decrementResult);
    }

    // Your main program continues here
    
    printf("The magic number is: %d\n", i);
    return 0;
}
