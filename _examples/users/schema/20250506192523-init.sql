-- +migrate Up
CREATE TABLE `users` (
    `id` uuid NOT NULL,
    `age` integer NOT NULL,
    `name` text NOT NULL DEFAULT 'unknown',
    `token` text NOT NULL DEFAULT 'abab',
    PRIMARY KEY (`id`)
);

-- Create a view public_users without the token column
CREATE VIEW `public_users` AS
SELECT `id`,
    `age`,
    `name`
FROM `users`;

-- +migrate Down
DROP TABLE `users`;