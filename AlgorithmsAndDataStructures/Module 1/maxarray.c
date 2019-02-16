#include <stdio.h>

int maxarray(void *base, unsigned long nel, unsigned long width, int (*compare)(void *a, void *b))
{
    void *max = base;
    unsigned long max_i = 0;
    for (unsigned long i = 0; i < nel ; i++){
        if (compare((void*)((char*)base + width * i) , max) > 0) {
            max = (void *)((char *)base + width * i);
            max_i = i;
        }
    }
    return max_i;
}
