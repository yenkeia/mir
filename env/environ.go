package env

import (
	"mir/orm"
	"github.com/jinzhu/gorm"
)

type Environ struct {
	DB   *gorm.DB
	Maps *map[uint32]Map
	AOI  *map[uint32][]AOIEntity
}

func InitEnviron() *Environ {
	db := orm.GetDB()

	db.AutoMigrate(&orm.AccountInfo{}, &orm.CharacterInfo{})

	maps := GetMaps(MapFilesPath)

	aoi := make(map[uint32][]AOIEntity)

	for i, m := range *maps {
		m.Index = i
		m.LoadNPC()
		m.LoadMonster()

		aoi[i] = m.GetAOIEntities()
	}

	return &Environ{
		DB:   db,
		Maps: maps,
		AOI:  &aoi,
	}
}
