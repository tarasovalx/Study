#include <stdio.h>
#include <malloc.h>
#include <string.h>
#define MaxStringLen 100

struct Elem {
    struct Elem  *next;
    char *word;
};

struct Elem *InitList() {
    struct Elem *list = malloc(sizeof(struct Elem));
    list->word = malloc(sizeof(char) * MaxStringLen);
    list->next = NULL;
    memset(list->word, 0, MaxStringLen);
    return list;
}

void Insert(struct Elem *last, struct Elem *elem){
    last->next = elem;
}

void bsort(struct Elem *list) {
    struct Elem *t = list;
    while (t != NULL){
        t = NULL;
        for(struct Elem *p = list; p->next != NULL; p = p->next){
            if(strlen(p->word) > strlen(p->next->word)){
                char *tmp = p->word;
                p->word = p->next->word;
                p->next->word = tmp;
                t = p;
            }
        }
    }
}

void ListFree(struct Elem *list){
    struct Elem *elem = list;
    while (elem != NULL){
        struct Elem *rec = elem;
        elem = elem->next;
        free(rec->word);
        free(rec);
    }
}

int main() {
    char buf[MaxStringLen] = {0};
    struct Elem *list = InitList();
    char tmp[10000] = {0};
    gets(tmp);
    struct Elem *last = list;
    for (int i = 0, j = 0; tmp[i]; i++){
        if (tmp[i] != ' '){
            buf[j] = tmp[i];
            j++;
        }
        if((tmp[i] == ' ' || !tmp[i + 1]) && j > 0 ){
            struct Elem *e = InitList();
            strcpy(e->word, buf);
            Insert(last, e);
            last = last->next;
            j = 0;
            memset(buf, 0, MaxStringLen);
        }
    }
    struct Elem *p = list;
    bsort(list);
    p = list;
    for (; p != NULL; p = p->next) printf("%s ", p->word);
    ListFree(list);
}
