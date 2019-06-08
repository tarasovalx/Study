#include <stdio.h>
#define abc_len 26
#define abc_threshold 97

void ABC_DistributionSort(char *base){
    int abc[abc_len]= {0};
    int i = 0, j = 0;
    for(; base[i]; i++)
        abc[base[i] - abc_threshold]++;
    for(i = 0; base[i] && (j <= abc_len);){
        if (abc[j]) {
            base[i] = (char) (j + abc_threshold);
            abc[j]--;
            i++;
        }
        else
            j++;
    }
}

int main(){
    char arr[1000001];
    gets(arr);
    ABC_DistributionSort(arr);
    printf("%s", arr);
    return 0;
}
