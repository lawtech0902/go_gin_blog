create database `blog` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

grant all privileges on `blog`.* to root@'%' identified by '12';

flush privileges;

use `blog`;

create table `blog_user`
(
    `id`           int(10) unsigned not null auto_increment,
    `username`     varchar(16)      not null,
    `password`     varchar(255)     not null,
    `avatar`       varchar(255) default null,
    `introduction` varchar(255) default null,
    `nickname`     varchar(32)  default null,
    `about`        text,
    primary key (`id`)
) engine = innodb
  auto_increment = 2
  default charset = utf8mb4;

insert into `blog_user` (`username`, `password`, `avatar`, `introduction`, `nickname`, `about`)
values ('admin', '123456',
        'https://ss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=2884107401,3797902000&fm=26&gp=0.jpg', 'a gopher',
        'gopher', '## hello world');

create table `blog_category`
(
    `id`            int(10) unsigned not null auto_increment,
    `category_name` varchar(16) default null,
    primary key (`id`)
) engine = innodb
  auto_increment = 4
  default charset = utf8mb4;

create table `blog_tag`
(
    `id`       int(10) unsigned not null auto_increment,
    `tag_name` varchar(16)      not null,
    primary key (`id`)
) engine = innodb
  auto_increment = 6
  default charset = utf8mb4;

create table `blog_article`
(
    `id`           int(10) unsigned not null auto_increment,
    `title`        varchar(32)      not null,
    `content`      text             not null,
    `html`         text             not null,
    `category_id`  int(10) unsigned not null,
    `created_time` varchar(32)      not null,
    `updated_time` varchar(32)               default '',
    `status`       varchar(16)      not null default 'published',
    primary key (`id`),
    unique key `title` (`title`) using btree,
    key `category` (`category_id`),
    constraint `category` foreign key (`category_id`) references `blog_category` (`id`) on delete cascade on update cascade
) engine = innodb
  auto_increment = 35
  default charset = utf8mb4;

create table `blog_tag_article`
(
    `id`         int(10) unsigned not null auto_increment,
    `tag_id`     int(10) unsigned not null,
    `article_id` int(10) unsigned not null,
    primary key (`id`),
    key `tag_id` (`tag_id`),
    key `article_id` (`article_id`),
    constraint `article_id` foreign key (`article_id`) references `blog_article` (`id`),
    constraint `tag_id` foreign key (`tag_id`) references `blog_tag` (`id`)
) engine = innodb
  auto_increment = 35
  default charset = utf8mb4;

create table `blog_soup`
(
    `id`      int(10) unsigned not null auto_increment,
    `content` tinytext         not null,
    primary key (`id`)
) engine = innodb
  auto_increment = 4
  default charset = utf8mb4;

insert into `blog_soup` (`content`)
values ('你全心全力做到的最好\r\n可能还不如别人的随便搞搞');

create table `blog_comment`
(
    `id`           int(10) unsigned not null auto_increment,
    `username`     varchar(16)      not null,
    `is_author`    tinyint(1)       not null default '0',
    `parent_id`    int(10) unsigned          default null,
    `root_id`      int(10) unsigned          default null,
    `article_id`   int(10) unsigned not null,
    `content`      varchar(255)     not null,
    `created_time` varchar(255)     not null,
    primary key (`id`),
    key `article` (`article_id`),
    key `parent` (`parent_id`),
    key `root` (`root_id`),
    constraint `root` foreign key (`root_id`) references `blog_comment` (`id`) on delete cascade on update cascade,
    constraint `article` foreign key (`article_id`) references `blog_article` (`id`) on delete cascade on update cascade,
    constraint `parent` foreign key (`parent_id`) references `blog_comment` (`id`) on delete cascade on update cascade
) engine = innodb
  auto_increment = 43
  default charset = utf8mb4;