package main

import (
    "testing"
)

// Table Test 1
func TestCommentType(t *testing.T) {
    if commentNum != type int {
        t.Error("Expected that to be a integer!")
    }
}

// Table Test 2
func TestScoreType(t *testing.T) {
    if scoreNum != type int {
        t.Error("Expected that to be a integer!")
    }
}

// Benchmark Test
func Benchmark
