package model

/*
int8 => -128 -> 127
int16 => -32768 -> 32767
int32 => -2 miliar -> + 2 miliar
int64 => -9 miliar miliar -> 9 miliar miliar

uint8 => 0 -> 255
uint16 => 0 -> 65535
unit32 => 0 -> 4 miliar
unit64 => 0 -> 18 miliar miliar
*/

type User struct {
	ID         string `gorm:"primaryKey;autoIncrement:true"`
	Name       string `gorm:"column:name"`
	Age        uint8  `gorm:"column:age"`
	Email      string `gorm:"type:varchar(50);unique;not null"`
	IsVerified bool   `gorm:"column:is_verified"`
}
