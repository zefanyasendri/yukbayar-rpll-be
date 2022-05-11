CREATE TABLE `diskon` (
  `id` varchar(6) NOT NULL,
  `kode` varchar(255) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `id_pengguna` varchar(12) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `kategorilayanan` (
  `id` varchar(6) NOT NULL,
  `kategori` varchar(255) NOT NULL,
  `id_layanan` varchar(6) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `layanan` (
  `id` varchar(6) NOT NULL,
  `id_mitra` varchar(6) NOT NULL,
  `nama` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `mitra` (
  `id` varchar(6) NOT NULL,
  `alamat` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `noTelpon` varchar(255) NOT NULL,
  `pemilikBisnis` varchar(255) NOT NULL,
  `bidangUsaha` varchar(255) NOT NULL,
  `badanUsaha` varchar(255) NOT NULL,
  `namaPerusahaan` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `pengguna` (
  `id` varchar(32) NOT NULL,
  `email` varchar(255) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `noTelpon` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `tglLahir` varchar(10) NOT NULL,
  `gender` varchar(20) NOT NULL,
  `saldoYukPay` int(10) NOT NULL,
  `tipepengguna` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `transaksi` (
  `id_transaksi` varchar(6) NOT NULL,
  `id_pengguna` varchar(32) NOT NULL,
  `totalHarga` int(10) NOT NULL,
  `tanggal` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `id_varian` varchar(6) NOT NULL,
  `status` varchar(12) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `varian` (
  `id` varchar(6) NOT NULL,
  `harga` int(10) NOT NULL,
  `jenis` varchar(255) NOT NULL,
  `id_kategori` varchar(6) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `wallettopup` (
  `id` varchar(6) NOT NULL,
  `kodeYukPay` varchar(255) NOT NULL,
  `metode` varchar(255) NOT NULL,
  `nominal` float NOT NULL,
  `tanggal` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `id_pengguna` varchar(32) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE `diskon`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_pengguna` (`id_pengguna`);

ALTER TABLE `kategorilayanan`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_varian` (`id_layanan`);

ALTER TABLE `layanan`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_mitra` (`id_mitra`);

ALTER TABLE `mitra`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `pengguna`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `transaksi`
  ADD PRIMARY KEY (`id_transaksi`),
  ADD KEY `transaksi_ibfk_1` (`id_varian`),
  ADD KEY `transaksi_ibfk_2` (`id_pengguna`);

ALTER TABLE `varian`
  ADD PRIMARY KEY (`id`),
  ADD KEY `varian_ibfk_1` (`id_kategori`);

ALTER TABLE `wallettopup`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_pengguna` (`id_pengguna`);

ALTER TABLE `diskon`
  ADD CONSTRAINT `diskon_ibfk_1` FOREIGN KEY (`id_pengguna`) REFERENCES `pengguna` (`id`);

ALTER TABLE `kategorilayanan`
  ADD CONSTRAINT `kategorilayanan_ibfk_1` FOREIGN KEY (`id_layanan`) REFERENCES `layanan` (`id`);

ALTER TABLE `layanan`
  ADD CONSTRAINT `layanan_ibfk_1` FOREIGN KEY (`id_mitra`) REFERENCES `mitra` (`id`);

ALTER TABLE `transaksi`
  ADD CONSTRAINT `transaksi_ibfk_1` FOREIGN KEY (`id_varian`) REFERENCES `varian` (`id`),
  ADD CONSTRAINT `transaksi_ibfk_2` FOREIGN KEY (`id_pengguna`) REFERENCES `pengguna` (`id`);

ALTER TABLE `varian`
  ADD CONSTRAINT `varian_ibfk_1` FOREIGN KEY (`id_kategori`) REFERENCES `kategorilayanan` (`id`);

ALTER TABLE `wallettopup`
  ADD CONSTRAINT `wallettopup_ibfk_1` FOREIGN KEY (`id_pengguna`) REFERENCES `pengguna` (`id`);
COMMIT;