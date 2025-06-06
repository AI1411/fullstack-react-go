// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
)

func newAssessment(db *gorm.DB, opts ...gen.DOOption) assessment {
	_assessment := assessment{}

	_assessment.assessmentDo.UseDB(db, opts...)
	_assessment.assessmentDo.UseModel(&model.Assessment{})

	tableName := _assessment.assessmentDo.TableName()
	_assessment.ALL = field.NewAsterisk(tableName)
	_assessment.ID = field.NewInt64(tableName, "id")
	_assessment.DisasterID = field.NewString(tableName, "disaster_id")
	_assessment.UserID = field.NewString(tableName, "user_id")
	_assessment.AssessmentDate = field.NewTime(tableName, "assessment_date")
	_assessment.AssessmentType = field.NewString(tableName, "assessment_type")
	_assessment.Status = field.NewString(tableName, "status")
	_assessment.AssessmentMethod = field.NewString(tableName, "assessment_method")
	_assessment.AssessmentSummary = field.NewString(tableName, "assessment_summary")
	_assessment.DamageAmount = field.NewFloat64(tableName, "damage_amount")
	_assessment.ApprovedAmount = field.NewFloat64(tableName, "approved_amount")
	_assessment.ApprovalDate = field.NewTime(tableName, "approval_date")
	_assessment.ApprovedBy = field.NewString(tableName, "approved_by")
	_assessment.Notes = field.NewString(tableName, "notes")
	_assessment.CreatedAt = field.NewTime(tableName, "created_at")
	_assessment.UpdatedAt = field.NewTime(tableName, "updated_at")
	_assessment.DeletedAt = field.NewField(tableName, "deleted_at")

	_assessment.fillFieldMap()

	return _assessment
}

type assessment struct {
	assessmentDo

	ALL               field.Asterisk
	ID                field.Int64   // 査定ID - 主キー
	DisasterID        field.String  // 災害ID - 査定対象の災害ID
	UserID            field.String  // 査定者ID - 査定を行ったユーザーのID
	AssessmentDate    field.Time    // 査定日 - 査定が行われた日付
	AssessmentType    field.String  // 査定種別 - 現地査定、リモート査定など
	Status            field.String  // 状態 - 査定の進行状況
	AssessmentMethod  field.String  // 査定方法 - 査定の実施方法
	AssessmentSummary field.String  // 査定概要 - 査定結果の概要
	DamageAmount      field.Float64 // 被害金額 - 査定された被害金額
	ApprovedAmount    field.Float64 // 承認金額 - 承認された支援金額
	ApprovalDate      field.Time    // 承認日時 - 査定が承認された日時
	ApprovedBy        field.String  // 承認者ID - 査定を承認したユーザーのID
	Notes             field.String  // 備考 - 査定に関する備考やメモ
	CreatedAt         field.Time    // 作成日時 - レコード作成日時
	UpdatedAt         field.Time    // 更新日時 - レコード最終更新日時
	DeletedAt         field.Field   // 削除日時 - 論理削除用のタイムスタンプ

	fieldMap map[string]field.Expr
}

func (a assessment) Table(newTableName string) *assessment {
	a.assessmentDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a assessment) As(alias string) *assessment {
	a.assessmentDo.DO = *(a.assessmentDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *assessment) updateTableName(table string) *assessment {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.DisasterID = field.NewString(table, "disaster_id")
	a.UserID = field.NewString(table, "user_id")
	a.AssessmentDate = field.NewTime(table, "assessment_date")
	a.AssessmentType = field.NewString(table, "assessment_type")
	a.Status = field.NewString(table, "status")
	a.AssessmentMethod = field.NewString(table, "assessment_method")
	a.AssessmentSummary = field.NewString(table, "assessment_summary")
	a.DamageAmount = field.NewFloat64(table, "damage_amount")
	a.ApprovedAmount = field.NewFloat64(table, "approved_amount")
	a.ApprovalDate = field.NewTime(table, "approval_date")
	a.ApprovedBy = field.NewString(table, "approved_by")
	a.Notes = field.NewString(table, "notes")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")

	a.fillFieldMap()

	return a
}

func (a *assessment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *assessment) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 16)
	a.fieldMap["id"] = a.ID
	a.fieldMap["disaster_id"] = a.DisasterID
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["assessment_date"] = a.AssessmentDate
	a.fieldMap["assessment_type"] = a.AssessmentType
	a.fieldMap["status"] = a.Status
	a.fieldMap["assessment_method"] = a.AssessmentMethod
	a.fieldMap["assessment_summary"] = a.AssessmentSummary
	a.fieldMap["damage_amount"] = a.DamageAmount
	a.fieldMap["approved_amount"] = a.ApprovedAmount
	a.fieldMap["approval_date"] = a.ApprovalDate
	a.fieldMap["approved_by"] = a.ApprovedBy
	a.fieldMap["notes"] = a.Notes
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
}

func (a assessment) clone(db *gorm.DB) assessment {
	a.assessmentDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a assessment) replaceDB(db *gorm.DB) assessment {
	a.assessmentDo.ReplaceDB(db)
	return a
}

type assessmentDo struct{ gen.DO }

type IAssessmentDo interface {
	gen.SubQuery
	Debug() IAssessmentDo
	WithContext(ctx context.Context) IAssessmentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAssessmentDo
	WriteDB() IAssessmentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAssessmentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAssessmentDo
	Not(conds ...gen.Condition) IAssessmentDo
	Or(conds ...gen.Condition) IAssessmentDo
	Select(conds ...field.Expr) IAssessmentDo
	Where(conds ...gen.Condition) IAssessmentDo
	Order(conds ...field.Expr) IAssessmentDo
	Distinct(cols ...field.Expr) IAssessmentDo
	Omit(cols ...field.Expr) IAssessmentDo
	Join(table schema.Tabler, on ...field.Expr) IAssessmentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAssessmentDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAssessmentDo
	Group(cols ...field.Expr) IAssessmentDo
	Having(conds ...gen.Condition) IAssessmentDo
	Limit(limit int) IAssessmentDo
	Offset(offset int) IAssessmentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAssessmentDo
	Unscoped() IAssessmentDo
	Create(values ...*model.Assessment) error
	CreateInBatches(values []*model.Assessment, batchSize int) error
	Save(values ...*model.Assessment) error
	First() (*model.Assessment, error)
	Take() (*model.Assessment, error)
	Last() (*model.Assessment, error)
	Find() ([]*model.Assessment, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Assessment, err error)
	FindInBatches(result *[]*model.Assessment, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Assessment) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAssessmentDo
	Assign(attrs ...field.AssignExpr) IAssessmentDo
	Joins(fields ...field.RelationField) IAssessmentDo
	Preload(fields ...field.RelationField) IAssessmentDo
	FirstOrInit() (*model.Assessment, error)
	FirstOrCreate() (*model.Assessment, error)
	FindByPage(offset int, limit int) (result []*model.Assessment, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAssessmentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a assessmentDo) Debug() IAssessmentDo {
	return a.withDO(a.DO.Debug())
}

func (a assessmentDo) WithContext(ctx context.Context) IAssessmentDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a assessmentDo) ReadDB() IAssessmentDo {
	return a.Clauses(dbresolver.Read)
}

func (a assessmentDo) WriteDB() IAssessmentDo {
	return a.Clauses(dbresolver.Write)
}

func (a assessmentDo) Session(config *gorm.Session) IAssessmentDo {
	return a.withDO(a.DO.Session(config))
}

func (a assessmentDo) Clauses(conds ...clause.Expression) IAssessmentDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a assessmentDo) Returning(value interface{}, columns ...string) IAssessmentDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a assessmentDo) Not(conds ...gen.Condition) IAssessmentDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a assessmentDo) Or(conds ...gen.Condition) IAssessmentDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a assessmentDo) Select(conds ...field.Expr) IAssessmentDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a assessmentDo) Where(conds ...gen.Condition) IAssessmentDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a assessmentDo) Order(conds ...field.Expr) IAssessmentDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a assessmentDo) Distinct(cols ...field.Expr) IAssessmentDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a assessmentDo) Omit(cols ...field.Expr) IAssessmentDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a assessmentDo) Join(table schema.Tabler, on ...field.Expr) IAssessmentDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a assessmentDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAssessmentDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a assessmentDo) RightJoin(table schema.Tabler, on ...field.Expr) IAssessmentDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a assessmentDo) Group(cols ...field.Expr) IAssessmentDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a assessmentDo) Having(conds ...gen.Condition) IAssessmentDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a assessmentDo) Limit(limit int) IAssessmentDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a assessmentDo) Offset(offset int) IAssessmentDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a assessmentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAssessmentDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a assessmentDo) Unscoped() IAssessmentDo {
	return a.withDO(a.DO.Unscoped())
}

func (a assessmentDo) Create(values ...*model.Assessment) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a assessmentDo) CreateInBatches(values []*model.Assessment, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a assessmentDo) Save(values ...*model.Assessment) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a assessmentDo) First() (*model.Assessment, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Assessment), nil
	}
}

func (a assessmentDo) Take() (*model.Assessment, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Assessment), nil
	}
}

func (a assessmentDo) Last() (*model.Assessment, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Assessment), nil
	}
}

func (a assessmentDo) Find() ([]*model.Assessment, error) {
	result, err := a.DO.Find()
	return result.([]*model.Assessment), err
}

func (a assessmentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Assessment, err error) {
	buf := make([]*model.Assessment, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a assessmentDo) FindInBatches(result *[]*model.Assessment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a assessmentDo) Attrs(attrs ...field.AssignExpr) IAssessmentDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a assessmentDo) Assign(attrs ...field.AssignExpr) IAssessmentDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a assessmentDo) Joins(fields ...field.RelationField) IAssessmentDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a assessmentDo) Preload(fields ...field.RelationField) IAssessmentDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a assessmentDo) FirstOrInit() (*model.Assessment, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Assessment), nil
	}
}

func (a assessmentDo) FirstOrCreate() (*model.Assessment, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Assessment), nil
	}
}

func (a assessmentDo) FindByPage(offset int, limit int) (result []*model.Assessment, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a assessmentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a assessmentDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a assessmentDo) Delete(models ...*model.Assessment) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *assessmentDo) withDO(do gen.Dao) *assessmentDo {
	a.DO = *do.(*gen.DO)
	return a
}
