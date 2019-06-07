#include <stdio.h>
#include <malloc.h>

int Parent(int i) {
    return (i - 1) / 2;
}

int LeftChild(int i) {
    return 2 * i + 1;
}

int RightChild(int i) {
    return 2 * i + 2;
}

struct QueueElem {
    int t1, t2, sum;
};

int Compare(struct QueueElem *a, struct QueueElem *b) {
    return a->sum - b->sum;
}

struct PriorityQueue {
    struct QueueElem *heap;
    int cap;
    int count;
};

void HeapSwap(struct QueueElem *a, struct QueueElem *b) {
    struct QueueElem tmp;
    tmp = *a;
    *a = *b;
    *b = tmp;
}

void InitPriorityQueue(struct PriorityQueue *q, int n) {
    q->heap = malloc(sizeof(struct QueueElem) * (n + 1));
    q->cap = n;
    q->count = 0;
}

void Up(struct PriorityQueue *q, int i) {
    while (i != 0 && Compare(&q->heap[i], &q->heap[Parent(i)]) < 0) {
        HeapSwap(&q->heap[i], &q->heap[Parent(i)]);
        i = Parent(i);
    }
}

void Insert(struct PriorityQueue *q, struct QueueElem a) {
    int i = q->count;
    if (i == q->cap) printf("Error: overflow");
    q->count++;
    q->heap[i] = a;
    Up(q, i);
}

void Down(struct PriorityQueue *q, int i) {
    while (i < q->count / 2) {
        int max_index = LeftChild(i);
        if (RightChild(i) < q->count && Compare(&q->heap[RightChild(i)], &q->heap[LeftChild(i)]) < 0)
            max_index = RightChild(i);
        if (Compare(&q->heap[i], &q->heap[max_index]) <= 0)
            return;
        HeapSwap(&q->heap[i], &q->heap[max_index]);
        i = max_index;
    }
}

struct QueueElem ExtractMin(struct PriorityQueue *q) {
    if (!q->count) printf("Error: queue is empty\n");
    struct QueueElem res = q->heap[0];
    if (q->count > 0){
        HeapSwap(&q->heap[--q->count], &q->heap[0]);
        Down(q, 0);
    }
    return res;
}

int main() {
    int n, m, max;
    scanf("%d\n%d", &n, &m);
    struct PriorityQueue q;
    InitPriorityQueue(&q, n);
    struct QueueElem *elems = malloc(sizeof(struct QueueElem) * (m + 1));
    for (int i = 0; i < m; i++) {
        scanf("%d %d", &elems[i].t1, &elems[i].t2);
        elems[i].sum = elems[i].t1 + elems[i].t2;
    }
    for (int i = 0; i < n; i++) Insert(&q, elems[i]);
    struct QueueElem buf;
    for (int i = n; q.count;) {
        buf = ExtractMin(&q);
        if (i < m) {
            if (elems[i].t1 > buf.sum)
                max = elems[i].t1;
            else max = buf.sum;
            elems[i].sum = max + elems[i].t2;
            Insert(&q, elems[i]);
            i++;
        }
    }
    printf("%d", buf.sum);
    free(q.heap);
    free(elems);
}
