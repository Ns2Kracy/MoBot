CREATE TABLE user(
    id INT NOT NULL,
    access_token longtext NULL ,
    refresh_token longtext NULL ,
    join_date datetime NULL ,
    main_mode    int null ,
    osu_id bigint null ,
    qq  bigint null ,
    expire_in bigint null ,
    PRIMARY KEY (id)
);

CREATE index bind_oid
    on user(osu_id);
CREATE index bind_qid
    on user(qq);