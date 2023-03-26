package _map

import (
	"reflect"
	"testing"
)

func Test_convertBuilding(t *testing.T) {
	type args struct {
		buildings []string
	}
	tests := []struct {
		name string
		args args
		want []*Building
	}{
		{"square", args{buildings: []string{"**\n**"}}, []*Building{
			{
				structure: [][][]int{{{1, 1}, {1, 1}}},
				size:      4,
			},
		}},
		{"L-shape", args{buildings: []string{"***\n*  "}}, []*Building{
			{
				structure: [][][]int{
					{{1, 1, 1}, {1, 0, 0}},
					{{1, 1}, {0, 1}, {0, 1}},
					{{0, 0, 1}, {1, 1, 1}},
					{{1, 0}, {1, 0}, {1, 1}},
					{{1, 1, 1}, {0, 0, 1}},
					{{0, 1}, {0, 1}, {1, 1}},
					{{1, 0, 0}, {1, 1, 1}},
					{{1, 1}, {1, 0}, {1, 0}},
				},
				size: 4,
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertBuilding(tt.args.buildings); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertBuilding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateUniqueArrays(t *testing.T) {

	expected1 := [][][]int{{{1, 1}, {1, 1}}}
	if output1 := generateUniqueArrays([][]int{{1, 1}, {1, 1}}); !reflect.DeepEqual(output1, expected1) {
		t.Errorf("Test case 1 failed: expected %v but got %v", expected1, output1)
	}

	expected2 := [][][]int{{{1}}}
	if output2 := generateUniqueArrays([][]int{{1}}); !reflect.DeepEqual(output2, expected2) {
		t.Errorf("Test case 2 failed: expected %v but got %v", expected2, output2)
	}

	input3 := [][]int{{1, 2}, {3, 4}}
	expected3 := [][][]int{
		{{1, 2}, {3, 4}}, {{3, 1}, {4, 2}}, {{4, 3}, {2, 1}}, {{2, 4}, {1, 3}},
		{{2, 1}, {4, 3}}, {{4, 2}, {3, 1}}, {{3, 4}, {1, 2}}, {{1, 3}, {2, 4}},
	}

	if output3 := generateUniqueArrays(input3); !reflect.DeepEqual(output3, expected3) {
		t.Errorf("Test case 3 failed: expected %v but got %v", expected3, output3)
	}

}
