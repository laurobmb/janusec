/*
 * @Copyright Reserved By Janusec (https://www.janusec.com/).
 * @Author: U2
 * @Date: 2023-05-28 12:45:10
 */

package models

type CookieType int64

const (
	Cookie_Necessary    CookieType = 1
	Cookie_Analytics    CookieType = 1 << 1
	Cookie_Marketing    CookieType = 1 << 2
	Cookie_Unclassified CookieType = 1 << 9 // 512
)

// Cookie used by applications
type Cookie struct {
	ID     int64      `json:"id,string"`
	AppID  int64      `json:"app_id,string"`
	Name   string     `json:"name"`
	Vendor string     `json:"vendor"`
	Type   CookieType `json:"type"`
}

// CookieBase used for classification automatically
type CookieBase struct {
	ID     int64      `json:"id,string"`
	Name   string     `json:"name"`
	Vendor string     `json:"vendor"`
	Type   CookieType `json:"type"`
}
