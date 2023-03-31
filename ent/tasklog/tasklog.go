// Code generated by ent, DO NOT EDIT.

package tasklog

import (
	"time"
)

const (
	// Label holds the string label denoting the tasklog type in the database.
	Label = "task_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStartedAt holds the string denoting the started_at field in the database.
	FieldStartedAt = "started_at"
	// FieldFinishedAt holds the string denoting the finished_at field in the database.
	FieldFinishedAt = "finished_at"
	// FieldResult holds the string denoting the result field in the database.
	FieldResult = "result"
	// EdgeTasks holds the string denoting the tasks edge name in mutations.
	EdgeTasks = "tasks"
	// Table holds the table name of the tasklog in the database.
	Table = "sys_task_logs"
	// TasksTable is the table that holds the tasks relation/edge.
	TasksTable = "sys_task_logs"
	// TasksInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TasksInverseTable = "sys_tasks"
	// TasksColumn is the table column denoting the tasks relation/edge.
	TasksColumn = "task_task_logs"
)

// Columns holds all SQL columns for tasklog fields.
var Columns = []string{
	FieldID,
	FieldStartedAt,
	FieldFinishedAt,
	FieldResult,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "sys_task_logs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"task_task_logs",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultStartedAt holds the default value on creation for the "started_at" field.
	DefaultStartedAt func() time.Time
)
