package stats

import (
	"go/project_go/pkg/db"
	"time"

	"gorm.io/datatypes"
)

type StatsRepository struct {
	datebase *db.Db
}

func NewStatsRepository(db *db.Db) *StatsRepository {
	return &StatsRepository{
		datebase: db,
	}
}

func (repo *StatsRepository) AddClick(linkId uint) {
	var stat Stat

	currentDate := datatypes.Date(time.Now())
	repo.datebase.DB.Table("stats").Find(&stat, "link_id = ? and date = ?", linkId, currentDate)
	
	if stat.Id == 0 {
		repo.datebase.DB.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Date:   currentDate,
		})
	} else {
		stat.Clicks += 1
		repo.datebase.DB.Save(&stat)
	}
}

func (repo *StatsRepository) GetStat(from, to time.Time, by string) []StatResponse {
	var stats []StatResponse
	var selectByObject string

	switch (by){
		case GroupByDay:
			selectByObject = "to_char(date, YYYY-MMM-DD) AS period, SUM(clicks)"
		case GroupByMonth:
			selectByObject = "to_char(date, YYYY-MMM) AS period, SUM(clicks)"	
	}

	repo.datebase.DB.Table("stats").
	Select(selectByObject).
	Where("date BETWEEN ? AND ?", from, to).
	Group("period").
	Order("period").
	Scan(&stats)

	return stats
}