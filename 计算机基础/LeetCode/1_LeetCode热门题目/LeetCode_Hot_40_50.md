[96. 不同的二叉搜索树 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/unique-binary-search-trees/)

```java
class Solution {
    public int numTrees(int n) {
        // 动态规划
        int[] dp = new int[n+1];
        if(n == 1) return 1;
        if(n == 0) return 0;
        if(n == 2) return 2;
        dp[0] = 1;
        dp[1] = 1;
        dp[2] = 2;
        for(int i=3; i<=n; i++){
            for(int j=1; j<=i; j++){
                dp[i] += dp[j-1]*dp[i-j]; // 全排列，所以用乘
            }
        }
        return dp[n];
    }
    
}
```

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

```java
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

[114. 二叉树展开为链表 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/)

```java
class Solution {

    public void flatten(TreeNode root) {
        dfs(root);
    }
    TreeNode pre = null;
    public void dfs(TreeNode root){
        // terminate
        if(root == null) return;

        dfs(root.right);
        dfs(root.left);
        root.right = pre;
        root.left = null; //这里一定要有， 否则会形成环
        pre = root;
    }
}
```

[121. 买卖股票的最佳时机 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)

```java
class Solution {
    public int maxProfit(int[] prices) {
        // 贪心算法
        int cnt = 0;
        int min = prices[0];
        for(int i=1; i<prices.length; i++){
            cnt = Math.max(prices[i]-min, cnt);
            min = Math.min(min, prices[i]);
        }
        return cnt;
    }
}
```

[124. 二叉树中的最大路径和 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/)

```java
class Solution {
    public int maxPathSum(TreeNode root) {
        dfs(root);
        return max;
    }

    public int max = Integer.MIN_VALUE;

    public int dfs(TreeNode root){
        // terminate
        if(root == null) return 0;

        // 左子树中包含左子节点的路径的最大值, 如果包含左子节点的路径最大值小于0， 则不使用，取0
        int leftMax = Math.max(dfs(root.left), 0);
        // 右子树中包含右子节点的路径的最大值
        int rightMax = Math.max(dfs(root.right), 0);
        // 更新全局最大值
        max = Math.max(max, root.val+leftMax+rightMax);
        return Math.max(leftMax, rightMax)+root.val;
    }
}
```

[128. 最长连续序列 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/longest-consecutive-sequence/)

```java
class Solution {
    public int longestConsecutive(int[] nums) {
        if(nums.length == 0) return 0;
        HashSet<Integer> set = new HashSet<>();
        // HashSet 去掉重复数字
        for(int n : nums){
            if(set.contains(n)) continue;
            set.add(n);
        }

        int maxCnt = 0;
        for(int n: set){
            if(!set.contains(n-1)){
                int cuN = n;
                int cuCnt = 1;
                // 当前值开始， 连续子序列的长度
                while (set.contains(cuN+1)){
                    cuN ++;
                    cuCnt ++;
                }
                maxCnt = Math.max(cuCnt, maxCnt);
            }
        }
        return maxCnt;
    }
}
```

[136. 只出现一次的数字 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/single-number/)

```java
class Solution {
    public int singleNumber(int[] nums) {
        // 亦或 
        int cnt = 0;
        for(int n: nums){
            cnt ^= n;
        }
        return cnt;
    }
}
```

[139. 单词拆分 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/word-break/)

```java
class Solution {
     public boolean wordBreak(String s, List<String> wordDict) {
        Set<String> set = new HashSet<String>();
        for(String str: wordDict){
            set.add(str);
        }

        // 动态规划
        boolean[] dp = new boolean[s.length()+1];
        dp[0] = true; // 前i个字符是否在字典中
        for(int i=0; i<s.length(); i++){
            for(int j=i+1; j<s.length()+1; j++){
                if(dp[i] && set.contains(s.substring(i, j))){
                    dp[j] = true;
                }
            }
        }
        return dp[s.length()];
    }
}
```

[141. 环形链表 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/linked-list-cycle/)

```java
public class Solution {
    public boolean hasCycle(ListNode head) {
        if(head == null) return false;
        // 快慢指针
        ListNode f = head.next;
        ListNode s = head;
        while(f != null && s != null && f.next != null && s.next !=null){
            if(f == s) return true;
            f = f.next.next;
            s = s.next;
        }
        return false;
    }
}
```



























