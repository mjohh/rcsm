#include <stdio.h>
#include <assert.h>
#include <stdlib.h>
#include <memory.h>

char* s = "*2\r\n*3\r\n:1\r\n:2\r\n:3\r\n*2\r\n+Foo\r\n-Bar\r\n";

struct node {
    int type;
    struct node * parent;
    int nchildren;// children number
    int cnt;      // children count
    struct node ** children;
    void* val;
};

struct node *stack[9];
int top = -1;

enum nodetype{
    ARRAY,
    INTEGER,
    STATUS,
    ERROR,
};

void prints(char* bgn, char* end){
    char str[64];
    memcpy(str, bgn, end-bgn+1);
    str[end-bgn+1]=0;
    printf("%s", str);
}
char* getoneline(char* s, char** nextline){
    if (NULL == s)
        return NULL;

    char* p = s;
    while(*p){
        if (*p == '\r' && *(p+1) == '\n'){
            *nextline = p+2;
            return s;
        }
        p++;
    }
    return NULL;
}

long readlong(char* line){
    char* p = line+1; //skip '*'
    long l = 0;
    while(*p != '\r'){
       char digit = *p-'0';
       assert(digit<=9 && digit>=0);
       l = l*10 + digit;
       p++;
    }
    return l;
}

char* readstatus(char* line, int* len){
    char* p = line+1; //skip '-'
    int l = 0;
    while(*p != '\r'){
        p++;
        l++;
    }
    *len = l;
    return line+1;
}

void _add_to_parent(struct node* parent, struct node* node){
    //set parent
    node->parent = parent;
    //add to parent
    int cnt = parent->cnt;
    parent->children[cnt] = node;
    parent->cnt++;
    assert(parent->cnt<=parent->nchildren);
}

void add_to_parent(struct node* node){
    struct node* parent = stack[top];
    assert(parent->type == ARRAY);
    _add_to_parent(parent, node);
    if(parent->cnt == parent->nchildren){
        //pop the parent ARRAY node
        top--;
    }
}

struct node* newarraynode(int nchildren){
    struct node* node = malloc(sizeof(struct node));
    memset(node, 0, sizeof(*node));
    node->type = ARRAY;
    node->nchildren = nchildren;
    node->children = malloc(sizeof(struct node*)*nchildren);
    return node;
}

struct node* newintegernode(int val){
    struct node* node = malloc(sizeof(struct node));
    memset(node, 0, sizeof(*node));
    node->type = INTEGER;
    node->val = (void*)(long)val;
    return node;
}

struct node* newstatusnode(char* s, int len){
    struct node* node = malloc(sizeof(struct node));
    memset(node, 0, sizeof(*node));
    node->type = STATUS;
    node->val = malloc(len+1);
    memcpy(node->val, s, len);
    *((char*)node->val+len) = 0;
    return node;
}

struct node* newerrornode(char* s, int len){
    struct node* node = malloc(sizeof(struct node));
    memset(node, 0, sizeof(*node));
    node->type = ERROR;
    node->val = malloc(len+1);
    memcpy(node->val, s, len);
    *((char*)node->val+len) = 0;
    return node;
}

struct node * parse(char * s){
    struct node* root = NULL;

    do{
    char* nextline = NULL; 
    char* line = getoneline(s, &nextline);

    if (line == NULL){
        printf("\nline is NULL!");
        assert(0);
        return NULL;
    }

    struct node* node = NULL;
    switch(*line){
    case '*':{
        int nchildren = readlong(line); 
        if (nchildren == 0){
            //TODO:free root             
            return NULL;
        }
        node = newarraynode(nchildren);
    break;
    }
    case ':':{
        int val = readlong(line);
        node = newintegernode(val);
    break;
    }
    case '-':{
        int len = 0;
        char* stat = readstatus(line, &len);
        assert(len>0 && stat);
        node = newstatusnode(stat, len);
    break;
    }
    case '+':{
        int len = 0;
        char * error = readstatus(line, &len);
        assert(len>0 && error);
        node = newerrornode(error, len);
    break;
    }
    default:
        //TODO: free root
        return NULL;
    break;
    }

    // if there is parent in stack
    if(top > -1){
        add_to_parent(node);
    }else{
        root = node;
    }
    if(*line == '*'){
        //push ARRAY node
        stack[++top] = node;
    }
    s = nextline;
    }while(top>-1 || *s!=0);
    printf("\nout of loop, top=%d, *nextline=%d", top, (int)*s);
    return root;
}


struct node* parse2(char* s, char** nextline){
    //char* nextline = NULL;
    char* line = NULL;
    struct node* node = NULL;

    line = getoneline(s, nextline);
    switch(*line){
    case '\0':
        return NULL;
    break;
    case '*' :{
        int nchildren = readlong(line);
        if(nchildren == 0){
            //TODO:fail handle
            return NULL;
        }
        node = newarraynode(nchildren);
        int i;
        for (i=0; i<nchildren; i++){
            s = *nextline;
            //line = getoneline(s, &nextline);
            node->children[i] = parse2(s, nextline);
            if (node->children[i]==NULL){
                //TODO:fail handle
                return NULL;
            }
        }
    break;
    }
    case ':':{
        int val = readlong(line);
        node = newintegernode(val);
    break;
    }
    case '-':{
        int len = 0;
        char* stat = readstatus(line, &len);
        assert(len>0 && stat);
        node = newstatusnode(stat, len);
    break;
    }
    case '+':{
        int len = 0;
        char * error = readstatus(line, &len);
        assert(len>0 && error);
        node = newerrornode(error, len);
    break;
    }
    default:
        //TODO: free root
        return NULL;
    break;
    }
    return node;
}

void printreply(struct node* node){
    if (node->type==ARRAY){
        printf("\ntype=%d,nchildren=%d", node->type, node->nchildren);
        int i;
        for(i=0; i<node->nchildren; i++)
            printreply(node->children[i]);
    }else if (node->type==INTEGER){
        printf("\ntype=%d,val=%d",node->type, (long)node->val);
    }else if (node->type==ERROR || node->type==STATUS){
        printf("\ntype=%d, val=%s", node->type, (char*)node->val);
    }
}

void printreply2(struct node* node){
    switch(node->type){
    case ARRAY:
        printf("\ntype=%d,nchildren=%d", node->type, node->nchildren);
        int i;
        for(i=0; i<node->nchildren; i++)
            printreply2(node->children[i]);
    break;
    case INTEGER:
        printf("\ntype=%d,val=%d",node->type, (long)node->val);
    break;
    case ERROR:
    case STATUS:
        printf("\ntype=%d, val=%s", node->type, (char*)node->val);
    break;
    }
}

void freereply(struct node* node){
    switch(node->type){
    case INTEGER:
        free(node);
    break;
    case ERROR:
    case STATUS:
        free(node->val);
        free(node);
    break;
    case ARRAY:{
        int i;
        for(i=0; i<node->nchildren; i++)
            freereply(node->children[i]);
        free(node->children);
        free(node);
    break;}
    }
}

int main(){
    printf("\n hello world!");
    struct node* reply = parse(s);
    printreply2(reply);
    freereply(reply);
    
    printf("\n----------------------------");
    char* nextline = NULL;
    struct node* reply2 = parse2(s, &nextline);
    printreply2(reply2);
    freereply(reply2);
    return 0;
}
