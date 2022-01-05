CREATE TABLE card_issuer
(
  card_issuer_key numeric(16,0) NOT NULL,
  "version" numeric(5,0),
  created_by numeric(16,0) NOT NULL,
  created_on timestamp without time zone NOT NULL,
  updated_by numeric(16,0) NOT NULL,
  updated_on timestamp without time zone NOT NULL,
  first_name text not null,
  last_name text not null,
  email_id text not null,
  mobile text not null,
  password text not null,
  access_key text ,
  CONSTRAINT card_issuer_pkey PRIMARY KEY (card_issuer_key),
  CONSTRAINT mobile_email_ukey UNIQUE (mobile,email_id)
);
--DROP DATABASE message_routing;