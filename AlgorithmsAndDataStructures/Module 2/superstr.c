#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int is_suffix(char const *str, char const *suf, int str_len, int suf_len){
    for(int i = 0; i < str_len; i++){
        if (strcmp(str + i, suf) == 0)
            return 1;
    }
    return 0;
}

int overlap(char const *a, char const *b){
    int a_len = strlen(a);
    int b_len = strlen(b);
    int overlap = 0;
    char suf[b_len + 1];

    for(int i = 1; i < b_len + 1; i++){
        suf[0] = 0;
        strncat(suf, b, i);
        if (is_suffix(a, suf, a_len, b_len)){
            overlap = i;
        }
    }
    return overlap;
}

int main(int argc, char const *argv[])
{
    int n;
    scanf("%d ", &n);
    char **strarr = malloc(sizeof(char*)*n);
    for (int i = 0; i<n; i++)
        strarr[i] = malloc(sizeof(char)*100);
    for(int i = 0; i < n; i++){
        memset(strarr[i], 0, 100);
        gets(strarr[i]);
    }
    int arr_res[n][n];
    for(int i = 0; i < n; i++){
        for(int j = 0; j < n; j++){
            if (i != j){
                arr_res[i][j] = overlap(strarr[i], strarr[j]);
            }
            else{
                arr_res[i][j] = 0;
            }
        }
    }

    int len = 0;
    int total_overlapped = 0;
    int notallow[n][n];

    for(int i = 0; i < n; i++){
        for(int j = 0; j < n; j++){
            if(i==j) notallow[i][j] = 1;
            else notallow[i][j] = 0;
        }
    }
    int cnt = n - 1;
    while (cnt){
        int max = -1;
        int max_i = -1;
        int max_j = -1;

        for(int i = 0; i < n; i++){
            for(int j = 0; j < n; j++){
                if(notallow[i][j]) continue;
                if(arr_res[i][j] > max){
                    max = arr_res[i][j];
                    max_i = i;
                    max_j = j;
                }
            }
        }

        if(max == -1) break;

        total_overlapped += max;
        for(int i = 0; i < n; i++){
            notallow[max_i][i] = 1;
            notallow[i][max_j] = 1;
        }
        cnt--;
    }

    for(int i = 0; i < n; i++){
        len+=strlen(strarr[i]);
    }
    printf("%d", len - total_overlapped);

    for (int i = 0; i < n; i++)
        free(strarr[i]);
    free(strarr);
}
