package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID, Parent int
}

func (r *Record) String() string {
	return fmt.Sprint("ID: %d, Parent: %d", r.ID, r.Parent)
}

type byID []Record

func (ids byID) Len() int           { return len(ids) }
func (ids byID) Swap(i, j int)      { ids[i], ids[j] = ids[j], ids[i] }
func (ids byID) Less(i, j int) bool { return ids[i].ID < ids[j].ID }

// Node is a node in tree structure
type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	sort.Sort(byID(records))
	nodes := make([]*Node, len(records))
	for r, rec := range records {
		nodes[r] = &Node{ID: rec.ID}
		if r == 0 && (rec.ID != 0 || rec.Parent != 0) {
			return nil, fmt.Errorf("Not a valid root record %s", rec.String())
		} else if r == 0 {
			continue
		} else if r != rec.ID || rec.ID <= rec.Parent {
			return nil, fmt.Errorf("Not a valid record: %s", rec.String())
		}

		if parent := nodes[rec.Parent]; parent != nil {
			parent.Children = append(parent.Children, nodes[r])
		}
	}
	return nodes[0], nil
}
