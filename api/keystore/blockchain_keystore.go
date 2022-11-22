// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package keystore

import (
	"go.uber.org/zap"

	"github.com/coinflect/coinflectchain/database"
	"github.com/coinflect/coinflectchain/database/encdb"
	"github.com/coinflect/coinflectchain/ids"
	"github.com/coinflect/coinflectchain/utils/logging"
)

var _ BlockchainKeystore = (*blockchainKeystore)(nil)

type BlockchainKeystore interface {
	// Get a database that is able to read and write unencrypted values from the
	// underlying database.
	GetDatabase(username, password string) (*encdb.Database, error)

	// Get the underlying database that is able to read and write encrypted
	// values. This Database will not perform any encrypting or decrypting of
	// values and is not recommended to be used when implementing a VM.
	GetRawDatabase(username, password string) (database.Database, error)
}

type blockchainKeystore struct {
	blockchainID ids.ID
	ks           *keystore
}

func (bks *blockchainKeystore) GetDatabase(username, password string) (*encdb.Database, error) {
	bks.ks.log.Debug("Keystore: GetDatabase called",
		logging.UserString("username", username),
		zap.Stringer("blockchainID", bks.blockchainID),
	)

	return bks.ks.GetDatabase(bks.blockchainID, username, password)
}

func (bks *blockchainKeystore) GetRawDatabase(username, password string) (database.Database, error) {
	bks.ks.log.Debug("Keystore: GetRawDatabase called",
		logging.UserString("username", username),
		zap.Stringer("blockchainID", bks.blockchainID),
	)

	return bks.ks.GetRawDatabase(bks.blockchainID, username, password)
}