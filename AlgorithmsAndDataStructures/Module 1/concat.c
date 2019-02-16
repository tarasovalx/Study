#include <stdio.h>
#include <string.h>
#include <malloc.h>
#define string_len 101U

char* concat(char **s, int n){
    int str_len = 0;
    for(int i = 0; i < n; i++) str_len += strnlen(*(s + i), sizeof(char) *string_len);
    char *res = malloc(str_len* sizeof(char) + 1);
    memset(res, 0, str_len* sizeof(char) + 1);
    for(int i = 0; i < n; i++) strncat(res, *(s + i), sizeof(char)* string_len);
    return res;
}

int main(void) {
    int n = 0;
    scanf("%d ", &n);
    char**s = malloc(n * sizeof(char*));
    for (int i = 0; i<n; i++)
        s[i] = malloc(string_len * sizeof(char));
    for (int i = 0; i<n; i++)
        gets(s[i]);
    char *res = concat(s, n);
    printf("%s",res);

    free(res);
    for (int i = 0; i<n; i++)
        free(*(s + i));
    free(s);
    return 0;
}
