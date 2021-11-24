--
-- 由SQLiteStudio v3.2.1 产生的文件 周二 11月 23 15:22:53 2021
--
-- 文本编码：UTF-8
--
PRAGMA foreign_keys = off;
BEGIN TRANSACTION;

-- 表：biz_backing_browser_r
DROP TABLE IF EXISTS biz_backing_browser_r;

CREATE TABLE biz_backing_browser_r (
    vm_backing_id INTEGER,
    browser_id    INTEGER,
    PRIMARY KEY (
        vm_backing_id,
        browser_id
    ),
    CONSTRAINT fk_biz_backing_browser_r_vm_backing FOREIGN KEY (
        vm_backing_id
    )
    REFERENCES biz_vm_baking (id),
    CONSTRAINT fk_biz_backing_browser_r_browser FOREIGN KEY (
        browser_id
    )
    REFERENCES biz_browser (id) 
);

INSERT INTO biz_backing_browser_r (
                                      vm_backing_id,
                                      browser_id
                                  )
                                  VALUES (
                                      1,
                                      1
                                  );

INSERT INTO biz_backing_browser_r (
                                      vm_backing_id,
                                      browser_id
                                  )
                                  VALUES (
                                      2,
                                      2
                                  );


-- 表：biz_browser
DROP TABLE IF EXISTS biz_browser;

CREATE TABLE biz_browser (
    id         INTEGER,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    deleted    NUMERIC  DEFAULT false,
    disabled   NUMERIC  DEFAULT false,
    name       TEXT,
    type       TEXT,
    version    TEXT,
    lang       TEXT,
    PRIMARY KEY (
        id
    )
);

INSERT INTO biz_browser (
                            id,
                            created_at,
                            updated_at,
                            deleted_at,
                            deleted,
                            disabled,
                            name,
                            type,
                            version,
                            lang
                        )
                        VALUES (
                            1,
                            NULL,
                            NULL,
                            NULL,
                            0,
                            0,
                            'chrome-91',
                            'chrome',
                            '91',
                            NULL
                        );

INSERT INTO biz_browser (
                            id,
                            created_at,
                            updated_at,
                            deleted_at,
                            deleted,
                            disabled,
                            name,
                            type,
                            version,
                            lang
                        )
                        VALUES (
                            2,
                            NULL,
                            NULL,
                            NULL,
                            0,
                            0,
                            'firefox-85',
                            'firefox',
                            '85',
                            NULL
                        );


-- 表：biz_build
DROP TABLE IF EXISTS biz_build;

CREATE TABLE biz_build (
    id                INTEGER,
    created_at        DATETIME,
    updated_at        DATETIME,
    deleted_at        DATETIME,
    deleted           NUMERIC  DEFAULT false,
    disabled          NUMERIC  DEFAULT false,
    queue_id          INTEGER,
    priority          INTEGER,
    serial            TEXT,
    vm_id             INTEGER,
    build_type        TEXT,
    os_platform       TEXT,
    os_type           TEXT,
    os_lang           TEXT,
    browser_type      TEXT,
    browser_version   TEXT,
    script_url        TEXT,
    scm_address       TEXT,
    scm_account       TEXT,
    scm_password      TEXT,
    app_url           TEXT,
    build_commands    TEXT,
    env_vars          TEXT,
    result_files      TEXT,
    keep_result_files NUMERIC,
    progress          TEXT,
    status            TEXT,
    retry             INTEGER  DEFAULT 0,
    task_name         TEXT,
    user_name         TEXT,
    start_time        DATETIME,
    pending_time      DATETIME,
    result_time       DATETIME,
    timeout_time      DATETIME,
    task_id           INTEGER,
    group_id          INTEGER,
    node_ip           TEXT,
    node_port         INTEGER,
    appium_port       INTEGER,
    complete_time     DATETIME,
    result_msg        TEXT,
    result_path       TEXT,
    end_time          DATETIME,
    PRIMARY KEY (
        id
    )
);


-- 表：biz_casbin_rules
DROP TABLE IF EXISTS biz_casbin_rules;

CREATE TABLE biz_casbin_rules (
    id     INTEGER,
    p_type TEXT,
    v0     TEXT,
    v1     TEXT,
    v2     TEXT,
    v3     TEXT,
    v4     TEXT,
    v5     TEXT,
    PRIMARY KEY (
        id
    )
);

INSERT INTO biz_casbin_rules (
                                 id,
                                 p_type,
                                 v0,
                                 v1,
                                 v2,
                                 v3,
                                 v4,
                                 v5
                             )
                             VALUES (
                                 1,
                                 'p',
                                 '1',
                                 '/api/v1/admin/logout',
                                 'POST',
                                 '',
                                 '',
                                 ''
                             );

INSERT INTO biz_casbin_rules (
                                 id,
                                 p_type,
                                 v0,
                                 v1,
                                 v2,
                                 v3,
                                 v4,
                                 v5
                             )
                             VALUES (
                                 2,
                                 'p',
                                 '1',
                                 '/api/v1/platform/view',
                                 'GET',
                                 '',
                                 '',
                                 ''
                             );

INSERT INTO biz_casbin_rules (
                                 id,
                                 p_type,
                                 v0,
                                 v1,
                                 v2,
                                 v3,
                                 v4,
                                 v5
                             )
                             VALUES (
                                 3,
                                 'p',
                                 '1',
                                 '/api/v1/test/tasks',
                                 'GET',
                                 '',
                                 '',
                                 ''
                             );

INSERT INTO biz_casbin_rules (
                                 id,
                                 p_type,
                                 v0,
                                 v1,
                                 v2,
                                 v3,
                                 v4,
                                 v5
                             )
                             VALUES (
                                 4,
                                 'p',
                                 '1',
                                 '/api/v1/test/tasks/{id:uint}',
                                 'GET',
                                 '',
                                 '',
                                 ''
                             );

INSERT INTO biz_casbin_rules (
                                 id,
                                 p_type,
                                 v0,
                                 v1,
                                 v2,
                                 v3,
                                 v4,
                                 v5
                             )
                             VALUES (
                                 5,
                                 'p',
                                 '1',
                                 '/api/v1/test/tasks',
                                 'POST',
                                 '',
                                 '',
                                 ''
                             );

INSERT INTO biz_casbin_rules (
                                 id,
                                 p_type,
                                 v0,
                                 v1,
                                 v2,
                                 v3,
                                 v4,
                                 v5
                             )
                             VALUES (
                                 6,
                                 'p',
                                 '1',
                                 '/api/v1/test/tasks/{id:uint}',
                                 'PUT',
                                 '',
                                 '',
                                 ''
                             );

INSERT INTO biz_casbin_rules (
                                 id,
                                 p_type,
                                 v0,
                                 v1,
                                 v2,
                                 v3,
                                 v4,
                                 v5
                             )
                             VALUES (
                                 7,
                                 'p',
                                 '1',
                                 '/api/v1/test/tasks/{id:uint}',
                                 'DELETE',
                                 '',
                                 '',
                                 ''
                             );

INSERT INTO biz_casbin_rules (
                                 id,
                                 p_type,
                                 v0,
                                 v1,
                                 v2,
                                 v3,
                                 v4,
                                 v5
                             )
                             VALUES (
                                 8,
                                 'p',
                                 '1',
                                 '/api/v1/test/envs',
                                 'POST',
                                 '',
                                 '',
                                 ''
                             );


-- 表：biz_container
DROP TABLE IF EXISTS biz_container;

CREATE TABLE biz_container (
    id          INTEGER,
    created_at  DATETIME,
    updated_at  DATETIME,
    deleted_at  DATETIME,
    deleted     NUMERIC  DEFAULT false,
    disabled    NUMERIC  DEFAULT false,
    ident       TEXT,
    host_id     INTEGER,
    os_version  TEXT,
    image_name  TEXT,
    [desc]      TEXT,
    os_lang     TEXT,
    node_ip     TEXT,
    os_category TEXT,
    node_port   INTEGER,
    name        TEXT,
    host_name   TEXT,
    os_type     TEXT,
    PRIMARY KEY (
        id
    )
);


-- 表：biz_environment
DROP TABLE IF EXISTS biz_environment;

CREATE TABLE biz_environment (
    id          INTEGER,
    created_at  DATETIME,
    updated_at  DATETIME,
    deleted_at  DATETIME,
    deleted     NUMERIC  DEFAULT false,
    disabled    NUMERIC  DEFAULT false,
    os_category TEXT,
    os_type     TEXT,
    os_version  TEXT,
    os_lang     TEXT,
    task_id     INTEGER,
    image_src   TEXT,
    image_name  TEXT,
    PRIMARY KEY (
        id
    ),
    CONSTRAINT fk_biz_task_environments FOREIGN KEY (
        task_id
    )
    REFERENCES biz_task (id) 
);


-- 表：biz_history
DROP TABLE IF EXISTS biz_history;

CREATE TABLE biz_history (
    id         INTEGER,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    deleted    NUMERIC  DEFAULT false,
    disabled   NUMERIC  DEFAULT false,
    queue_id   INTEGER,
    type       TEXT,
    progress   TEXT,
    status     TEXT,
    owner_id   INTEGER,
    owner_type TEXT,
    PRIMARY KEY (
        id
    )
);


-- 表：biz_host
DROP TABLE IF EXISTS biz_host;

CREATE TABLE biz_host (
    id                 INTEGER,
    name               TEXT,
    os_type            TEXT,
    os_lang            TEXT,
    os_version         TEXT,
    os_build           TEXT,
    os_bits            TEXT,
    ip                 TEXT,
    port               INTEGER,
    work_dir           TEXT,
    ssh_port           INTEGER,
    vnc_port           INTEGER,
    status             TEXT,
    last_register_date DATETIME,
    os_category        TEXT,
    task_count         INTEGER,
    max_vm_num         INTEGER,
    last_register_time DATETIME,
    priority           INTEGER,
    vnc_address        TEXT,
    platform           TEXT,
    cloud_key          TEXT,
    cloud_region       TEXT,
    cloud_namespace    TEXT,
    cloud_secret       TEXT,
    cloud_user         TEXT,
    cloud_iam_user     TEXT,
    cloud_iam_password TEXT,
    cloud_iam_key      TEXT,
    vpc_id             TEXT,
    bridge             TEXT,
    created_at         DATETIME,
    updated_at         DATETIME,
    deleted_at         DATETIME,
    deleted            NUMERIC  DEFAULT false,
    disabled           NUMERIC  DEFAULT false,
    PRIMARY KEY (
        id
    )
);

INSERT INTO biz_host (
                         id,
                         name,
                         os_type,
                         os_lang,
                         os_version,
                         os_build,
                         os_bits,
                         ip,
                         port,
                         work_dir,
                         ssh_port,
                         vnc_port,
                         status,
                         last_register_date,
                         os_category,
                         task_count,
                         max_vm_num,
                         last_register_time,
                         priority,
                         vnc_address,
                         platform,
                         cloud_key,
                         cloud_region,
                         cloud_namespace,
                         cloud_secret,
                         cloud_user,
                         cloud_iam_user,
                         cloud_iam_password,
                         cloud_iam_key,
                         vpc_id,
                         bridge,
                         created_at,
                         updated_at,
                         deleted_at,
                         deleted,
                         disabled
                     )
                     VALUES (
                         1,
                         '192.168.0.56',
                         '',
                         '',
                         NULL,
                         NULL,
                         NULL,
                         '192.168.0.56',
                         8086,
                         NULL,
                         NULL,
                         NULL,
                         'online',
                         NULL,
                         'linux',
                         NULL,
                         3,
                         NULL,
                         400,
                         NULL,
                         'vm,docker,native',
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         '2021-07-29 21:10:09.90093+08:00',
                         '2021-10-12 17:20:23.755882+08:00',
                         NULL,
                         0,
                         0
                     );

INSERT INTO biz_host (
                         id,
                         name,
                         os_type,
                         os_lang,
                         os_version,
                         os_build,
                         os_bits,
                         ip,
                         port,
                         work_dir,
                         ssh_port,
                         vnc_port,
                         status,
                         last_register_date,
                         os_category,
                         task_count,
                         max_vm_num,
                         last_register_time,
                         priority,
                         vnc_address,
                         platform,
                         cloud_key,
                         cloud_region,
                         cloud_namespace,
                         cloud_secret,
                         cloud_user,
                         cloud_iam_user,
                         cloud_iam_password,
                         cloud_iam_key,
                         vpc_id,
                         bridge,
                         created_at,
                         updated_at,
                         deleted_at,
                         deleted,
                         disabled
                     )
                     VALUES (
                         2,
                         '华为云ECS',
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         'online',
                         NULL,
                         NULL,
                         NULL,
                         10,
                         NULL,
                         300,
                         NULL,
                         'vm,cloud,huawei',
                         'TVZFYALVUODSL6S8FTNX',
                         'cn-east-3',
                         NULL,
                         'rF81YZEE9slggcqDBO5aBHtIh63mWi2g8saeUIe5',
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         '2021-07-29 21:10:09.90093+08:00',
                         '2021-08-9 21:10:09.90093+08:00',
                         NULL,
                         0,
                         0
                     );

INSERT INTO biz_host (
                         id,
                         name,
                         os_type,
                         os_lang,
                         os_version,
                         os_build,
                         os_bits,
                         ip,
                         port,
                         work_dir,
                         ssh_port,
                         vnc_port,
                         status,
                         last_register_date,
                         os_category,
                         task_count,
                         max_vm_num,
                         last_register_time,
                         priority,
                         vnc_address,
                         platform,
                         cloud_key,
                         cloud_region,
                         cloud_namespace,
                         cloud_secret,
                         cloud_user,
                         cloud_iam_user,
                         cloud_iam_password,
                         cloud_iam_key,
                         vpc_id,
                         bridge,
                         created_at,
                         updated_at,
                         deleted_at,
                         deleted,
                         disabled
                     )
                     VALUES (
                         3,
                         '华为云CCI',
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         'online',
                         NULL,
                         NULL,
                         NULL,
                         10,
                         NULL,
                         300,
                         NULL,
                         'docker,cloud,huawei',
                         'TVZFYALVUODSL6S8FTNX',
                         'cn-east-3',
                         'cci-com-deeptest',
                         'rF81YZEE9slggcqDBO5aBHtIh63mWi2g8saeUIe5',
                         'aaronchen2k',
                         'aaronchen2k-iam',
                         'z%T!cG"q!vpCOJj!m4hI',
                         NULL,
                         NULL,
                         NULL,
                         '2021-07-29 21:10:09.90093+08:00',
                         '2021-07-29 21:10:09.90093+08:00',
                         NULL,
                         0,
                         0
                     );

INSERT INTO biz_host (
                         id,
                         name,
                         os_type,
                         os_lang,
                         os_version,
                         os_build,
                         os_bits,
                         ip,
                         port,
                         work_dir,
                         ssh_port,
                         vnc_port,
                         status,
                         last_register_date,
                         os_category,
                         task_count,
                         max_vm_num,
                         last_register_time,
                         priority,
                         vnc_address,
                         platform,
                         cloud_key,
                         cloud_region,
                         cloud_namespace,
                         cloud_secret,
                         cloud_user,
                         cloud_iam_user,
                         cloud_iam_password,
                         cloud_iam_key,
                         vpc_id,
                         bridge,
                         created_at,
                         updated_at,
                         deleted_at,
                         deleted,
                         disabled
                     )
                     VALUES (
                         4,
                         '阿里云ECS',
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         'online',
                         NULL,
                         NULL,
                         NULL,
                         10,
                         NULL,
                         300,
                         NULL,
                         'vm,cloud,ali',
                         'LTAI5t8sfNV3M5a7VWJfzGZ2',
                         'cn-hangzhou',
                         NULL,
                         'Y8znXiZ5ZANjoyDjztR0bmm2xSfOfD',
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         'vpc-bp1p8nzox2o3kvfh4o0wo',
                         NULL,
                         '2021-07-29 21:10:09.90093+08:00',
                         '2021-07-29 21:10:09.90093+08:00',
                         NULL,
                         0,
                         0
                     );

INSERT INTO biz_host (
                         id,
                         name,
                         os_type,
                         os_lang,
                         os_version,
                         os_build,
                         os_bits,
                         ip,
                         port,
                         work_dir,
                         ssh_port,
                         vnc_port,
                         status,
                         last_register_date,
                         os_category,
                         task_count,
                         max_vm_num,
                         last_register_time,
                         priority,
                         vnc_address,
                         platform,
                         cloud_key,
                         cloud_region,
                         cloud_namespace,
                         cloud_secret,
                         cloud_user,
                         cloud_iam_user,
                         cloud_iam_password,
                         cloud_iam_key,
                         vpc_id,
                         bridge,
                         created_at,
                         updated_at,
                         deleted_at,
                         deleted,
                         disabled
                     )
                     VALUES (
                         5,
                         '阿里云ECI',
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         'online',
                         NULL,
                         NULL,
                         NULL,
                         10,
                         NULL,
                         200,
                         NULL,
                         'docker,cloud,ali',
                         'LTAI5t8sfNV3M5a7VWJfzGZ2',
                         'cn-hangzhou',
                         NULL,
                         'Y8znXiZ5ZANjoyDjztR0bmm2xSfOfD',
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         'vpc-bp1p8nzox2o3kvfh4o0wo',
                         NULL,
                         '2021-07-29 21:10:09.90093+08:00',
                         '2021-07-29 21:10:09.90093+08:00',
                         NULL,
                         0,
                         0
                     );

INSERT INTO biz_host (
                         id,
                         name,
                         os_type,
                         os_lang,
                         os_version,
                         os_build,
                         os_bits,
                         ip,
                         port,
                         work_dir,
                         ssh_port,
                         vnc_port,
                         status,
                         last_register_date,
                         os_category,
                         task_count,
                         max_vm_num,
                         last_register_time,
                         priority,
                         vnc_address,
                         platform,
                         cloud_key,
                         cloud_region,
                         cloud_namespace,
                         cloud_secret,
                         cloud_user,
                         cloud_iam_user,
                         cloud_iam_password,
                         cloud_iam_key,
                         vpc_id,
                         bridge,
                         created_at,
                         updated_at,
                         deleted_at,
                         deleted,
                         disabled
                     )
                     VALUES (
                         6,
                         '192.168.0.56',
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         '192.168.0.56',
                         8086,
                         NULL,
                         NULL,
                         NULL,
                         'online',
                         NULL,
                         NULL,
                         NULL,
                         10,
                         NULL,
                         100,
                         NULL,
                         'vm,cloud,virtualbox',
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         '',
                         'aaron',
                         'P2ssw0rd',
                         NULL,
                         NULL,
                         'br0',
                         '2021-07-29 21:10:09.90093+08:00',
                         '2021-10-12 17:20:23.755882+08:00',
                         NULL,
                         0,
                         0
                     );

INSERT INTO biz_host (
                         id,
                         name,
                         os_type,
                         os_lang,
                         os_version,
                         os_build,
                         os_bits,
                         ip,
                         port,
                         work_dir,
                         ssh_port,
                         vnc_port,
                         status,
                         last_register_date,
                         os_category,
                         task_count,
                         max_vm_num,
                         last_register_time,
                         priority,
                         vnc_address,
                         platform,
                         cloud_key,
                         cloud_region,
                         cloud_namespace,
                         cloud_secret,
                         cloud_user,
                         cloud_iam_user,
                         cloud_iam_password,
                         cloud_iam_key,
                         vpc_id,
                         bridge,
                         created_at,
                         updated_at,
                         deleted_at,
                         deleted,
                         disabled
                     )
                     VALUES (
                         7,
                         '127.0.0.1',
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         '127.0.0.1',
                         8086,
                         NULL,
                         NULL,
                         NULL,
                         'online',
                         NULL,
                         NULL,
                         NULL,
                         10,
                         NULL,
                         50,
                         NULL,
                         'vm,docker,native',
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         '2021-07-29 21:10:09.90093+08:00',
                         '2021-09-10 14:27:06.003862+08:00',
                         NULL,
                         0,
                         0
                     );

INSERT INTO biz_host (
                         id,
                         name,
                         os_type,
                         os_lang,
                         os_version,
                         os_build,
                         os_bits,
                         ip,
                         port,
                         work_dir,
                         ssh_port,
                         vnc_port,
                         status,
                         last_register_date,
                         os_category,
                         task_count,
                         max_vm_num,
                         last_register_time,
                         priority,
                         vnc_address,
                         platform,
                         cloud_key,
                         cloud_region,
                         cloud_namespace,
                         cloud_secret,
                         cloud_user,
                         cloud_iam_user,
                         cloud_iam_password,
                         cloud_iam_key,
                         vpc_id,
                         bridge,
                         created_at,
                         updated_at,
                         deleted_at,
                         deleted,
                         disabled
                     )
                     VALUES (
                         8,
                         '192.168.0.56',
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         '192.168.0.56',
                         8086,
                         NULL,
                         NULL,
                         NULL,
                         'online',
                         NULL,
                         NULL,
                         NULL,
                         3,
                         NULL,
                         30,
                         NULL,
                         'vm,cloud,vmware',
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         NULL,
                         'aaron',
                         'P2ssw0rd@',
                         NULL,
                         NULL,
                         'br0',
                         '2021-07-29 21:10:09.90093+08:00',
                         '2021-10-12 17:20:23.755882+08:00',
                         NULL,
                         0,
                         0
                     );


-- 表：biz_host_backing_r
DROP TABLE IF EXISTS biz_host_backing_r;

CREATE TABLE biz_host_backing_r (
    host_id       INTEGER,
    vm_backing_id INTEGER,
    PRIMARY KEY (
        host_id,
        vm_backing_id
    ),
    CONSTRAINT fk_biz_host_backing_r_host FOREIGN KEY (
        host_id
    )
    REFERENCES biz_host (id),
    CONSTRAINT fk_biz_host_backing_r_vm_backing FOREIGN KEY (
        vm_backing_id
    )
    REFERENCES biz_vm_baking (id) 
);

INSERT INTO biz_host_backing_r (
                                   host_id,
                                   vm_backing_id
                               )
                               VALUES (
                                   1,
                                   1
                               );

INSERT INTO biz_host_backing_r (
                                   host_id,
                                   vm_backing_id
                               )
                               VALUES (
                                   2,
                                   1
                               );

INSERT INTO biz_host_backing_r (
                                   host_id,
                                   vm_backing_id
                               )
                               VALUES (
                                   4,
                                   2
                               );

INSERT INTO biz_host_backing_r (
                                   host_id,
                                   vm_backing_id
                               )
                               VALUES (
                                   6,
                                   1
                               );

INSERT INTO biz_host_backing_r (
                                   host_id,
                                   vm_backing_id
                               )
                               VALUES (
                                   8,
                                   1
                               );


-- 表：biz_iso
DROP TABLE IF EXISTS biz_iso;

CREATE TABLE biz_iso (
    id                  INTEGER,
    created_at          DATETIME,
    updated_at          DATETIME,
    deleted_at          DATETIME,
    deleted             NUMERIC  DEFAULT false,
    disabled            NUMERIC  DEFAULT false,
    name                TEXT,
    path                TEXT,
    size                INTEGER,
    os_platform         TEXT,
    os_type             TEXT,
    os_lang             TEXT,
    os_version          TEXT,
    os_build            TEXT,
    os_bits             TEXT,
    resolution_height   INTEGER,
    resolution_width    INTEGER,
    suggest_memory_size INTEGER,
    suggest_disk_size   INTEGER,
    PRIMARY KEY (
        id
    )
);


-- 表：biz_project
DROP TABLE IF EXISTS biz_project;

CREATE TABLE biz_project (
    id           INTEGER,
    created_at   DATETIME,
    updated_at   DATETIME,
    deleted_at   DATETIME,
    deleted      NUMERIC  DEFAULT false,
    disabled     NUMERIC  DEFAULT false,
    name         TEXT,
    [desc]       TEXT,
    path         TEXT,
    is_default   NUMERIC,
    service_port INTEGER,
    user_id      INTEGER,
    PRIMARY KEY (
        id
    )
);


-- 表：biz_queue
DROP TABLE IF EXISTS biz_queue;

CREATE TABLE biz_queue (
    id                INTEGER,
    created_at        DATETIME,
    updated_at        DATETIME,
    deleted_at        DATETIME,
    deleted           NUMERIC  DEFAULT false,
    disabled          NUMERIC  DEFAULT false,
    queue_id          INTEGER,
    priority          INTEGER,
    serial            TEXT,
    vm_id             INTEGER,
    build_type        TEXT,
    os_category       TEXT,
    os_type           TEXT,
    os_lang           TEXT,
    browser_type      TEXT,
    browser_version   TEXT,
    script_url        TEXT,
    scm_address       TEXT,
    scm_account       TEXT,
    scm_password      TEXT,
    app_url           TEXT,
    env_vars          TEXT,
    build_commands    TEXT,
    result_files      TEXT,
    keep_result_files NUMERIC,
    progress          TEXT,
    status            TEXT,
    retry             INTEGER,
    task_name         TEXT,
    user_name         TEXT,
    start_time        DATETIME,
    pending_time      DATETIME,
    result_time       DATETIME,
    timeout_time      DATETIME,
    task_id           INTEGER,
    group_id          INTEGER,
    node_ip           TEXT,
    node_port         INTEGER,
    appium_port       INTEGER,
    complete_time     DATETIME,
    res_launched_time DATETIME,
    run_time          DATETIME,
    res_pending_time  DATETIME,
    docker_image      TEXT,
    PRIMARY KEY (
        id
    )
);


-- 表：biz_task
DROP TABLE IF EXISTS biz_task;

CREATE TABLE biz_task (
    id                INTEGER,
    created_at        DATETIME,
    updated_at        DATETIME,
    deleted_at        DATETIME,
    deleted           NUMERIC  DEFAULT false,
    disabled          NUMERIC  DEFAULT false,
    name              TEXT,
    priority          INTEGER,
    serials           TEXT,
    environments      TEXT,
    build_type        TEXT,
    script_url        TEXT,
    scm_address       TEXT,
    scm_account       TEXT,
    scm_password      TEXT,
    app_url           TEXT,
    build_commands    TEXT,
    result_files      TEXT,
    keep_result_files NUMERIC,
    progress          TEXT,
    status            TEXT,
    start_time        DATETIME,
    pending_time      DATETIME,
    result_time       DATETIME,
    user_name         TEXT,
    user_id           INTEGER,
    group_id          INTEGER,
    env_vars          TEXT,
    [desc]            TEXT,
    browser_version   TEXT,
    browser_type      TEXT,
    end_time          DATETIME,
    is_docker_native  NUMERIC,
    PRIMARY KEY (
        id
    )
);


-- 表：biz_vm
DROP TABLE IF EXISTS biz_vm;

CREATE TABLE biz_vm (
    id                  INTEGER,
    created_at          DATETIME,
    updated_at          DATETIME,
    deleted_at          DATETIME,
    deleted             NUMERIC  DEFAULT false,
    disabled            NUMERIC  DEFAULT false,
    base_id             INTEGER,
    host_id             INTEGER,
    name                TEXT,
    src                 TEXT,
    base                TEXT,
    image_path          TEXT,
    base_path           TEXT,
    os_category         TEXT,
    os_type             TEXT,
    os_version          TEXT,
    os_lang             TEXT,
    status              TEXT,
    destroy_at          DATETIME,
    first_detected_time DATETIME,
    node_ip             TEXT,
    node_port           INTEGER,
    mac_address         TEXT,
    rpc_port            INTEGER,
    ssh_port            INTEGER,
    vnc_port            INTEGER,
    work_dir            TEXT,
    def_path            TEXT,
    disk_size           INTEGER,
    memory_size         INTEGER,
    cdrom_sys           TEXT,
    cdrom_driver        TEXT,
    resolution_height   INTEGER,
    resolution_width    INTEGER,
    tmpl                TEXT,
    backing             TEXT,
    backing_path        TEXT,
    backing_id          INTEGER,
    [desc]              TEXT,
    host_name           TEXT,
    tmpl_id             INTEGER,
    tmpl_name           TEXT,
    last_register_time  DATETIME,
    vnc_address         TEXT,
    cloud_inst_id       TEXT,
    cloud_eip_id        TEXT,
    PRIMARY KEY (
        id
    )
);


-- 表：biz_vm_baking
DROP TABLE IF EXISTS biz_vm_baking;

CREATE TABLE biz_vm_baking (
    id                  INTEGER,
    created_at          DATETIME,
    updated_at          DATETIME,
    deleted_at          DATETIME,
    deleted             NUMERIC  DEFAULT false,
    disabled            NUMERIC  DEFAULT false,
    name                TEXT,
    path                TEXT,
    size                INTEGER,
    os_category         TEXT,
    os_type             TEXT,
    os_lang             TEXT,
    os_version          TEXT,
    os_build            TEXT,
    os_bits             TEXT,
    resolution_height   INTEGER,
    resolution_width    INTEGER,
    suggest_cpu_count   INTEGER,
    suggest_memory_size INTEGER,
    suggest_disk_size   INTEGER,
    sys_iso_id          INTEGER,
    driver_iso_id       INTEGER,
    PRIMARY KEY (
        id
    )
);

INSERT INTO biz_vm_baking (
                              id,
                              created_at,
                              updated_at,
                              deleted_at,
                              deleted,
                              disabled,
                              name,
                              path,
                              size,
                              os_category,
                              os_type,
                              os_lang,
                              os_version,
                              os_build,
                              os_bits,
                              resolution_height,
                              resolution_width,
                              suggest_cpu_count,
                              suggest_memory_size,
                              suggest_disk_size,
                              sys_iso_id,
                              driver_iso_id
                          )
                          VALUES (
                              1,
                              NULL,
                              NULL,
                              NULL,
                              0,
                              0,
                              'win10-pro-x64-zh_cn',
                              'backing/windows/win10/pro-x64-zh_cn.qcow2',
                              NULL,
                              'windows',
                              'win10',
                              'zh_cn',
                              'pro-x64',
                              NULL,
                              NULL,
                              NULL,
                              NULL,
                              2,
                              4000,
                              40000,
                              NULL,
                              NULL
                          );

INSERT INTO biz_vm_baking (
                              id,
                              created_at,
                              updated_at,
                              deleted_at,
                              deleted,
                              disabled,
                              name,
                              path,
                              size,
                              os_category,
                              os_type,
                              os_lang,
                              os_version,
                              os_build,
                              os_bits,
                              resolution_height,
                              resolution_width,
                              suggest_cpu_count,
                              suggest_memory_size,
                              suggest_disk_size,
                              sys_iso_id,
                              driver_iso_id
                          )
                          VALUES (
                              2,
                              NULL,
                              NULL,
                              NULL,
                              0,
                              0,
                              'ubuntu-20-desktop-x64-zh_cn',
                              'backing/linux/ubuntu/20-desktop-x64-zh_cn.qcow2',
                              NULL,
                              'linux',
                              'ubuntu',
                              'zh_cn',
                              '20-desktop-x64',
                              NULL,
                              NULL,
                              NULL,
                              NULL,
                              2,
                              4000,
                              30000,
                              NULL,
                              NULL
                          );


-- 表：biz_vm_tmpl
DROP TABLE IF EXISTS biz_vm_tmpl;

CREATE TABLE biz_vm_tmpl (
    id          INTEGER,
    created_at  DATETIME,
    updated_at  DATETIME,
    deleted_at  DATETIME,
    deleted     NUMERIC  DEFAULT false,
    disabled    NUMERIC  DEFAULT false,
    host_id     INTEGER,
    name        TEXT,
    os_category TEXT,
    os_type     TEXT,
    os_version  TEXT,
    os_lang     TEXT,
    status      TEXT,
    PRIMARY KEY (
        id
    )
);

INSERT INTO biz_vm_tmpl (
                            id,
                            created_at,
                            updated_at,
                            deleted_at,
                            deleted,
                            disabled,
                            host_id,
                            name,
                            os_category,
                            os_type,
                            os_version,
                            os_lang,
                            status
                        )
                        VALUES (
                            1,
                            NULL,
                            NULL,
                            NULL,
                            0,
                            0,
                            1,
                            'win10-x64-pro-zh_cn',
                            'windows',
                            'win10',
                            'x64-pro',
                            'zh_cn',
                            'created'
                        );


-- 表：sys_permission
DROP TABLE IF EXISTS sys_permission;

CREATE TABLE sys_permission (
    id           INTEGER,
    created_at   DATETIME,
    updated_at   DATETIME,
    deleted_at   DATETIME,
    deleted      NUMERIC       DEFAULT false,
    disabled     NUMERIC       DEFAULT false,
    name         VARCHAR (256) NOT NULL,
    display_name VARCHAR (256),
    description  VARCHAR (256),
    act          VARCHAR (256),
    PRIMARY KEY (
        id
    )
);

INSERT INTO sys_permission (
                               id,
                               created_at,
                               updated_at,
                               deleted_at,
                               deleted,
                               disabled,
                               name,
                               display_name,
                               description,
                               act
                           )
                           VALUES (
                               1,
                               '2021-06-29 17:31:36.705549+08:00',
                               '2021-06-29 17:31:36.705549+08:00',
                               NULL,
                               0,
                               0,
                               '/api/v1/admin/logout',
                               '退出',
                               '退出',
                               'POST'
                           );

INSERT INTO sys_permission (
                               id,
                               created_at,
                               updated_at,
                               deleted_at,
                               deleted,
                               disabled,
                               name,
                               display_name,
                               description,
                               act
                           )
                           VALUES (
                               2,
                               '2021-06-29 17:31:36.707018+08:00',
                               '2021-06-29 17:31:36.707018+08:00',
                               NULL,
                               0,
                               0,
                               '/api/v1/platform/view',
                               '平台浏览',
                               '',
                               'GET'
                           );

INSERT INTO sys_permission (
                               id,
                               created_at,
                               updated_at,
                               deleted_at,
                               deleted,
                               disabled,
                               name,
                               display_name,
                               description,
                               act
                           )
                           VALUES (
                               3,
                               '2021-06-29 17:31:36.707922+08:00',
                               '2021-06-29 17:31:36.707922+08:00',
                               NULL,
                               0,
                               0,
                               '/api/v1/test/tasks',
                               '任务列表',
                               '',
                               'GET'
                           );

INSERT INTO sys_permission (
                               id,
                               created_at,
                               updated_at,
                               deleted_at,
                               deleted,
                               disabled,
                               name,
                               display_name,
                               description,
                               act
                           )
                           VALUES (
                               4,
                               '2021-06-29 17:31:36.708817+08:00',
                               '2021-06-29 17:31:36.708817+08:00',
                               NULL,
                               0,
                               0,
                               '/api/v1/test/tasks/{id:uint}',
                               '任务详情',
                               '',
                               'GET'
                           );

INSERT INTO sys_permission (
                               id,
                               created_at,
                               updated_at,
                               deleted_at,
                               deleted,
                               disabled,
                               name,
                               display_name,
                               description,
                               act
                           )
                           VALUES (
                               5,
                               '2021-06-29 17:31:36.709703+08:00',
                               '2021-06-29 17:31:36.709703+08:00',
                               NULL,
                               0,
                               0,
                               '/api/v1/test/tasks',
                               '任务创建',
                               '',
                               'POST'
                           );

INSERT INTO sys_permission (
                               id,
                               created_at,
                               updated_at,
                               deleted_at,
                               deleted,
                               disabled,
                               name,
                               display_name,
                               description,
                               act
                           )
                           VALUES (
                               6,
                               '2021-06-29 17:31:36.710591+08:00',
                               '2021-06-29 17:31:36.710591+08:00',
                               NULL,
                               0,
                               0,
                               '/api/v1/test/tasks/{id:uint}',
                               '任务更新',
                               '',
                               'PUT'
                           );

INSERT INTO sys_permission (
                               id,
                               created_at,
                               updated_at,
                               deleted_at,
                               deleted,
                               disabled,
                               name,
                               display_name,
                               description,
                               act
                           )
                           VALUES (
                               7,
                               '2021-06-29 17:31:36.711538+08:00',
                               '2021-06-29 17:31:36.711538+08:00',
                               NULL,
                               0,
                               0,
                               '/api/v1/test/tasks/{id:uint}',
                               '任务删除',
                               '',
                               'DELETE'
                           );

INSERT INTO sys_permission (
                               id,
                               created_at,
                               updated_at,
                               deleted_at,
                               deleted,
                               disabled,
                               name,
                               display_name,
                               description,
                               act
                           )
                           VALUES (
                               8,
                               '2021-06-29 17:31:36.712389+08:00',
                               '2021-06-29 17:31:36.712389+08:00',
                               NULL,
                               0,
                               0,
                               '/api/v1/test/envs',
                               '测试环境',
                               '',
                               'POST'
                           );


-- 表：sys_role
DROP TABLE IF EXISTS sys_role;

CREATE TABLE sys_role (
    id           INTEGER,
    created_at   DATETIME,
    updated_at   DATETIME,
    deleted_at   DATETIME,
    deleted      NUMERIC       DEFAULT false,
    disabled     NUMERIC       DEFAULT false,
    name         VARCHAR (256) NOT NULL
                               UNIQUE,
    display_name VARCHAR (256),
    description  VARCHAR (256),
    PRIMARY KEY (
        id
    )
);

INSERT INTO sys_role (
                         id,
                         created_at,
                         updated_at,
                         deleted_at,
                         deleted,
                         disabled,
                         name,
                         display_name,
                         description
                     )
                     VALUES (
                         1,
                         '2021-06-29 17:31:36.713709+08:00',
                         '2021-06-29 17:31:36.713709+08:00',
                         NULL,
                         0,
                         0,
                         'administrator',
                         '超级管理员',
                         '超级管理员'
                     );


-- 表：sys_user
DROP TABLE IF EXISTS sys_user;

CREATE TABLE sys_user (
    id                 INTEGER,
    created_at         DATETIME,
    updated_at         DATETIME,
    deleted_at         DATETIME,
    name               VARCHAR (60)  NOT NULL,
    username           VARCHAR (60)  NOT NULL
                                     UNIQUE,
    password           VARCHAR (100),
    intro              VARCHAR (512) NOT NULL,
    avatar             LONGTEXT,
    token              VARCHAR (128),
    token_updated_time DATETIME,
    project_id         INTEGER,
    PRIMARY KEY (
        id
    )
);

INSERT INTO sys_user (
                         id,
                         created_at,
                         updated_at,
                         deleted_at,
                         name,
                         username,
                         password,
                         intro,
                         avatar,
                         token,
                         token_updated_time,
                         project_id
                     )
                     VALUES (
                         1,
                         '2021-06-25 10:12:28.742755+08:00',
                         '2021-11-19 13:16:03.315327+08:00',
                         NULL,
                         'admin',
                         'admin',
                         '$2a$10$MYfPM5RNRJKuSwH.T6fDOOiErjV46gua.s72ARU4OVJ/ze5wrNfEm',
                         '檀越',
                         'https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTIPbZRufW9zPiaGpfdXgU7icRL1licKEicYyOiace8QQsYVKvAgCrsJx1vggLAD2zJMeSXYcvMSkw9f4pw/132',
                         'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzQwMzE5MDQsImlhdCI6MTYzNDAyODMwNH0.ZE9wcuUWMYoFnhPUbc3xf0sGptKdkQVd1DfkJzwbR94',
                         '2021-10-12 16:45:04.401401+08:00',
                         0
                     );


-- 索引：idx_sys_user_deleted_at
DROP INDEX IF EXISTS idx_sys_user_deleted_at;

CREATE INDEX idx_sys_user_deleted_at ON sys_user (
    deleted_at
);


-- 索引：unique_index
DROP INDEX IF EXISTS unique_index;

CREATE UNIQUE INDEX unique_index ON biz_casbin_rules (
    p_type,
    v0,
    v1,
    v2,
    v3,
    v4,
    v5
);


COMMIT TRANSACTION;
PRAGMA foreign_keys = on;
