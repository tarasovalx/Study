#include <stdio.h>
#include <malloc.h>
#include <string.h>
#include <math.h>

#define ANSI_threshold 97
#define MAX(a, b) (((a)>(b))?(a):(b))

int *Delta1(char *S, int size) {
    int *d = malloc(sizeof(int) * size);
    int slen = strlen(S);
    for (int a = 0; a < size; a++) d[a] = slen;
    for (int j = 0; j < slen; j++) d[S[j] - ANSI_threshold] = slen - j - 1;
    return d;
}

int *Suffix(char *S) {
    int slen = strlen(S);
    int *res = malloc(sizeof(int) * slen);
    int t = slen - 1;
    res[slen - 1] = t;
    for (int i = slen - 2; i >= 0; i--) {
        for (; (t < slen - 1) && S[t] != S[i]; t = res[t + 1]);
        if (S[t] == S[i]) t--;
        res[i] = t;
    }
    return res;
}

int *Delta2(char *S) {
    int slen = strlen(S);
    int *res = malloc(sizeof(int) * slen);
    int *d = Suffix(S);
    int t = d[0], i = 0;
    for (; i < slen; i++) {
        for (; t < i; t = d[t + 1]);
        res[i] = -i + t + slen;
    }
    for (i = 0; i < slen - 1; i++) {
        for (t = i; t < slen - 1;) {
            t = d[t + 1];
            if (S[i] != S[t]) {
                res[t] = -i  - 1 + slen;
            }
        }
    }
    free(d);
    return res;
}

void BMSubst(char *S, int size, char *T) {
    int *d = Delta1(S, size);
    int *d2 = Delta2(S);
    int slen = strlen(S), i, k;
    int tlen = strlen(T);
    for (k = slen - 1; k < tlen;) {
        for (i = slen - 1; T[k] == S[i]; k-- ,i--) {
            if (!i){
                printf("%d ", k);
                break;
            }
        }
        k+= MAX(d[T[k] - ANSI_threshold], d2[i]);
    }
    free(d);
    free(d2);
}

int main(int argc, char **argv) {
    BMSubst(argv[1], 26, argv[2]);
    return 0;
}
