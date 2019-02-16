#include <stdio.h>
#include <string.h>

void csort(char*, char*);

int main (void){
    char s[1000];
    char dest[1000] = {0};
    gets(s);
    csort(s, dest);
    printf("%s", dest);
}

void csort(char *src, char *dest) {
    unsigned long n = 0;
    for(unsigned long locked = 0, i = 0; src[i]; i++){
        if (src[i] != ' ' && !locked) {
            n++;
            locked = 1;
        }
        else if (src[i] == ' ') locked = 0;
    }
    unsigned long count[n];
    unsigned long count_buf[n];

    for(unsigned long locked = 0, j = 0, i = 0; src[i]; i++){
        if (src[i] != ' ' && !locked) {
            count[j] = 0;
            count_buf[j] = i;
            locked = 1;
            j++;
        }
        else if (src[i] == ' ') locked = 0;
    }

    for(unsigned long j = 0; j < n - 1; j++){
        char a[100] = {0};
        for (unsigned long k = 0; src[count_buf[j]+k] != ' ' && src[count_buf[j]+k]; k++) a[k] = src[count_buf[j]+k];
        for(unsigned long i = j + 1; i < n; i++){
            char b[100] = {0};
            for (unsigned long k = 0; src[count_buf[i]+k] != ' ' && src[count_buf[i]+k]; k++) b[k] = src[count_buf[i]+k];
            if(strlen(a) > strlen(b)) count[j] += strlen(b) + 1;
            else count[i] += strlen(a) + 1;
        }
    }

    for(unsigned long i = 0; i < n; i++){
        for(unsigned long k = 0, j = count[i]; ; k++, j++) {
            if (src[count_buf[i] + k] == ' ' || !src[count_buf[i] + k]){
                dest[j] = ' ';
                break;
            }
            else dest[j] = src[count_buf[i] + k];
        }
    }
}
