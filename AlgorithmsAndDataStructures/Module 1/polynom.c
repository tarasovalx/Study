#include <stdio.h>

int main() {
    long n, x;
    scanf("%ld%ld", &n, &x);

    long long arr[n+1];
    for (long long *p = arr; p - arr < n+1; p++) scanf("%ld", p);

    long long res = arr[0];
    long long der = 0;
    for (long i = 1; i <= n; ++i) {
        res *= x;
        res += arr[i];
        der = der*x +arr[i-1]*(n-i+1);
    }
    printf("%lld %lld", res, der);
    return 0;
}
