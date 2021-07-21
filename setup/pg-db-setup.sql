CREATE TABLE task (
	id varchar PRIMARY KEY,
	user_id VARCHAR ( 50 )  NOT NULL,
	created_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        modified_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       task varchar(100),
       status varchar(20),
      is_active boolean
);