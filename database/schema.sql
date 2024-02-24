PRAGMA foreign_keys = ON;
-- create "users" table
CREATE TABLE `users` (
    `id` VARCHAR(26) NOT NULL PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    -- TODO: we need to change collation to make email case insensitive
    `email` VARCHAR(255) NOT NULL UNIQUE,
    `passwordHash` VARCHAR(255) NOT NULL,
    `createdAt` TEXT NOT NULL
);

-- create "organization" table
CREATE TABLE `organizations` (
    `id` VARCHAR(26) NOT NULL PRIMARY KEY,
    `name` VARCHAR(255) not NULL,
    `deletedAt` TEXT
);

-- create "organizationMemberRoleType" table
CREATE TABLE `organizationMemberRoleTypes` (
    `memberType`    VARCHAR(25)      PRIMARY KEY NOT NULL,
    `sequence`     INTEGER
);

-- create "organizationMembers" table
CREATE TABLE `organizationMembers` (
    `id` VARCHAR(26) NOT NULL PRIMARY KEY,
    `orgId` VARCHAR(26) NOT NULL,
    `userId` VARCHAR(26) NOT NULL,
    `role` TEXT NOT NULL DEFAULT ('member') REFERENCES organizationMemberRoleTypes(memberType),
    FOREIGN KEY (`orgId`) REFERENCES `organizations`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`userId`) REFERENCES `users`(`id`) ON DELETE CASCADE
);

-- create "transaction" table
CREATE TABLE `transactions` (
    `id` VARCHAR(26) NOT NULL,
    `orgId` VARCHAR(26) NOT NULL,
    `createdUserId` VARCHAR(26) NOT NULL,
    `name` VARCHAR(255) not NULL,
    `description` TEXT,
    `createdAt` DATETIME NOT NULL,
    `deletedAt` DATETIME,
    `centsValue` INTEGER, 
    PRIMARY KEY(`id`),
    FOREIGN KEY (`orgId`) REFERENCES `organizations`(id) ON DELETE RESTRICT
    FOREIGN KEY (`createdUserId`) REFERENCES `users`(id) ON DELETE RESTRICT
);

CREATE TABLE `transactionDetails` (
    `id` VARCHAR(26) NOT NULL,
    `orgId` VARCHAR(26) NOT NULL,
    `transactionId` VARCHAR(26) NOT NULL,
    `createdUserId` VARCHAR(26) NOT NULL,

    `name` VARCHAR(255) not NULL,
    `description` TEXT,
    `product` TEXT,
    `quantity` REAL,
    `centsValue` INTEGER,  

    `createdAt` DATETIME NOT NULL,
    `deletedAt` DATETIME,
    PRIMARY KEY(`id`),
    FOREIGN KEY (`orgId`) REFERENCES `organizations`(id) ON DELETE RESTRICT
    FOREIGN KEY (`transactionId`) REFERENCES `transactions`(id) ON DELETE RESTRICT
    FOREIGN KEY (`createdUserId`) REFERENCES `users`(id) ON DELETE RESTRICT
);