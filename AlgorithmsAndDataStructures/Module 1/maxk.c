#include <stdio.h>

int main() {
    int n, k, max, current_sum  = 0;
    scanf("%d", &n);
    int arr[n];
    for (int*p = arr; p - arr < n; p++) scanf("%d", p);
    scanf("%d", &k);
    int *p0 = arr;
    int *p1 = arr + k - 1;
    for(int *p = p0; p <= p1; p++) current_sum+= *p;
    max = current_sum;
    while ((p1 - arr < n - 1)){
        current_sum = current_sum - *p0 + *(p1 + 1);
        if (current_sum> max) max = current_sum;
        p0++;
        p1++;
    }
    printf("%d", max);
    return 0;
}
