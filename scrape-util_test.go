package main

import (
    "testing"
)

func TestCommentType(t *testing.T) {
    if commentNum != 4 {
        t.Error("Expected that to be a integer!")
    }
}

func TestScoreType(t *testing.T) {
    if scoreNum != 4 {
        t.Error("Expected that to be a integer!")
    }
}