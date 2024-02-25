

-- create "users" table
CREATE TABLE users (
    id VARCHAR(26) NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    -- TODO: we need to change collation to make email case insensitive
    email VARCHAR(255) NOT NULL UNIQUE,
    passwordHash VARCHAR(255) NOT NULL,
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL
);

-- create "organization" table
CREATE TABLE organizations (
    id VARCHAR(26) NOT NULL PRIMARY KEY,
    name VARCHAR(255) not NULL,
    deletedAt TIMESTAMP,
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL
);

-- create "organizationMembers" table
CREATE TABLE organizationMembers (
    id VARCHAR(26) NOT NULL PRIMARY KEY,
    orgId VARCHAR(26) NOT NULL,
    userId VARCHAR(26) NOT NULL,
    -- 30 is default member role
    role INT NOT NULL DEFAULT 30,
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL,
    FOREIGN KEY (orgId) REFERENCES organizations(id) ON DELETE CASCADE,
    FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE
);

-- create "accounts" table
CREATE TABLE accounts (
    id VARCHAR(26) NOT NULL PRIMARY KEY,
    orgId VARCHAR(26) NOT NULL,
    name VARCHAR(255) not NULL,
    description TEXT,
    owned BOOLEAN NOT NULL,
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL,
    FOREIGN KEY (orgId) REFERENCES organizations(id) ON DELETE RESTRICT
);

-- create "transaction" table
CREATE TABLE transactions (
    id VARCHAR(26) NOT NULL PRIMARY KEY,
    orgId VARCHAR(26) NOT NULL,
    createdUserId VARCHAR(26) NOT NULL,
    name VARCHAR(255) not NULL,
    description TEXT,
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL,
    deletedAt TIMESTAMP,

    FOREIGN KEY (orgId) REFERENCES organizations(id) ON DELETE RESTRICT,
    FOREIGN KEY (createdUserId) REFERENCES users(id) ON DELETE RESTRICT
);

CREATE TABLE transactionDetails (
    id VARCHAR(26) NOT NULL PRIMARY KEY,
    orgId VARCHAR(26) NOT NULL,
    transactionId VARCHAR(26) NOT NULL,
    createdUserId VARCHAR(26) NOT NULL,

    name VARCHAR(255) not NULL,
    description TEXT,
    product TEXT,
    quantity DECIMAL,
    centsValue INTEGER,

    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL,
    deletedAt TIMESTAMP,

    FOREIGN KEY (orgId) REFERENCES organizations(id) ON DELETE RESTRICT,
    FOREIGN KEY (transactionId) REFERENCES transactions(id) ON DELETE RESTRICT,
    FOREIGN KEY (createdUserId) REFERENCES users(id) ON DELETE RESTRICT
);