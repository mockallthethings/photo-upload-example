CREATE DATABASE IF NOT EXISTS photos;
USE photos;

CREATE TABLE IF NOT EXISTS albums (
  album_name TEXT,
  photo_id CHAR(36)
);
