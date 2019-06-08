#include <stdio.h>

unsigned long binsearch(unsigned long nel, int (*compare)(unsigned long i)){
    unsigned long start = 0;
    unsigned long end = nel - 1;
    unsigned long mid = (end + start) / 2;
    int com = compare(mid);
    while (start <= end){
        mid = (end + start) / 2;
        com = compare(mid);
        if (com > 0)
            end = mid - 1;
        else if (com < 0)
            start = mid +1;
        else return mid;
    }
    return  nel;
}
