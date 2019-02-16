#include <stdio.h>

void shellsort(unsigned long nel, int (*compare)(unsigned long i, unsigned long j), void (*swap)(unsigned long i, unsigned long j))
{
    unsigned long i,j,k;
    long a = 1, b = 2;
    while(a + b < nel){
        long t = a;
        a = b;
        b += t;
    }
    for(k = (unsigned long)b; b > 0; k = (unsigned long)b){
        for(i = k; i < nel; i++)
            for(j = i; j >= k && (compare(j, j - k) == -1); j-=k)
                swap(j, j - k);
        long t = b;
        b = a;
        a = t - b;
        if (k == 1)
            break;
    }
}
