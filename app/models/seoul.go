//It will be deleted
package models
import (
  "bytes"
  "database/sql/driver"
	"encoding/binary"
	"encoding/hex"
  "fmt"
)
type MultiPolygon struct {
	Lng float64
	Lat float64
}
func (p *MultiPolygon) String() string {
	return fmt.Sprintf("SRID=4326;POINT(%v %v)", p.Lng, p.Lat)
}
func (p *MultiPolygon) Scan(val interface{}) error {
  b, err := hex.DecodeString(string(val.([]uint8)))
	if err != nil {
		return err
	}
	r := bytes.NewReader(b)
	var wkbByteOrder uint8
	if err := binary.Read(r, binary.LittleEndian, &wkbByteOrder); err != nil {
		return err
	}

	var byteOrder binary.ByteOrder
	switch wkbByteOrder {
	case 0:
		byteOrder = binary.BigEndian
	case 1:
		byteOrder = binary.LittleEndian
	default:
		return fmt.Errorf("Invalid byte order %d", wkbByteOrder)
	}

	var wkbGeometryType uint64
	if err := binary.Read(r, byteOrder, &wkbGeometryType); err != nil {
		return err
	}

	if err := binary.Read(r, byteOrder, p); err != nil {
		return err
	}
	return nil
}
func (p MultiPolygon) Value() (driver.Value, error) {
	return p.String(), nil
}
type Seoul struct {
  gid         int64 //`gorm:"primary_key"`
  Code        string
  First_do    string
  First_gu    string
  First_dong  string
  Geom        MultiPolygon
}
