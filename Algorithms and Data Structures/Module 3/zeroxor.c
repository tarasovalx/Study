#include <stdio.h>
#include <malloc.h>
#include <memory.h>
#include <math.h>
#define TableSize 101111   //Prime number
#define HashBase 571       //Prime number
#define HashModule  101111 //Prime number
#define dist 6
#define x 143
#define y 7

unsigned int Hash(int key){
    unsigned int res = 1;
    unsigned int Key = (unsigned int)key;
    for(int i = 32; i >= 0; i--) {
        res *= HashBase;
        res += (Key & (1 << i))? x : y;
        res %= HashModule;
    }
    return (res % TableSize);
}

struct Pair{
    int key;
    int value;
};

int main() {
    int n, a;
    scanf("%d", &n);
    struct Pair **table = malloc(sizeof(struct Pair*) * TableSize * dist);
    memset(table, 0, sizeof(struct Pair*) * TableSize * dist);
    int PrefixXOR = 0;
    for(int i = 0; i < n; i++){
        scanf("%d",&a);
        PrefixXOR ^= a;
        if (table[Hash(PrefixXOR)* dist] == NULL){
            table[Hash(PrefixXOR)* dist] = malloc(sizeof(struct Pair));
            table[Hash(PrefixXOR)* dist]->value = (PrefixXOR)? 0 : 1;
            table[Hash(PrefixXOR)* dist]->key = PrefixXOR;
        }
        else if (table[Hash(PrefixXOR)* dist]->key == PrefixXOR){
            table[Hash(PrefixXOR)* dist]->value++;
        }
        else{
            int j = Hash(PrefixXOR);
            int k = 1;
            for(; table[j * dist + k] != NULL; k++){
                if (table[j * dist + k]->key == PrefixXOR) break;
            };
            //if (k > 3) printf("PANIC\n");
            if(table[j * dist + k] == NULL){
                table[j * dist + k] = malloc(sizeof(struct Pair));
                table[j * dist + k]->key = PrefixXOR;
                table[j * dist + k]->value = (PrefixXOR) ? 0 : 1;
            } else {
                table[j * dist + k]->value++;
            }
        }
    }
    int res = 0;
    for (int i = 0; i < TableSize * dist; i++){
        struct Pair *elem = table[i];
        if (elem == NULL) continue;
        int val = elem->value;
        res += val * (val + 1) / 2;
        free(elem);
    }
    free(table);
    printf("%d\n", res);
}
