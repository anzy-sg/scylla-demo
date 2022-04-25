CREATE KEYSPACE mailsettings WITH replication = {'class': 'NetworkTopologyStrategy', 'mdw1':3, 'las1':3, 'iad2':3};

use mailsettings;


-- user_id, require_tls, require_valid_cert, version
create table user_tls_policy (
                                 user_id bigint,
                                 require_tls boolean,
                                 require_valid_cert boolean,
                                 version double,
                                 PRIMARY KEY (user_id)
);



insert into user_tls_policy (user_id, require_tls, require_valid_cert, version) values (1, true, true, 1.0);
insert into user_tls_policy (user_id, require_tls, require_valid_cert, version) values (2, false, false, 1.2);
insert into user_tls_policy (user_id, require_tls, require_valid_cert, version) values (3, true, false, 1.3);

select user_id, require_tls, require_valid_cert, version from user_tls_policy
where user_id = 1;

SELECT user_id, require_tls, require_valid_cert, version
FROM mailsettings.user_tls_policy;