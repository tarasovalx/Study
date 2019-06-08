#include <stdio.h>
#include <malloc.h>

#define MAX(a, b) (((a)>(b))?(a):(b))

void BuildTree(int *arr, int *T, int v, int a, int b) {
    if (a == b) T[v] = arr[a];
    else {
        int m = (a + b) / 2;
        BuildTree(arr, T, v * 2, a, m);
        BuildTree(arr, T, v * 2 + 1, m + 1, b);
        T[v] = MAX(T[v * 2], T[v * 2 + 1]);
    }
}

void UpdateTree(int v, int *T, int a, int b, int value, int i) {
    if (a == b) T[v] = value;
    else {
        int m = (a + b) / 2;
        if (i <= m) UpdateTree(v * 2, T, a, m, value, i);
        else UpdateTree(v * 2 + 1, T, m + 1, b, value, i);
        T[v] = MAX(T[v * 2], T[v * 2 + 1]);
    }
}

int MaxQ(int *T, int v, int l, int r, int a, int b) {
    if ((l == a) && (r == b))
        return T[v];
    int m = (a + b) / 2;
    if (r <= m)
        return MaxQ(T, 2 * v, l, r, a, m);
    if (l > m)
        return MaxQ(T, 2 * v + 1, l, r, m + 1, b);
    else
        return MAX(MaxQ(T, 2 * v, l, m, a, m), MaxQ(T, 2 * v + 1, m + 1, r, m + 1, b));
}

int main() {
    int n, k;
    scanf("%d", &n);
    int *arr = calloc(n ,sizeof(int)), res_cnt = 0;
    for (int i = 0; i < n; i++) scanf("%d ", arr + i);
    int *tree = malloc(sizeof(int) * n * 4 + 1);
    BuildTree(arr, tree, 1, 0, n - 1);
    scanf("%d ", &k);
    int *res = calloc(k ,sizeof(int));
    char cmd[4];
    int a, b;
    for (int i = 0; i < k; i++) {
        scanf("%s %d %d", cmd, &a, &b);
        if (cmd[0] == 'M') {
            res[res_cnt] = MaxQ(tree, 1, a, b, 0, n - 1);
            res_cnt++;
        } else UpdateTree(1, tree, 0, n - 1, b, a);
    }
    for (int i = 0; i < res_cnt; i++) printf("%d ", res[i]);
    free(tree);
    free(arr);
    free(res);
    return 0;
}
