#include <stdio.h>

void revarray(void *base, unsigned long nel, unsigned long width)
{
    char* p = (char*)base;
    for (unsigned long i = 0; i < nel/2; i++) {
        char *p1 = p + i*width;
        char *p2 = p + nel*width - (i+1)*width;
        for (unsigned long k = 0; k < width; k++) {
            char tmp = p1[k];
            p1[k] = p2[k];
            p2[k] = tmp;
        }
    }
}
