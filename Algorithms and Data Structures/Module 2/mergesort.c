#include <stdlib.h>
#include <stdio.h>
#include <math.h>
#define threshold 5

int compare(int a, int b){
    return (abs(a) - abs(b));
}

void Merge(int *base, int (*compare)(int a, int b), int k, int l, int m) {
    int t[m - k + 1];
    int i = k, j = l + 1, h;
    for(h = 0; h < m - k + 1; h++){
        if (j <= m && (i == l + 1 || (compare(base[j], base[i]) < 0))){
            t[h] = base[j];
            j++;
        }else{
            t[h] = base[i];
            i++;
        }
    }
    for(h--; h >= 0 ; h--) base[h + k] = t[h];
}

void InsertSort(int *base, int nel, int (*compare)(int a, int b)) {
    int i, j, elem;
    for(i = 0; i < nel; i++) {
        elem = base[i];
        for(j = i - 1; j >= 0 && compare(base[j], elem) > 0; j--)
            base[j + 1] = base[j];
        base[j + 1] = elem;
    }
}

void MergeSort(int *base, int (*compare)(int a, int b), int low, int high){
    if (high - low > threshold){
        int med = (low + high) / 2;
        MergeSort(base, compare, low, med);
        MergeSort(base, compare, med + 1, high);
        Merge(base, compare, low, med, high);
    } else InsertSort(base + low, high - low + 1, compare);
}

int main(){
    int n;
    scanf("%d", &n);
    int arr[n];
    for(int i = 0; i < n; i++) scanf("%d", arr + i);
    MergeSort(arr, compare, 0, n - 1);
    for(int i = 0; i < n; i++) printf("%d ", arr[i]);
    return 0;
}
