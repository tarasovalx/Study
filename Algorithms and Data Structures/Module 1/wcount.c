#include <stdio.h>

int wcount(char* s){
    int cnt = 0;
    int flag = 1;
    for(int i = 0; s[i]; i++){
        if(s[i] == 32)
            flag = 1;
        if(flag && (s[i] != 32)){
            flag = 0;
            cnt++;
        }
    }
    return cnt;
}

int main() {
    char s[100] = "";
    gets(s);
    printf("%d", wcount(s));
    return 0;
}
