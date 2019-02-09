#include <stdio.h>
#include <malloc.h>
#include <memory.h>
#include <math.h>
#include <stdlib.h>

struct ListElement{
    long key;
    long value;
    struct ListElement *next;
};

struct ListElement* Search(struct ListElement *elem, long key){
    for(;elem != NULL; elem = elem->next){
        if(elem->key == key) return elem;
    }
    return NULL;
}

void Insert(struct ListElement **list, long key, long value){
    struct ListElement *elem = malloc(sizeof(struct ListElement));
    elem->key = key;
    elem->value = value;
    elem->next = *list;
    *list = elem;
}

struct ListElement* TableSearch(long key, struct ListElement **table, long m){
    struct ListElement *res = Search(table[labs(key) % m], key);
    if (res != NULL) return res;
    return NULL;
}

void ListFree(struct ListElement *list){
    struct ListElement *elem = list;
    while (elem != NULL){
        struct ListElement *rec = elem;
        elem = elem->next;
        free(rec);
    }
}

void Delete(long key, struct ListElement **list){
    struct ListElement *buf = *list;
    for(; buf->next != NULL && buf->next->key != key; buf = buf->next);
    if (buf->next != NULL){
        struct ListElement *bufb = buf->next;
        buf->next = buf->next->next;
        free(bufb);
    }
}

void TableFree(struct ListElement **table, long m){
    for(int i = 0; i < m; i++) {
        if(table[i] != NULL) ListFree(table[i]);
    }
}

void Add (long key, long val, struct ListElement **table, long m){
    if (val) Insert(&table[labs(key) % m], key, val);
    else if (table[labs(key) % m] != NULL){
        Delete(key, &table[labs(key) % m]);
    }
}

int main(){
    long n, m, a, b;
    char cmd[10];
    scanf("%ld\n%ld", &n, &m);
    struct ListElement **table = malloc(sizeof(struct ListElement*) * m);
    memset(table, NULL, sizeof(struct ListElement*) * m);
    for (int i = 0; i < n; i++) {
        scanf("%s", cmd);
        if (cmd[1] == 'T'){
            scanf("%ld", &a);
            struct ListElement *res = TableSearch(a, table, m);
            if (res != NULL) printf("%ld\n", res->value);
            else printf("0\n");
        } else{
            scanf("%ld %ld", &a, &b);
            Add(a, b, table, m);
        }
    }
    TableFree(table, m);
    free(table);
}
