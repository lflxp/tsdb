/*
  https://blog.csdn.net/recall_yesterday/article/details/51476588
*/
#include<stdio.h>
#include<stdlib.h>
#include<string.h>

const int maxsize = 5;//定义b+树的阶数
const int maxelem = maxsize-1;//非叶节点元素的最大数
const int minsize = maxsize / 2;//非叶节点元素最小个数
const int leafmax = maxsize;
const int leafmin = maxsize / 2 + 1;
typedef enum elemtype
{
    Leaf, Brch
}elemtype;//定义节点类型

typedef struct bplus_treenode
{
    char data[maxsize + 1];
    struct bplus_treenode *parent;
    union
    {
        struct bplus_treenode *sub[maxsize + 1];
        struct
        {
            char  *p[maxsize + 1];
            struct bplus_treenode *next;
            struct bplus_treenode *front;
        };
    };
    elemtype nodetype;
    int num;
}node;//定义树的节点

typedef struct bptree
{
    node* root;
    node* first;
}Bptree;

typedef struct Result
{
    node * pnode;
    int pos;
    bool tag;
}Result;//用来存放查找结果

node * buyleaf()
{
    node * ret = (node*)malloc(sizeof(node));
    ret->num = 0;
    ret->nodetype = Leaf;
    ret->parent = NULL;
    ret->next = ret->front = NULL;
    memset(ret->sub, NULL, sizeof(node*)*(maxsize + 1));
    return ret;
}

node * buybrch()
{
    node * ret = (node*)malloc(sizeof(node));
    ret->nodetype = Brch;
    ret->parent = NULL;
    memset(ret->sub, 0, sizeof(node*)*(maxsize + 1));
    ret->next = ret->front = NULL;
    ret->num = 0;
    return ret;
}

Result find(node * root, char key)
{
    node * ptr = root;
    Result res = { NULL,0,false };
    if (root == NULL)
        return res;

    while (ptr != NULL && ptr->nodetype != Leaf)
    {
        int i = ptr->num;
        for (; i > 0 && key < ptr->data[i]; i--) {}
        res.pnode = ptr;
        res.pos = i;
        ptr = ptr->sub[i];
    }
    if (ptr != NULL)
    {
        int i = ptr->num - 1;
        for (; i >= 0 && key < ptr->data[i]; i--) {}
        res.pnode = ptr;
        res.pos = i;
        if (ptr->data[i] == key)
        {
            res.tag = true;
        }
    }
    return res;
}

node * makeroot(char key, char *p)
{
    node * root = (node*)malloc(sizeof(node));
    root->data[0] = key;
    memset(root->sub, NULL, sizeof(node*)*maxsize + 1);
    root->p[0] = p;
    root->num = 1;
    root->parent = NULL;
    root->nodetype = Leaf;
    root->next = root->front = NULL;
    return root;
}

void insert_leaf(node *ptr, int pos, char key, char *record)
{
    int i = ptr->num;
    for (; i - 1 > pos; --i)
    {
        ptr->data[i] = ptr->data[i - 1];
        ptr->p[i] = ptr->p[i - 1];
    }
    ptr->data[i] = key;
    ptr->p[i] = record;
    ptr->num++;
}


//从sourc中pos位置之后移动数据到des中
void move_leaf(node * sourc, node *des, int pos)
{
    int i = pos + 1;
    for (int j = 0; i < sourc->num; j++, i++)
    {
        des->data[j] = sourc->data[i];
        des->p[j] = sourc->p[i];
        sourc->p[i] = NULL;
    }
    des->num = sourc->num = leafmin;
}
void move_brch(node * sourc, node * des, int pos)
{
    int i = pos + 1;
    for (int j = 0; i <= sourc->num; j++, i++)
    {
        des->data[j] = sourc->data[i];
        des->sub[j] = sourc->sub[i];
        des->sub[j]->parent = des;
        sourc->sub[i] = NULL;
    }
    des->num = sourc->num = minsize;
}
void insert_brch(node *ptr, int pos, char kx, node *right)
{
    int i = ++ptr->num;
    for (; i - 1 > pos; --i)
    {
        ptr->data[i] = ptr->data[i - 1];
        ptr->sub[i] = ptr->sub[i - 1];
    }
    ptr->data[i] = kx;
    ptr->sub[i] = right;
    if (right != NULL)
    {
        right->parent = ptr;
    }
}

node * adjust(node * ptr)
{//插入时的调整
    node * par = ptr->parent;
    node * newnode = NULL;

    if (ptr->nodetype == Leaf)
    {
        newnode = buyleaf();
        newnode->next = ptr->next;
        ptr->next = newnode;
        newnode->front = ptr;
        if (newnode->next != NULL)
            newnode->next->front = newnode;
    }
    else
        newnode = buybrch();
    if (newnode->nodetype == Leaf)
        move_leaf(ptr, newnode, leafmin - 1);
    else
        move_brch(ptr, newnode, minsize);

    if (par == NULL)
    {
        node * newroot = buybrch();
        newroot->sub[0] = ptr;
        newroot->sub[1] = newnode;
        newroot->data[1] = newnode->data[0];
        newnode->parent = newroot;
        ptr->parent = newroot;
        newroot->num++;
        return newroot;
    }
    else
    {
        int pos = par->num;
        for (; pos > 0 && ptr != par->sub[pos]; --pos) {}

        insert_brch(par, pos, newnode->data[0], newnode);
        if (par->num > maxelem)
        {
            return adjust(par);
        }
    }
    return NULL;
}

void insert(node * &ptr, char key, char *reco)
{
    if (ptr == NULL)
    {
        ptr = makeroot(key, reco);
        return;
    }

    Result res = find(ptr, key);
    if (res.tag == true) return;
    else
    {
        node * pnode = res.pnode;
        unsigned pos = res.pos;
        node *par = pnode->parent;
        insert_leaf(pnode, pos, key, reco);

        if (pnode->num > leafmax)
        {
            node * temp = adjust(pnode);
            if (temp != NULL)
                ptr = temp;
        }
    }
}

node * set_first(node *root)
{
    node * first = NULL;
    while (root != NULL && root->nodetype != Leaf)
    {
        first = root;
        root = root->sub[0];
    }
    if (root->nodetype == Leaf)
        return root;
    return first;
}


void dele_brch(node * ptr, int pos)
{
    int i = pos;
    for (; i < ptr->num; ++i)
    {
        ptr->data[i] = ptr->data[i + 1];
        ptr->sub[i] = ptr->sub[i + 1];
    }
    --ptr->num;
}
void merge_lebrch(node * left, node *&ptr, node *par, int parpos)
{
    int i = ++left->num;
    ptr->data[0] = par->data[parpos];
    for (int j = 0; j < ptr->num; ++j, ++i)
    {
        left->data[i] = ptr->data[j];
        left->sub[i] = ptr->sub[j];
        ptr->sub[j]->parent = left;
    }
    left->num += ptr->num;
    free(ptr);
    dele_brch(par, parpos);
    ptr = left;
}
void merge_leaf(node * des, node* &sorc, node *par, int parpos)
{
    int i = des->num;
    for (int j = 0; j < sorc->num; ++j, ++i)
    {
        des->data[i] = sorc->data[j];
        des->p[i] = sorc->p[j];
    }
    des->num += sorc->num;
    free(sorc);
    dele_brch(par, parpos);
    sorc = des;
}
void dele_leaf(node *ptr, int pos, int parpos)
{

    --ptr->num;
    for (int i = pos; i < ptr->num; ++i)
    {
        ptr->data[i] = ptr->data[i+1];
        ptr->p[i] = ptr->p[i+1];
    }
    if (pos == 0)
    {
        ptr->parent->data[parpos] = ptr->data[0];
    }
}
node *adjust_brch(node *ptr)
{
    if (ptr == NULL) return ptr;
    node *par = ptr->parent;

    if (par != NULL)
    {
        int parpos = par->num;
        for (; parpos >= 0 && par->sub[parpos] != ptr; --parpos) {}

        node *left = par->sub[parpos - 1];
        node *right = par->sub[parpos + 1];

        if (left != NULL && left->num > minsize)
        {
            int i = ++ptr->num;
            ptr->data[0] = par->data[parpos];

            for (; i > 0; i--)
            {
                ptr->data[i] = ptr->data[i - 1];
                ptr->sub[i] = ptr->sub[i - 1];
            }
            ptr->sub[0] = left->sub[left->num];
            par->data[parpos] = left->data[left->num];
            left->num--;
        }
        else if (right != NULL && right->num > minsize)
        {
            node * ptnode = ptr;

            ptnode->data[++ptnode->num] = par->data[parpos + 1];
            ptnode->sub[ptnode->num] = right->sub[0];
            right->sub[0]->parent = ptr;

            for (int i = 0; i < right->num; i--)
            {
                right->data[i] = right->data[i + 1];
                right->sub[i] = right->sub[i + 1];
            }
            par->data[parpos + 1] = right->data[0];
            --right->num;
        }
        else if (left != NULL)
        {
            merge_lebrch(left, ptr, par, parpos);
        }
        else if (right != NULL)
        {
            merge_lebrch(ptr, right, par, parpos + 1);
        }
        if (par->parent == NULL && par->num == 0)
        {
            free(par);
            return ptr;
        }
        else if (par->parent != NULL && par->num < minsize)
        {
            return adjust_brch(par);
        }
    }
    else
    {
        if (ptr->num == 0)
        {
            node * temp = ptr->sub[0];
            free(ptr);
            return temp;
        }
    }
    return NULL;
}
node * adjust_tree(node *ptnode)
{//删除节点时用到的调整函数
    node * left = ptnode->front;
    node * right = ptnode->next;
    node * pare = ptnode->parent;

    int parpos = ptnode->parent->num;
    for (; parpos >= 0 && ptnode->parent->sub[parpos] != ptnode; --parpos) {}

    if (left != NULL &&left->num > leafmin)
    {
        int i = ++ptnode->num;
        for (; i > 0; i--)
        {
            ptnode->data[i] = ptnode->data[i - 1];
            ptnode->p[i] = ptnode->p[i - 1];
        }
        left->num--;
        ptnode->data[0] = left->data[left->num];
        ptnode->parent->data[parpos] = ptnode->data[0];
        ptnode->p[0] = left->p[left->num];

    }
    else if (right != NULL && right->num > leafmin)
    {
        ptnode->data[ptnode->num] = right->data[0];
        ptnode->p[ptnode->num] = right->p[0];
        ++ptnode->num;
        --right->num;
        for (int i = 0; i < right->num; i--)
        {
            right->data[i] = right->data[i + 1];
            right->p[i] = right->p[i + 1];
        }
        ptnode->parent->data[parpos + 1] = right->data[0];
    }
    else if (left != NULL)
    {
        merge_leaf(left, ptnode, ptnode->parent, parpos);
    }
    else if (right != NULL)
    {
        merge_leaf(ptnode, right, ptnode->parent, parpos + 1);
    }
    if (pare->num < minsize)
    {
        return adjust_brch(pare);
    }
    else
    {
        return NULL;
    }
}
void dele_tree(node* &root, char key)
{
    if (root == NULL)return;
    Result res = find(root, key);
    node * ptnode = res.pnode;
    int pos = res.pos;

    if (res.tag == false)
        return;
    int parpos = ptnode->parent->num;
    for (; parpos >= 0 && ptnode->parent->sub[parpos] != ptnode; --parpos) {}
    dele_leaf(ptnode, pos, parpos);

    if (ptnode->num < leafmin)
    {
        if (ptnode->parent == NULL && ptnode->num == 0)//整个树只有一个节点时，又恰好被删掉唯一的数据
        {
            free(root);
            root = NULL;
        }
        else if (ptnode->parent != NULL && ptnode->parent->num < minsize)//不平衡，调整分支
        {
            node *temp = adjust_tree(ptnode);
            if (temp != NULL)
                root = temp;
        }
    }
    return;
}
void show(node * fi)
{
    while (fi != NULL)
    {
        char ch;
        for (int i = 0; i < fi->num; ++i)
        {
            ch = fi->data[i];
            printf("%c ", ch);
        }
        fi = fi->next;
    }
    printf("\n");
}
int main(int argc, char **argv)
{
    Bptree tree = { NULL,NULL };
    char ar[] = { "qwe23rt9yui8opa5sdf1ghj2k0lz4xcv6bn8m" };
    int n = sizeof(ar) / sizeof(ar[0]) - 1;
    char *rec = (char*)0x00008888;
    for (int i = 0; i < n; ++i)
    {
        insert(tree.root, ar[i], rec);
        tree.first = set_first(tree.root);
        //show(tree.first);
    }
    show(tree.first);
    for (int i = 0; i < n; i++)
    {
        char cha = getchar();
        dele_tree(tree.root,cha);
        tree.first = set_first(tree.root);
        show(tree.first);
    }
    show(tree.first);
    return 0;
}