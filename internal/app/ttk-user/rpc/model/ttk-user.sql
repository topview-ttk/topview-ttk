-- 用户表
CREATE TABLE ttk_user_info
(
    id                    BIGINT(20) AUTO_INCREMENT PRIMARY KEY COMMENT '用户ID (主键)',
    ttk_id                VARCHAR(255) UNIQUE NOT NULL COMMENT '用户名（唯一）',
    password              VARCHAR(255)        NOT NULL COMMENT '密码（加密存储）',
    nick_name             VARCHAR(255) COMMENT '昵称（可修改,用于显示和@提及）',
    real_name             VARCHAR(255) COMMENT '真实姓名（实名认证）',
    id_card               VARCHAR(255) COMMENT '身份证ID （实名认证）',
    gender                ENUM('Male', 'Female', 'Other') COMMENT '性别',
    birthdate             DATE COMMENT '生日',
    avatar_path           VARCHAR(255) COMMENT '头像路径（存储在对象存储中，如S3）',
    bio                   TEXT COMMENT '个人简介',
    country               VARCHAR(255) COMMENT '国家/地区',
    city                  VARCHAR(255) COMMENT '城市',
    followers             INT            DEFAULT 0 COMMENT '关注数',
    following             INT            DEFAULT 0 COMMENT '粉丝数',
    videos                INT            DEFAULT 0 COMMENT '视频数',
    private_account       BOOLEAN        DEFAULT 0 COMMENT '私密账号设置（公开、仅好友可见、私密等）',
    push_notifications    BOOLEAN        DEFAULT 1 COMMENT '推送通知设置（点赞、评论、关注等通知）',
    email                 VARCHAR(255) COMMENT '邮箱',
    phone                 VARCHAR(255) COMMENT '手机号',
    verification_status   ENUM('Verified', 'Unverified') COMMENT '验证状态（已验证、未验证）',
    account_status        ENUM('Normal', 'Banned', 'Restricted') COMMENT '账号状态（正常、封禁、限制等）',
    registration_source   VARCHAR(255) COMMENT '注册来源（iOS App、Android App、Web等）',
    registration_ip       VARCHAR(255) COMMENT '注册IP地址',
    last_active           DATETIME COMMENT '最近活跃时间',
    wallet_balance        DECIMAL(10, 2) DEFAULT 0.00 COMMENT '钱包余额',
    messaging_permission  ENUM('Public', 'Friends Only', 'Private') COMMENT '私信权限设置',
    tfa_enable            INT            DEFAULT 0 COMMENT '是否启用双因素认证',
    social_activity_score INT            DEFAULT 0 COMMENT '社交活跃度分析',
    created_at            DATETIME       DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at            DATETIME       DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at            DATETIME COMMENT '删除时间',
    PRIMARY KEY (`id`),
    CONSTRAINT chk_gender CHECK (gender IN ('Male', 'Female', 'Other')),
    CONSTRAINT chk_verification_status CHECK (verification_status IN ('Verified', 'Unverified')),
    CONSTRAINT chk_account_status CHECK (account_status IN ('Normal', 'Banned', 'Restricted'))
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 认证授权表
CREATE TABLE ttk_authorization
(
    id                 BIGINT(20) PRIMARY KEY COMMENT '认证授权ID',
    user_id            BIGINT(20),
    client_id          VARCHAR(255) COMMENT '客户端ID',
    authorization_code VARCHAR(255) COMMENT '授权码',
    access_token       VARCHAR(255) COMMENT '访问令牌',
    refresh_token      VARCHAR(255) COMMENT '刷新令牌',
    expires_at         DATETIME COMMENT '过期时间',
    created_at         DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at         DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at         DATETIME COMMENT '删除时间',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 第三方登录绑定表
CREATE TABLE ttk_third_party_binding
(
    id                       BIGINT(20) AUTO_INCREMENT PRIMARY KEY COMMENT '第三方登录绑定ID',
    user_id                  BIGINT(20),
    third_party_binding_type ENUM('Google', 'WeChat', 'QQ', 'Alipay', 'GitHub') COMMENT '第三方登录类型',
    third_party_id           VARCHAR(255) COMMENT '第三方登录ID',
    created_at               DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at               DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at               DATETIME COMMENT '删除时间',
    PRIMARY KEY (`id`),
    CONSTRAINT chk_third_party_binding_type CHECK (third_party_binding_type IN
                                                   ('Google', 'WeChat', 'QQ', 'Alipay', 'GitHub'))
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE ttk_login_status
(
    id                       BIGINT(20) AUTO_INCREMENT PRIMARY KEY COMMENT '登录状态ID (主键)',
    user_id                  BIGINT(20)          NOT NULL COMMENT '用户ID (外键关联到ttk_user表)',
    access_token             VARCHAR(255) NOT NULL COMMENT '访问令牌 (用于验证登录状态)',
    expires_at               DATETIME     NOT NULL COMMENT '过期时间 (登录状态的有效期)',
    login_method             ENUM('Email', 'Phone', 'ThirdParty') NOT NULL COMMENT '登录方式 (邮箱、手机号、第三方登录)',
    third_party_type         ENUM('Google', 'WeChat', 'QQ', 'Alipay', 'GitHub', 'Origin') COMMENT '第三方登录类型 (仅当登录方式为ThirdParty时使用)',
    third_party_access_token VARCHAR(255) COMMENT '第三方登录访问令牌 (仅当登录方式为ThirdParty时使用)',
    created_at               DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at               DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at               DATETIME COMMENT '删除时间',
    PRIMARY KEY (`id`),
    CONSTRAINT chk_third_party_type CHECK (third_party_type IN ('Google', 'WeChat', 'QQ', 'Alipay', 'GitHub', 'Origin'))
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- 用户设置表
CREATE TABLE ttk_user_settings
(
    id                       BIGINT(20) AUTO_INCREMENT PRIMARY KEY COMMENT '设置ID',
    user_id                  BIGINT(20),
    notification_preferences JSON COMMENT '通知首选项（以JSON格式存储）',
    privacy_settings         JSON COMMENT '隐私设置（以JSON格式存储）',
    created_at               DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at               DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at               DATETIME COMMENT '删除时间',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 用户统计表
CREATE TABLE ttk_user_statistics
(
    id             BIGINT(20) AUTO_INCREMENT PRIMARY KEY COMMENT '统计ID',
    user_id        BIGINT(20),
    posts_count    INT      DEFAULT 0 COMMENT '发表的帖子数量',
    comments_count INT      DEFAULT 0 COMMENT '发表的评论数量',
    likes_count    INT      DEFAULT 0 COMMENT '获得的点赞数',
    created_at     DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at     DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at     DATETIME COMMENT '删除时间',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 用户地理位置表
CREATE TABLE ttk_user_locations
(
    id            BIGINT(20) AUTO_INCREMENT PRIMARY KEY COMMENT '地理位置ID',
    user_id       BIGINT(20),
    latitude      DECIMAL(10, 8) COMMENT '纬度',
    longitude     DECIMAL(11, 8) COMMENT '经度',
    location_name VARCHAR(255) COMMENT '位置名称',
    created_at    DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at    DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at    DATETIME COMMENT '删除时间',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 用户反馈表
CREATE TABLE ttk_user_feedback
(
    id            BIGINT(20) AUTO_INCREMENT PRIMARY KEY COMMENT '反馈ID',
    user_id       BIGINT(20),
    feedback_text TEXT COMMENT '反馈内容',
    timestamp     DATETIME COMMENT '时间戳',
    created_at    DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at    DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at    DATETIME COMMENT '删除时间',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 用户日志表
CREATE TABLE ttk_user_logs
(
    id          BIGINT(20) AUTO_INCREMENT PRIMARY KEY COMMENT '日志ID',
    user_id     BIGINT(20),
    log_type    ENUM('Login', 'Operate', 'Activity', 'Other') COMMENT '日志类型',
    log_details JSON COMMENT '日志详情（以JSON格式存储）',
    timestamp   DATETIME COMMENT '时间戳',
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at  DATETIME COMMENT '删除时间',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;