package model

type Money int32
type Timestamp int64

type Team struct {
	Name string `json:"name"`
}

type TeamRename struct {
	OldTeamName string `json:"old_team_name"`
	NewTeamName string `json:"new_team_name"`
}

type Member struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Active bool   `json:"active"`
	Secret string `json:"secret"`
}

type Order struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Timestamp    int64  `json:"timestamp"`
	FounderId    int64  `json:"founder_id"`
	DeliveryCost Money  `json:"delivery_cost"`
	TipCost      Money  `json:"tip_cost"`
}

type OrderMember struct {
	OrderId  int64 `json:"order_id"`
	MemberId int64 `json:"member_id"`
	Amount   Money `json:"amount"`
}

type Debt struct {
	DebtorId        int64     `json:"debtor_id"`
	CreditorId      int64     `json:"creditor_id"`
	Amount          Money     `json:"amount"`
	ReturnTimestamp Timestamp `json:"return_timestamp"`
}
