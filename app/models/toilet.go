package models
import (
  "github.com/nferruzzi/gormGIS"
  "time"
)

type Toilet struct {
  Gid         int64 //`gorm:"primary_key"`
  Gu_nm       string
  Hnr_nam     string
  Mtc_at      string
  Masterno    string
  Slaveno     string
  Neadres_nm  string
  Wc_nam      string
  Wc_gbn      string
  Hd_wc_yno   string
  Creat_de    time.Time
  Po_fe_nm    string
  Geom        gormGIS.GeoPoint
}
