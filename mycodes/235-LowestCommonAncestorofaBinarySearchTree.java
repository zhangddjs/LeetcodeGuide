/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */

class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        List<TreeNode> list1 = new ArrayList<>(), list2 = new ArrayList<>();
        dfs(root, p, list1);
        dfs(root, q, list2);
        int i = list1.size() - 1, j = list2.size() - 1;
        for (i = i, j = j; i >= 1 && j >= 1 && list1.get(i - 1) == list2.get(j - 1); i--, j--);
        return list1.get(i);
    }

    public void dfs (TreeNode root, TreeNode desc, List<TreeNode> list) {
        if (root == null) return;
        if (list.isEmpty() && root != desc) dfs(root.left, desc, list);
        if (!list.isEmpty() || root == desc) list.add(root);
        if (list.isEmpty() && root != desc) dfs(root.right, desc, list);
        if (!list.isEmpty() || root == desc) list.add(root);
    }
}

class Solution2 {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        List<TreeNode> list1 = new ArrayList<>(), list2 = new ArrayList<>();
        dfs(root, p, list1);
        dfs(root, q, list2);
        int i = list1.size() - 1, j = list2.size() - 1;
        for (i = i, j = j; i >= 1 && j >= 1 && list1.get(i - 1) == list2.get(j - 1); i--, j--);
        return list1.get(i);
    }

    public void dfs (TreeNode root, TreeNode desc, List<TreeNode> list) {
        if (root == null) return;
        if (list.isEmpty() && root != desc) dfs(root.val > desc.val ? root.left : root.right, desc, list);
        if (!list.isEmpty() || root == desc) list.add(root);
    }
}

class Solution3 {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        while ((root.val > p.val && root.val > q.val) || (root.val < p.val && root.val < q.val))
            root = root.val > p.val ? root.left : root.right;
        return root;
    }
}

class Solution4 {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if (root.val > p.val && root.val > q.val) return lowestCommonAncestor(root.left, p, q);
        else if (root.val < p.val && root.val < q.val) return lowestCommonAncestor(root.right, p, q);
        else return root;
    }
}