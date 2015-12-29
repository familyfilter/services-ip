package db

import (
    "testing"
)

func TestInsertRange(t *testing.T) {
    _, err := InsertRange(nil, nil)
    if err != nil {
        t.Error("InsertRange threw error, ", err)
    }
}

func TestRetrieveRange(t *testing.T) {
    _, _, err := RetrieveRange()
    if err != nil {
        t.Error("RetrieveRange threw error, ", err)
    }
}

func TestDeleteRange(t *testing.T) {
    err := DeleteRange()
    if err != nil {
        t.Error("RetrieveRange threw error, ", err)
    }
}
