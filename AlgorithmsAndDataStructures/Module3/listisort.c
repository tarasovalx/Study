#include <stdio.h>
#include <malloc.h>

struct Elem {
    struct Elem *prev, *next;
    int v;
};

struct Elem *InitList() {
    struct Elem *list = malloc(sizeof(struct Elem));
    list->next = list;
    list->prev = list;
    return list;
}

void InsertAfter(struct Elem *list, struct Elem *adding) {
    struct Elem *buf = list->next;
    list->next = adding;
    adding->prev = list;
    adding->next = buf;
    buf->prev = adding;
}

void Delete(struct Elem *elem) {
    struct Elem *bufPrev = elem->prev;
    struct Elem *bufNext = elem->next;
    bufPrev->next = bufNext;
    bufNext->prev = bufPrev;
}

void InsertionSort(struct Elem *list) {
    struct Elem *i = list->next->next, *j;
    for (; i != list; i = i->next) {
        for (j = i->prev; j != list && i->v < j->v; j = j->prev);
        Delete(i);
        InsertAfter(j, i);
    }
}

int main() {
    int n;
    scanf("%d", &n);
    struct Elem *list = InitList();
    for (int i = 0; i < n; i++) {
        struct Elem *e = malloc(sizeof(struct Elem));
        scanf("%d", &(e->v));
        InsertAfter(list, e);
    }
    InsertionSort(list);
    struct Elem *p = list->next;
    for (; p != list; p = p->next) {
        printf("%d ", p->v);
    }
    p = list->next;
    while (p != list) {
        p = p->next;
        free(p->prev);
    }
    free(list);
}
