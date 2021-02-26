package src

import "testing"

func TestNewDAG(t *testing.T) {

	test_map := make(map[string][]string, 0)

	id_one := "1"

	example_adjacency_list := make([]string, 0, 5)

	example_adjacency_list = append(example_adjacency_list, "arara")

	test_map[id_one] = example_adjacency_list

	t.Error(test_map)

}
