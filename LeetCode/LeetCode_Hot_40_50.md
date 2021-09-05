[104. 二叉树的最大深度 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)

```java
class Solution {
    // 层次遍历，找到二叉树得最大深度。。
    public int maxDepth(TreeNode root) {
      if(root == null) return 0;

        Queue<TreeNode> queue = new LinkedList<>();
        queue.offer(root);
        int height = 0;
        while (queue.size() != 0){
            height ++;
            Queue<TreeNode> tmp = new LinkedList<>();
            while (queue.size() != 0){
                TreeNode node = queue.poll();
                if(node.left != null) tmp.offer(node.left);
                if(node.right != null) tmp.offer(node.right);
            }
           
            queue = tmp;
        }
        return height;
    }
}

```



```java
class Solution {
    // 递归返回左右子树的最大高度
    public int maxDepth(TreeNode root) {
        if(root == null) return 0;
        int left = maxDepth(root.left);
        int right = maxDepth(root.right);
        return Math.max(left, right)+1;
    }
}
```



[105. 从前序与中序遍历序列构造二叉树 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)

```jaba
class Solution {
    public TreeNode buildTree(int[] preorder, int[] inorder) {
        return build(preorder, 0, preorder.length-1, inorder, 0, inorder.length-1);
    }
     // 定义; lp， 左子序列的最左边的值的位置， rp 位最有边的位置， 都包括当前值
    public TreeNode build(int[] preorder, int lp, int rp, int[] inorder, int li, int ri){
        if(rp < lp || li> ri) return null;

        TreeNode node = new TreeNode(preorder[lp]); // 跟节点
        // 找到根节点在中序序列中的位置
        int idx = li;
        while (idx<=ri && inorder[idx] != preorder[lp] ) idx++;
        
        int leftLen = idx-li;

        node.left = build(preorder, lp+1, lp+leftLen+1, inorder, li, idx-1);
        node.right = build(preorder, lp+leftLen+1, rp, inorder, idx+1, ri);
        return node;
    }
}
```





