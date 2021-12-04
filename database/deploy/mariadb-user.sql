create database toy;
CREATE USER 'kimdh'@'%' IDENTIFIED BY 'rlaejrgud';
GRANT ALL PRIVILEGES ON toy.* TO 'kimdh'@'%';
FLUSH PRIVILEGES;
