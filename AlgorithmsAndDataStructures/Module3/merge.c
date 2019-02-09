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
    int *data;
    int i;
};

int Compare(struct QueueElem *a, struct QueueElem *b) {
    return a->data[a->i] - b->data[b->i];
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

void InitPriorityQueue(struct PriorityQueue *q, int n, int (*compare)(struct QueueElem *, struct QueueElem *)) {
    q->heap = malloc(sizeof(struct QueueElem) * n);
    q->cap = n;
    q->count = 0;
}

void Up(struct PriorityQueue *q, int i) {
    while (i != 0 && Compare(&q->heap[i], &q->heap[Parent(i)]) > 0) {
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
        if (RightChild(i) < q->count && Compare(&q->heap[RightChild(i)], &q->heap[LeftChild(i)]) > 0)
            max_index = RightChild(i);
        if (Compare(&q->heap[i], &q->heap[max_index]) >= 0)
            return;
        HeapSwap(&q->heap[i], &q->heap[max_index]);
        i = max_index;
    }
}

int ExtractMax(struct PriorityQueue *q) {
    if (!q->count) printf("Error: queue is empty\n");
    int res;

    if (q->heap[0].i == 0) {
        res = q->heap[0].data[q->heap[0].i];
        HeapSwap(&q->heap[--q->count], &q->heap[0]);
        Down(q, 0);
    } else {
        res = q->heap[0].data[q->heap->i];
        q->heap[0].i--;
        Down(q, 0);
    }
    return res;
}

int main() {
    int n;
    scanf("%d", &n);
    int sizes[n], res_cap = 0;
    struct PriorityQueue q;
    InitPriorityQueue(&q, n, &Compare);
    for (int i = 0; i < n; i++) scanf("%d", sizes + i);
    int **arr = malloc(sizeof(int *) * n);

    for (int i = 0; i < n; i++) {
        if (sizes[i]) {
            int j;
            arr[i] = malloc(sizeof(int) * sizes[i]);
            for (j = 0; j < sizes[i]; j++) scanf("%d", &arr[i][j]);
            struct QueueElem Inserting;
            Inserting.data = arr[i];
            Inserting.i = j - 1;
            Insert(&q, Inserting);
            res_cap += j;
        }
    }
    int res[res_cap];

    for (int i = res_cap - 1; i >= 0; i--) {
        res[i] = ExtractMax(&q);
    }
    for (int i = 0; i < res_cap; i++) printf("%d ", res[i]);
    free(q.heap);
    for (int i = 0; i < n; i++) {
        if (sizes[i]) {
            free(arr[i]);
        }
    }
    free(arr);
}
