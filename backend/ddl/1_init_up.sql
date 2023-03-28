SET GLOBAL time_zone = 'Asia/Tokyo';

SET time_zone = 'Asia/Tokyo';

CREATE USER IF NOT EXISTS 'dab_user' @'%';

SET PASSWORD FOR 'dab_user' @'%' = 'passw0rd';

GRANT ALL ON *.* TO 'dab_user';
