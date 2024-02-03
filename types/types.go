package types

// struct defined for GPS data
type OBUData struct {
	OBUID int     `json:"ubuID"`
	Lat   float64 `json:"lat"`
	Long  float64 `josn:"long"`
}
