-- 用户表
CREATE TABLE ttk_user_info
(
    id                  BIGINT(20) AUTO_INCREMENT COMMENT '用户ID (主键)',
    ttk_id              VARCHAR(16) UNIQUE NOT NULL COMMENT '用户名（唯一）',
    password            VARCHAR(255)       NOT NULL COMMENT '密码（加密存储）',
    salt                VARBINARY(16)      NOT NULL COMMENT '密码（加密存储）',
    nick_name           VARCHAR(50)        NOT NULL COMMENT '昵称（可修改,用于显示和@提及）',
    real_name           VARCHAR(20)        NOT NULL COMMENT '真实姓名（实名认证）',
    id_card             VARCHAR(18)        NOT NULL COMMENT '身份证ID （实名认证）',
    gender              TINYINT(1) NOT NULL DEFAULT 1 COMMENT '性别 1：未设置；2：男性；3：女性',
    birthdate           DATE               NOT NULL COMMENT '生日',
    avatar_path         VARCHAR(255)       NOT NULL COMMENT '头像路径（存储在对象存储中，如S3）',
    bio                 TEXT               NOT NULL COMMENT '个人简介',
    country             VARCHAR(255)       NOT NULL COMMENT '国家/地区',
    city                VARCHAR(255)       NOT NULL COMMENT '城市',
    email               VARCHAR(255)       NOT NULL COMMENT '邮箱',
    phone               VARCHAR(255)       NOT NULL COMMENT '手机号',
    account_status      TINYINT(1) NOT NULL DEFAULT 0 COMMENT '账号状态（0正常、1封禁、2限制）',
    registration_source VARCHAR(50)        NOT NULL COMMENT '注册来源（iOS App、Android App、Web等）',
    registration_ip     VARCHAR(32)        NOT NULL COMMENT '注册IP地址,ipv4 or ipv6',
    last_active         DATETIME(0) NOT NULL COMMENT '最近活跃时间',
    wallet_balance      DECIMAL(10, 2)     NOT NULL DEFAULT 0.00 COMMENT '钱包余额',
    created_at          DATETIME(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at          DATETIME(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at          DATETIME(0) NULL DEFAULT NULL COMMENT '删除时间',
    deleted_flag        TINYINT(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否删除 1：正常  2：已删除',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 第三方登录绑定表
CREATE TABLE ttk_third_party_binding
(
    id                       BIGINT(20) AUTO_INCREMENT COMMENT '第三方登录绑定ID',
    user_id                  BIGINT(20) NOT NULL COMMENT '用户ID',
    third_party_binding_type TINYINT(1) NOT NULL COMMENT '第三方登录类型,GitHub为0，Google为1，FaceBook为2',
    third_party_id           VARCHAR(100) NOT NULL COMMENT '第三方登录ID',
    created_at               DATETIME(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at               DATETIME(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at               DATETIME(0) NULL DEFAULT NULL COMMENT '删除时间',
    deleted_flag             TINYINT(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否删除 1：正常  2：已删除',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- 用户设置表
CREATE TABLE ttk_user_settings
(
    id                       BIGINT(20) AUTO_INCREMENT COMMENT '设置ID',
    user_id                  BIGINT(20) NOT NULL,
    notification_preferences JSON COMMENT '通知首选项（以JSON格式存储）',
    privacy_settings         JSON COMMENT '隐私设置（以JSON格式存储）',
    created_at               DATETIME(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at               DATETIME(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at               DATETIME(0) NULL DEFAULT NULL COMMENT '删除时间',
    deleted_flag             TINYINT(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否删除 1：正常  2：已删除',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 用户统计表
CREATE TABLE ttk_user_statistics
(
    id             BIGINT(20) AUTO_INCREMENT COMMENT '统计ID',
    user_id        BIGINT(20),
    posts_count    INT DEFAULT 0 COMMENT '发表的帖子数量',
    comments_count INT DEFAULT 0 COMMENT '发表的评论数量',
    likes_count    INT DEFAULT 0 COMMENT '获得的点赞数',
    created_at     DATETIME(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at     DATETIME(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at     DATETIME(0) NULL DEFAULT NULL COMMENT '删除时间',
    deleted_flag   TINYINT(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否删除 1：正常  2：已删除',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 用户地理位置表
CREATE TABLE ttk_user_locations
(
    id            BIGINT(20) AUTO_INCREMENT COMMENT '地理位置ID',
    user_id       BIGINT(20) NOT NULL,
    latitude      DECIMAL(10, 8) NOT NULL COMMENT '纬度',
    longitude     DECIMAL(11, 8) NOT NULL COMMENT '经度',
    location_name VARCHAR(255)   NOT NULL COMMENT '位置名称',
    created_at    DATETIME(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at    DATETIME(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at    DATETIME(0) NULL DEFAULT NULL COMMENT '删除时间',
    deleted_flag  TINYINT(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否删除 1：正常  2：已删除',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 用户反馈表
CREATE TABLE ttk_user_feedback
(
    id            BIGINT(20) AUTO_INCREMENT COMMENT '反馈ID',
    user_id       BIGINT(20) NOT NULL,
    feedback_text TEXT     NOT NULL COMMENT '反馈内容',
    created_at    DATETIME(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at    DATETIME(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at    DATETIME(0) NULL DEFAULT NULL COMMENT '删除时间',
    deleted_flag  TINYINT(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否删除 1：正常  2：已删除',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 用户日志表
CREATE TABLE ttk_user_logs
(
    id           BIGINT(20) AUTO_INCREMENT COMMENT '日志ID',
    user_id      BIGINT(20) NOT NULL,
    log_type     TINYINT(1) NOT NULL COMMENT '日志类型0登录，1操作，2活动，3其他',
    log_details  JSON     NOT NULL COMMENT '日志详情（以JSON格式存储）',
    timestamp    DATETIME NOT NULL COMMENT '时间戳',
    created_at   DATETIME(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at   DATETIME(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at   DATETIME(0) NULL DEFAULT NULL COMMENT '删除时间',
    deleted_flag TINYINT(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否删除 1：正常  2：已删除',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;