// compile with:  gcc -g main.c ringbuf.c -lpthread

#include <pthread.h>
#include <semaphore.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#include "ringbuf.h"

struct BoundedBuffer {
  struct RingBuffer *buf;
  pthread_mutex_t mtx;
  sem_t capacity;
  sem_t numElements;
};

struct BoundedBuffer *buf_new(int size) {
  struct BoundedBuffer *buf = malloc(sizeof(struct BoundedBuffer));
  buf->buf = rb_new(size);

  pthread_mutex_init(&buf->mtx, NULL);
  // TODO: initialize semaphores
  sem_init(&buf->capacity, 0, size);
  sem_init(&buf->numElements, 0, 0);

  return buf;
}

void buf_destroy(struct BoundedBuffer *buf) {
  rb_destroy(buf->buf);
  pthread_mutex_destroy(&buf->mtx);
  sem_destroy(&buf->capacity);
  sem_destroy(&buf->numElements);
  free(buf);
}

void buf_push(struct BoundedBuffer *buf, int val) {
  // printf("buf_push: waiting for buffer capacity...\n");

  sem_wait(&buf->capacity);

  pthread_mutex_lock(&buf->mtx);

  // printf("buf_push: locked mutex, pushing value %d\n", val);

  rb_push(buf->buf, val);

  sem_post(&buf->numElements);
  pthread_mutex_unlock(&buf->mtx);
  // printf("Buffer length: %d\n", buf->buf->length);
}

int buf_pop(struct BoundedBuffer *buf) {
  // printf("buf_pop: waiting for available elements...\n");

  sem_wait(&buf->numElements); // Check if there are elements in the buffer
  pthread_mutex_lock(&buf->mtx);

  // printf("buf_pop: locked mutex, popping value\n");

  int val = rb_pop(buf->buf);

  sem_post(&buf->capacity);
  pthread_mutex_unlock(&buf->mtx);
  // printf("Buffer capacity: %d\n", buf->buf->capacity);
  // printf("buf_pop: value popped and mutex unlocked, incremented capacity\n");

  return val;
}

void *producer(void *args) {
  struct BoundedBuffer *buf = (struct BoundedBuffer *)(args);

  for (int i = 0; i < 10; i++) {
    nanosleep(&(struct timespec){0, 100 * 1000 * 1000}, NULL);
    printf("[producer]: pushing %d\n", i);
    buf_push(buf, i);
  }
  return NULL;
}

void *consumer(void *args) {
  struct BoundedBuffer *buf = (struct BoundedBuffer *)(args);

  // give the producer a 1-second head start
  nanosleep(&(struct timespec){1, 0}, NULL);
  while (1) {
    int val = buf_pop(buf);
    printf("[consumer]: %d\n", val);
    nanosleep(&(struct timespec){0, 50 * 1000 * 1000}, NULL);
  }
}

int main() {

  struct BoundedBuffer *buf = buf_new(5);

  pthread_t producer_thr;
  pthread_t consumer_thr;
  pthread_create(&producer_thr, NULL, producer, buf);
  pthread_create(&consumer_thr, NULL, consumer, buf);

  pthread_join(producer_thr, NULL);
  pthread_cancel(consumer_thr);

  buf_destroy(buf);

  return 0;
}
