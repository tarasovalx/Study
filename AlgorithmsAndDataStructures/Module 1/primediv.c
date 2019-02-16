#include <stdio.h>
#include <stdlib.h>
#include <math.h>

int main() {
    int x;
    scanf("%d", &x);
    if (x < 0) x = abs(x);
    int length = (int)(sqrt((double) x) - 1);
    int arr[length];
    int max_prime_div = 0;

    for (int i = 0; i < length; i++) {
        arr[i] = i + 2;
    }

    for (int i = 0; i < length; i++) {
        if (arr[i]) {
            for (int j = (arr[i] * arr[i])-2; j < length; j += arr[i]) {
                arr[j] = 0;
            }
        }
    }

    for (int i = 0; i < length; i++){
        if (arr[i] && x % arr[i] == 0 && arr[i] > max_prime_div) {
            max_prime_div = arr[i];
            int tmp = x;
            int buf = 1;
            for (int j = 0; j < length; j++) {
                while (arr[j] && !(tmp % arr[j])) {
                    buf *= arr[j];
                    tmp /= arr[j];
                }
            }
            if ((buf != 1) && ((x / buf) > max_prime_div)) max_prime_div = x / buf;
        }
    }

    if (max_prime_div) printf("%d", max_prime_div);
    else printf("%d", x);
}
