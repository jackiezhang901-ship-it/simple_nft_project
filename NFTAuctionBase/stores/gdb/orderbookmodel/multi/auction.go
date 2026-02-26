package multi

import (
	"fmt"
	"github.com/shopspring/decimal"
	"time"
)

const (
	MarketAuction = iota
)

type Auction struct {
	AuctionId       uint64          `gorm:"column:auction_id;AUTO_INCREMENT;primary_key" json:"auction_id"`
	TokenId         string          `gorm:"column:token_id" json:"token_id"`
	Seller          string          `gorm:"column:seller" json:"seller"`
	AuctionKey      string          `gorm:"column:auction_key" json:"auction_key"`
	AuctionType     string          `gorm:"column:auction_type" json:"auction_type"`
	StartPrice      decimal.Decimal `gorm:"column:start_price" json:"start_price"`
	CurrentPrice    decimal.Decimal `gorm:"column:current_price" json:"current_price"`
	MinBidIncrement decimal.Decimal `gorm:"column:min_bid_increment" json:"min_bid_increment"`
	StartTime       time.Time       `gorm:"column:start_time" json:"start_time"`
	EndTime         time.Time       `gorm:"column:end_time" json:"end_time"`
	BidCount        int             `gorm:"column:bid_count" json:"bid_count"`
	Status          string          `gorm:"column:status;NOT NULL" json:"status"`
	Winner          string          `gorm:"column:winner" json:"winner"`
	CreatedAt       time.Time       `gorm:"column:created_at;NOT NULL" json:"created_at"`
	UpdatedAt       time.Time       `gorm:"column:updated_at;NOT NULL" json:"updated_at"`
}

type AuctionBid struct {
	BidId           uint64    `gorm:"column:bid_id;AUTO_INCREMENT;primary_key" json:"bid_id"`
	AuctionId       uint64    `gorm:"column:auction_id;NOT NULL" json:"auction_id"`
	Bidder          string    `gorm:"column:bidder;NOT NULL" json:"bidder"`
	BidAmount       float32   `gorm:"column:bid_amount;NOT NULL" json:"bid_amount"`
	TransactionHash string    `gorm:"column:transaction_hash" json:"transaction_hash"`
	BidStatus       string    `gorm:"column:status" json:"status"`
	CreatedAt       time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;NOT NULL" json:"updated_at"`
}

func AuctionTableName(chainName string) string {
	return fmt.Sprintf("aunctions_%s", chainName)
}
