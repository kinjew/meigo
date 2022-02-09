//Package common is useful methods based on gorm
package common

import (
	"meigo/library/db"
)

// Create insert the value into database
func Create(value interface{}) error {

	return db.DB.Create(value).Error
}

// Save update value in database, if the value doesn't have primary key, will insert it
func Save(value interface{}) error {
	return db.DB.Save(value).Error
}

// Updates update attributes with callbacks
func Updates(where interface{}, value interface{}) error {
	return db.DB.Model(where).Updates(value).Error
}

// DeleteByModel delete by model
func DeleteByModel(model interface{}) (count int64, err error) {
	mdb := db.DB.Delete(model)
	err = mdb.Error
	if err != nil {
		return
	}
	count = mdb.RowsAffected
	return
}

// DeleteByWhere delete by where
func DeleteByWhere(model, where interface{}) (count int64, err error) {
	mdb := db.DB.Where(where).Delete(model)
	err = mdb.Error
	if err != nil {
		return
	}
	count = mdb.RowsAffected
	return
}

// DeleteByID delete by id
func DeleteByID(model interface{}, id uint64) (count int64, err error) {
	mdb := db.DB.Where("id=?", id).Delete(model)
	err = mdb.Error
	if err != nil {
		return
	}
	count = mdb.RowsAffected
	return
}

// DeleteByIDS delete by id set
func DeleteByIDS(model interface{}, ids []uint64) (count int64, err error) {
	mdb := db.DB.Where("id in (?)", ids).Delete(model)
	err = mdb.Error
	if err != nil {
		return
	}
	count = mdb.RowsAffected
	return
}

/*
// FirstByID find first record that match given conditions by id, order by primary key
func FirstByID(out interface{}, id int) (notFound bool, err error) {
	err = db.DB.First(out, id).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
	}
	return
}

// First find first record that match given conditions, order by primary key
func First(where interface{}, out interface{}) (notFound bool, err error) {
	err = db.DB.Where(where).First(out).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
	}
	return
}


*/
// Find find records that match given conditions
func Find(where interface{}, out interface{}, orders ...string) error {
	mdb := db.DB.Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			mdb = mdb.Order(order)
		}
	}
	return mdb.Find(out).Error
}

/*
// Scan scan value to a struct
func Scan(model, where interface{}, out interface{}) (notFound bool, err error) {
	err = db.DB.Model(model).Where(where).Scan(out).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
	}
	return
}

*/

// ScanList scan value to a struct and order
func ScanList(model, where interface{}, out interface{}, orders ...string) error {
	mdb := db.DB.Model(model).Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			mdb = mdb.Order(order)
		}
	}
	return mdb.Scan(out).Error
}

// GetPage is get page.
func GetPage(model, where interface{}, out interface{}, pageIndex, pageSize int, totalCount *int64, whereOrder ...PageWhereOrder) error {
	mdb := db.DB.Model(model).Where(where)
	if len(whereOrder) > 0 {
		for _, wo := range whereOrder {
			if wo.Order != "" {
				mdb = mdb.Order(wo.Order)
			}
			if wo.Where != "" {
				mdb = mdb.Where(wo.Where, wo.Value...)
			}
		}
	}
	err := mdb.Count(totalCount).Error
	if err != nil {
		return err
	}
	if *totalCount == 0 {
		return nil
	}
	return mdb.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(out).Error
}

// PluckList  used to query single column from a model as a map
func PluckList(model, where interface{}, out interface{}, fieldName string) error {
	return db.DB.Model(model).Where(where).Pluck(fieldName, out).Error
}
