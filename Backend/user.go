//Schema for user

type User struct {
    ID uint `gorm:"primaryKey"`
    Name string `gorm:"not null"`
    Email string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
    Role string `gorm:"default: 'viewer'"` //Admin, Editor, Viewer
}

