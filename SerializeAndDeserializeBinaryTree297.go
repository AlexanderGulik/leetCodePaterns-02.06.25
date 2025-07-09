/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    var result []string
    var preorder func(node *TreeNode)
    preorder = func(node *TreeNode) {
        if node == nil {
            result = append(result, "N")
            return
        }
        result = append(result, strconv.Itoa(node.Val))
        preorder(node.Left)
          preorder(node.Right)
    }
    preorder(root)
    return strings.Join(result, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    nodeValues := strings.Split(data, ",")
    index := 0
    var buildTree func() *TreeNode
    buildTree = func() *TreeNode {
        if nodeValues[index] == "N" {
            index++
            return nil
        }
        val, _ := strconv.Atoi(nodeValues[index])
        index++
        node := &TreeNode{Val: val}
        node.Left = buildTree()
        node.Right = buildTree()
        return node
    }
    return buildTree()
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
