// mautrix-slack - A Matrix-Slack puppeting bridge.
// Copyright (C) 2022 Tulir Asokan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package database

import (
	_ "embed"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"maunium.net/go/mautrix/util/dbutil"

	"github.com/mautrix/slack/database/upgrades"
)

type Database struct {
	*dbutil.Database

	User       *UserQuery
	Portal     *PortalQuery
	Puppet     *PuppetQuery
	Message    *MessageQuery
	Reaction   *ReactionQuery
	Attachment *AttachmentQuery
}

func New(baseDB *dbutil.Database) *Database {
	db := &Database{Database: baseDB}

	db.UpgradeTable = upgrades.Table
	db.User = &UserQuery{
		db:  db,
		log: db.Log.Sub("User"),
	}
	db.Portal = &PortalQuery{
		db:  db,
		log: db.Log.Sub("Portal"),
	}
	db.Puppet = &PuppetQuery{
		db:  db,
		log: db.Log.Sub("Puppet"),
	}
	db.Message = &MessageQuery{
		db:  db,
		log: db.Log.Sub("Message"),
	}
	db.Reaction = &ReactionQuery{
		db:  db,
		log: db.Log.Sub("Reaction"),
	}
	db.Attachment = &AttachmentQuery{
		db:  db,
		log: db.Log.Sub("Attachment"),
	}

	return db
}
