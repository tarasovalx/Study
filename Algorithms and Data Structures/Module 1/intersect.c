#include <stdio.h>

int main() {
    int a, b, n, res, inp = 0;
    res = 0;
    a = 0;
    b = 0;
    scanf("%d", &n);

    for (int  i = 0; i < n; i++){
        scanf("%d ", &inp);
        a = a | (1 << inp);
    }
    scanf("%d", &n);

    for (int  i = 0; i < n; i++){
        scanf("%d", &inp);
        b = b | (1 << inp);
    }
    res = a & b;

    for (int i = 0; i < 32; i ++){
        if (res & (1 << i)){
            printf("%d ", i);
        }
    }
    return 0;
}
