#include <stdio.h>
#include <malloc.h>

void BuildTree(int *arr, int *T, int v, int a, int b) {
    if (a == b) T[v] = arr[a];
    else {
        int m = (a + b) / 2;
        BuildTree(arr, T, v * 2, a, m);
        BuildTree(arr, T, v * 2 + 1, m + 1, b);
        T[v] = T[v * 2] + T[v * 2 + 1];
    }
}

void UpdateTree(int v, int *T, int a, int b, int value, int i) {
    if (a == b) T[v] = value;
    else {
        int m = (a + b) / 2;
        if (i <= m) UpdateTree(v * 2, T, a, m, value, i);
        else UpdateTree(v * 2 + 1, T, m + 1, b, value, i);
        T[v] = T[v * 2] + T[v * 2 + 1];
    }
}

int PeakQ(int *T, int v, int l, int r, int a, int b) {
    if ((l == a) && (r == b))
        return T[v];
    int m = (a + b) / 2;
    if (r <= m)
        return PeakQ(T, 2 * v, l, r, a, m);
    if (l > m)
        return PeakQ(T, 2 * v + 1, l, r, m + 1, b);
    else
        return PeakQ(T, 2 * v, l, m, a, m) + PeakQ(T, 2 * v + 1, m + 1, r, m + 1, b);
}

int main(){
    int n, k, a, b;
    scanf("%d", &n);
    int *arr = malloc(sizeof(int) * n);
    int *tree = malloc(sizeof(int) * (n * 4 + 1));
    int *peaks = malloc(sizeof(int) * n);
    char cmd[4];
    for (int i = 0; i < n; i++) scanf("%d", arr + i);
    for (int i = 0; i < n; i++){
        if (!i && (i == n - 1))
            peaks[i] = 1;
        else if (!i)
            peaks[i] = arr[i] >= arr[i + 1];
        else if(i == n - 1)
            peaks[i] = arr[i] >= arr[i - 1];
        else
            peaks[i] = (arr[i] >= arr[i + 1]) && (arr[i] >= arr[i - 1]);
    }
    BuildTree(peaks, tree, 1, 0, n - 1);
    scanf("%d ", &k);
    int res[k], cnt = 0;
    for (int i = 0; i < k; i++){
        scanf("%s %d %d", cmd, &a, &b);
        if (cmd[0] == 'P') {
            res[cnt] = PeakQ(tree, 1, a, b, 0, n - 1);
            cnt++;
        } else {
            arr[a] = b;
            for (int i = a - 1; (i <= a + 1) && (i < n); i++){
                if (i < 0) continue;
                if (!i && (i == n - 1))
                    peaks[i] = 1;
                else if (!i)
                    peaks[i] = arr[i] >= arr[i + 1];
                else if(i == n - 1)
                    peaks[i] = arr[i] >= arr[i - 1];
                else
                    peaks[i] = (arr[i] >= arr[i + 1]) && (arr[i] >= arr[i - 1]);
                UpdateTree(1, tree, 0, n - 1, peaks[i], i);
            }
        }
    }
    for(int i = 0; i < cnt; i++) {
        printf("%d ", res[i]);
    }
    free(arr);
    free(peaks);
    free(tree);
    return 0;
}
