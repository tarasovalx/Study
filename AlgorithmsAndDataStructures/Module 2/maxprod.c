#include <stdio.h>
#include <malloc.h>
#include <math.h>
#include <string.h>
#include <stdlib.h>

int* Kadane(double const *base, int n){
    int l = 0, r = 0, start = 0;
    double maxsum = base[0], sum = 0;
    for (int i = 0; i < n;){
        sum += base[i];
        if ((sum - maxsum) > 0.001){
            maxsum = sum;
            l = start;
            r = i;
        }
        i++;
        if (sum < 0){
            sum = 0;
            start = i;
        }
    }
    int *res = malloc(sizeof(int) * 2);
    res[0] = l;
    res[1] = r;
    return res;
}

int main(){
    int n;
    scanf("%d ", &n);
    double *values = malloc(sizeof(double) * n);
    int a = 0;
    int b = 0;
    for (int j = 0; j < n; j++){
        scanf("%d/%d", &a, &b);
        values[j] = log((double)a/(double)b);
    }
    int *res = Kadane(values, n);
    printf("%d %d", res[0], res[1]);
    free(res);
    free(values);
}
