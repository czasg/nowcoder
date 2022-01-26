package nowcoder

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

/* 前序遍历 - VLR - 根左右
二叉树的深度优先算法可细分为：VLR、LDR、LRD
*/
func ReadVLRTree(tree *TreeNode) []int {
    val := []int{}
    if tree == nil {
        return val
    }
    val = append(val, tree.Val)
    val = append(val, ReadVLRTree(tree.Left)...)
    val = append(val, ReadVLRTree(tree.Right)...)
    return val
}

/* 中序遍历 - LDR - 左根右
 */
func ReadLDRTree(tree *TreeNode) []int {
    val := []int{}
    if tree == nil {
        return val
    }
    val = append(val, ReadLDRTree(tree.Left)...)
    val = append(val, tree.Val)
    val = append(val, ReadLDRTree(tree.Right)...)
    return val
}

/* 后序遍历 - LRD - 左右根
 */
func ReadLRDTree(tree *TreeNode) []int {
    val := []int{}
    if tree == nil {
        return val
    }
    val = append(val, ReadLRDTree(tree.Left)...)
    val = append(val, ReadLRDTree(tree.Right)...)
    val = append(val, tree.Val)
    return val
}

/*  深度优先求和
1,2,3 => 12 + 13 = 25
*/
func SumDFSTree(tree *TreeNode) int {
    var dfs func(tree *TreeNode, rootVal int) int
    dfs = func(tree *TreeNode, rootVal int) int {
        if tree == nil {
            return 0
        }
        if tree.Left == nil && tree.Right == nil {
            return rootVal + tree.Val
        }
        rootVal = (rootVal + tree.Val) * 10
        return dfs(tree.Left, rootVal) + dfs(tree.Right, rootVal)
    }
    return dfs(tree, 0)
}

func MaxSumDFSTree(tree *TreeNode) int {
    return 0
}
