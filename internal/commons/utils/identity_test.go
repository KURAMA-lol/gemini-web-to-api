package utils

import (
    "strings"
    "testing"
)

func TestAppendDefaultGeminiIdentity(t *testing.T) {
    prompt := AppendDefaultGeminiIdentity("System: You are a helpful coding assistant.\n\nUser: Hello")
    if !strings.HasPrefix(prompt, "System: You are a helpful coding assistant.") {
        t.Fatalf("caller system context was not preserved: %q", prompt)
    }
    if !strings.HasSuffix(prompt, DefaultGeminiIdentity) {
        t.Fatalf("default identity must be the final instruction")
    }
    if got := AppendDefaultGeminiIdentity(prompt); got != prompt {
        t.Fatalf("identity was appended more than once")
    }
    if got := AppendDefaultGeminiIdentity(" "); got != "" {
        t.Fatalf("empty prompts must remain empty, got %q", got)
    }
}
