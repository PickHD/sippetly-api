package models

import "time"

type Quote struct{
	Id          int		 `json:"id"`
	QuoteName		string `json:"quoteName" binding:"required"`
	CreatedBy		string `json:"createdBy" binding:"required"`
	CreatedAt		time.Time `json:"createdAt"`
	UpdatedAt		time.Time	`json:"updatedAt"`
}