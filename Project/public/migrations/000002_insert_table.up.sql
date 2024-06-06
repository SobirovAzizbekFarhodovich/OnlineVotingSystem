INSERT INTO party (id, name, slogan, opened_date, description) VALUES
('f47ac10b-58cc-4372-a567-0e02b2c3d479', 'Party A', 'For a Better Future', '2020-01-01', 'A party dedicated to economic growth.'),
('3bbcee75-cecc-5b56-8031-b6641c1ed1f1', 'Party B', 'Unity and Progress', '2021-05-15', 'A party focusing on social unity.'),
('8b4494b3-08d8-46fa-9040-5a42a8388f9e', 'Party C', 'Strength and Honor', '2019-07-20', 'A party with a focus on national defense.'),
('3ba6d49c-4795-4baf-8fda-e4a4ec9f9d9a', 'Party D', 'Innovation and Change', '2022-03-10', 'A party promoting technological advancement.'),
('091e5d28-5d1d-4d7b-9d49-5a0b9af76f5b', 'Party E', 'Green Earth', '2018-11-25', 'A party advocating for environmental sustainability.');

INSERT INTO public (id, first_name, last_name, birthday, gender, nation, party_id) VALUES
('f47ac10b-58cc-4372-a567-0e02b2c3d471', 'John', 'Doe', '1990-04-12', 'M', 'USA', 'f47ac10b-58cc-4372-a567-0e02b2c3d479'),
('3bbcee75-cecc-5b56-8031-b6641c1ed1f2', 'Jane', 'Smith', '1985-06-24', 'F', 'Canada', '3bbcee75-cecc-5b56-8031-b6641c1ed1f1'),
('8b4494b3-08d8-46fa-9040-5a42a8388f9f', 'Alice', 'Johnson', '1992-09-30', 'F', 'UK', '8b4494b3-08d8-46fa-9040-5a42a8388f9e'),
('3ba6d49c-4795-4baf-8fda-e4a4ec9f9d9f', 'Bob', 'Brown', '1978-02-15', 'M', 'Australia', '3ba6d49c-4795-4baf-8fda-e4a4ec9f9d9a'),
('091e5d28-5d1d-4d7b-9d49-5a0b9af76f5f', 'Charlie', 'Davis', '1980-12-10', 'M', 'USA', '091e5d28-5d1d-4d7b-9d49-5a0b9af76f5b'),
('d8a6e6e5-2ba7-4f60-8c24-4ed62c70ec2f', 'Diana', 'Garcia', '1995-03-21', 'F', 'Mexico', 'f47ac10b-58cc-4372-a567-0e02b2c3d479'),
('6dcb5a47-7644-48f1-9984-3c0d7a02e536', 'Edward', 'Martinez', '1988-07-08', 'M', 'Spain', '3bbcee75-cecc-5b56-8031-b6641c1ed1f1'),
('2e9b2c3b-1d72-45da-8d5a-2b5b2f2e3d8f', 'Fiona', 'Wilson', '1975-05-18', 'F', 'New Zealand', '8b4494b3-08d8-46fa-9040-5a42a8388f9e'),
('9b6d7d60-2a9d-4c7d-9e42-5e4a5b6c7d8e', 'George', 'Miller', '1999-11-03', 'M', 'South Africa', '3ba6d49c-4795-4baf-8fda-e4a4ec9f9d9a'),
('7f8c4e32-3c8f-4b8e-9b5e-8e5a4c7f3d6e', 'Hannah', 'Taylor', '1983-08-29', 'F', 'USA', '091e5d28-5d1d-4d7b-9d49-5a0b9af76f5b');
