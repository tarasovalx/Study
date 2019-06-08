#include "stdio.h"

int main() {
    unsigned long a, b, m;
    scanf("%ld%ld%ld", &a, &b, &m);
    unsigned long res = 0;

    for (int i = 63; i > 0; i--){
        res += (a % m * (b >> i & 1)) % m;
        res *= 2;
        res = res % m;
    }

    if(b % 2) res += a;
    printf("%ld", res % m);
    return 0;
}
