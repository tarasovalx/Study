#include <stdio.h>
#include <malloc.h>
#include <math.h>

void swap(void *a, void *b, size_t width) {
    char tmp;
    for (int i = 0; i < width; i++) {
        tmp = ((char *) a)[i];
        ((char *) a)[i] = ((char *) b)[i];
        ((char *) b)[i] = tmp;
    }
}

struct Task {
    int left;
    int right;
};

struct Stack {
    struct Task *data;
    int cap;
    int top;
};

void Push(struct Stack *s, struct Task value) {
    if (s->top == s->cap) printf("Error: overflow");
    s->data[s->top] = value;
    s->top += 1;
}

struct Task Pop(struct Stack *s) {
    if (!s->top) printf("Error: devastation");
    s->top--;
    return s->data[s->top];
}

int Partition(int *base, int low, int high) {
    int i = low;
    for (int j = low; j < high; j++) {
        if (base[j] < base[high]) {
            swap(base + i, base + j, sizeof(int));
            i++;
        }
    }
    swap(base + i, base + high, sizeof(int));
    return i;
}

void Init(struct Stack *s, int n) {
    s->data = malloc(sizeof(struct Task) * (n + 1));
    s->cap = n;
    s->top = 0;
}

void QuickSort(int *base, int l, int r) {
    int left, right;
    struct Task CurrentTask;
    CurrentTask.left = l;
    CurrentTask.right = r;
    struct Stack s;
    Init(&s, r + 1);
    Push(&s, CurrentTask);
    while (s.top) {
        CurrentTask = Pop(&s);
        left = CurrentTask.left;
        right = CurrentTask.right;
        if (right <= left) continue;
        int i = Partition(base, left, right);
        struct Task Pushable;
        if (i - left > right - i) {
            Pushable.right = i - 1;
            Pushable.left = left;
            Push(&s, Pushable);
            Pushable.left = i + 1;
            Pushable.right = right;
            Push(&s, Pushable);
        } else {
            Pushable.right = right;
            Pushable.left = i + 1;
            Push(&s, Pushable);
            Pushable.left = left;
            Pushable.right = i - 1;
            Push(&s, Pushable);
        }
    }
    free(s.data);
}

int main() {
    int n;
    scanf("%d", &n);
    int *arr = malloc(sizeof(int) * (n + 1));
    for (int i = 0; i < n; i++) scanf("%d", arr + i);
    QuickSort(arr, 0, n - 1);
    for (int i = 0; i < n; i++) printf("%d ", arr[i]);
    free(arr);
    return 0;
}
