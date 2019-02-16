#include <stdlib.h>
#include <stdio.h>
#include <malloc.h>
#include <memory.h>

struct Date{
    int Day, Month, Year
};

int keyDay(struct Date elem){
    return elem.Day - 1;
}

int keyMonth(struct Date elem) {
    return elem.Month - 1;
}

int keyYear(struct Date elem) {
    return elem.Year - 1970;
}

struct Date* DistributionSort(int (*key)(struct Date), struct Date *base, int m, int n){
    int count[m];
    struct Date *dest = malloc(sizeof(struct Date) * n);
    memset(dest, 0, sizeof(struct Date) * n);
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

struct Date* RadixSort(int (*keys[])(struct Date), int* m, int q, struct Date *base, int n) {
    for(int i = 0; i < q ; i++) base = DistributionSort(keys[i], base, m[i], n);
    return base;
}

int main(){
    int n;
    scanf("%d", &n);
    struct Date *arr = malloc(sizeof(struct Date) * n);
    for(int i =0; i < n; i++){
        scanf("%d", &arr[i].Year);
        scanf("%d", &arr[i].Month);
        scanf("%d", &arr[i].Day);
    }
    int (*keys[3])(struct Date) = {keyDay, keyMonth, keyYear};
    int m[3] = {32, 13, 61};
    arr = RadixSort(keys, m, 3, arr, n);
    for(int i = 0; i < n; i++){
        printf("%d ", arr[i].Year);
        printf("%d ", arr[i].Month);
        printf("%d\n", arr[i].Day);
    }
    free(arr);
    return 0;
}
