#include <stdio.h>

unsigned long peak(unsigned long nel, int (*less)(unsigned long i, unsigned long j)) {
    unsigned long start = 0;
    unsigned long end = nel - 1;
    unsigned long mid = end / 2 + start / 2;
    while (start < end){
        mid = end / 2 + start / 2;
        if (less(mid, mid + 1) > 0)
            start = mid + 1;
        else if(less(mid, mid -1))
            end = mid -1;
        else
            return mid;
    }
    return start;
}
