INSERT INTO Country (cname, population) VALUES
('Kazakhstan', 19000000),
('USA', 331000000),
('India', 1393409038),
('Russia', 146000000),
('China', 1440000000),
('Germany', 83000000),
('France', 67000000),
('Brazil', 213000000),
('Japan', 126000000),
('Canada', 38000000);

INSERT INTO Users (email, name, surname, salary, phone, cname) VALUES
('syrymbek@nu.edu.kz', 'Syrymbek', 'Akhmetov', 60000, '87015555555', 'Kazakhstan'),
('gulnara@nu.edu.kz', 'Gulnara', 'Iskakova', 45000, '87017777777', 'Kazakhstan'),
('bekzhan@nu.edu.kz', 'Bekzhan', 'Nurtayev', 55000, '87018888888', 'Kazakhstan'),
('aida@nu.edu.kz', 'Aida', 'Tursynova', 50000, '87027777777', 'USA'),
('yerzhan@nu.edu.kz', 'Yerzhan', 'Tleukhan', 60000, '87014444444', 'Russia'),
('dias@nu.edu.kz', 'Dias', 'Aytbayev', 72000, '87017775555', 'India'),
('aliya@nu.edu.kz', 'Aliya', 'Mukasheva', 40000, '87016666666', 'Germany'),
('adil@nu.edu.kz', 'Adil', 'Karimov', 80000, '87015551111', 'Kazakhstan'),
('sanzhar@nu.edu.kz', 'Sanzhar', 'Omarov', 48000, '87019999999', 'China'),
('marzhan@nu.edu.kz', 'Marzhan', 'Zhanibekova', 46000, '87012225555', 'France');

INSERT INTO Patients (email) VALUES
('syrymbek@nu.edu.kz'),
('gulnara@nu.edu.kz'),
('bekzhan@nu.edu.kz'),
('aida@nu.edu.kz'),
('yerzhan@nu.edu.kz'),
('dias@nu.edu.kz'),
('aliya@nu.edu.kz'),
('adil@nu.edu.kz'),
('sanzhar@nu.edu.kz'),
('marzhan@nu.edu.kz');

INSERT INTO DiseaseType (id, description) VALUES
(1, 'Infectious Diseases'),
(2, 'Virology'),
(3, 'Bacterial Infections'),
(4, 'Chronic Illnesses'),
(5, 'Cancer'),
(6, 'Neurological Diseases'),
(7, 'Cardiovascular Diseases'),
(8, 'Autoimmune Diseases'),
(9, 'Endocrine Disorders'),
(10, 'Parasitic Diseases');

INSERT INTO Disease (disease_code, pathogen, description, id) VALUES
('TB', 'bacteria', 'Tuberculosis is a bacterial infection', 3),
('CHOL', 'bacteria', 'Cholera is a severe bacterial infection', 3),
('FLU', 'virus', 'Seasonal flu', 2),
('COVID-19', 'virus', 'Coronavirus disease', 2),
('HIV', 'virus', 'AIDS-causing virus', 1),
('Ebola', 'virus', 'Hemorrhagic fever', 2),
('Asthma', 'none', 'Chronic respiratory disease', 4),
('Diabetes', 'none', 'Blood sugar disorder', 9),
('Hypertension', 'none', 'High blood pressure', 7),
('Malaria', 'parasite', 'Mosquito-borne parasitic disease', 10);

INSERT INTO Discover (cname, disease_code, first_enc_date) VALUES
('Kazakhstan', 'TB', '2015-06-01'),
('USA', 'CHOL', '2000-01-01'),
('India', 'FLU', '1998-11-20'),
('Russia', 'COVID-19', '2019-12-01'),
('China', 'HIV', '1985-07-10'),
('France', 'Ebola', '2014-08-01'),
('Germany', 'Asthma', '2005-06-30'),
('Brazil', 'Diabetes', '1995-09-01'),
('Japan', 'Hypertension', '1990-03-01'),
('Canada', 'Malaria', '2010-07-15');

INSERT INTO PatientDisease (email, disease_code) VALUES
('syrymbek@nu.edu.kz', 'COVID-19'),
('gulnara@nu.edu.kz', 'TB'),
('bekzhan@nu.edu.kz', 'CHOL'),
('aida@nu.edu.kz', 'FLU'),
('yerzhan@nu.edu.kz', 'Ebola'),
('dias@nu.edu.kz', 'HIV'),
('aliya@nu.edu.kz', 'Asthma'),
('adil@nu.edu.kz', 'Diabetes'),
('sanzhar@nu.edu.kz', 'Hypertension'),
('marzhan@nu.edu.kz', 'Malaria');

INSERT INTO PublicServant (email, department) VALUES
('syrymbek@nu.edu.kz', 'Health'),
('gulnara@nu.edu.kz', 'Research'),
('bekzhan@nu.edu.kz', 'Education'),
('aida@nu.edu.kz', 'Defense'),
('yerzhan@nu.edu.kz', 'Health'),
('dias@nu.edu.kz', 'Finance'),
('aliya@nu.edu.kz', 'Education'),
('adil@nu.edu.kz', 'Research'),
('sanzhar@nu.edu.kz', 'Health'),
('marzhan@nu.edu.kz', 'Defense');

INSERT INTO Doctor (email, degree) VALUES
('syrymbek@nu.edu.kz', 'MD'),
('gulnara@nu.edu.kz', 'PhD'),
('bekzhan@nu.edu.kz', 'MBBS'),
('aida@nu.edu.kz', 'PhD'),
('yerzhan@nu.edu.kz', 'MD'),
('dias@nu.edu.kz', 'PhD'),
('aliya@nu.edu.kz', 'MBBS'),
('adil@nu.edu.kz', 'MD'),
('sanzhar@nu.edu.kz', 'PhD'),
('marzhan@nu.edu.kz', 'MD');

INSERT INTO Specialize (id, email) VALUES
(1, 'syrymbek@nu.edu.kz'),
(2, 'syrymbek@nu.edu.kz'),
(3, 'syrymbek@nu.edu.kz'),
(2, 'gulnara@nu.edu.kz'),
(3, 'gulnara@nu.edu.kz'),
(1, 'bekzhan@nu.edu.kz'),
(7, 'aida@nu.edu.kz'),
(2, 'yerzhan@nu.edu.kz'),
(9, 'dias@nu.edu.kz'),
(6, 'aliya@nu.edu.kz');

INSERT INTO Record (email, cname, disease_code, total_deaths, total_patients) VALUES
('syrymbek@nu.edu.kz', 'Kazakhstan', 'COVID-19', 500, 5000),
('gulnara@nu.edu.kz', 'Kazakhstan', 'TB', 200, 3000),
('bekzhan@nu.edu.kz', 'USA', 'CHOL', 100, 2000),
('aida@nu.edu.kz', 'India', 'FLU', 50, 1000),
('yerzhan@nu.edu.kz', 'Russia', 'Ebola', 800, 4000),
('dias@nu.edu.kz', 'Germany', 'Asthma', 30, 1200),
('aliya@nu.edu.kz', 'France', 'Hypertension', 100, 800),
('adil@nu.edu.kz', 'China', 'Diabetes', 200, 1500),
('sanzhar@nu.edu.kz', 'Japan', 'Malaria', 300, 900),
('marzhan@nu.edu.kz', 'Brazil', 'HIV', 400, 2500);