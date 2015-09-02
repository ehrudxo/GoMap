package models

type User struct {
  Seq               int64 //`gorm:"primary_key"`
  Username          string
  Mail              string
  Password          string
  Description       string
}
