CREATE TABLE user(
    id              INT NOT NULL,
    access_token    LONGTEXT NULL ,
    refresh_token   LONGTEXT NULL ,
    join_date       DATETIME NULL ,
    main_mode       INT NULL DEFAULT 0,
    osu_id          BIGINT NULL ,
    qq              BIGINT NULL ,
    expire_in       BIGINT NULL ,
    PRIMARY KEY (id)
);

CREATE INDEX bind_oid
    ON user(osu_id);
CREATE INDEX bind_qid
    ON user(qq);