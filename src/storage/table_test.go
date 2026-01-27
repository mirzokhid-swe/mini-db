package storage

import (
	"testing"
)

func TestCreateTable(t *testing.T) {
	tm, err := NewTableManager(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}

	schema := &Schema{
		Columns: []Column{
			{Name: "id", Type: TypeInt},
			{Name: "name", Type: TypeVarchar, Length: 50},
		},
	}

	err = tm.CreateTable("users", schema)
	if err != nil {
		t.Fatal(err)
	}

	err = tm.CreateTable("users", schema)
	if err == nil {
		t.Fatal("expected error for duplicate table")
	}
}

func TestInsert(t *testing.T) {
	tm, err := NewTableManager(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}

	schema := &Schema{
		Columns: []Column{
			{Name: "id", Type: TypeInt},
			{Name: "name", Type: TypeVarchar, Length: 50},
		},
	}
	tm.CreateTable("users", schema)

	record := Record{
		Items: []Item{
			{Literal: int64(1)},
			{Literal: "Hotz"},
		},
	}

	err = tm.Insert("users", record)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAllData(t *testing.T) {
	tm, err := NewTableManager(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}

	schema := &Schema{
		Columns: []Column{
			{Name: "id", Type: TypeInt},
			{Name: "name", Type: TypeVarchar, Length: 50},
		},
	}
	tm.CreateTable("users", schema)

	tm.Insert("users", Record{Items: []Item{{Literal: int64(1)}, {Literal: "John"}}})
	tm.Insert("users", Record{Items: []Item{{Literal: int64(2)}, {Literal: "Bob"}}})

	data, err := tm.GetAllData("users", []Filter{}, SelectedColumns{Columns: []string{"id", "name"}})
	if err != nil {
		t.Fatal(err)
	}

	if len(data) != 2 {
		t.Fatalf("expected 2 rows, got %d", len(data))
	}
}

func TestGetAllDataWithFilter(t *testing.T) {
	tm, err := NewTableManager(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}

	schema := &Schema{
		Columns: []Column{
			{Name: "id", Type: TypeInt},
			{Name: "name", Type: TypeVarchar, Length: 50},
		},
	}
	tm.CreateTable("users", schema)

	tm.Insert("users", Record{Items: []Item{{Literal: int64(1)}, {Literal: "Alice"}}})
	tm.Insert("users", Record{Items: []Item{{Literal: int64(2)}, {Literal: "Bob"}}})

	filters := []Filter{{Column: "name", Operator: "=", Value: "Alice"}}
	data, err := tm.GetAllData("users", filters, SelectedColumns{Columns: []string{"id", "name"}})
	if err != nil {
		t.Fatal(err)
	}

	if len(data) != 1 {
		t.Fatalf("expected 1 row, got %d", len(data))
	}
	if data[0]["name"] != "Alice" {
		t.Fatalf("expected Alice, got %v", data[0]["name"])
	}
}

func TestDelete(t *testing.T) {
	tm, err := NewTableManager(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}

	schema := &Schema{
		Columns: []Column{
			{Name: "id", Type: TypeInt},
			{Name: "name", Type: TypeVarchar, Length: 50},
		},
	}
	tm.CreateTable("users", schema)

	tm.Insert("users", Record{Items: []Item{{Literal: int64(1)}, {Literal: "Alice"}}})
	tm.Insert("users", Record{Items: []Item{{Literal: int64(2)}, {Literal: "Bob"}}})

	deleted, err := tm.Delete("users", []Filter{{Column: "name", Operator: "=", Value: "Alice"}})
	if err != nil {
		t.Fatal(err)
	}
	if deleted != 1 {
		t.Fatalf("expected 1 deleted, got %d", deleted)
	}

	data, _ := tm.GetAllData("users", []Filter{}, SelectedColumns{Columns: []string{"id", "name"}})
	if len(data) != 1 {
		t.Fatalf("expected 1 row remaining, got %d", len(data))
	}
}

func TestUpdate(t *testing.T) {
	tm, err := NewTableManager(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}

	schema := &Schema{
		Columns: []Column{
			{Name: "id", Type: TypeInt},
			{Name: "name", Type: TypeVarchar, Length: 50},
		},
	}
	tm.CreateTable("users", schema)

	tm.Insert("users", Record{Items: []Item{{Literal: int64(1)}, {Literal: "Alice"}}})

	updated, err := tm.Update("users", map[string]any{"name": "Alice_updated"}, []Filter{{Column: "id", Operator: "=", Value: int64(1)}})
	if err != nil {
		t.Fatal(err)
	}
	if updated != 1 {
		t.Fatalf("expected 1 updated, got %d", updated)
	}

	data, _ := tm.GetAllData("users", []Filter{}, SelectedColumns{Columns: []string{"name"}})
	if len(data) != 1 {
		t.Fatalf("expected 1 row, got %d", len(data))
	}
	if data[0]["name"] != "Alice_updated" {
		t.Fatalf("expected Alice_updated, got %v", data[0]["name"])
	}
}
