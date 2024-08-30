package db


type SQLiteOptions struct {
    Name string
    Path string
}

func NewSQLiteConnection(opts *SQLiteOptions) *gorm.DB {
    return nil
}
