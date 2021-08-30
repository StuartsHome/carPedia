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

DROP TABLE IF EXISTS car_pedia.carMake;
CREATE TABLE car_pedia.carMake (
  id INT PRIMARY KEY AUTO_INCREMENT,
  make VARCHAR(256) NOT NULL,
  description VARCHAR(256)
  -- FOREIGN KEY (make) REFERENCES cars(make)
);

INSERT INTO car_pedia.carMake (make, description) VALUES
('citroen', 'french car manafacturer'),
('ford', 'american car manafacturer'),
('vw', 'german car manafacturer'),
('land rover', 'british car manafacturer'),
('seat', 'spanish car manafacturer');
