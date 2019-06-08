#include <stdio.h>
#include <malloc.h>
#include <string.h>

char *fibstr(int n)
{
    unsigned long fib_el0 = 1;
    unsigned long fib_el1 = 2;

    for (int i = 2; i < n; i++) {
        unsigned long tmp = fib_el1;
        fib_el1 += fib_el0;
        fib_el0 = tmp;
    }

    char *s1 = malloc(sizeof(char)*(fib_el0 + 1));
    memset(s1, 0, sizeof(char)*fib_el0 + 1);
    s1[0] = 'a';
    char *s2 = malloc(sizeof(char)*(fib_el1 + 1));
    memset(s2, 0, sizeof(char)*fib_el1 + 1);
    s2[0] = 'b';
    char *s3 = malloc((fib_el0 + fib_el1 + 1)* sizeof(char));

    for (int i = 1; i < n - 1; i++) {
        memset(s3, 0, (fib_el0 + fib_el1 + 1)* sizeof(char));
        strncat(s3, s1, fib_el0);
        strncat(s3, s2, fib_el1);
        strncpy(s1,s2, fib_el0);
        strncpy(s2,s3, fib_el1);
    }
    free(s1);
    free(s2);

    if (n == 1){
        s3[0] = 'a';
        s3[1] = '\0';
    }
    else if (n == 2){
        s3[0] = 'b';
        s3[1] = '\0';
    }
    return s3;
}

int main() {
    int x = 0;
    scanf("%d", &x);
    char *res = fibstr(x);
    printf("%s", res);
    free(res);
    return 0;
}
