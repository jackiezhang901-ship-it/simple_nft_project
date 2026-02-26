package multi

import (
	"fmt"
	"time"
)

type Nft struct {
	NftId             int64     `gorm:"column:nft_id;AUTO_INCREMENT;primary_key" json:"nft_id"`       // 主键
	TokenId           string    `gorm:"column:token_id;NOT NULL" json:"token_id"`                     // 区块链上的Token ID
	ContractAddress   string    `gorm:"column:contract_address;NOT NULL" json:"contract_address"`     // NFT合约地址
	ChainId           int64     `gorm:"column:chain_id;NOT NULL" json:"chain_id"`                     // 所属区块链网络ID
	Category          string    `gorm:"column:category;NOT NULL" json:"category"`                     // 类别id
	Name              string    `gorm:"column:name;NOT NULL" json:"name"`                             // NFT名称
	Description       string    `gorm:"column:description;NOT NULL" json:"description"`               // NFT描述
	ImagUrl           string    `gorm:"column:image_url;NOT NULL" json:"image_url"`                   // NFT图片URL
	ThumbnailUrl      string    `gorm:"column:thumbnail_url;NOT NULL" json:"thumbnail_url"`           // NFT缩略图URL
	MetadataUrl       string    `gorm:"column:metadata_url;NOT NULL" json:"metadata_url"`             // NFT元数据URL
	CreatorId         string    `gorm:"column:creator_id;NOT NULL" json:"creator_id"`                 // 创作者用户ID
	OwnerId           string    `gorm:"column:owner_id;NOT NULL" json:"owner_id"`                     // 当前拥有者用户ID
	RoyaltyPercentage string    `gorm:"column:royalty_percentage;NOT NULL" json:"royalty_percentage"` // 版税百分比，如10.00表示10%
	TokenStandard     string    `gorm:"column:token_standard;NOT NULL" json:"token_standard"`         // Token标准：ERC721（非同质化）、ERC1155（半同质化）、其他
	TotalSupply       int64     `gorm:"column:total_supply;NOT NULL" json:"total_supply"`             // 总供应量，ERC721为1，ERC1155可以多个
	IsMinted          int64     `gorm:"column:is_minted;NOT NULL" json:"is_minted"`                   // 是否已在区块链上铸造-0：否，1：是
	MintedAt          time.Time `gorm:"column:minted_at;NOT NULL" json:"minted_at"`                   // 铸造时间
	CreatedAt         time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`                 // 创建时间
	UpdatedAt         time.Time `gorm:"column:updated_at;NOT NULL" json:"updated_at"`                 // 更新时间
	Status            string    `gorm:"column:status;NOT NULL" json:"status"`                         // NFT状态:active-inactive-deleted：活跃、非活跃、已删除
}

type NftAttributes struct {
	AttributeId      int64     `gorm:"column:attribute_id;AUTO_INCREMENT;primary_key" json:"attribute_id"`
	TokenId          string    `gorm:"column:token_id;NOT NULL" json:"token_id"`
	TraitType        string    `gorm:"column:trait_type;NOT NULL" json:"trait_type"`
	TraitValue       string    `gorm:"column:trait_value;NOT NULL" json:"trait_value"`
	DisplayType      string    `gorm:"column:display_type" json:"display_type"`
	RarityPercentage float32   `gorm:"column:rarity_percentage" json:"rarity_percentage"`
	CreatedAt        time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`
}

func NftTableName() string {
	return fmt.Sprintf("nfts")
}
