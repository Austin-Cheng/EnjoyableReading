use papers;

create table if not exists `t_tag`(
    `id` int auto_increment DEFAULT NULL,
    `name` varchar(255) DEFAULT NULL,
    `parent_id` int DEFAULT NULL,
    `created_at` int64 not null,
    `updated_at` int64 not null
)


