package query

// Model provides functionality to access properties of models.
type Model interface {
	// Dim returns the dimension of the model. This value is expected to be
	// always positive.
	Dim() int
	// Nodes returns a slice composed of the model's nodes as Node type. The
	// node at index 0 must correspond to the model's root.
	Nodes() []Node
}

// Node represents a node in the model, be it a leaf or an internal one.
// This type is meant to be used for passing a model's nodes a fixed slice and
// thus the 'pointers' to a node's children should be their index in the slice.
type Node struct {
	// Value of a leaf node.
	Value bool
	// Feat index of the feature that an internal node decides on. This value
	// is expected to always be positive.
	Feat int
	// ZChild is the index on the node slice corresponding to the node's zero
	// child. Use NoChild value to indicate the node has no child.
	ZChild int
	// OChild is the index on the node slice corresponding to the node's one
	// child. Use NoChild value to indicate the node has no child.
	OChild int
}

// NoChild is used in a Node's ZChild or OChild fields to indicate that it does
// not have a child.
const NoChild int = -1

// IsLeaf returns true if the node does not have a child and thus is a leaf.
func (n Node) IsLeaf() bool {
	return n.ZChild < 0 || n.OChild < 0
}
