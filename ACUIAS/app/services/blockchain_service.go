package services

// 操作区块链逻辑

import "test.com/hello/blockchain"

// 区块链服务接口
type BlockchainService struct {
	blockchainClient blockchain.BlockchainClient
	// ...
}

// 创建新的区块链服务
func NewBlockchainService(blockchainClient blockchain.BlockchainClient) *BlockchainService {
	return &BlockchainService{
		blockchainClient: blockchainClient,
	}

}

// 实现其他方法，如写入操作记录到联盟链等
