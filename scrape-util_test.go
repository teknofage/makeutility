package main

import (
    "testing"
)

func TestCommentType(t *testing.T) {
    if commentNum != type int {
        t.Error("Expected that to be a integer!")
    }
}

func TestScoreType(t *testing.T) {
    if scoreNum != type int {
        t.Error("Expected that to be a integer!")
    }
}