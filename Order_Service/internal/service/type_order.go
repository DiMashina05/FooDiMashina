package service

import "time"

type Order struct {
    ID         string
    UserID     string
    Status     string
    Amount     int64
    CreatedAt  time.Time
    UpdatedAt  time.Time
}