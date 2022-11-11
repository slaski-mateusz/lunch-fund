package model

type Money int32
type Timestamp int64

type Team struct {
	TeamName string `json:"team_name"`
}

type TeamRename struct {
	OldTeamName string `json:"old_team_name"`
	NewTeamName string `json:"new_team_name"`
}

type Member struct {
	Team
	Id         int64  `json:"id"`
	MemberName string `json:"member_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	IsAdmin    int64  `json:"is_admin"`
	IsActive   int64  `json:"is_active"`
	Secret     string `json:"secret"`
}

type Order struct {
	Team
	Id           int64  `json:"id"`
	OrderName    string `json:"order_name"`
	Timestamp    int64  `json:"timestamp"`
	DeliveryCost Money  `json:"delivery_cost"`
	TipCost      Money  `json:"tip_cost"`
}

type OrderDetail struct {
	Team
	OrderId   int64 `json:"order_id"`
	MemberId  int64 `json:"member_id"`
	IsFounder int64 `json:"is_founder"`
	Amount    Money `json:"amount"`
}

type Debt struct {
	Team
	DebtorId        int64     `json:"debtor_id"`
	CreditorId      int64     `json:"creditor_id"`
	Amount          Money     `json:"amount"`
	ReturnTimestamp Timestamp `json:"return_timestamp"`
}
