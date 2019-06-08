#include <stdio.h>

int strdiff(char *a, char *b)
{
    int i = 0;
    while (a[i] || b[i]) {
        if (a[i] != b[i]) {
            for(int j = 0; j<8; j++)
                if ((a[i] & (1 << j)) != (b[i] & (1 << j)))
                    return (i * 8 + j);
        }
        if(!a[i] || !b[i])
            break;
        i++;
    }
    return -1;
}
