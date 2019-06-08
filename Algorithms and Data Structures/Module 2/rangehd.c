#include <stdio.h>
#include <malloc.h>
#include <string.h>

#define MIN(x, y) (((x) < (y)) ? (x) : (y))

int BuildTree(char const *arr, int l, int r, int *T, int n) {
    int cnt = 0, m, bound = MIN(r, n);
    for (; l < bound; l = m + 1) {
        m = (l + r) / 2;
        cnt = cnt ^ BuildTree(arr, l, m, T, n);
    }
    if (r < n) {
        cnt = cnt ^ (1 << (arr[r] - 97));
        T[r] = cnt;
    }
    return cnt;
}

void UpdateTree(int *T, int i, int value, int n) {
    for (; i < n; i = (i | (i + 1))) {
        T[i] = T[i] ^ value;
    }
}

int Query(int const *T, int i) {
    int res = 0;
    for (; i >= 0; i = ((i & (i + 1)) - 1)) {
        res = res ^ T[i];
    }
    return res;
}

int QueryXOR(int a, int b, int const *T) {
    return Query(T, b) ^ Query(T, a - 1);
}

int RoundToNextPowOf2(int x) {
   x = x - 1;
   x = x | (x >> 1);
   x = x | (x >> 2);
   x = x | (x >> 4);
   x = x | (x >> 8);
   x = x | (x >> 16);
   return x + 1;
}

int main() {
    int k, n, a, b, value = 0;
    char *arr = malloc(2000001), cmd[4], *buf = malloc(1000001);
    gets(arr);
    n = strlen(arr);
    int *tree = malloc(sizeof(int) * n);
    memset(tree, 0, sizeof(int) * n);
    BuildTree(arr, 0, RoundToNextPowOf2(n) - 1, tree, n);
    scanf("%d ", &k);
    for (int i = 0; i < k; i++) {
        scanf("%s", cmd);
        if (cmd[0] == 'H') {
            scanf("%d %d", &a, &b);
            int res = QueryXOR(a, b, tree);
            if (!(res & (res - 1)))
                printf("YES\n");
            else
                printf("NO\n");
        } else {
            scanf("%d ", &a);
            gets(buf);
            for (int i = 0; buf[i]; a++, i++) {
                value = (1 << (buf[i] - 97)) ^ (1 << (arr[a] - 97));
                UpdateTree(tree, a, value, n);
                arr[a] = buf[i];
            }
        }
    }
    free(tree);
    free(arr);
    free(buf);
    return 0;
}
