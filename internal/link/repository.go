package link

import (
	"go/project_go/pkg/db"

	"gorm.io/gorm/clause"
)

type LinkRepository struct{
	Database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository{
	return &LinkRepository{
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error){
	result := repo.Database.DB.Table("links").Create(link)
	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error){
	var link Link
	result := repo.Database.DB.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	
	return &link, nil
}

func (repo *LinkRepository) Upload(link *Link) (*Link, error){
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(link)

	if result.Error != nil {
		return nil, result.Error
	}
	
	return link, nil
}

func (repo *LinkRepository) Delete(id uint) error {
	result := repo.Database.DB.Delete(&Link{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *LinkRepository) GetById(id uint) (*Link, error){
	var link Link
	result := repo.Database.DB.First(&link, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}

func (repo *LinkRepository) Count() int64 {
	var count int64

	repo.Database.DB.Table("links").
	Where("deleted_at is null").
	Count(&count)

	return count
}


func (repo *LinkRepository) GetAll(limit, offset int) []Link{
	var link []Link

	repo.Database.DB.Table("links").
	Where("deleted_at is null").
	Order("id ASC").
	Limit(limit).
	Offset(offset).
	Scan(&link)

	return link
}