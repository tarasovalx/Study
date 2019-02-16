#include <stdio.h>

void SelectionSort(int *base, int nel){
    int k, i, tmp;
    for(int j = nel - 1; j > 0; j--){
        k = j;
        for (i = j -1; i >= 0; i--){
            if(base[k] < base[i])
                k = i;
        }
        tmp = base[j];
        base[j] = base[k];
        base[k] = tmp;
    }
}

void swap(void *a, void *b, size_t width) {
    char tmp;
    for(int i = 0; i < width; i++){
        tmp = ((char*)a)[i];
        ((char*)a)[i] = ((char*)b)[i];
        ((char*)b)[i] = tmp;
    }
}

int Partition(int *base, int low, int high){
    int i = low;
    for(int j = low; j < high; j++){
        if (base[j] < base[high]){
            swap(base + i, base + j, sizeof(int));
            i++;
        }
    }
    swap(base + i, base + high, sizeof(int));
    return i;
}


void QuickSort(int *base, int low, int high, int m){
    if (high - low < m)
        SelectionSort(base + low, high - low + 1);
    else{
        while(low < high){
            int q = Partition(base, low, high);
            if (q - low < high - q){
                QuickSort(base, low, q - 1, m);
                low = q + 1;      
            } 
            else{
                QuickSort(base, q + 1, high, m);
                high = q - 1;   
            }
        }
    }
}
int main(){
    int n, m;
    scanf("%d", &n);
    scanf("%d", &m);
    int arr[n];
    for (int i = 0; i < n; i++) scanf("%d", arr + i);
    QuickSort(arr, 0, n - 1, 4);
    for (int i = 0; i < n; i++) printf("%d ", arr[i]);
    return 0;
}
