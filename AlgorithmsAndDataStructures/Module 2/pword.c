#include <stdio.h>
#include <malloc.h>
#include <string.h>

int *Prefix(char *s){
    int slen = strlen(s);
    int *p = malloc(sizeof(int)* slen);
    p[0] = 0;
    int t = 0, i = 1;
    for (i; i < slen; i++){
        for (t; t > 0 && (s[t] != s[i]); t = p[t - 1]);
        if (s[t] == s[i])
            t++;
        p[i] = t;
    }
    return p;
}

void KMP(char* S, char* T){
    int slen = strlen(S);
    int tlen = strlen(T);
    int *prefs = Prefix(S);
    int q = 0, k = 0;
    for(;k < tlen; k++){
        for(;q > 0 && S[q] != T[k]; q = prefs[q - 1]);
        if (S[q] == T[k]) {
            q++;
        }
        else{
            printf("%s", "no");
            free(prefs);
            return;
        }
        if (q == slen) {
            k = k - slen + 1;
            k += slen - 1;
        }
    }
    free(prefs);
    printf("%s", "yes");
}

int main(int argc, char **argv) {
    KMP(argv[1], argv[2]);
}
