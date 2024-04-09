CREATE DATABASE IF NOT EXISTS TokoBuku;

CREATE TABLE Buku (
    ID INT PRIMARY KEY,
    Judul VARCHAR(255),
    Pengarang VARCHAR(255),
    Harga DECIMAL(10, 2)
);
