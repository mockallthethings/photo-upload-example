CREATE DATABASE IF NOT EXISTS photos;
USE photos;

CREATE TABLE IF NOT EXISTS dim_album (
  id CHAR(36),
  name TEXT
);

CREATE TABLE IF NOT EXISTS fact_photo_album (
  album_id CHAR(36),
  photo_id CHAR(36)
);
