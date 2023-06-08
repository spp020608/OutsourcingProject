// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"Outsourcing/cmd/gorm/model"
)

func newWorkState(db *gorm.DB, opts ...gen.DOOption) workState {
	_workState := workState{}

	_workState.workStateDo.UseDB(db, opts...)
	_workState.workStateDo.UseModel(&model.WorkState{})

	tableName := _workState.workStateDo.TableName()
	_workState.ALL = field.NewAsterisk(tableName)
	_workState.ID = field.NewInt64(tableName, "id")
	_workState.StateType = field.NewString(tableName, "state_type")
	_workState.StateExplain = field.NewString(tableName, "state_explain")
	_workState.CreateTime = field.NewTime(tableName, "create_time")

	_workState.fillFieldMap()

	return _workState
}

type workState struct {
	workStateDo

	ALL          field.Asterisk
	ID           field.Int64
	StateType    field.String // 工作状态类型
	StateExplain field.String // 工作状态说明
	CreateTime   field.Time   // 记录创建时间

	fieldMap map[string]field.Expr
}

func (w workState) Table(newTableName string) *workState {
	w.workStateDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w workState) As(alias string) *workState {
	w.workStateDo.DO = *(w.workStateDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *workState) updateTableName(table string) *workState {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.StateType = field.NewString(table, "state_type")
	w.StateExplain = field.NewString(table, "state_explain")
	w.CreateTime = field.NewTime(table, "create_time")

	w.fillFieldMap()

	return w
}

func (w *workState) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *workState) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 4)
	w.fieldMap["id"] = w.ID
	w.fieldMap["state_type"] = w.StateType
	w.fieldMap["state_explain"] = w.StateExplain
	w.fieldMap["create_time"] = w.CreateTime
}

func (w workState) clone(db *gorm.DB) workState {
	w.workStateDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w workState) replaceDB(db *gorm.DB) workState {
	w.workStateDo.ReplaceDB(db)
	return w
}

type workStateDo struct{ gen.DO }

type IWorkStateDo interface {
	gen.SubQuery
	Debug() IWorkStateDo
	WithContext(ctx context.Context) IWorkStateDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWorkStateDo
	WriteDB() IWorkStateDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWorkStateDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWorkStateDo
	Not(conds ...gen.Condition) IWorkStateDo
	Or(conds ...gen.Condition) IWorkStateDo
	Select(conds ...field.Expr) IWorkStateDo
	Where(conds ...gen.Condition) IWorkStateDo
	Order(conds ...field.Expr) IWorkStateDo
	Distinct(cols ...field.Expr) IWorkStateDo
	Omit(cols ...field.Expr) IWorkStateDo
	Join(table schema.Tabler, on ...field.Expr) IWorkStateDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWorkStateDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWorkStateDo
	Group(cols ...field.Expr) IWorkStateDo
	Having(conds ...gen.Condition) IWorkStateDo
	Limit(limit int) IWorkStateDo
	Offset(offset int) IWorkStateDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWorkStateDo
	Unscoped() IWorkStateDo
	Create(values ...*model.WorkState) error
	CreateInBatches(values []*model.WorkState, batchSize int) error
	Save(values ...*model.WorkState) error
	First() (*model.WorkState, error)
	Take() (*model.WorkState, error)
	Last() (*model.WorkState, error)
	Find() ([]*model.WorkState, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WorkState, err error)
	FindInBatches(result *[]*model.WorkState, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WorkState) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWorkStateDo
	Assign(attrs ...field.AssignExpr) IWorkStateDo
	Joins(fields ...field.RelationField) IWorkStateDo
	Preload(fields ...field.RelationField) IWorkStateDo
	FirstOrInit() (*model.WorkState, error)
	FirstOrCreate() (*model.WorkState, error)
	FindByPage(offset int, limit int) (result []*model.WorkState, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWorkStateDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w workStateDo) Debug() IWorkStateDo {
	return w.withDO(w.DO.Debug())
}

func (w workStateDo) WithContext(ctx context.Context) IWorkStateDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w workStateDo) ReadDB() IWorkStateDo {
	return w.Clauses(dbresolver.Read)
}

func (w workStateDo) WriteDB() IWorkStateDo {
	return w.Clauses(dbresolver.Write)
}

func (w workStateDo) Session(config *gorm.Session) IWorkStateDo {
	return w.withDO(w.DO.Session(config))
}

func (w workStateDo) Clauses(conds ...clause.Expression) IWorkStateDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w workStateDo) Returning(value interface{}, columns ...string) IWorkStateDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w workStateDo) Not(conds ...gen.Condition) IWorkStateDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w workStateDo) Or(conds ...gen.Condition) IWorkStateDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w workStateDo) Select(conds ...field.Expr) IWorkStateDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w workStateDo) Where(conds ...gen.Condition) IWorkStateDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w workStateDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IWorkStateDo {
	return w.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (w workStateDo) Order(conds ...field.Expr) IWorkStateDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w workStateDo) Distinct(cols ...field.Expr) IWorkStateDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w workStateDo) Omit(cols ...field.Expr) IWorkStateDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w workStateDo) Join(table schema.Tabler, on ...field.Expr) IWorkStateDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w workStateDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWorkStateDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w workStateDo) RightJoin(table schema.Tabler, on ...field.Expr) IWorkStateDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w workStateDo) Group(cols ...field.Expr) IWorkStateDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w workStateDo) Having(conds ...gen.Condition) IWorkStateDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w workStateDo) Limit(limit int) IWorkStateDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w workStateDo) Offset(offset int) IWorkStateDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w workStateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWorkStateDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w workStateDo) Unscoped() IWorkStateDo {
	return w.withDO(w.DO.Unscoped())
}

func (w workStateDo) Create(values ...*model.WorkState) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w workStateDo) CreateInBatches(values []*model.WorkState, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w workStateDo) Save(values ...*model.WorkState) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w workStateDo) First() (*model.WorkState, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WorkState), nil
	}
}

func (w workStateDo) Take() (*model.WorkState, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WorkState), nil
	}
}

func (w workStateDo) Last() (*model.WorkState, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WorkState), nil
	}
}

func (w workStateDo) Find() ([]*model.WorkState, error) {
	result, err := w.DO.Find()
	return result.([]*model.WorkState), err
}

func (w workStateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WorkState, err error) {
	buf := make([]*model.WorkState, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w workStateDo) FindInBatches(result *[]*model.WorkState, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w workStateDo) Attrs(attrs ...field.AssignExpr) IWorkStateDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w workStateDo) Assign(attrs ...field.AssignExpr) IWorkStateDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w workStateDo) Joins(fields ...field.RelationField) IWorkStateDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w workStateDo) Preload(fields ...field.RelationField) IWorkStateDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w workStateDo) FirstOrInit() (*model.WorkState, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WorkState), nil
	}
}

func (w workStateDo) FirstOrCreate() (*model.WorkState, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WorkState), nil
	}
}

func (w workStateDo) FindByPage(offset int, limit int) (result []*model.WorkState, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w workStateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w workStateDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w workStateDo) Delete(models ...*model.WorkState) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *workStateDo) withDO(do gen.Dao) *workStateDo {
	w.DO = *do.(*gen.DO)
	return w
}
