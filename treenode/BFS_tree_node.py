class Node(object):
    def __init__(self, item):
        self.elem = item
        self.rchild = None
        self.lchild = None
class Tree(object):
    def __init__(self):
        self.root = None
    def add(self,item):
        node = Node(item)
        if self.root is None:
            self.root = node
            return
        queue = [self.root] # 用來存放遍處理的東西
        
        while queue:
            cur_node = queue.pop(0)
            if cur_node.lchild is None:
                cur_node.lchild = node
                return
            else:
                queue.append(cur_node.lchild)
            if cur_node.rchild is None:
                cur_node.rchild = node
                return
            else:
                queue.append(cur_node.rchild)
    def breadth_travel(self):
        """廣度遍歷"""
        if self.root is None:
            return
        queue = [self.root]
        while queue:
            cur_node = queue.pop(0)
            print(cur_node.elem, end = " ")
            if cur_node.lchild is not None:
                queue.append(cur_node.lchild)
            if cur_node.rchild is not None:
                queue.append(cur_node.rchild)
    """前序"""
    def preorder(self,node):
        if node is None:
            return
        print(node.elem, end = " ")
        self.preorder(node.lchild)
        self.preorder(node.rchild)
    """中序"""
    def inorder(self,node):
        if node is None:
            return
        
        self.inorder(node.lchild)
        print(node.elem, end = " ")
        self.inorder(node.rchild)
    """後序"""
    def postorder(self,node):
        if node is None:
            return
        
        self.postorder(node.lchild)
        self.postorder(node.rchild)
        print(node.elem, end = " ")    
if __name__ =="__main__":
    tree = Tree()
    tree.add(1)
    tree.add(2)
    tree.add(3)
    tree.add(4)
    tree.add(5)
    tree.breadth_travel()
    print("\nPreorder Traversal:")
    tree.preorder(tree.root)

    print("\ninorder Traversal:")
    tree.inorder(tree.root)

    print("\nPostorder Traversal:")
    tree.postorder(tree.root)