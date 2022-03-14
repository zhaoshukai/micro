package cmd

import (
	"reflect"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/micro-community/micro/v3/service/registry"
)

type parseCase struct {
	args     []string
	values   *registry.Value
	expected map[string]interface{}
}

func TestDynamicFlagParsing(t *testing.T) {
	cases := []parseCase{
		{
			args: []string{"--ss=a,b"},
			values: &registry.Value{
				Values: []*registry.Value{
					{
						Name: "ss",
						Type: "[]string",
					},
				},
			},
			expected: map[string]interface{}{
				"ss": []interface{}{"a", "b"},
			},
		},
		{
			args: []string{"--ss", "a,b"},
			values: &registry.Value{
				Values: []*registry.Value{
					{
						Name: "ss",
						Type: "[]string",
					},
				},
			},
			expected: map[string]interface{}{
				"ss": []interface{}{"a", "b"},
			},
		},
		{
			args: []string{"--ss=a", "--ss=b"},
			values: &registry.Value{
				Values: []*registry.Value{
					{
						Name: "ss",
						Type: "[]string",
					},
				},
			},
			expected: map[string]interface{}{
				"ss": []interface{}{"a", "b"},
			},
		},
		{
			args: []string{"--ss", "a", "--ss", "b"},
			values: &registry.Value{
				Values: []*registry.Value{
					{
						Name: "ss",
						Type: "[]string",
					},
				},
			},
			expected: map[string]interface{}{
				"ss": []interface{}{"a", "b"},
			},
		},
		{
			args: []string{"--bs=true,false"},
			values: &registry.Value{
				Values: []*registry.Value{
					{
						Name: "bs",
						Type: "[]bool",
					},
				},
			},
			expected: map[string]interface{}{
				"bs": []interface{}{true, false},
			},
		},
		{
			args: []string{"--bs", "true,false"},
			values: &registry.Value{
				Values: []*registry.Value{
					{
						Name: "bs",
						Type: "[]bool",
					},
				},
			},
			expected: map[string]interface{}{
				"bs": []interface{}{true, false},
			},
		},
		{
			args: []string{"--bs=true", "--bs=false"},
			values: &registry.Value{
				Values: []*registry.Value{
					{
						Name: "bs",
						Type: "[]bool",
					},
				},
			},
			expected: map[string]interface{}{
				"bs": []interface{}{true, false},
			},
		},
		{
			args: []string{"--bs", "true", "--bs", "false"},
			values: &registry.Value{
				Values: []*registry.Value{
					{
						Name: "bs",
						Type: "[]bool",
					},
				},
			},
			expected: map[string]interface{}{
				"bs": []interface{}{true, false},
			},
		},
		{
			args: []string{"--is=10,20"},
			values: &registry.Value{
				Values: []*registry.Value{
					{
						Name: "is",
						Type: "[]int32",
					},
				},
			},
			expected: map[string]interface{}{
				"is": []interface{}{int32(10), int32(20)},
			},
		},
		{
			args: []string{"--is", "10,20"},
			values: &registry.Value{
				Values: []*registry.Value{
					{
						Name: "is",
						Type: "[]int32",
					},
				},
			},
			expected: map[string]interface{}{
				"is": []interface{}{int32(10), int32(20)},
			},
		},
		{
			args: []string{"--is=10", "--is=20"},
			values: &registry.Value{
				Values: []*registry.Value{
					{
						Name: "is",
						Type: "[]int32",
					},
				},
			},
			expected: map[string]interface{}{
				"is": []interface{}{int32(10), int32(20)},
			},
		},
		{
			args: []string{"--is", "10", "--is", "20"},
			values: &registry.Value{
				Values: []*registry.Value{
					{
						Name: "is",
						Type: "[]int32",
					},
				},
			},
			expected: map[string]interface{}{
				"is": []interface{}{int32(10), int32(20)},
			},
		},
		{
			args: []string{"--is=10,20"},
			values: &registry.Value{
				Values: []*registry.Value{
					{
						Name: "is",
						Type: "[]int64",
					},
				},
			},
			expected: map[string]interface{}{
				"is": []interface{}{int64(10), int64(20)},
			},
		},
		{
			args: []string{"--is", "10,20"},
			values: &registry.Value{
				Values: []*registry.Value{
					{
						Name: "is",
						Type: "[]int64",
					},
				},
			},
			expected: map[string]interface{}{
				"is": []interface{}{int64(10), int64(20)},
			},
		},
		{
			args: []string{"--is=10", "--is=20"},
			values: &registry.Value{
				Values: []*registry.Value{
					{
						Name: "is",
						Type: "[]int64",
					},
				},
			},
			expected: map[string]interface{}{
				"is": []interface{}{int64(10), int64(20)},
			},
		},
		{
			args: []string{"--is", "10", "--is", "20"},
			values: &registry.Value{
				Values: []*registry.Value{
					{
						Name: "is",
						Type: "[]int64",
					},
				},
			},
			expected: map[string]interface{}{
				"is": []interface{}{int64(10), int64(20)},
			},
		},
		{
			args: []string{"--fs=10.1,20.2"},
			values: &registry.Value{
				Values: []*registry.Value{
					{
						Name: "fs",
						Type: "[]float64",
					},
				},
			},
			expected: map[string]interface{}{
				"fs": []interface{}{float64(10.1), float64(20.2)},
			},
		},
		{
			args: []string{"--fs", "10.1,20.2"},
			values: &registry.Value{
				Values: []*registry.Value{
					{
						Name: "fs",
						Type: "[]float64",
					},
				},
			},
			expected: map[string]interface{}{
				"fs": []interface{}{float64(10.1), float64(20.2)},
			},
		},
		{
			args: []string{"--fs=10.1", "--fs=20.2"},
			values: &registry.Value{
				Values: []*registry.Value{
					{
						Name: "fs",
						Type: "[]float64",
					},
				},
			},
			expected: map[string]interface{}{
				"fs": []interface{}{float64(10.1), float64(20.2)},
			},
		},
		{
			args: []string{"--fs", "10.1", "--fs", "20.2"},
			values: &registry.Value{
				Values: []*registry.Value{
					{
						Name: "fs",
						Type: "[]float64",
					},
				},
			},
			expected: map[string]interface{}{
				"fs": []interface{}{float64(10.1), float64(20.2)},
			},
		},
		{
			args: []string{"--user_email=someemail"},
			values: &registry.Value{
				Values: []*registry.Value{
					{

						Name: "user_email",
						Type: "string",
					},
				},
			},
			expected: map[string]interface{}{

				"user_email": "someemail",
			},
		},
		{
			args: []string{"--user_email=someemail", "--user_name=somename"},
			values: &registry.Value{
				Values: []*registry.Value{
					{
						Name: "user_email",
						Type: "string",
					},
					{
						Name: "user_name",
						Type: "string",
					},
				},
			},
			expected: map[string]interface{}{
				"user_email": "someemail",
				"user_name":  "somename",
			},
		},
		{
			args: []string{"--b"},
			values: &registry.Value{
				Values: []*registry.Value{
					{
						Name: "b",
						Type: "bool",
					},
				},
			},
			expected: map[string]interface{}{
				"b": true,
			},
		},
		{
			args: []string{"--user_friend_email=hi"},
			values: &registry.Value{
				Values: []*registry.Value{
					{

						Name: "user_friend_email",
						Type: "string",
					},
				},
			},
			expected: map[string]interface{}{
				"user_friend_email": "hi",
			},
		},
	}
	for _, c := range cases {
		t.Run(strings.Join(c.args, " "), func(t *testing.T) {
			_, flags, err := splitCmdArgs(c.args)
			if err != nil {
				t.Fatal(err)
			}
			req, err := flagsToRequest(flags, c.values)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(c.expected, req) {
				spew.Dump("Expected:", c.expected, "got: ", req)
				t.Fatalf("Expected %v, got %v", c.expected, req)
			}
		})

	}
}
