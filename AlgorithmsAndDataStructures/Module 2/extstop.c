#include <stdio.h>
#include <malloc.h>
#include <string.h>
#include <math.h>

int **Delta1(char* S, int size){
    int slen = strlen(S);
    int** d = malloc(sizeof(int*) * slen);
    for(int i = 0; i < slen; i++){
        d[i] = malloc(sizeof(int) * size);
        for(int j = 0; j < size; j++) d[i][j] = slen;
    }
    for (int i = 0; i < slen; i++){
        for(int j = 0; j < slen; j++)
            d[i][S[j] - 97] = slen - j - 1;
    }
    return d;
}

int BMSubst(char* S, int size, char* T){
    int** d = Delta1(S, size);
    int slen = strlen(S), i, k;
    int tlen = strlen(T);
    for(k = slen - 1; k < tlen;){
        for(i = slen - 1; T[k] == S[i]; i--){
            if (!i){
                for(int q = 0; q < slen; q++) free(d[q]);
                free(d);
                return k;
            }
            k--;
        }
        if (d[i][T[k] - 97] > slen - i) k += d[i][T[k] - 97];
        else k+= slen - i;
    }
    for(int q = 0; q < slen; q++) free(d[q]);
    free(d);
    return tlen;
}

int main(int argc, char **argv) {
    printf("%d", BMSubst(argv[1], 26, argv[2]));
}
