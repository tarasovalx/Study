#include <stdio.h>
#include <string.h>
#include <malloc.h>

struct NODE {
    struct NODE *arcs[26];
    int IsTerminal, cnt;
};

struct NODE *InitNODE() {
    struct NODE *elem = malloc(sizeof(struct NODE));
    elem->IsTerminal = 0;
    elem->cnt = 0;
    memset(elem->arcs, NULL, sizeof(struct NODE *) * 26);
    return elem;
}

void Insert(struct NODE *root, char *k) {
    struct NODE *elem = root;
    int slen = strlen(k);
    for (int i = 0; i < slen; i++) {
        if (elem->arcs[k[i] - 'a'] == NULL) elem->arcs[k[i] - 'a'] = InitNODE();
        elem->cnt++;
        elem = elem->arcs[k[i] - 'a'];
    }
    elem->cnt++;
    if (!elem->IsTerminal) elem->IsTerminal = 1;
    else {
        elem = root;
        for (int i = 0; i < slen; i++) {
            elem->cnt--;
            elem = elem->arcs[k[i] - 'a'];
        }
        elem->cnt--;
    }
}

void Prefix(struct NODE *root, char *k) {
    struct NODE *elem = root;
    int slen = strlen(k);
    for (int i = 0; i < slen; i++) {
        if (elem->arcs[k[i] - 'a'] == NULL) {
            printf("%d\n", 0);
            return;
        }
        elem = elem->arcs[k[i] - 'a'];
    }
    printf("%d\n", elem->cnt);
}

void Delete(struct NODE *root, char *k) {
    struct NODE *elem = root;
    int slen = strlen(k);
    for (int i = 0; i < slen; i++) {
        elem->cnt--;
        elem = elem->arcs[k[i] - 'a'];
    }
    elem->cnt--;
    elem->IsTerminal = 0;
}

void FreeTrie(struct NODE *root){
    if (root == NULL) return;
    for (int i = 0; i < 26; i++) FreeTrie(root->arcs[i]);
    free(root);
}

int main() {
    struct NODE *trie = InitNODE();
    int n;
    scanf("%d", &n);
    char cmd[10];
    char k[100000];
    for (int i = 0; i < n; i++) {
        scanf("%s %s", cmd, k);
        if (cmd[0] == 'I') Insert(trie, k);
        if (cmd[0] == 'P') Prefix(trie, k);
        if (cmd[0] == 'D') Delete(trie, k);
    }
    FreeTrie(trie);
}
