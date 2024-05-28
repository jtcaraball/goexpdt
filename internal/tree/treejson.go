package tree

import (
	"encoding/json"
	"errors"
	"slices"
)

type nodeJSON struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	// Leaf fields
	Class string `json:"class"`
	// Internal fields
	FeatIdx int `json:"feature_index"`
	LeftID  int `json:"id_left"`
	RightID int `json:"id_right"`
}

type treeJSON struct {
	ClassNames []string                   `json:"class_names"`
	Positive   string                     `json:"positive"`
	Features   []string                   `json:"feature_names"`
	RawNodes   map[string]json.RawMessage `json:"nodes"`
	Nodes      map[int]*nodeJSON          `json:"-"`
}

func newTreeJSON() *treeJSON {
	treeJSON := new(treeJSON)
	treeJSON.Nodes = make(map[int]*nodeJSON)
	return treeJSON
}

func unmarhsalTree(jsonBytes []byte) (*treeJSON, error) {
	treeJSON := newTreeJSON()
	if err := json.Unmarshal(jsonBytes, treeJSON); err != nil {
		return nil, err
	}
	for _, nodeBytes := range treeJSON.RawNodes {
		nodeJSON := &nodeJSON{
			ID:      -1,
			FeatIdx: -1,
			LeftID:  -1,
			RightID: -1,
		}
		if err := json.Unmarshal(nodeBytes, nodeJSON); err != nil {
			return nil, err
		}
		treeJSON.Nodes[nodeJSON.ID] = nodeJSON
	}
	if err := treeJSON.Validate(); err != nil {
		return nil, err
	}
	return treeJSON, nil
}

func (tj treeJSON) Validate() error {
	// Validate fields
	if len(tj.ClassNames) != 2 {
		return errors.New(
			"Tree encoding error: must have exactly two class_names",
		)
	}
	if !slices.Contains(tj.ClassNames, tj.Positive) {
		return errors.New(
			"Tree encoding error: positive must be contained in class_names",
		)
	}
	if len(tj.Features) == 0 {
		return errors.New(
			"Tree encoding error: must have at least one feature_name",
		)
	}
	// Validate nodes
	for _, node := range tj.Nodes {
		if err := node.Validate(
			len(tj.Features),
			len(tj.Nodes),
			tj.ClassNames,
		); err != nil {
			return err
		}
	}
	return nil
}

func (nj nodeJSON) Validate(
	featCount, nodeCount int,
	validClasses []string,
) error {
	if nj.Type != "internal" && nj.Type != "leaf" {
		return errors.New("Tree encoding error: invalid or missing node's type")
	}
	if nj.ID < 0 || nj.ID > nodeCount {
		return errors.New(
			"Tree encoding error: invalid or missing node's id",
		)
	}
	if nj.Type == "internal" {
		if nj.FeatIdx < 0 || nj.FeatIdx >= featCount {
			return errors.New(
				"Tree encoding error: invalid or missing node's feature_index",
			)
		}
		if nj.LeftID < 0 {
			return errors.New(
				"Tree encoding error: invalid node's id_left",
			)
		}
		if nj.RightID < 0 {
			return errors.New(
				"Tree encoding error: invalid node's id_right",
			)
		}
	} else {
		if !slices.Contains(validClasses, nj.Class) {
			return errors.New(
				"Tree encoding error: invalid node's class",
			)
		}
	}
	return nil
}
