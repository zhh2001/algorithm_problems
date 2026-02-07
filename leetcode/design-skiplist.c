#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define MAX(a, b) (a) > (b) ? (a) : (b)
#define MAX_LV 32
#define P_FAC (RAND_MAX >> 2)

typedef struct SkiplistNode {
    int val, lv;
    struct SkiplistNode** next;
} SkiplistNode;

typedef struct {
    SkiplistNode* head;
    int level;
} Skiplist;

static SkiplistNode* createNode(int val, int lv) {
    SkiplistNode* node = malloc(sizeof(SkiplistNode));
    node->val = val;
    node->lv = lv;
    node->next = malloc(sizeof(SkiplistNode*) * lv);
    for (int i = 0; i < lv; node->next[i++] = NULL) {
    }
    return node;
}

static void freeNode(SkiplistNode* node) {
    free(node->next);
    free(node);
}

Skiplist* skiplistCreate() {
    Skiplist* sl = malloc(sizeof(Skiplist));
    sl->head = createNode(-1, MAX_LV);
    sl->level = 0;
    srand(time(NULL));
    return sl;
}

static int randLv() {
    int lv = 1;
    while (rand() < P_FAC && lv < MAX_LV)
        lv++;
    return lv;
}

bool skiplistSearch(Skiplist* sl, int target) {
    SkiplistNode* curr = sl->head;
    for (int i = sl->level - 1; i >= 0; i--)
        while (curr->next[i] && curr->next[i]->val < target)
            curr = curr->next[i];
    curr = curr->next[0];
    return curr && curr->val == target;
}

void skiplistAdd(Skiplist* sl, int num) {
    SkiplistNode *update[MAX_LV], *curr = sl->head;
    for (int i = sl->level - 1; i >= 0; i--) {
        while (curr->next[i] && curr->next[i]->val < num)
            curr = curr->next[i];
        update[i] = curr;
    }
    int lv = randLv();
    if (lv > sl->level) {
        for (int i = sl->level; i < lv; update[i++] = sl->head)
            ;
        sl->level = lv;
    }
    SkiplistNode* newNode = createNode(num, lv);
    for (int i = 0; i < lv; i++) {
        newNode->next[i] = update[i]->next[i];
        update[i]->next[i] = newNode;
    }
}

bool skiplistErase(Skiplist* obj, int num) {
    SkiplistNode* update[MAX_LV];
    SkiplistNode* curr = obj->head;
    for (int i = obj->level - 1; i >= 0; i--) {
        while (curr->next[i] && curr->next[i]->val < num)
            curr = curr->next[i];
        update[i] = curr;
    }
    curr = curr->next[0];
    if (!curr || curr->val != num)
        return false;
    for (int i = 0; i < obj->level && update[i]->next[i] == curr; i++)
        update[i]->next[i] = curr->next[i];
    freeNode(curr);
    while (obj->level > 1 && !obj->head->next[obj->level - 1])
        obj->level--;
    return true;
}

void skiplistFree(Skiplist* sl) {
    for (SkiplistNode *curr = sl->head, *tmp; curr; curr = tmp) {
        tmp = curr->next[0];
        freeNode(curr);
    }
    free(sl);
}

int main() {
    Skiplist* obj = skiplistCreate();
    skiplistAdd(obj, 1);
    skiplistAdd(obj, 2);
    skiplistAdd(obj, 3);
    bool param_1 = skiplistSearch(obj, 0);
    skiplistAdd(obj, 4);
    bool param_2 = skiplistSearch(obj, 1);
    bool param_3 = skiplistErase(obj, 0);
    bool param_4 = skiplistErase(obj, 1);
    bool param_5 = skiplistSearch(obj, 1);
    skiplistFree(obj);
    printf("%s\n", param_1 ? "true" : "false");
    printf("%s\n", param_2 ? "true" : "false");
    printf("%s\n", param_3 ? "true" : "false");
    printf("%s\n", param_4 ? "true" : "false");
    printf("%s\n", param_5 ? "true" : "false");
    return 0;
}
