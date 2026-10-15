package main

import "testing"

func TestRecallCompletionDispatch(t *testing.T) {
	checkCompletionDispatch(t, []completionDispatchCase{
		{[]string{"deja", "recall", "--"}, []string{"--project", "--harness", "--limit"}, []string{"--re", "--role", "--json"}},
		{[]string{"deja", "recall", "--harness", ""}, []string{"claude", "codex"}, []string{"--re", "--limit"}},
		{[]string{"deja", "search", "--"}, []string{"--re", "--role"}, nil},
	})
}
