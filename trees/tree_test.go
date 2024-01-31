package trees

import (
	"fmt"
	"os"
	"slices"
	"testing"
)

// =========================== //
//           HELPERS           //
// =========================== //

func writeNewTree(t *testing.T, treeJSONBytes []byte) (string, error) {
	t.Helper()
	f, err := os.CreateTemp("", "tree_tmp")
	if err != nil {
		return "", err
	}
	if _, err = f.Write(treeJSONBytes); err != nil {
		return "", err
	}
	defer f.Close()
	t.Cleanup(func() {
		os.Remove(f.Name())
	})
	return f.Name(), nil
}

func reprTree(t *Tree) []string {
	treeRepr := []string{}
	for _, node := range t.Nodes() {
		if node.IsLeaf() {
			treeRepr = append(
				treeRepr,
				fmt.Sprintf("%d:%t", node.ID, node.Value),
			)
			continue
		}
		treeRepr = append(
			treeRepr,
			fmt.Sprintf("%d:%d", node.ID, node.Feat),
		)
	}
	return treeRepr
}

// =========================== //
//            TESTS            //
// =========================== //

func TestLoadTree_ValidTree(t *testing.T) {
	tJSONBytes := []byte(`
		{
			"class_names": ["pos", "neg"],
			"positive": "pos",
			"nodes": {
				"0": {
				  "id": 0,
				  "type": "internal",
				  "feature_name": "ft2",
				  "feature_index": 6,
				  "threshold": 0.5,
				  "id_left": 1,
				  "id_right": 2
				},
				"1": {
				  "id": 1,
				  "type": "internal",
				  "feature_name": "ft5",
				  "feature_index": 5,
				  "threshold": 0.5,
				  "id_left": 3,
				  "id_right": 4
				},
				"2": {
				  "id": 2,
				  "type": "leaf",
				  "class": "pos"
				},
				"3": {
				  "id": 3,
				  "type": "internal",
				  "feature_name": "ft3",
				  "feature_index": 3,
				  "threshold": 0.5,
				  "id_left": 5,
				  "id_right": 6
				},
				"5": {
				  "id": 5,
				  "type": "leaf",
				  "class": "pos"
				},
				"6": {
				  "id": 6,
				  "type": "leaf",
				  "class": "neg"
				},
				"4": {
				  "id": 4,
				  "type": "internal",
				  "feature_name": "ft7",
				  "feature_index": 7,
				  "threshold": 0.5,
				  "id_left": 7,
				  "id_right": 8
				},
				"7": {
				  "id": 7,
				  "type": "leaf",
				  "class": "neg"
				},
				"8": {
				  "id": 8,
				  "type": "internal",
				  "feature_name": "ft4",
				  "feature_index": 4,
				  "threshold": 0.5,
				  "id_left": 9,
				  "id_right": 10
				},
				"9": {
				  "id": 9,
				  "type": "leaf",
				  "class": "neg"
				},
				"10": {
				  "id": 10,
				  "type": "leaf",
				  "class": "neg"
				}
			}
		}
	`)
	expTreeRepr := []string{
		"0:6",
		"1:5",
		"3:3",
		"5:true",
		"6:false",
		"4:7",
		"7:false",
		"8:4",
		"9:false",
		"10:false",
		"2:true",
	}
	path, err := writeNewTree(t, tJSONBytes)
	if err != nil {
		t.Errorf("Failed to write tree file: %s", err.Error())
		return
	}
	tree, err := LoadTree(path)
	if err != nil {
		t.Errorf("Failed to load tree: %s", err.Error())
		return
	}
	treeRepr := reprTree(tree)
	if !slices.Equal[[]string](expTreeRepr, treeRepr) {
		t.Errorf(
			"Trees not equal. Expected %s but got %s",
			expTreeRepr,
			treeRepr,
		)
	}
}
