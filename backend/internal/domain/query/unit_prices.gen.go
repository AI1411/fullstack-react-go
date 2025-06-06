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

func newUnitPrice(db *gorm.DB, opts ...gen.DOOption) unitPrice {
	_unitPrice := unitPrice{}

	_unitPrice.unitPriceDo.UseDB(db, opts...)
	_unitPrice.unitPriceDo.UseModel(&model.UnitPrice{})

	tableName := _unitPrice.unitPriceDo.TableName()
	_unitPrice.ALL = field.NewAsterisk(tableName)
	_unitPrice.ID = field.NewInt64(tableName, "id")
	_unitPrice.CategoryID = field.NewInt32(tableName, "category_id")
	_unitPrice.PrefectureCode = field.NewString(tableName, "prefecture_code")
	_unitPrice.UnitPrice = field.NewFloat64(tableName, "unit_price")
	_unitPrice.UnitType = field.NewString(tableName, "unit_type")
	_unitPrice.ValidFrom = field.NewTime(tableName, "valid_from")
	_unitPrice.ValidTo = field.NewTime(tableName, "valid_to")
	_unitPrice.Notes = field.NewString(tableName, "notes")
	_unitPrice.CreatedAt = field.NewTime(tableName, "created_at")
	_unitPrice.UpdatedAt = field.NewTime(tableName, "updated_at")

	_unitPrice.fillFieldMap()

	return _unitPrice
}

type unitPrice struct {
	unitPriceDo

	ALL            field.Asterisk
	ID             field.Int64   // 単価ID（主キー、自動掲番）
	CategoryID     field.Int32   // 工種区分ID（外部キー、work_categoriesテーブルのID）
	PrefectureCode field.String  // 都道府県コード（外部キー、prefecturesテーブルのコード）
	UnitPrice      field.Float64 // 単価（円）
	UnitType       field.String  // 単位タイプ（"per_meter", "per_sqm", "per_unit"）
	ValidFrom      field.Time    // 有効開始日
	ValidTo        field.Time    // 有効終了日
	Notes          field.String  // 備考
	CreatedAt      field.Time    // 作成日時
	UpdatedAt      field.Time    // 更新日時

	fieldMap map[string]field.Expr
}

func (u unitPrice) Table(newTableName string) *unitPrice {
	u.unitPriceDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u unitPrice) As(alias string) *unitPrice {
	u.unitPriceDo.DO = *(u.unitPriceDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *unitPrice) updateTableName(table string) *unitPrice {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.CategoryID = field.NewInt32(table, "category_id")
	u.PrefectureCode = field.NewString(table, "prefecture_code")
	u.UnitPrice = field.NewFloat64(table, "unit_price")
	u.UnitType = field.NewString(table, "unit_type")
	u.ValidFrom = field.NewTime(table, "valid_from")
	u.ValidTo = field.NewTime(table, "valid_to")
	u.Notes = field.NewString(table, "notes")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")

	u.fillFieldMap()

	return u
}

func (u *unitPrice) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *unitPrice) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 10)
	u.fieldMap["id"] = u.ID
	u.fieldMap["category_id"] = u.CategoryID
	u.fieldMap["prefecture_code"] = u.PrefectureCode
	u.fieldMap["unit_price"] = u.UnitPrice
	u.fieldMap["unit_type"] = u.UnitType
	u.fieldMap["valid_from"] = u.ValidFrom
	u.fieldMap["valid_to"] = u.ValidTo
	u.fieldMap["notes"] = u.Notes
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
}

func (u unitPrice) clone(db *gorm.DB) unitPrice {
	u.unitPriceDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u unitPrice) replaceDB(db *gorm.DB) unitPrice {
	u.unitPriceDo.ReplaceDB(db)
	return u
}

type unitPriceDo struct{ gen.DO }

type IUnitPriceDo interface {
	gen.SubQuery
	Debug() IUnitPriceDo
	WithContext(ctx context.Context) IUnitPriceDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUnitPriceDo
	WriteDB() IUnitPriceDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUnitPriceDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUnitPriceDo
	Not(conds ...gen.Condition) IUnitPriceDo
	Or(conds ...gen.Condition) IUnitPriceDo
	Select(conds ...field.Expr) IUnitPriceDo
	Where(conds ...gen.Condition) IUnitPriceDo
	Order(conds ...field.Expr) IUnitPriceDo
	Distinct(cols ...field.Expr) IUnitPriceDo
	Omit(cols ...field.Expr) IUnitPriceDo
	Join(table schema.Tabler, on ...field.Expr) IUnitPriceDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUnitPriceDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUnitPriceDo
	Group(cols ...field.Expr) IUnitPriceDo
	Having(conds ...gen.Condition) IUnitPriceDo
	Limit(limit int) IUnitPriceDo
	Offset(offset int) IUnitPriceDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUnitPriceDo
	Unscoped() IUnitPriceDo
	Create(values ...*model.UnitPrice) error
	CreateInBatches(values []*model.UnitPrice, batchSize int) error
	Save(values ...*model.UnitPrice) error
	First() (*model.UnitPrice, error)
	Take() (*model.UnitPrice, error)
	Last() (*model.UnitPrice, error)
	Find() ([]*model.UnitPrice, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UnitPrice, err error)
	FindInBatches(result *[]*model.UnitPrice, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UnitPrice) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUnitPriceDo
	Assign(attrs ...field.AssignExpr) IUnitPriceDo
	Joins(fields ...field.RelationField) IUnitPriceDo
	Preload(fields ...field.RelationField) IUnitPriceDo
	FirstOrInit() (*model.UnitPrice, error)
	FirstOrCreate() (*model.UnitPrice, error)
	FindByPage(offset int, limit int) (result []*model.UnitPrice, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUnitPriceDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u unitPriceDo) Debug() IUnitPriceDo {
	return u.withDO(u.DO.Debug())
}

func (u unitPriceDo) WithContext(ctx context.Context) IUnitPriceDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u unitPriceDo) ReadDB() IUnitPriceDo {
	return u.Clauses(dbresolver.Read)
}

func (u unitPriceDo) WriteDB() IUnitPriceDo {
	return u.Clauses(dbresolver.Write)
}

func (u unitPriceDo) Session(config *gorm.Session) IUnitPriceDo {
	return u.withDO(u.DO.Session(config))
}

func (u unitPriceDo) Clauses(conds ...clause.Expression) IUnitPriceDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u unitPriceDo) Returning(value interface{}, columns ...string) IUnitPriceDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u unitPriceDo) Not(conds ...gen.Condition) IUnitPriceDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u unitPriceDo) Or(conds ...gen.Condition) IUnitPriceDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u unitPriceDo) Select(conds ...field.Expr) IUnitPriceDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u unitPriceDo) Where(conds ...gen.Condition) IUnitPriceDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u unitPriceDo) Order(conds ...field.Expr) IUnitPriceDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u unitPriceDo) Distinct(cols ...field.Expr) IUnitPriceDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u unitPriceDo) Omit(cols ...field.Expr) IUnitPriceDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u unitPriceDo) Join(table schema.Tabler, on ...field.Expr) IUnitPriceDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u unitPriceDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUnitPriceDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u unitPriceDo) RightJoin(table schema.Tabler, on ...field.Expr) IUnitPriceDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u unitPriceDo) Group(cols ...field.Expr) IUnitPriceDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u unitPriceDo) Having(conds ...gen.Condition) IUnitPriceDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u unitPriceDo) Limit(limit int) IUnitPriceDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u unitPriceDo) Offset(offset int) IUnitPriceDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u unitPriceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUnitPriceDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u unitPriceDo) Unscoped() IUnitPriceDo {
	return u.withDO(u.DO.Unscoped())
}

func (u unitPriceDo) Create(values ...*model.UnitPrice) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u unitPriceDo) CreateInBatches(values []*model.UnitPrice, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u unitPriceDo) Save(values ...*model.UnitPrice) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u unitPriceDo) First() (*model.UnitPrice, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnitPrice), nil
	}
}

func (u unitPriceDo) Take() (*model.UnitPrice, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnitPrice), nil
	}
}

func (u unitPriceDo) Last() (*model.UnitPrice, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnitPrice), nil
	}
}

func (u unitPriceDo) Find() ([]*model.UnitPrice, error) {
	result, err := u.DO.Find()
	return result.([]*model.UnitPrice), err
}

func (u unitPriceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UnitPrice, err error) {
	buf := make([]*model.UnitPrice, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u unitPriceDo) FindInBatches(result *[]*model.UnitPrice, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u unitPriceDo) Attrs(attrs ...field.AssignExpr) IUnitPriceDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u unitPriceDo) Assign(attrs ...field.AssignExpr) IUnitPriceDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u unitPriceDo) Joins(fields ...field.RelationField) IUnitPriceDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u unitPriceDo) Preload(fields ...field.RelationField) IUnitPriceDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u unitPriceDo) FirstOrInit() (*model.UnitPrice, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnitPrice), nil
	}
}

func (u unitPriceDo) FirstOrCreate() (*model.UnitPrice, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnitPrice), nil
	}
}

func (u unitPriceDo) FindByPage(offset int, limit int) (result []*model.UnitPrice, count int64, err error) {
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

func (u unitPriceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u unitPriceDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u unitPriceDo) Delete(models ...*model.UnitPrice) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *unitPriceDo) withDO(do gen.Dao) *unitPriceDo {
	u.DO = *do.(*gen.DO)
	return u
}
