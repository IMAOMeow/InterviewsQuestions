// written during interview, can't promise the correctness, just for draft reference

// 1. serialization of binary tree
// preorder, inorder, postorder

// preorder, inorder [root, ] [, root, ]
// inorder, postorder [,root, ] [, ,root]
// preorder, postorder [root, ,] [, , root]

// level order [root, ] <-> flag: layer, empty spot

// 1. root -> string
// 2. de-serial
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrderSeq(root *TreeNode) []int{
	// input validation
	if root == nil {
		return nil
	}

	res := []int{}
	s := []*TreeNode{root}
	l := 1

	for {
		flag := false

		for i:=0; i<l; i++ {
			// pop up
			node := s[0]
			s = s[1:]

			// invalid val for null node
			if node != nil {
				res = append(res, node.Val)
			}else {
				// fake value for invalid node
				res = append(res, math.MinInt32-1)
			}
			
			
			// insert   - > full binary tree
			if node.Left != nil {
				flag = true
				s = append(s, node.Left)
			}else{
				s = append(s, nil)
			}

			if node.Right != nil {
				flag = true
				s = append(s, node.Right)
			}else{
				s = append(s, nil)
			}
		}

		// update layer count
		l = len(s)

		// end condition
		if flag {
			break
		}
	}

	return res[len(res)-l]
}

func levelOrderDeSeq(input []int) *TreeNode {
	// input validation
	if len(input) == 0 {
		return nil
	}

	l := 1
	root := &TreeNode{Val: input[0]}
	s := []*TreeNode{root}
	count := 1

	
	for count < len(input) {
		// construct tree -> insert into stack -> count += 2
		// full Binay tree, input validation skip

		// add value check
		for i:=0; i<l; i++{
			// pop
			node := s[0]
			s := s[1:]

			leftVal := input[count]
			rightVal := input[count+1]

			if leftVal != math.MinInt32-1 {
				node.Left = &TreeNode{Val: leftVal}
				s = append(s, node.Left)
			}

			if rightVal != math.MinInt32-1 {
				node.Right = &TreeNode{Val: rightVal}
				s = append(s, node.Right)
			}

			count += 2

			l := len(s)
		}
	}

	return root
}



// 2. how to design a buffer pool
// this can refer to SQL innodb engine's related logic, not an application level design, more like a infra level stuff

// following are the drafts written during interview
// a. apply for new buffer
// b. free buffer
// c. list applications

user {
	id string
	application string
}

application {
	id string
	space -?>
}

allocationSerivce 1. allocates 2. free


// pool : buffers length 
// 1. space > buffer len 
// 2. re-allocate detail  -> contact space?
// 3. lifespan mechanism  

// user/ buffer space/location


buffer {
	id string
	lifespan []timestamp
	userId string
	lastUsed timestamp
}

pool {
	buffers []buffer
}

scale up 
1. request conflict/peak management ->
2. service: 
	1. scale up -> space usage 
	2. scale down -> space usage


cleaner service -> notification service
1. over lifesan 
2. lastUsed

