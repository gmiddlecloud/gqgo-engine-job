// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/gmiddlecloud/gqgo-engine-job/ent/task"
	"github.com/gmiddlecloud/gqgo-engine-job/ent/tasklog"
)

const errInvalidPage = "INVALID_PAGE"

const (
	listField     = "list"
	pageNumField  = "pageNum"
	pageSizeField = "pageSize"
)

type PageDetails struct {
	Page  uint64 `json:"page"`
	Size  uint64 `json:"size"`
	Total uint64 `json:"total"`
}

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

const errInvalidPagination = "INVALID_PAGINATION"

type TaskPager struct {
	Order  OrderFunc
	Filter func(*TaskQuery) (*TaskQuery, error)
}

// TaskPaginateOption enables pagination customization.
type TaskPaginateOption func(*TaskPager)

// DefaultTaskOrder is the default ordering of Task.
var DefaultTaskOrder = Desc(task.FieldID)

func newTaskPager(opts []TaskPaginateOption) (*TaskPager, error) {
	pager := &TaskPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultTaskOrder
	}
	return pager, nil
}

func (p *TaskPager) ApplyFilter(query *TaskQuery) (*TaskQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// TaskPageList is Task PageList result.
type TaskPageList struct {
	List        []*Task      `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (t *TaskQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...TaskPaginateOption,
) (*TaskPageList, error) {

	pager, err := newTaskPager(opts)
	if err != nil {
		return nil, err
	}

	if t, err = pager.ApplyFilter(t); err != nil {
		return nil, err
	}

	ret := &TaskPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := t.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		t = t.Order(pager.Order)
	} else {
		t = t.Order(DefaultTaskOrder)
	}

	t = t.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := t.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type TaskLogPager struct {
	Order  OrderFunc
	Filter func(*TaskLogQuery) (*TaskLogQuery, error)
}

// TaskLogPaginateOption enables pagination customization.
type TaskLogPaginateOption func(*TaskLogPager)

// DefaultTaskLogOrder is the default ordering of TaskLog.
var DefaultTaskLogOrder = Desc(tasklog.FieldID)

func newTaskLogPager(opts []TaskLogPaginateOption) (*TaskLogPager, error) {
	pager := &TaskLogPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultTaskLogOrder
	}
	return pager, nil
}

func (p *TaskLogPager) ApplyFilter(query *TaskLogQuery) (*TaskLogQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// TaskLogPageList is TaskLog PageList result.
type TaskLogPageList struct {
	List        []*TaskLog   `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (tl *TaskLogQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...TaskLogPaginateOption,
) (*TaskLogPageList, error) {

	pager, err := newTaskLogPager(opts)
	if err != nil {
		return nil, err
	}

	if tl, err = pager.ApplyFilter(tl); err != nil {
		return nil, err
	}

	ret := &TaskLogPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := tl.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		tl = tl.Order(pager.Order)
	} else {
		tl = tl.Order(DefaultTaskLogOrder)
	}

	tl = tl.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := tl.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}