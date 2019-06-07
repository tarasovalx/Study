#include <stdio.h>
#include <malloc.h>
#include <string.h>

#define Max(a, b) (((a)>(b))?(a):(b))

enum Comands {
    ENQ, DEQ, EMPTY, MAX
};

int GetCMD(char *cmd) {
    const char comands[4][6] = {"ENQ", "DEQ", "EMPTY", "MAX"};
    for (int i = 0; i < 4; i++) {
        if (strcmp(cmd, comands[i]) == 0) return i;
    }
    return -1;
}

struct StackElement {
    int value;
    int PrefixMax;
};

struct DoubleStack {
    struct StackElement *data;
    int cap;
    int top1;
    int top2;
};

void InitDoubleStack(struct DoubleStack *s, int n) {
    s->data = malloc(sizeof(struct StackElement) * n);
    s->cap = n;
    s->top1 = 0;
    s->top2 = n - 1;
}

int IsStackEmpty1(struct DoubleStack *s) {
    return !s->top1;
}

int IsStackEmpty2(struct DoubleStack *s) {
    return (s->top2 == s->cap - 1);
}

void Push1(struct DoubleStack *s, int elem) {
    if (s->top2 < s->top1) printf("Error: overflow");
    struct StackElement Pushable;
    Pushable.value = elem;
    if (IsStackEmpty1(s)) Pushable.PrefixMax = elem;
    else Pushable.PrefixMax = (elem > s->data[s->top1 - 1].PrefixMax) ? elem : s->data[s->top1 - 1].PrefixMax;
    s->data[s->top1] = Pushable;
    s->top1++;
}

void Push2(struct DoubleStack *s, struct StackElement elem) {
    if (s->top2 < s->top1) printf("Error: overflow");
    s->data[s->top2] = elem;
    s->top2--;
}

struct StackElement Pop1(struct DoubleStack *s) {
    if (IsStackEmpty1(s)) printf("Error: underflow");
    s->top1--;
    return s->data[s->top1];
}

struct StackElement Pop2(struct DoubleStack *s) {
    if (IsStackEmpty2(s)) printf("Error: underflow");
    s->top2++;
    return s->data[s->top2];
}

void InitQueueOnStack(struct DoubleStack *s, int n) {
    InitDoubleStack(s, n);
}

int QueueEmpty(struct DoubleStack *s) {
    return (IsStackEmpty1(s) && IsStackEmpty2(s));
}

void Enqueue(struct DoubleStack *s, int elem) {
    Push1(s, elem);
}

struct StackElement Dequeue(struct DoubleStack *s) {
    if (IsStackEmpty2(s)) {
        while (!IsStackEmpty1(s)) {
            struct StackElement Pushable;
            Pushable = Pop1(s);
            if (IsStackEmpty2(s))
                Pushable.PrefixMax = Pushable.value;
            else
                Pushable.PrefixMax = (Pushable.value > s->data[s->top2 + 1].PrefixMax) ? Pushable.value : s->data[s->top2 + 1].PrefixMax;
            Push2(s, Pushable);
        }
    }
    return Pop2(s);
}

int main() {
    int n;
    struct StackElement j;
    scanf("%d", &n);
    struct DoubleStack q;
    InitQueueOnStack(&q, n);
    char cmd[6];
    int a, b;
    for (int i = 0; i < n; i++) {
        scanf("%s", cmd);
        enum Comands res = GetCMD(cmd);
        switch (res) {
            case ENQ:
                scanf("%d", &a);
                Enqueue(&q, a);
                break;
            case DEQ:
                printf("%d\n", Dequeue(&q).value);
                break;
            case EMPTY:
                if (QueueEmpty(&q)) printf("true\n");
                else printf("false\n");
                break;
            case MAX:
                if (IsStackEmpty1(&q)) printf("%d\n", q.data[q.top2 + 1].PrefixMax);
                else if (IsStackEmpty2(&q)) printf("%d\n", q.data[q.top1 - 1].PrefixMax);
                else printf("%d\n",Max(q.data[q.top1 - 1].PrefixMax, q.data[q.top2 + 1].PrefixMax));
                break;
        }
    }
    free(q.data);
}
