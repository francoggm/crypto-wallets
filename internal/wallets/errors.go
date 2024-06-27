package wallets

import "errors"

var (
	FailedGettingWallets = "Failed getting wallets!"
	FailedGettingWallet  = "Failed getting wallet!"
	WalletNotFound       = "Wallet not found!"
	InsufficientAmount   = "User don't have sufficient amount!"
	InvalidAmount        = "Invalid amount!"
	InvalidBodyValues    = "Invalid values!"
	FailedBuyingAsset    = "Error buying asset!"
)

var (
	ErrFailedGettingUserWallets = errors.New(FailedGettingWallets)
	ErrFailedGettingUserWallet  = errors.New(FailedGettingWallet)
	ErrWalletNotFound           = errors.New(WalletNotFound)
	ErrInsufficientAmount       = errors.New(InsufficientAmount)
	ErrInvalidAmount            = errors.New(InvalidAmount)
	ErrInvalidBodyValues        = errors.New(InvalidBodyValues)
	ErrFailedBuyingAsset        = errors.New(FailedBuyingAsset)
)
