#include <stdio.h>
#include <malloc.h>
#include <string.h>

int Max(int a, int b) {
    return (a > b) ? a : b;
}

int Min(int a, int b) {
    return (a < b) ? a : b;
}

struct Stack {
    int *data;
    int cap;
    int top;
};

void Init(struct Stack *s, int n) {
    s->data = malloc(sizeof(int) * (n + 1));
    s->top = 0;
    s->cap = n;
}

void Push(struct Stack *s, int elem) {
    if (s->cap == s->top) printf("Error: overflow");
    s->data[s->top] = elem;
    s->top++;
}

int Pop(struct Stack *s) {
    if (!s->top) printf("Error: devastation");
    s->top--;
    return s->data[s->top];
}

enum Comands {
    CONST, ADD, SUB, MUL, DIV, MAX, MIN, NEG, DUP, SWAP
};

int GetCMD(char *cmd) {
    char comands[10][6] = {"CONST", "ADD", "SUB", "MUL", "DIV", "MAX", "MIN", "NEG", "DUP", "SWAP"};
    for (int i = 0; i < 10; i++) {
        if (strcmp(cmd, comands[i]) == 0) return i;
    }
    return -1;
}

int main() {
    int n;
    scanf("%d", &n);
    char cmd[10];
    int a, b;
    struct Stack s;
    Init(&s, n);
    for (int i = 0; i < n; i++) {
        scanf("%s", cmd);
        enum Comands res = GetCMD(cmd);
        switch (res) {
            case CONST:
                scanf("%d", &a);
                Push(&s, a);
                break;
            case ADD:
                Push(&s, Pop(&s) + Pop(&s));
                break;
            case SUB:
                Push(&s, Pop(&s) - Pop(&s));
                break;
            case MUL:
                Push(&s, Pop(&s) * Pop(&s));
                break;
            case DIV:
                Push(&s, Pop(&s) / Pop(&s));
                break;
            case MAX:
                Push(&s, Max(Pop(&s), Pop(&s)));
                break;
            case MIN:
                Push(&s, Min(Pop(&s), Pop(&s)));
                break;
            case NEG:
                Push(&s, -Pop(&s));
                break;
            case DUP:
                b = Pop(&s);
                Push(&s, b);
                Push(&s, b);
                break;
            case SWAP:
                a = Pop(&s);
                b = Pop(&s);
                Push(&s, a);
                Push(&s, b);
                break;
        }
    }
    for (int i = 0; i < s.top; i++) printf("%d\n", s.data[i]);
    free(s.data);
    return 0;
}
