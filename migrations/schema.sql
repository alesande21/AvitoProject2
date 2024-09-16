-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS employee (
    id UUID PRIMARY KEY ,
    username VARCHAR(50) UNIQUE NOT NULL,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

-- CREATE TYPE organization_type AS ENUM (
--     'IE',
--     'LLC',
--     'JSC'
-- );

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'organization_type') THEN
CREATE TYPE organization_type AS ENUM ('IE', 'LLC', 'JSC');
END IF;
END $$;

CREATE TABLE IF NOT EXISTS organization (
    id UUID PRIMARY KEY ,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500),
    type organization_type,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS organization_responsible (
    id UUID PRIMARY KEY,
    organization_id UUID REFERENCES organization(id) ON DELETE CASCADE,
    user_id UUID REFERENCES employee(id) ON DELETE CASCADE
);

-- CREATE TYPE service_type AS ENUM (
--     'Construction',
--     'Delivery',
--     'Manufacture'
--     );

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'service_type') THEN
CREATE TYPE service_type AS ENUM ('Construction', 'Delivery', 'Manufacture');
END IF;
END $$;

-- CREATE TYPE tender_status AS ENUM (
--     'Created',
--     'Published',
--     'Closed'
--     );

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tender_status') THEN
CREATE TYPE tender_status AS ENUM ('Created', 'Published', 'Closed');
END IF;
END $$;

CREATE TABLE IF NOT EXISTS  tender (
    id UUID PRIMARY KEY ,
    status tender_status,
    organization_id UUID REFERENCES organization(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
--     проверить какое время использовать серверное или бд
);

CREATE TABLE IF NOT EXISTS  tender_condition (
    id SERIAL PRIMARY KEY ,
    tender_id UUID REFERENCES tender(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL ,
    description VARCHAR(500) NOT NULL ,
    type service_type ,
    version INT DEFAULT 1 CHECK ( version >= 1 ),
    UNIQUE (tender_id, version)
);


-- CREATE TYPE bid_status AS ENUM (
--     'Created',
--     'Published',
--     'Canceled'
--     );

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'bid_status') THEN
CREATE TYPE bid_status AS ENUM ('Created', 'Published', 'Canceled');
END IF;
END $$;


-- CREATE TYPE bid_decision AS ENUM (
--     'Approved',
--     'Rejected'
--     );

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'bid_decision') THEN
CREATE TYPE bid_decision AS ENUM ('Approved', 'Rejected');
END IF;
END $$;

-- CREATE TYPE bid_author_type AS ENUM (
--     'Organization',
--     'User'
--     );

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'bid_author_type') THEN
CREATE TYPE bid_author_type AS ENUM ('Organization', 'User');
END IF;
END $$;

CREATE TABLE IF NOT EXISTS bid (
    id UUID PRIMARY KEY ,
    status bid_status,
    tender_id UUID REFERENCES tender(id) ON DELETE CASCADE,
    author_type bid_author_type,
    author_id UUID NOT NULL ,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
    -- проверить какое время использовать серверное или бд
    -- разделить биды
);

CREATE TABLE IF NOT EXISTS bid_condition (
    id SERIAL PRIMARY KEY ,
    bid_id UUID REFERENCES bid(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL ,
    description VARCHAR(500) NOT NULL ,
    version INT DEFAULT 1 CHECK ( version >= 1 ),
    UNIQUE (bid_id, version)
);

CREATE TABLE IF NOT EXISTS bid_feedback (
    bid_id UUID REFERENCES bid(id) ON DELETE CASCADE,
    feedback VARCHAR(1000) NOT NULL ,
    username VARCHAR(50) NOT NULL
);

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'bid_response_type') THEN
CREATE TYPE bid_response_type AS ENUM ('Approved', 'Rejected');
END IF;
END $$;

CREATE TABLE IF NOT EXISTS bid_response (
    id SERIAL PRIMARY KEY ,
    bid_id UUID REFERENCES bid(id) ON DELETE CASCADE,
    response bid_response_type
    );