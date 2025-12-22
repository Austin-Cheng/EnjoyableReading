use paper;

create table if not exists `t_tags`(
    `id` int auto_increment NOT NULL COMMENT '主键',
    `name`  varchar(255) NOT NULL COMMENT '名称',
    `parent_id` int DEFAULT NULL,
    `created_at` datetime(3) NOT NULL default current_timestamp(3) COMMENT '创建时间',
    `updated_at` datetime(3) NOT NULL default current_timestamp(3) on update current_timestamp(3) COMMENT '更新时间',
    PRIMARY KEY (`id`)
) COMMENT='标签表';


-- papers.papers definition

CREATE TABLE if  not exists  `t_papers` (
    `id` int auto_increment NOT NULL COMMENT '主键',
    `address` varchar(1024) DEFAULT NULL COMMENT  '文章地址',
    `title` varchar(255) NOT NULL  COMMENT  '文章英文标题',
    `title_ch` varchar(512) NOT NULL  COMMENT  '文章中文标题',
    `authors` varchar(512) DEFAULT NULL  COMMENT  '作者',
    `published_time` datetime NOT NULL   COMMENT  '发布时间',
    `summary` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT  '英文文概括',
    `summary_ch` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT  '中文概括',
    `categories` varchar(512) DEFAULT NULL COMMENT  '阅读数量',
    `filepath` varchar(512) DEFAULT NULL COMMENT  '文件名',
    `read` int  NOT NULL DEFAULT 0 COMMENT  '阅读数量' ,
    `favorite` int  NOT NULL DEFAULT 0 COMMENT  '收藏数量',
    `fulltext_ch`  varchar(512)  COMMENT  '我也不知道是啥',
    `created_at` datetime(3) NOT NULL default current_timestamp(3) COMMENT '创建时间',
    `updated_at` datetime(3) NOT NULL default current_timestamp(3) on update current_timestamp(3) COMMENT '更新时间',
    PRIMARY KEY (`id`)
) COMMENT='文章表';



-- papers.paper_tags definition

CREATE TABLE if  not exists  `t_paper_tags` (
    `id` int NOT NULL AUTO_INCREMENT  COMMENT '主键',
    `paper_id` int  NOT NULL COMMENT 'paper主键',
    `tag_id` int NOT NULL COMMENT '主键',
    PRIMARY KEY (`id`),
    index(`paper_id`, `tag_id`)
) COMMENT='文章和标签关联表';