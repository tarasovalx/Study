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

int main(int argc, char **argv){
    int slen = strlen(argv[1]);
    int* res = Prefix(argv[1]);
    for(int i = 2; i <= slen; i++)
        if(res[i -1] && i%(i-res[i - 1]) == 0)
            printf("%d %d\n", i, i/(i-res[i - 1]));
    free(res);
    return 0;
}
