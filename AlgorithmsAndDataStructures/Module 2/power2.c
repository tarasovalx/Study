#include <stdio.h>
#include <math.h>

int Check_applied_mask_sum(int mask, int n, int* q){
    long long sum = 0;
    for (int i = 0; i < n; i++)
        if(mask & (1 << i))
            sum+=q[i];
    return  (sum && !(sum &(sum - 1)));
}

int Comb(int n, int m, int *q){
    int cnt = 0;
    for (int i = 0; i < m; i++) cnt += Check_applied_mask_sum(i, n, q);
    return cnt;
}

int main(){
    int n;
    scanf("%d", &n);
    int arr[n];
    for (int *p = arr; p - arr < n; p++) scanf("%d", p);
    printf("%d",Comb(n, pow(2,n), arr));
    return 0;
}
