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

func newProjectScheduleRelation(db *gorm.DB, opts ...gen.DOOption) projectScheduleRelation {
	_projectScheduleRelation := projectScheduleRelation{}

	_projectScheduleRelation.projectScheduleRelationDo.UseDB(db, opts...)
	_projectScheduleRelation.projectScheduleRelationDo.UseModel(&model.ProjectScheduleRelation{})

	tableName := _projectScheduleRelation.projectScheduleRelationDo.TableName()
	_projectScheduleRelation.ALL = field.NewAsterisk(tableName)
	_projectScheduleRelation.ID = field.NewInt64(tableName, "id")
	_projectScheduleRelation.ProjectID = field.NewInt64(tableName, "project_id")
	_projectScheduleRelation.ScheduleID = field.NewInt64(tableName, "schedule_id")
	_projectScheduleRelation.CreateTime = field.NewTime(tableName, "create_time")

	_projectScheduleRelation.fillFieldMap()

	return _projectScheduleRelation
}

type projectScheduleRelation struct {
	projectScheduleRelationDo

	ALL        field.Asterisk
	ID         field.Int64
	ProjectID  field.Int64 // 项目id
	ScheduleID field.Int64 // 状态id
	CreateTime field.Time  // 创建时间

	fieldMap map[string]field.Expr
}

func (p projectScheduleRelation) Table(newTableName string) *projectScheduleRelation {
	p.projectScheduleRelationDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p projectScheduleRelation) As(alias string) *projectScheduleRelation {
	p.projectScheduleRelationDo.DO = *(p.projectScheduleRelationDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *projectScheduleRelation) updateTableName(table string) *projectScheduleRelation {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.ProjectID = field.NewInt64(table, "project_id")
	p.ScheduleID = field.NewInt64(table, "schedule_id")
	p.CreateTime = field.NewTime(table, "create_time")

	p.fillFieldMap()

	return p
}

func (p *projectScheduleRelation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *projectScheduleRelation) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 4)
	p.fieldMap["id"] = p.ID
	p.fieldMap["project_id"] = p.ProjectID
	p.fieldMap["schedule_id"] = p.ScheduleID
	p.fieldMap["create_time"] = p.CreateTime
}

func (p projectScheduleRelation) clone(db *gorm.DB) projectScheduleRelation {
	p.projectScheduleRelationDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p projectScheduleRelation) replaceDB(db *gorm.DB) projectScheduleRelation {
	p.projectScheduleRelationDo.ReplaceDB(db)
	return p
}

type projectScheduleRelationDo struct{ gen.DO }

type IProjectScheduleRelationDo interface {
	gen.SubQuery
	Debug() IProjectScheduleRelationDo
	WithContext(ctx context.Context) IProjectScheduleRelationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IProjectScheduleRelationDo
	WriteDB() IProjectScheduleRelationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IProjectScheduleRelationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IProjectScheduleRelationDo
	Not(conds ...gen.Condition) IProjectScheduleRelationDo
	Or(conds ...gen.Condition) IProjectScheduleRelationDo
	Select(conds ...field.Expr) IProjectScheduleRelationDo
	Where(conds ...gen.Condition) IProjectScheduleRelationDo
	Order(conds ...field.Expr) IProjectScheduleRelationDo
	Distinct(cols ...field.Expr) IProjectScheduleRelationDo
	Omit(cols ...field.Expr) IProjectScheduleRelationDo
	Join(table schema.Tabler, on ...field.Expr) IProjectScheduleRelationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IProjectScheduleRelationDo
	RightJoin(table schema.Tabler, on ...field.Expr) IProjectScheduleRelationDo
	Group(cols ...field.Expr) IProjectScheduleRelationDo
	Having(conds ...gen.Condition) IProjectScheduleRelationDo
	Limit(limit int) IProjectScheduleRelationDo
	Offset(offset int) IProjectScheduleRelationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IProjectScheduleRelationDo
	Unscoped() IProjectScheduleRelationDo
	Create(values ...*model.ProjectScheduleRelation) error
	CreateInBatches(values []*model.ProjectScheduleRelation, batchSize int) error
	Save(values ...*model.ProjectScheduleRelation) error
	First() (*model.ProjectScheduleRelation, error)
	Take() (*model.ProjectScheduleRelation, error)
	Last() (*model.ProjectScheduleRelation, error)
	Find() ([]*model.ProjectScheduleRelation, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProjectScheduleRelation, err error)
	FindInBatches(result *[]*model.ProjectScheduleRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ProjectScheduleRelation) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IProjectScheduleRelationDo
	Assign(attrs ...field.AssignExpr) IProjectScheduleRelationDo
	Joins(fields ...field.RelationField) IProjectScheduleRelationDo
	Preload(fields ...field.RelationField) IProjectScheduleRelationDo
	FirstOrInit() (*model.ProjectScheduleRelation, error)
	FirstOrCreate() (*model.ProjectScheduleRelation, error)
	FindByPage(offset int, limit int) (result []*model.ProjectScheduleRelation, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IProjectScheduleRelationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p projectScheduleRelationDo) Debug() IProjectScheduleRelationDo {
	return p.withDO(p.DO.Debug())
}

func (p projectScheduleRelationDo) WithContext(ctx context.Context) IProjectScheduleRelationDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p projectScheduleRelationDo) ReadDB() IProjectScheduleRelationDo {
	return p.Clauses(dbresolver.Read)
}

func (p projectScheduleRelationDo) WriteDB() IProjectScheduleRelationDo {
	return p.Clauses(dbresolver.Write)
}

func (p projectScheduleRelationDo) Session(config *gorm.Session) IProjectScheduleRelationDo {
	return p.withDO(p.DO.Session(config))
}

func (p projectScheduleRelationDo) Clauses(conds ...clause.Expression) IProjectScheduleRelationDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p projectScheduleRelationDo) Returning(value interface{}, columns ...string) IProjectScheduleRelationDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p projectScheduleRelationDo) Not(conds ...gen.Condition) IProjectScheduleRelationDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p projectScheduleRelationDo) Or(conds ...gen.Condition) IProjectScheduleRelationDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p projectScheduleRelationDo) Select(conds ...field.Expr) IProjectScheduleRelationDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p projectScheduleRelationDo) Where(conds ...gen.Condition) IProjectScheduleRelationDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p projectScheduleRelationDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IProjectScheduleRelationDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p projectScheduleRelationDo) Order(conds ...field.Expr) IProjectScheduleRelationDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p projectScheduleRelationDo) Distinct(cols ...field.Expr) IProjectScheduleRelationDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p projectScheduleRelationDo) Omit(cols ...field.Expr) IProjectScheduleRelationDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p projectScheduleRelationDo) Join(table schema.Tabler, on ...field.Expr) IProjectScheduleRelationDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p projectScheduleRelationDo) LeftJoin(table schema.Tabler, on ...field.Expr) IProjectScheduleRelationDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p projectScheduleRelationDo) RightJoin(table schema.Tabler, on ...field.Expr) IProjectScheduleRelationDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p projectScheduleRelationDo) Group(cols ...field.Expr) IProjectScheduleRelationDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p projectScheduleRelationDo) Having(conds ...gen.Condition) IProjectScheduleRelationDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p projectScheduleRelationDo) Limit(limit int) IProjectScheduleRelationDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p projectScheduleRelationDo) Offset(offset int) IProjectScheduleRelationDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p projectScheduleRelationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IProjectScheduleRelationDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p projectScheduleRelationDo) Unscoped() IProjectScheduleRelationDo {
	return p.withDO(p.DO.Unscoped())
}

func (p projectScheduleRelationDo) Create(values ...*model.ProjectScheduleRelation) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p projectScheduleRelationDo) CreateInBatches(values []*model.ProjectScheduleRelation, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p projectScheduleRelationDo) Save(values ...*model.ProjectScheduleRelation) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p projectScheduleRelationDo) First() (*model.ProjectScheduleRelation, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectScheduleRelation), nil
	}
}

func (p projectScheduleRelationDo) Take() (*model.ProjectScheduleRelation, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectScheduleRelation), nil
	}
}

func (p projectScheduleRelationDo) Last() (*model.ProjectScheduleRelation, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectScheduleRelation), nil
	}
}

func (p projectScheduleRelationDo) Find() ([]*model.ProjectScheduleRelation, error) {
	result, err := p.DO.Find()
	return result.([]*model.ProjectScheduleRelation), err
}

func (p projectScheduleRelationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProjectScheduleRelation, err error) {
	buf := make([]*model.ProjectScheduleRelation, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p projectScheduleRelationDo) FindInBatches(result *[]*model.ProjectScheduleRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p projectScheduleRelationDo) Attrs(attrs ...field.AssignExpr) IProjectScheduleRelationDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p projectScheduleRelationDo) Assign(attrs ...field.AssignExpr) IProjectScheduleRelationDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p projectScheduleRelationDo) Joins(fields ...field.RelationField) IProjectScheduleRelationDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p projectScheduleRelationDo) Preload(fields ...field.RelationField) IProjectScheduleRelationDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p projectScheduleRelationDo) FirstOrInit() (*model.ProjectScheduleRelation, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectScheduleRelation), nil
	}
}

func (p projectScheduleRelationDo) FirstOrCreate() (*model.ProjectScheduleRelation, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectScheduleRelation), nil
	}
}

func (p projectScheduleRelationDo) FindByPage(offset int, limit int) (result []*model.ProjectScheduleRelation, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p projectScheduleRelationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p projectScheduleRelationDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p projectScheduleRelationDo) Delete(models ...*model.ProjectScheduleRelation) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *projectScheduleRelationDo) withDO(do gen.Dao) *projectScheduleRelationDo {
	p.DO = *do.(*gen.DO)
	return p
}