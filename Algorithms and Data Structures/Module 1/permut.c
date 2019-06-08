
#include <stdio.h>

int main() {
    long long a1[8];
    long long a2[8];
    unsigned char mask = 0;

    for (long long *p = a1; p - a1 < 8; p++) scanf("%ld", p);
    for (long long *p = a2; p - a2 < 8; p++) scanf("%ld", p);

    for (unsigned short i = 0; i < 8;i++) {
        for (unsigned short j = 0; j < 8;j++){
            if (a1[i] == a2[j] && ((mask & (1 << j)) | (255 - (1 << j))) != 255){
                mask = mask | (1 << j);
                break;
            }
        }
    }
    if (mask == 255) printf("yes");
    else printf("no");
}
