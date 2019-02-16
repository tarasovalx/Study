
#include <stdio.h>
#include <stdlib.h>
#include "elem.h"

struct Elem *searchlist(struct Elem *list, int k) {
    if (list->tag == INTEGER) {
        if (list->value.i == k)
            return list;
        else
            goto  next_el;
    }
    else if (list->tag == LIST){
        struct Elem * res = NULL;
        if (list->value.list)
            res = searchlist(list->value.list, k);
        if (res) 
            return res;
        else
            goto  next_el;
    }
    else{
        next_el:    if (list->tail)
            return searchlist(list->tail, k);
        else
            return NULL;
    }
}
