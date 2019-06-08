#include <stdio.h>
#include <malloc.h>
#include <math.h>
#include <stdlib.h>

int GCD(int a, int b) {
    a = abs(a);
    b = abs(b);
    return (a == 0 ? b : GCD(b % a, a));
}

void CompLogs(int *lg, int m) {
    lg[1] = 0;
    for(int i = 2; i < m; i++) {
        lg[i] = lg[i / 2] + 1;
    }
}

int STableBuild(int *lg, int const *arr, int **ST, int n, int m) {
    for (int i = 0; i < n; i++)
        ST[i][0] = arr[i];
    for (int j = 1; j < m; j++){
        for (int i = 0; i <= (n - (1 << j)); i++){
            ST[i][j] = GCD(ST[i][j - 1], ST[i + (1 << (j - 1))][j - 1]);
        }
    }
}

int STGCD(int **ST, int l, int r, int const *lg) {
    int j = lg[r - l + 1];
    return GCD(ST[l][j], ST[r - (1 << j) + 1][j]);
}

int main() {
    int n;
    scanf("%d", &n);
    int *arr = malloc(sizeof(int) * n);
    int m = (int)log2(n) + 1;
    int *lg = malloc(sizeof(int) * (n + 1));
    CompLogs(lg, n + 1);
    for (int i = 0; i < n; i++) scanf("%d ", arr + i);
    int **ST = malloc(sizeof(int*) * n);
    for (int i = 0; i < n; i++) ST[i] = malloc(sizeof(int) * m);
    STableBuild(lg, arr, ST, n, m);
    int k;
    scanf("%d", &k);
    for (int i = 0; i < k; i++) {
        int l, r;
        scanf("%d %d", &l, &r);
        printf("%d\n", STGCD(ST, l, r, lg));
    }
    for (int i = 0; i < n; i++) free(ST[i]);
    free(ST);
    free(lg);
    free(arr);
    return 0;
}
