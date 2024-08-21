package migration

const CreateCampaignTable = `
	CREATE TABLE IF NOT EXISTS campaigns(
		id TEXT PRIMARY KEY, 
		title TEXT
	);
`

const DropCampaignTable = `
	DROP TABLE IF EXISTS campaigns;
`

const CreateMailTable = `
	CREATE TABLE IF NOT EXISTS mails (
		id TEXT PRIMARY KEY,
		from TEXT NOT NULL,
		to TEXT NOT NULL,
		sent_time TIMESTAMP,
		open_time TIMESTAMP
	);
`

const DropMailTable = `
	DROP TABLE IF EXISTS mails;
`
