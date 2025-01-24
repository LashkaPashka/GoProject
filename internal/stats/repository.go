package stats

import (
	"go/project_go/pkg/db"
	"time"
	"gorm.io/datatypes"
)

type StatsRepository struct {
	*db.Db
}

func NewStatsRepository(db *db.Db) *StatsRepository {
	return &StatsRepository{
		Db: db,
	}
}

func (repo *StatsRepository) AddClick(linkId uint) {
	var stat Stat

	currentDate := datatypes.Date(time.Now())
	repo.Db.Find(stat, "link_id = ? and date = ?", linkId, currentDate)
	if stat.ID == 0 {
		repo.Db.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Date:   currentDate,
		})
	} else {
		stat.Clicks += 1
		repo.Db.Save(&stat)
	}
}

func (repo *StatsRepository) GetStat(from, to time.Time, by string) []StatResponse {
	var stats []StatResponse
	var selectByObject string

	switch (by){
		case GroupByDay:
			selectByObject = "char_to(date, YYYY-MMM-DD) as period, sum(clicks)"

		case GroupByMonth:
			selectByObject = "char_to(date, YYYY-MMM) as period, sum(clicks)"	
	}

	repo.DB.Table("links").
	Select(selectByObject).
	Where("date BETWEEN ? AND ?", from, to).
	Group("period").
	Order("period").
	Scan(&stats)

	return stats
}