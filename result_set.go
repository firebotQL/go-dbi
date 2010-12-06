package dbi

import "os"

type ResultSet interface {
    Len() uint64
    Next() bool
    GetStringByCol(col int) (string, os.Error)
    GetString(col string) (string, os.Error)
    Close() os.Error
}

