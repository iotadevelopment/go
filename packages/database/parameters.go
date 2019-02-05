package database

import (
    "github.com/iotadevelopment/go/packages/parameter"
)

var DIRECTORY = parameter.IXI().AddString("DATABASE/DIRECTORY", "mainnetdb", "path to the database folder")
