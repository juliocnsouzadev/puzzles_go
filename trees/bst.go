package trees

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) FindClosestValue(target int) int {
	return tree.findClosestValue(target, -1, nil)
}

func (tree *BST) findClosestValue(target int, lowestDiff int, holder *BST) int {
	if tree.Value == target {
		return tree.Value
	}

	if holder == nil {
		lowestDiff = tree.diff(target)
		holder = tree
	}

	if tree.diff(target) < lowestDiff {
		lowestDiff = tree.diff(target)
		holder = tree
	}

	if tree.Left != nil && target < tree.Value {
		return tree.Left.findClosestValue(target, lowestDiff, holder)
	}

	if tree.Right != nil && target > tree.Value {
		return tree.Right.findClosestValue(target, lowestDiff, holder)
	}

	return holder.Value
}

func (tree *BST) diff(target int) int {
	diff := tree.Value - target
	if diff < 0 {
		diff = diff * -1
	}
	return diff
}

func (tree *BST) isLeaf() bool {
	return tree.Left == nil && tree.Right == nil
}
