package database

import (
    "github.com/iotadevelopment/go/packages/parameter"
)

var DIRECTORY = parameter.AddString("DATABASE/DIRECTORY", "mainnetdb", "path to the database folder")
