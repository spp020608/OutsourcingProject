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

func newProjectNotPlatform(db *gorm.DB, opts ...gen.DOOption) projectNotPlatform {
	_projectNotPlatform := projectNotPlatform{}

	_projectNotPlatform.projectNotPlatformDo.UseDB(db, opts...)
	_projectNotPlatform.projectNotPlatformDo.UseModel(&model.ProjectNotPlatform{})

	tableName := _projectNotPlatform.projectNotPlatformDo.TableName()
	_projectNotPlatform.ALL = field.NewAsterisk(tableName)
	_projectNotPlatform.ID = field.NewInt64(tableName, "id")
	_projectNotPlatform.UserID = field.NewInt64(tableName, "user_id")
	_projectNotPlatform.ProjectName = field.NewString(tableName, "project_name")
	_projectNotPlatform.ProjectIntro = field.NewString(tableName, "project_intro")
	_projectNotPlatform.StartTime = field.NewTime(tableName, "start_time")
	_projectNotPlatform.EndTime = field.NewTime(tableName, "end_time")
	_projectNotPlatform.CreateTime = field.NewTime(tableName, "create_time")

	_projectNotPlatform.fillFieldMap()

	return _projectNotPlatform
}

type projectNotPlatform struct {
	projectNotPlatformDo

	ALL          field.Asterisk
	ID           field.Int64
	UserID       field.Int64  // 用户id
	ProjectName  field.String // 非平台项目名称
	ProjectIntro field.String // 非平台项目介绍
	StartTime    field.Time   // 非平台项目开始时间
	EndTime      field.Time   // 非平台项目结束时间
	CreateTime   field.Time

	fieldMap map[string]field.Expr
}

func (p projectNotPlatform) Table(newTableName string) *projectNotPlatform {
	p.projectNotPlatformDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p projectNotPlatform) As(alias string) *projectNotPlatform {
	p.projectNotPlatformDo.DO = *(p.projectNotPlatformDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *projectNotPlatform) updateTableName(table string) *projectNotPlatform {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.UserID = field.NewInt64(table, "user_id")
	p.ProjectName = field.NewString(table, "project_name")
	p.ProjectIntro = field.NewString(table, "project_intro")
	p.StartTime = field.NewTime(table, "start_time")
	p.EndTime = field.NewTime(table, "end_time")
	p.CreateTime = field.NewTime(table, "create_time")

	p.fillFieldMap()

	return p
}

func (p *projectNotPlatform) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *projectNotPlatform) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 7)
	p.fieldMap["id"] = p.ID
	p.fieldMap["user_id"] = p.UserID
	p.fieldMap["project_name"] = p.ProjectName
	p.fieldMap["project_intro"] = p.ProjectIntro
	p.fieldMap["start_time"] = p.StartTime
	p.fieldMap["end_time"] = p.EndTime
	p.fieldMap["create_time"] = p.CreateTime
}

func (p projectNotPlatform) clone(db *gorm.DB) projectNotPlatform {
	p.projectNotPlatformDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p projectNotPlatform) replaceDB(db *gorm.DB) projectNotPlatform {
	p.projectNotPlatformDo.ReplaceDB(db)
	return p
}

type projectNotPlatformDo struct{ gen.DO }

type IProjectNotPlatformDo interface {
	gen.SubQuery
	Debug() IProjectNotPlatformDo
	WithContext(ctx context.Context) IProjectNotPlatformDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IProjectNotPlatformDo
	WriteDB() IProjectNotPlatformDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IProjectNotPlatformDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IProjectNotPlatformDo
	Not(conds ...gen.Condition) IProjectNotPlatformDo
	Or(conds ...gen.Condition) IProjectNotPlatformDo
	Select(conds ...field.Expr) IProjectNotPlatformDo
	Where(conds ...gen.Condition) IProjectNotPlatformDo
	Order(conds ...field.Expr) IProjectNotPlatformDo
	Distinct(cols ...field.Expr) IProjectNotPlatformDo
	Omit(cols ...field.Expr) IProjectNotPlatformDo
	Join(table schema.Tabler, on ...field.Expr) IProjectNotPlatformDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IProjectNotPlatformDo
	RightJoin(table schema.Tabler, on ...field.Expr) IProjectNotPlatformDo
	Group(cols ...field.Expr) IProjectNotPlatformDo
	Having(conds ...gen.Condition) IProjectNotPlatformDo
	Limit(limit int) IProjectNotPlatformDo
	Offset(offset int) IProjectNotPlatformDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IProjectNotPlatformDo
	Unscoped() IProjectNotPlatformDo
	Create(values ...*model.ProjectNotPlatform) error
	CreateInBatches(values []*model.ProjectNotPlatform, batchSize int) error
	Save(values ...*model.ProjectNotPlatform) error
	First() (*model.ProjectNotPlatform, error)
	Take() (*model.ProjectNotPlatform, error)
	Last() (*model.ProjectNotPlatform, error)
	Find() ([]*model.ProjectNotPlatform, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProjectNotPlatform, err error)
	FindInBatches(result *[]*model.ProjectNotPlatform, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ProjectNotPlatform) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IProjectNotPlatformDo
	Assign(attrs ...field.AssignExpr) IProjectNotPlatformDo
	Joins(fields ...field.RelationField) IProjectNotPlatformDo
	Preload(fields ...field.RelationField) IProjectNotPlatformDo
	FirstOrInit() (*model.ProjectNotPlatform, error)
	FirstOrCreate() (*model.ProjectNotPlatform, error)
	FindByPage(offset int, limit int) (result []*model.ProjectNotPlatform, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IProjectNotPlatformDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p projectNotPlatformDo) Debug() IProjectNotPlatformDo {
	return p.withDO(p.DO.Debug())
}

func (p projectNotPlatformDo) WithContext(ctx context.Context) IProjectNotPlatformDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p projectNotPlatformDo) ReadDB() IProjectNotPlatformDo {
	return p.Clauses(dbresolver.Read)
}

func (p projectNotPlatformDo) WriteDB() IProjectNotPlatformDo {
	return p.Clauses(dbresolver.Write)
}

func (p projectNotPlatformDo) Session(config *gorm.Session) IProjectNotPlatformDo {
	return p.withDO(p.DO.Session(config))
}

func (p projectNotPlatformDo) Clauses(conds ...clause.Expression) IProjectNotPlatformDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p projectNotPlatformDo) Returning(value interface{}, columns ...string) IProjectNotPlatformDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p projectNotPlatformDo) Not(conds ...gen.Condition) IProjectNotPlatformDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p projectNotPlatformDo) Or(conds ...gen.Condition) IProjectNotPlatformDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p projectNotPlatformDo) Select(conds ...field.Expr) IProjectNotPlatformDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p projectNotPlatformDo) Where(conds ...gen.Condition) IProjectNotPlatformDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p projectNotPlatformDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IProjectNotPlatformDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p projectNotPlatformDo) Order(conds ...field.Expr) IProjectNotPlatformDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p projectNotPlatformDo) Distinct(cols ...field.Expr) IProjectNotPlatformDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p projectNotPlatformDo) Omit(cols ...field.Expr) IProjectNotPlatformDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p projectNotPlatformDo) Join(table schema.Tabler, on ...field.Expr) IProjectNotPlatformDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p projectNotPlatformDo) LeftJoin(table schema.Tabler, on ...field.Expr) IProjectNotPlatformDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p projectNotPlatformDo) RightJoin(table schema.Tabler, on ...field.Expr) IProjectNotPlatformDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p projectNotPlatformDo) Group(cols ...field.Expr) IProjectNotPlatformDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p projectNotPlatformDo) Having(conds ...gen.Condition) IProjectNotPlatformDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p projectNotPlatformDo) Limit(limit int) IProjectNotPlatformDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p projectNotPlatformDo) Offset(offset int) IProjectNotPlatformDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p projectNotPlatformDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IProjectNotPlatformDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p projectNotPlatformDo) Unscoped() IProjectNotPlatformDo {
	return p.withDO(p.DO.Unscoped())
}

func (p projectNotPlatformDo) Create(values ...*model.ProjectNotPlatform) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p projectNotPlatformDo) CreateInBatches(values []*model.ProjectNotPlatform, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p projectNotPlatformDo) Save(values ...*model.ProjectNotPlatform) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p projectNotPlatformDo) First() (*model.ProjectNotPlatform, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectNotPlatform), nil
	}
}

func (p projectNotPlatformDo) Take() (*model.ProjectNotPlatform, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectNotPlatform), nil
	}
}

func (p projectNotPlatformDo) Last() (*model.ProjectNotPlatform, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectNotPlatform), nil
	}
}

func (p projectNotPlatformDo) Find() ([]*model.ProjectNotPlatform, error) {
	result, err := p.DO.Find()
	return result.([]*model.ProjectNotPlatform), err
}

func (p projectNotPlatformDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProjectNotPlatform, err error) {
	buf := make([]*model.ProjectNotPlatform, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p projectNotPlatformDo) FindInBatches(result *[]*model.ProjectNotPlatform, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p projectNotPlatformDo) Attrs(attrs ...field.AssignExpr) IProjectNotPlatformDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p projectNotPlatformDo) Assign(attrs ...field.AssignExpr) IProjectNotPlatformDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p projectNotPlatformDo) Joins(fields ...field.RelationField) IProjectNotPlatformDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p projectNotPlatformDo) Preload(fields ...field.RelationField) IProjectNotPlatformDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p projectNotPlatformDo) FirstOrInit() (*model.ProjectNotPlatform, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectNotPlatform), nil
	}
}

func (p projectNotPlatformDo) FirstOrCreate() (*model.ProjectNotPlatform, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectNotPlatform), nil
	}
}

func (p projectNotPlatformDo) FindByPage(offset int, limit int) (result []*model.ProjectNotPlatform, count int64, err error) {
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

func (p projectNotPlatformDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p projectNotPlatformDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p projectNotPlatformDo) Delete(models ...*model.ProjectNotPlatform) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *projectNotPlatformDo) withDO(do gen.Dao) *projectNotPlatformDo {
	p.DO = *do.(*gen.DO)
	return p
}
