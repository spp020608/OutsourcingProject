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

func newUserEducation(db *gorm.DB, opts ...gen.DOOption) userEducation {
	_userEducation := userEducation{}

	_userEducation.userEducationDo.UseDB(db, opts...)
	_userEducation.userEducationDo.UseModel(&model.UserEducation{})

	tableName := _userEducation.userEducationDo.TableName()
	_userEducation.ALL = field.NewAsterisk(tableName)
	_userEducation.ID = field.NewInt64(tableName, "id")
	_userEducation.UserID = field.NewInt64(tableName, "user_id")
	_userEducation.CollegeName = field.NewString(tableName, "college_name")
	_userEducation.MajorName = field.NewString(tableName, "major_name")
	_userEducation.DegreeName = field.NewString(tableName, "degree_name")
	_userEducation.EducationIntro = field.NewString(tableName, "education_intro")
	_userEducation.StartTime = field.NewTime(tableName, "start_time")
	_userEducation.EndTime = field.NewTime(tableName, "end_time")
	_userEducation.CreateTime = field.NewTime(tableName, "create_time")
	_userEducation.EntranceTime = field.NewString(tableName, "entrance_time")

	_userEducation.fillFieldMap()

	return _userEducation
}

type userEducation struct {
	userEducationDo

	ALL            field.Asterisk
	ID             field.Int64
	UserID         field.Int64  // 用户id
	CollegeName    field.String // 学校名称
	MajorName      field.String // 专业名称
	DegreeName     field.String // 学位名称
	EducationIntro field.String // 教育经历简介
	StartTime      field.Time   // 开始时间
	EndTime        field.Time   // 结束时间
	CreateTime     field.Time   // 记录创建时间
	EntranceTime   field.String // 入学时间

	fieldMap map[string]field.Expr
}

func (u userEducation) Table(newTableName string) *userEducation {
	u.userEducationDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userEducation) As(alias string) *userEducation {
	u.userEducationDo.DO = *(u.userEducationDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userEducation) updateTableName(table string) *userEducation {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.UserID = field.NewInt64(table, "user_id")
	u.CollegeName = field.NewString(table, "college_name")
	u.MajorName = field.NewString(table, "major_name")
	u.DegreeName = field.NewString(table, "degree_name")
	u.EducationIntro = field.NewString(table, "education_intro")
	u.StartTime = field.NewTime(table, "start_time")
	u.EndTime = field.NewTime(table, "end_time")
	u.CreateTime = field.NewTime(table, "create_time")
	u.EntranceTime = field.NewString(table, "entrance_time")

	u.fillFieldMap()

	return u
}

func (u *userEducation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userEducation) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 10)
	u.fieldMap["id"] = u.ID
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["college_name"] = u.CollegeName
	u.fieldMap["major_name"] = u.MajorName
	u.fieldMap["degree_name"] = u.DegreeName
	u.fieldMap["education_intro"] = u.EducationIntro
	u.fieldMap["start_time"] = u.StartTime
	u.fieldMap["end_time"] = u.EndTime
	u.fieldMap["create_time"] = u.CreateTime
	u.fieldMap["entrance_time"] = u.EntranceTime
}

func (u userEducation) clone(db *gorm.DB) userEducation {
	u.userEducationDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userEducation) replaceDB(db *gorm.DB) userEducation {
	u.userEducationDo.ReplaceDB(db)
	return u
}

type userEducationDo struct{ gen.DO }

type IUserEducationDo interface {
	gen.SubQuery
	Debug() IUserEducationDo
	WithContext(ctx context.Context) IUserEducationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserEducationDo
	WriteDB() IUserEducationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserEducationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserEducationDo
	Not(conds ...gen.Condition) IUserEducationDo
	Or(conds ...gen.Condition) IUserEducationDo
	Select(conds ...field.Expr) IUserEducationDo
	Where(conds ...gen.Condition) IUserEducationDo
	Order(conds ...field.Expr) IUserEducationDo
	Distinct(cols ...field.Expr) IUserEducationDo
	Omit(cols ...field.Expr) IUserEducationDo
	Join(table schema.Tabler, on ...field.Expr) IUserEducationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserEducationDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserEducationDo
	Group(cols ...field.Expr) IUserEducationDo
	Having(conds ...gen.Condition) IUserEducationDo
	Limit(limit int) IUserEducationDo
	Offset(offset int) IUserEducationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserEducationDo
	Unscoped() IUserEducationDo
	Create(values ...*model.UserEducation) error
	CreateInBatches(values []*model.UserEducation, batchSize int) error
	Save(values ...*model.UserEducation) error
	First() (*model.UserEducation, error)
	Take() (*model.UserEducation, error)
	Last() (*model.UserEducation, error)
	Find() ([]*model.UserEducation, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserEducation, err error)
	FindInBatches(result *[]*model.UserEducation, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserEducation) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserEducationDo
	Assign(attrs ...field.AssignExpr) IUserEducationDo
	Joins(fields ...field.RelationField) IUserEducationDo
	Preload(fields ...field.RelationField) IUserEducationDo
	FirstOrInit() (*model.UserEducation, error)
	FirstOrCreate() (*model.UserEducation, error)
	FindByPage(offset int, limit int) (result []*model.UserEducation, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserEducationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userEducationDo) Debug() IUserEducationDo {
	return u.withDO(u.DO.Debug())
}

func (u userEducationDo) WithContext(ctx context.Context) IUserEducationDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userEducationDo) ReadDB() IUserEducationDo {
	return u.Clauses(dbresolver.Read)
}

func (u userEducationDo) WriteDB() IUserEducationDo {
	return u.Clauses(dbresolver.Write)
}

func (u userEducationDo) Session(config *gorm.Session) IUserEducationDo {
	return u.withDO(u.DO.Session(config))
}

func (u userEducationDo) Clauses(conds ...clause.Expression) IUserEducationDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userEducationDo) Returning(value interface{}, columns ...string) IUserEducationDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userEducationDo) Not(conds ...gen.Condition) IUserEducationDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userEducationDo) Or(conds ...gen.Condition) IUserEducationDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userEducationDo) Select(conds ...field.Expr) IUserEducationDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userEducationDo) Where(conds ...gen.Condition) IUserEducationDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userEducationDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IUserEducationDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u userEducationDo) Order(conds ...field.Expr) IUserEducationDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userEducationDo) Distinct(cols ...field.Expr) IUserEducationDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userEducationDo) Omit(cols ...field.Expr) IUserEducationDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userEducationDo) Join(table schema.Tabler, on ...field.Expr) IUserEducationDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userEducationDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserEducationDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userEducationDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserEducationDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userEducationDo) Group(cols ...field.Expr) IUserEducationDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userEducationDo) Having(conds ...gen.Condition) IUserEducationDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userEducationDo) Limit(limit int) IUserEducationDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userEducationDo) Offset(offset int) IUserEducationDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userEducationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserEducationDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userEducationDo) Unscoped() IUserEducationDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userEducationDo) Create(values ...*model.UserEducation) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userEducationDo) CreateInBatches(values []*model.UserEducation, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userEducationDo) Save(values ...*model.UserEducation) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userEducationDo) First() (*model.UserEducation, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserEducation), nil
	}
}

func (u userEducationDo) Take() (*model.UserEducation, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserEducation), nil
	}
}

func (u userEducationDo) Last() (*model.UserEducation, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserEducation), nil
	}
}

func (u userEducationDo) Find() ([]*model.UserEducation, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserEducation), err
}

func (u userEducationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserEducation, err error) {
	buf := make([]*model.UserEducation, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userEducationDo) FindInBatches(result *[]*model.UserEducation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userEducationDo) Attrs(attrs ...field.AssignExpr) IUserEducationDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userEducationDo) Assign(attrs ...field.AssignExpr) IUserEducationDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userEducationDo) Joins(fields ...field.RelationField) IUserEducationDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userEducationDo) Preload(fields ...field.RelationField) IUserEducationDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userEducationDo) FirstOrInit() (*model.UserEducation, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserEducation), nil
	}
}

func (u userEducationDo) FirstOrCreate() (*model.UserEducation, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserEducation), nil
	}
}

func (u userEducationDo) FindByPage(offset int, limit int) (result []*model.UserEducation, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userEducationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userEducationDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userEducationDo) Delete(models ...*model.UserEducation) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userEducationDo) withDO(do gen.Dao) *userEducationDo {
	u.DO = *do.(*gen.DO)
	return u
}
