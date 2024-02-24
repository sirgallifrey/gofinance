-- Create "users" table
CREATE TABLE `users` (`id` varchar NOT NULL, `name` varchar NOT NULL, `email` varchar NOT NULL, `passwordHash` varchar NOT NULL, `createdAt` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "users_email" to table: "users"
CREATE UNIQUE INDEX `users_email` ON `users` (`email`);
-- Create "organizations" table
CREATE TABLE `organizations` (`id` varchar NOT NULL, `name` varchar NOT NULL, `deletedAt` text NULL, PRIMARY KEY (`id`));
-- Create "organizationMemberRoleTypes" table
CREATE TABLE `organizationMemberRoleTypes` (`memberType` varchar NOT NULL, `sequence` integer NULL, PRIMARY KEY (`memberType`));
-- Create "organizationMembers" table
CREATE TABLE `organizationMembers` (`id` varchar NOT NULL, `orgId` varchar NOT NULL, `userId` varchar NOT NULL, `role` text NOT NULL DEFAULT 'member', PRIMARY KEY (`id`), CONSTRAINT `0` FOREIGN KEY (`userId`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT `1` FOREIGN KEY (`orgId`) REFERENCES `organizations` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT `2` FOREIGN KEY (`role`) REFERENCES `organizationMemberRoleTypes` (`memberType`) ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create "transactions" table
CREATE TABLE `transactions` (`id` varchar NOT NULL, `orgId` varchar NOT NULL, `createdUserId` varchar NOT NULL, `name` varchar NOT NULL, `description` text NULL, `createdAt` datetime NOT NULL, `deletedAt` datetime NULL, `centsValue` integer NULL, PRIMARY KEY (`id`), CONSTRAINT `0` FOREIGN KEY (`createdUserId`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE RESTRICT, CONSTRAINT `1` FOREIGN KEY (`orgId`) REFERENCES `organizations` (`id`) ON UPDATE NO ACTION ON DELETE RESTRICT);
-- Create "transactionDetails" table
CREATE TABLE `transactionDetails` (`id` varchar NOT NULL, `orgId` varchar NOT NULL, `transactionId` varchar NOT NULL, `createdUserId` varchar NOT NULL, `name` varchar NOT NULL, `description` text NULL, `product` text NULL, `quantity` real NULL, `centsValue` integer NULL, `createdAt` datetime NOT NULL, `deletedAt` datetime NULL, PRIMARY KEY (`id`), CONSTRAINT `0` FOREIGN KEY (`createdUserId`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE RESTRICT, CONSTRAINT `1` FOREIGN KEY (`transactionId`) REFERENCES `transactions` (`id`) ON UPDATE NO ACTION ON DELETE RESTRICT, CONSTRAINT `2` FOREIGN KEY (`orgId`) REFERENCES `organizations` (`id`) ON UPDATE NO ACTION ON DELETE RESTRICT);
