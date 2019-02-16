#include <stdlib.h>
#include <stdio.h>
#include <malloc.h>
#include <memory.h>

union Int32 {
    int x;
    unsigned char bytes[4];
};

int keyHighByte(union Int32 elem) {
    //printf("%d %d \n", elem.x ,(int)elem.bytes[3]);
    if ((int)elem.bytes[3] > 127)
        return (int)elem.bytes[3] % 128;
    else return  (int)elem.bytes[3] + 127;
}
int key3Byte(union Int32 elem) {
    return (int)elem.bytes[2];
}
int key2Byte(union Int32 elem) {
    return (int)elem.bytes[1];
}
int key1Byte(union Int32 elem) {
    return (int)elem.bytes[0];
}

union Int32* DistributionSort(int (*key)(union Int32), union Int32 *base, int m, int n){
    int count[m];
    union Int32  *dest = malloc(sizeof(union Int32) * n);
    memset(dest, 0, sizeof(union Int32 ) * n);
    for (int i = 0; i < m; i++) count[i] = 0;
    for(int j = 0; j < n; j++){
        count[key(base[j])]++;
    }
    for(int i = 1; i < m; i++)
        count[i] += count[i - 1];
    for(int j = n - 1; j >= 0; j--){
        int i =  count[key(base[j])] - 1;
        count[key(base[j])] = i;
        dest[i] = base[j];
    }
    free(base);
    return dest;
}

union Int32* RadixSort(int (*keys[])(union Int32), int* m, int q, union Int32 *base, int n) {
    for(int i = 0; i < q ; i++) base = DistributionSort(keys[i], base, m[i], n);
    return base;
}

int main(){
    int n;
    scanf("%d", &n);
    union Int32 *arr = malloc(sizeof(union Int32) * n);
    for(int i = 0; i < n; i++) scanf("%d", &arr[i].x);
    int (*keys[4])(union Int32) = {key1Byte, key2Byte, key3Byte, keyHighByte};
    int m[4] = {256, 256, 256, 256};
    arr = RadixSort(keys, m, 4, arr, n);
    for (int i = 0; i < n; i++) printf("%d ", arr[i].x);
    free(arr);
    return 0;
}
