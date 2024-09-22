package earthbucks

import (
	"errors"
	"math"
)

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Hash  *FixedBuf
}

// CountNonNullLeaves counts the non-null leaves in the Merkle tree.
func (n *MerkleNode) CountNonNullLeaves() int {
	if n.Left != nil || n.Right != nil {
		return (countNonNull(n.Left) + countNonNull(n.Right))
	}
	if n.Hash != nil {
		return 1
	}
	return 0
}

// CountAllLeaves counts all leaves in the Merkle tree.
func (n *MerkleNode) CountAllLeaves() int {
	if n.Left != nil || n.Right != nil {
		return (countAll(n.Left) + countAll(n.Right))
	}
	return 1
}

// Helper function to count non-null leaves.
func countNonNull(n *MerkleNode) int {
	if n == nil {
		return 0
	}
	return n.CountNonNullLeaves()
}

// Helper function to count all leaves.
func countAll(n *MerkleNode) int {
	if n == nil {
		return 0
	}
	return n.CountAllLeaves()
}

// Concatenate two FixedBufs into a single byte slice.
func concat(left, right *FixedBuf) []byte {
	var leftBuf, rightBuf []byte
	if left != nil {
		leftBuf = left.buf[:]
	}
	if right != nil {
		rightBuf = right.buf[:]
	}
	return append(leftBuf, rightBuf...)
}

// Compute the hash for the Merkle node.
func (n *MerkleNode) ComputeHash() *FixedBuf {
	if n.Left != nil || n.Right != nil {
		var leftData, rightData *FixedBuf
		if n.Left != nil {
			leftData = n.Left.ComputeHash()
		}
		if n.Right != nil {
			rightData = n.Right.ComputeHash()
		}
		doubleBlake3Hash , _ := DoubleBlake3Hash(concat(leftData, rightData))
		return doubleBlake3Hash
	}
	return n.Hash
}

// Calculate left height of the node.
func (n *MerkleNode) LeftHeight() int {
	if n.Left != nil {
		return n.Left.LeftHeight() + 1
	} else if n.Hash != nil {
		return 1
	}
	return 0
}

// Calculate right height of the node.
func (n *MerkleNode) RightHeight() int {
	if n.Right != nil {
		return n.Right.RightHeight() + 1
	} else if n.Hash != nil {
		return 1
	}
	return 0
}

// Check if the tree is null balanced.
func (n *MerkleNode) IsNullBalanced() bool {
	if n.Left == nil && n.Right == nil {
		return true
	}
	if n.Left == nil || n.Right == nil {
		return false
	}
	return n.Left.IsNullBalanced() &&
		n.Right.IsNullBalanced() &&
		math.Abs(float64(n.Left.LeftHeight()-n.Right.LeftHeight())) <= 1 &&
		math.Abs(float64(n.Left.RightHeight()-n.Right.RightHeight())) <= 1
}

// Create a MerkleNode from leaf hashes.
func FromLeafHashes(hashes []*FixedBuf) *MerkleNode {
	if len(hashes) == 0 {
		return &MerkleNode{}
	}
	if len(hashes) == 1 {
		return &MerkleNode{Hash: hashes[0]}
	}
	if len(hashes) == 2 {
		left := &MerkleNode{Hash: hashes[0]}
		right := &MerkleNode{Hash: hashes[1]}
		doubleBlake3Hash, _ := DoubleBlake3Hash(concat(left.ComputeHash(), right.ComputeHash()))
		return &MerkleNode{
			Left:  left,
			Right: right,
			Hash:  doubleBlake3Hash,
		}
	}

	// Ensure balance by filling with nils
	for (len(hashes) & (len(hashes) - 1)) != 0 {
		hashes = append(hashes, nil)
	}

	left := FromLeafHashes(hashes[:len(hashes)/2])
	right := FromLeafHashes(hashes[len(hashes)/2:])

	doubleBlake3Hash, _ := DoubleBlake3Hash(concat(left.ComputeHash(), right.ComputeHash()))

	return &MerkleNode{
		Left:  left,
		Right: right,
		Hash:  doubleBlake3Hash,
	}
}

// Double the tree with nulls.
func (n *MerkleNode) DoubleWithNulls() (*MerkleNode, error) {
	count := n.CountAllLeaves()
	if math.Log2(float64(count)) != math.Floor(math.Log2(float64(count))) {
		return nil, errors.New("Cannot double a tree that is not a power of 2")
	}
	nullHashes := make([]*FixedBuf, count)
	nullTree := FromLeafHashes(nullHashes)
	return &MerkleNode{Left: n, Right: nullTree}, nil
}

// Update a balanced leaf hash at a specific position.
func (n *MerkleNode) UpdateBalancedLeafHash(pos int, hash *FixedBuf) (*MerkleNode, error) {
	if pos < 0 {
		return nil, errors.New("Position must be greater than or equal to 0")
	}
	countAll := n.CountAllLeaves()
	if pos >= countAll {
		return nil, errors.New("Position must be less than the number of leaves")
	}
	if math.Log2(float64(countAll)) != math.Floor(math.Log2(float64(countAll))) {
		return nil, errors.New("Cannot update a tree that is not a power of 2")
	}
	if countAll == 1 {
		return &MerkleNode{Hash: hash}, nil
	}
	if countAll == 2 {
		doubleBlake3Hash , _ :=  DoubleBlake3Hash(concat(hash, n.Right.Hash))
		if pos == 0 {
			return &MerkleNode{
				Left:  &MerkleNode{Hash: hash},
				Right: n.Right,
				Hash:  doubleBlake3Hash,
			}, nil
		}
		doubleBlake3Hash, _ = DoubleBlake3Hash(concat(n.Left.Hash, hash))
		return &MerkleNode{
			Left:  n.Left,
			Right: &MerkleNode{Hash: hash},
			Hash:  doubleBlake3Hash,
		}, nil
	}

	countLeft := countAll / 2
	if countLeft == 0 {
		return nil, errors.New("Left node must not be null")
	}
	if pos < countLeft {
		left, err := n.Left.UpdateBalancedLeafHash(pos, hash)
		if err != nil {
			return nil, err
		}
		doubelBlake3Hash, _ := DoubleBlake3Hash(concat(left.Hash, n.Right.Hash))
		return &MerkleNode{
			Left:  left,
			Right: n.Right,
			Hash:  doubelBlake3Hash,
		}, nil
	}
	right, err := n.Right.UpdateBalancedLeafHash(pos-countLeft, hash)
	if err != nil {
		return nil, err
	}

	doubleBlake3Hash , _ := DoubleBlake3Hash(concat(n.Left.Hash, right.Hash))
	return &MerkleNode{
		Left:  n.Left,
		Right: right,
		Hash:  doubleBlake3Hash,
	}, nil
}

// Add a leaf hash to the Merkle node.
func (n *MerkleNode) AddLeafHash(hash *FixedBuf) (*MerkleNode, error) {
	countNonNull := n.CountNonNullLeaves()
	countAll := n.CountAllLeaves()
	if countNonNull == countAll {
		nullTree, err := n.DoubleWithNulls()
		if err != nil {
			return nil, err
		}
		return nullTree.UpdateBalancedLeafHash(countNonNull, hash)
	}
	return n.UpdateBalancedLeafHash(countNonNull, hash)
}

// Update balanced leaf hashes starting from a specific position.
func (n *MerkleNode) UpdateBalancedLeafHashes(startPos int, hashes []*FixedBuf) (*MerkleNode, error) {
	tree := n
	for i, hash := range hashes {
		var err error
		tree, err = tree.UpdateBalancedLeafHash(startPos+i, hash)
		if err != nil {
			return nil, err
		}
	}
	return tree, nil
}

// Add multiple leaf hashes to the Merkle node.
func (n *MerkleNode) AddLeafHashes(hashes []*FixedBuf) (*MerkleNode, error) {
	tree := n
	for _, hash := range hashes {
		var err error
		tree, err = tree.AddLeafHash(hash)
		if err != nil {
			return nil, err
		}
	}
	return tree, nil
}