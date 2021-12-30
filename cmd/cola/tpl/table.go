package tpl

const UpTableTemp = `CREATE TABLE {{.TableName}} (
    id BIGINT(20) UNSIGNED PRIMARY KEY,
    created_at BIGINT(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
    updated_at BIGINT(20) NOT NULL DEFAULT 0 COMMENT '更新时间'
)ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='' COLLATE=utf8mb4_unicode_ci;
`

const DownTableTemp = "DROP table `admins`;"
