#include <stdio.h>

int main() {
    unsigned long x;
    scanf("%lld", &x);
    if (x == 0) {
        printf("%d", 0);
        return 0;
    }
    if (x == 1) {
        printf("%d", 1);
        return 0;
    }
    unsigned long fib_el0 = 1;
    unsigned long fib_el1 = 2;

    while (fib_el0 + fib_el1 <= x) {
        unsigned long tmp = fib_el1;
        fib_el1 += fib_el0;
        fib_el0 = tmp;
    }
    while (fib_el0 != 0){
        if (fib_el1 <= x) {
            printf("%d", 1);
            x -= fib_el1;
            unsigned long tmp = fib_el0;
            fib_el0 = fib_el1 - fib_el0;
            fib_el1 = tmp;
        }
        else {
            printf("%d", 0);
            unsigned long tmp = fib_el0;
            fib_el0 = fib_el1 - fib_el0;
            fib_el1 = tmp;
        }
    }
    return 0;
}
