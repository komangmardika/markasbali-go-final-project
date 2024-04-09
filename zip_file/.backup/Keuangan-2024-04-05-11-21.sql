CREATE DATABASE IF NOT EXISTS Keuangan;

CREATE TABLE Transaksi (
    ID INT PRIMARY KEY,
    Tanggal DATE,
    Keterangan VARCHAR(255),
    Jumlah DECIMAL(10, 2)
);