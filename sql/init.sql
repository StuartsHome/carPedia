CREATE DATABASE IF NOT EXISTS car_pedia;

USE car_pedia;

DROP TABLE IF EXISTS car_pedia.cars;
CREATE TABLE car_pedia.cars (
  id SERIAL PRIMARY KEY,
  make VARCHAR(256),
  model VARCHAR(256) 
);

INSERT INTO car_pedia.cars (make, model) VALUES
('citroen', 'c3'),
('ford', 'fiesta'),
('vw', 'golf'),
('vw', 'polo'),
('land rover', 'discovery sport'),
('ford', 'puma'),
('seat', 'ibiza');