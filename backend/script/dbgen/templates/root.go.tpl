// Code generated by "script/dbgen". DO NOT EDIT.
package daocore

import (
    "context"
    "database/sql"
    "strings"
    "time"

    "github.com/Masterminds/squirrel"

    "github.com/youcan-jpn/dab/backend/src/dberror"
    "github.com/youcan-jpn/dab/backend/src/util/filter"
)

{{ template "metadata" . }}
{{ template "typedef" . }}
{{ template "query" . }}
{{ template "command" . }}