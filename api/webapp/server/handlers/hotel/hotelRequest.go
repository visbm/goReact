package hotel

type hotelRequest struct {
	HotelID    int    `json:"hotelId"`
	Name       string `json:"nameId"`
	Address    string `json:"addressId"`
	RoomsID    []int  `json:"roomsIds"`
	BookingsID []int  `json:"bookingsIds"`
}
