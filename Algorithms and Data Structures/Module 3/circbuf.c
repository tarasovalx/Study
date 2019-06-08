#include <stdio.h>
#include <malloc.h>
#include <string.h>

enum Comands {
    ENQ, DEQ, EMPTY
};

int GetCMD(char *cmd) {
    const char comands[3][6] = {"ENQ", "DEQ", "EMPTY"};
    for (int i = 0; i < 3; i++) {
        if (strcmp(cmd, comands[i]) == 0) return i;
    }
    return -1;
}

struct Queue {
    int *data;
    int count;
    int head;
    int tail;
    int cap;
};

void InitQueue(struct Queue *q, int n) {
    q->data = malloc(sizeof(int) * n);
    q->count = 0, q->head = 0, q->tail = 0, q->cap = n;
}

int QueueEmpty(struct Queue *q) {
    return !q->count;
}

int Dequeue(struct Queue *q) {
    if (!q->count) printf("Error: devastation");
    int elem = q->data[q->head];
    q->head++;
    if (q->head == q->cap) q->head = 0;
    q->count--;
    return elem;
}

void ExpandQueue(struct Queue *q, int n) {
    int *buf = malloc(sizeof(int) * q->cap * n);
    int i = 0;
    for (int j = q->head; q->count; j = (j + 1) % (q->cap), i++) {
        buf[i] = Dequeue(q);
    }
    q->cap *= n;
    q->tail = i;
    q->count = i;
    q->head = 0;
    free(q->data);
    q->data = buf;
}

void Enqueue(struct Queue *q, int elem) {
    if (q->count == q->cap) ExpandQueue(q, 2);
    q->data[q->tail] = elem;
    q->tail++;
    if (q->tail == q->cap) q->tail = 0;
    q->count++;
}

int main() {
    int n, a;
    enum Comands res;
    scanf("%d", &n);
    struct Queue q;
    InitQueue(&q, 4);
    char cmd[5];
    for (int i = 0; i < n; i++) {
        scanf("%s", cmd);
        res = GetCMD(cmd);
        switch (res) {
            case ENQ:
                scanf("%d", &a);
                Enqueue(&q, a);
                break;
            case DEQ:
                printf("%d\n", Dequeue(&q));
                break;
            case EMPTY:
                if (QueueEmpty(&q)) printf("true\n");
                else printf("false\n");
                break;
        }
    }
    free(q.data);
    return 0;
}
