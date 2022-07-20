package database

import (
	"github.com/flashbots/boost-relay/common"
)

var (
	tableBase                  = common.GetEnv("DB_TABLE_PREFIX", "dev")
	TableEvent                 = tableBase + "_event"
	TableValidatorRegistration = tableBase + "_validator_registration"
	TableEpochSummary          = tableBase + "_epoch_summary"
)

var schema = `CREATE TABLE IF NOT EXISTS ` + TableEvent + ` (
	id bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
	inserted_at timestamp NOT NULL default current_timestamp,

	slot       bigint NOT NULL,
	epoch      bigint NOT NULL,
	event_type varchar(255) NOT NULL,
	event_data jsonb NOT NULL
);

CREATE TABLE IF NOT EXISTS ` + TableValidatorRegistration + ` (
	id bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
	inserted_at timestamp NOT NULL default current_timestamp,

	pubkey varchar(98) NOT NULL,
	registration jsonb NOT NULL,
	registration_timestamp timestamp NOT NULL,

	UNIQUE (pubkey, registration_timestamp)
);


CREATE TABLE IF NOT EXISTS ` + TableEpochSummary + ` (
	id bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
	inserted_at timestamp NOT NULL default current_timestamp,

	epoch      bigint NOT NULL UNIQUE,
	slot_first bigint NOT NULL,
	slot_last  bigint NOT NULL,

	validators_known_total          bigint NOT NULL,
	validator_registrations_total   bigint NOT NULL,
	validator_registrations_renewed bigint NOT NULL,
	validator_registrations_new     bigint NOT NULL,

	num_register_validator_requests bigint NOT NULL,
	num_get_header_requests         bigint NOT NULL,
	num_get_payload_requests        bigint NOT NULL,

	num_header_sent          bigint NOT NULL,
	num_header_no_content    bigint NOT NULL,
	num_payload_sent         bigint NOT NULL,
	num_builder_bid_received bigint NOT NULL
);
`
