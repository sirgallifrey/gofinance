CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- create "users" table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    -- TODO: we need to change collation to make email case insensitive
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc'),
    updated_at TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc')
);

-- create "organization" table
CREATE TABLE organizations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) not NULL,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc'),
    updated_at TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc')
);

-- create "organization_members" table
CREATE TABLE organization_members (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    org_id UUID NOT NULL,
    user_id UUID NOT NULL,
    -- 30 is default member role
    role INT NOT NULL DEFAULT 30,
    created_at TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc'),
    updated_at TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc'),
    FOREIGN KEY (org_id) REFERENCES organizations(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- create "projects" table
CREATE TABLE projects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- org_id UUID NOT NULL,
    -- user_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc'),
    updated_at TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc')
    -- FOREIGN KEY (org_id) REFERENCES organizations(id) ON DELETE CASCADE,
    -- FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- create "features" table
CREATE TABLE features (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    num INT NOT NULL UNIQUE,
    project_id UUID NOT NULL,
    -- org_id UUID NOT NULL,
    -- user_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc'),
    updated_at TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc'),
    FOREIGN KEY (projectId) REFERENCES projects(id) ON DELETE CASCADE

    -- FOREIGN KEY (org_id) REFERENCES organizations(id) ON DELETE CASCADE,
    -- FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- create "features" table
CREATE TABLE requirements (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    num INT NOT NULL UNIQUE,
    feature_id UUID NOT NULL,
    parent_id UUID,
    -- org_id UUID NOT NULL,
    -- user_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    non_functional_type VARCHAR(40),
    created_at TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc'),
    updated_at TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc'),
    FOREIGN KEY (feature_id) REFERENCES features(id) ON DELETE CASCADE,
    FOREIGN KEY (parent_id) REFERENCES requirements(id) ON DELETE CASCADE

    -- FOREIGN KEY (org_id) REFERENCES organizations(id) ON DELETE CASCADE,
    -- FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
