package types

import (
	"database/sql/driver"
	"fmt"
)

type ContractType string

const (
	OneInchRouter  ContractType = "oneinch_router"
	UniV2Pair      ContractType = "uni_v2_pair"
	UniV3Pool      ContractType = "uni_v3_pool"
	DmmPool        ContractType = "dmm_pool"
	SwaapVault     ContractType = "swaap_vault"
	DodoexPool     ContractType = "dodoex_pool"
	BentoBoxV1Pool ContractType = "bento_box_v1_pool"
)

func (ct ContractType) Value() (driver.Value, error) {
	return string(ct), nil
}

func (ct *ContractType) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		*ct = ContractType(v)
	case []byte:
		*ct = ContractType(string(v))
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
	return nil
}
